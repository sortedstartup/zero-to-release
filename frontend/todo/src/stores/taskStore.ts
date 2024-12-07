import { atom } from 'nanostores';
import axios from 'axios';

const API_BASE_URL = 'http://localhost:8000'; // Replace with your backend URL

// Define Task interface
export interface Task {
    id: number;
    title: string;
    description: string;
    created_at: string;
    updated_at: string;
}

export const taskStore = atom<Task[]>([]);

// Fetch tasks
export const fetchTasks = async (): Promise<void> => {
    try {
        const response = await axios.get<Task[]>(`${API_BASE_URL}/tasks`);
        taskStore.set(response.data);
    } catch (error) {
        console.error('Error fetching tasks:', error);
    }
};

// Add task
export const addTask = async (task: Omit<Task, 'id' | 'created_at' | 'updated_at'>): Promise<void> => {
    try {
        const response = await axios.post<Task>(`${API_BASE_URL}/tasks`, task);
        taskStore.set([...taskStore.get(), response.data]);
    } catch (error) {
        console.error('Error adding task:', error);
    }
};

// Update task
export const updateTask = async (id: number, updatedTask: Partial<Task>): Promise<void> => {
    try {
        await axios.put(`${API_BASE_URL}/tasks/${id}`, updatedTask);
        const updatedList = taskStore.get().map((task) =>
            task.id === id ? { ...task, ...updatedTask } : task
        );
        taskStore.set(updatedList);
    } catch (error) {
        console.error('Error updating task:', error);
    }
};

// Delete task
export const deleteTask = async (id: number): Promise<void> => {
    try {
        await axios.delete(`${API_BASE_URL}/tasks/${id}`);
        taskStore.set(taskStore.get().filter((task) => task.id !== id));
    } catch (error) {
        console.error('Error deleting task:', error);
    }
};
