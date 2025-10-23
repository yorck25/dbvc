import {useState} from "preact/hooks";
import {Button, ButtonType} from "../../components/button";
import {useProjectContext} from "../../contexts/projects.context";
import type {Project} from "../../models/projects.models";
import styles from "./style.module.scss";

export const ProjectsPage = () => {
    const {projects} = useProjectContext();

    const [isCreateProjectModalOpen, setIsCreateProjectModalOpen] = useState(false);

    const openCreateProjectModal = () => {
        setIsCreateProjectModalOpen(true);
    }

    return (
        <div>
            <div className={styles.create_button_wrapper}>
                <Button text={"Click Me!"} callback={() => openCreateProjectModal()}
                        ariaLabel={"new Button"} type={ButtonType.Default}>
                </Button>
            </div>


            {projects?.map((project: Project) => (
                <ProjectCell project={project}/>
            ))
            }
        </div>
    )
}

const ProjectCell = ({project}: { project: Project }) => {
    return (
        <div key={project.id}>
            <h2>{project.name}</h2>
            <p>{project.description}</p>
        </div>
    )
}