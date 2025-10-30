import styles from './style.module.scss';
import {useParams} from "react-router-dom";
import {useProjectContext} from "../../contexts/projects.context.tsx";
import {useEffect, useState} from "react";
import type {IProjectWithUsers} from "../../models/projects.models.ts";

export const ProjectDashboard = () => {
    const { projectId } = useParams();
    const { projects, getProjectById, ensureProjectLoaded } = useProjectContext();
    const [currentProject, setCurrentProject] = useState<IProjectWithUsers | undefined>();

    useEffect(() => {
        if (!projectId) return;

        const id = Number(projectId);

        if (projects === undefined) return;

        const existing = getProjectById(id);
        if (existing) {
            setCurrentProject(existing);
        } else {
            ensureProjectLoaded(id).then((fetched) => {
                if (fetched) setCurrentProject(fetched);
            });
        }
    }, [projectId, projects]);

    return (
        <div className={styles.project_dashboard_page}>
            <div className={styles.header}>
                <h1 className={styles.title}>Project Dashboard</h1>
            </div>

            {currentProject ? (
                <p>{currentProject.project?.name ?? currentProject.project.name ?? "No name"}</p>
            ) : (
                <p>Loading...</p>
            )}
        </div>
    );
};
