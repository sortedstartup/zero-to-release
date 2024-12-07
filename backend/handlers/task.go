package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"sortedstartup.com/zero-to-release/models"

	"github.com/gorilla/mux"
)

func RegisterTaskHandlers(r *mux.Router, db *sql.DB) {
	h := &TaskHandler{db: db}
	r.HandleFunc("/tasks", h.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", h.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", h.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", h.DeleteTask).Methods("DELETE")
}

type TaskHandler struct {
	db *sql.DB

	//temporary remove once db works
	tasks []models.Task
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task = []models.Task{}
	rows, err := h.db.Query("SELECT id, title, description, created_at, updated_at FROM tasks")
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(tasks)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt); err != nil {
			log.Println(err)
		}
		tasks = append(tasks, task)
	}
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	log.Println("I am in Create task handler")
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.db.QueryRow("INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id",
		task.Title, task.Description).Scan(&task.ID)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := h.db.Exec("UPDATE tasks SET title=$1, description=$2, updated_at=CURRENT_TIMESTAMP WHERE id=$3",
		task.Title, task.Description, id)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	_, err := h.db.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusNoContent)
}
