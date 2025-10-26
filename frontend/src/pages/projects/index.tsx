import {useState} from "preact/hooks";
import {Button, ButtonType} from "../../components/button";
import {useProjectContext} from "../../contexts/projects.context";
import type {ICreateProjectRequest} from "../../models/projects.models";
import styles from "./style.module.scss";
import {Modal} from "../../components/modal";
import {CreateProjectForm} from "../../components/projects/createProjectForm";
import {ProjectVisibilityType} from "../../enums/projects.enum.ts";
import {ProjectsTable} from "../../components/projects/projectsTable";

export interface ICreateProjectFormData {
    projectName: string;
    description: string;
    visibility: ProjectVisibilityType;
    connectionType: string;
}

export const ProjectsPage = () => {
    const {projects, createProject} = useProjectContext();

    const [isCreateProjectModalOpen, setIsCreateProjectModalOpen] = useState(false);

    const openCreateProjectModal = () => {
        setIsCreateProjectModalOpen(true);
    }

    const closeCreateProjectModal = () => {
        setIsCreateProjectModalOpen(false);
    }

    const [newProjectData, setNewProjectData] = useState<ICreateProjectFormData>({
        projectName: "",
        description: "",
        visibility: ProjectVisibilityType.PRIVATE,
        connectionType: "",
    });

    const handleSubmit = () => {
        const cpr: ICreateProjectRequest = {
            name: newProjectData.projectName,
            description: newProjectData.description,
            visibility: newProjectData.visibility,
            connectionType: Number(newProjectData.connectionType),
        }

        createProject(cpr).then((res) => {
            if(res){
                clearCreateForm();
                closeCreateProjectModal();
            }
        });
    }

    const clearCreateForm = () => {
        setNewProjectData({
            projectName: "",
            description: "",
            visibility: ProjectVisibilityType.PRIVATE,
            connectionType: "",
        });
    }

    return (
        <div className={styles.project_page}>
            <div className={styles.header}>
                <h1 className={styles.title}>Projects</h1>

                <div className={styles.button_wrapper}>
                    <Button
                        text="New Project +"
                        callback={() => openCreateProjectModal()}
                        ariaLabel="new project"
                        type={ButtonType.Default}
                    />
                </div>
            </div>

            <div className={styles.table_wrapper}>
                <ProjectsTable projects={projects}/>
            </div>

            {isCreateProjectModalOpen && (
                <Modal
                    title="Create New Project"
                    hint="Don’t worry, you can always refactor it later... probably."
                    content={
                        <CreateProjectForm
                            newProjectData={newProjectData}
                            setNewProjectData={setNewProjectData}
                        />
                    }
                    footerType="double"
                    isOpen
                    submitButtonText="Create Project"
                    onSubmit={handleSubmit}
                    cancelButtonText="Cancel"
                    onCancel={closeCreateProjectModal}
                    onClose={closeCreateProjectModal}
                />
            )}
        </div>
    )
}