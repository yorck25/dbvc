import styles from './style.module.scss';
import {useParams} from "react-router-dom";
import {useProjectContext} from "../../contexts/projects.context.tsx";
import {useEffect, useState} from "react";
import type {IProjectWithUsers} from "../../models/projects.models.ts";

export const ProjectDashboard = () => {
    const {projectId} = useParams();
    const {projects, getProjectById, ensureProjectLoaded} = useProjectContext();
    const [currentProject, setCurrentProject] = useState<IProjectWithUsers | undefined>();
    const [activeTab, setActiveTab] = useState<string>("Overview");


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

        console.log(currentProject);
    }, [projectId, projects]);

    const getNavigationItemClass = (id: string) => {
        return `${styles.navigation_bar_item_list_item} ${
            activeTab === id ? styles.active_tab : ""
        }`
    }

    return (
        <div className={styles.project_dashboard_page}>
            <div className={styles.header}>
                <h1 className={styles.title}>Project Dashboard</h1>
            </div>

            <nav className={styles.navigation_bar}>
                <ul className={styles.navigation_bar_item_list}>
                    <li
                        className={getNavigationItemClass("Overview")}
                        onClick={() => setActiveTab("Overview")}
                    >
                        Overview
                    </li>
                    <li
                        className={getNavigationItemClass("Database")}
                        onClick={() => setActiveTab("Database")}
                    >
                        Database
                    </li>
                    <li
                        className={getNavigationItemClass("Members")}
                        onClick={() => setActiveTab("Members")}
                    >
                        Members
                    </li>
                    <li
                        className={getNavigationItemClass("Tasks")}
                        onClick={() => setActiveTab("Tasks")}
                    >
                        Tasks (+99)
                    </li>
                    <li
                        className={getNavigationItemClass("Discussion")}
                        onClick={() => setActiveTab("Discussion")}
                    >
                        Discussion (05)
                    </li>
                    <li
                        className={getNavigationItemClass("Settings")}
                        onClick={() => setActiveTab("Settings")}
                    >
                        Settings
                    </li>
                </ul>
            </nav>

            <div className={styles.content_container}>
                {activeTab === "Overview" && (
                    <>Not implemented yet.</>
                )}

                {activeTab === "Database" && (
                    <ProjectDashboardDatabase />
                )}

                {activeTab === "Members" && (
                    <>Not implemented yet.</>
                )}

                {activeTab === "Tasks" && (
                    <>Not implemented yet.</>
                )}

                {activeTab === "Discussion" && (
                    <>Not implemented yet.</>
                )}

                {activeTab === "Settings" && (
                    <>Not implemented yet.</>
                )}
            </div>
        </div>
    );
};

export const ProjectDashboardDatabase = () => {
    return (
        <div>
        </div>
    )
}