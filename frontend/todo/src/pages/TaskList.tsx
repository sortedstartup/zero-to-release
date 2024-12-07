import React, { useEffect } from 'react';
import {
    IonPage,
    IonHeader,
    IonToolbar,
    IonTitle,
    IonContent,
    IonButton,
    IonList,
    IonItem,
    IonLabel,
    IonButtons,
} from '@ionic/react';
import { useStore } from '@nanostores/react';
import { taskStore, fetchTasks, deleteTask, Task } from '../stores/taskStore';
import { useHistory } from 'react-router-dom';

const TaskList: React.FC = () => {
    const tasks = useStore(taskStore);
    const history = useHistory();

    useEffect(() => {
        fetchTasks();
    }, []);

    return (
        <IonPage>
            <IonHeader>
                <IonToolbar>
                    <IonTitle>Task Manager</IonTitle>
                </IonToolbar>
            </IonHeader>
            <IonContent>
                <IonButton expand="block" onClick={() => history.push('/add')}>
                    Add Task
                </IonButton>
                <IonList>
                    {tasks.map((task: Task) => (
                        <IonItem key={task.id}>
                            <IonLabel>
                                <h2>{task.title}</h2>
                                <p>{task.description}</p>
                            </IonLabel>
                            <IonButtons slot="end">
                                <IonButton
                                    color="primary"
                                    onClick={() => history.push(`/edit/${task.id}`)}
                                >
                                    Edit
                                </IonButton>
                                <IonButton color="danger" onClick={() => deleteTask(task.id)}>
                                    Delete
                                </IonButton>
                            </IonButtons>
                        </IonItem>
                    ))}
                </IonList>
            </IonContent>
        </IonPage>
    );
};

export default TaskList;
