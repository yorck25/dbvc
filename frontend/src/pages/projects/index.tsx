import {useState} from "preact/hooks";
import {Button, ButtonType} from "../../components/button";
import {useProjectContext} from "../../contexts/projects.context";
import type {ICreateProjectRequest, IProject} from "../../models/projects.models";
import styles from "./style.module.scss";
import {Modal} from "../../components/modal";
import {CreateProjectForm} from "../../components/projects/createProjectForm";
import {ProjectVisibilityType} from "../../enums/projects.enum.ts";

export interface ICreateProjectFormData {
    projectName: string;
    description: string;
    visibility: ProjectVisibilityType;
    connectionType: string;
}

export const ProjectsPage = () => {
    const {projects, createProject} = useProjectContext();

    const [isCreateProjectModalOpen, setIsCreateProjectModalOpen] = useState(true);

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
        <div>
            <div className={styles.create_button_wrapper}>
                <Button text={"New Project +"} callback={() => openCreateProjectModal()}
                        ariaLabel={"new Button"} type={ButtonType.Default}>
                </Button>
            </div>

            {projects?.map((project: IProject) => (
                <ProjectCell project={project}/>
            ))}

            {isCreateProjectModalOpen && (
                <Modal
                    title="Create New Project"
                    hint="Don’t worry, you can always refactor it later... probably."
                    content={
                        <div>
                            <CreateProjectForm newProjectData={newProjectData} setNewProjectData={setNewProjectData}/>
                        </div>
                    }
                    footerType="double"
                    isOpen={true}
                    submitButtonText={"Create Project"}
                    onSubmit={handleSubmit}
                    cancelButtonText="Cancel"
                    onCancel={closeCreateProjectModal}
                    onClose={closeCreateProjectModal}
                    primaryButtonText="Got it"
                />
            )}
        </div>
    )
}

const ProjectCell = ({project}: { project: IProject }) => {
    return (
        <div key={project.id}>
            <h2>{project.name}</h2>
            <p>{project.description}</p>
        </div>
    )
}