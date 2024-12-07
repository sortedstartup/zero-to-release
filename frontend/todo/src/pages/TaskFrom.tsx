import React, { useState, useEffect } from 'react';
import {
    IonPage,
    IonHeader,
    IonToolbar,
    IonTitle,
    IonContent,
    IonInput,
    IonTextarea,
    IonButton,
} from '@ionic/react';
import { useParams, useHistory } from 'react-router-dom';
import { addTask, updateTask, taskStore, Task } from '../stores/taskStore';
import { useStore } from '@nanostores/react';

const TaskForm: React.FC = () => {
    const { id } = useParams<{ id?: string }>();
    const history = useHistory();
    const tasks = useStore(taskStore);

    const [title, setTitle] = useState<string>('');
    const [description, setDescription] = useState<string>('');

    useEffect(() => {
        if (id) {
            const task = tasks.find((task: Task) => task.id === parseInt(id));
            if (task) {
                setTitle(task.title);
                setDescription(task.description);
            }
        }
    }, [id, tasks]);

    const handleSubmit = async (): Promise<void> => {
        const taskData = { title, description };
        if (id) {
            await updateTask(parseInt(id), taskData);
        } else {
            await addTask(taskData);
        }
        history.push('/');
    };

    return (
        <IonPage>
            <IonHeader>
                <IonToolbar>
                    <IonTitle>{id ? 'Edit Task' : 'Add Task'}</IonTitle>
                </IonToolbar>
            </IonHeader>
            <IonContent>
                <IonInput
                    placeholder="Title"
                    value={title}
                    onIonChange={(e) => setTitle(e.detail.value!)}
                />
                <IonTextarea
                    placeholder="Description"
                    value={description}
                    onIonChange={(e) => setDescription(e.detail.value!)}
                />
                <IonButton expand="block" onClick={handleSubmit}>
                    Save
                </IonButton>
            </IonContent>
        </IonPage>
    );
};

export default TaskForm;
