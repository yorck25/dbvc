import {useState} from "preact/hooks";
import {Button, ButtonType} from "../../components/button";
import {useProjectContext} from "../../contexts/projects.context";
import type {
    ICreateProjectCredentialsRequest,
    ICreateProjectMembersRequest,
    ICreateProjectMetadataRequest,
    ICreateProjectRequest
} from "../../models/projects.models";
import styles from "./style.module.scss";
import {Modal} from "../../components/modal";
import {CreateProjectForm} from "../../components/projects/createProjectForm";
import {ProjectVisibilityType} from "../../enums/projects.enum.ts";
import {ProjectsTable} from "../../components/projects/projectsTable";
import React from "react";
import {CreateProjectCredentialsForm} from "../../components/projects/createProjectCredentailsForm";
import {AddMembersForm} from "../../components/projects/addMembersForm";
import type {IMemberRequest} from "../../models/user.models.ts";
import type {DatabaseAuthData} from "../../components/projects/createProjectCredentailsForm/databaseCredentialsForm";

export interface ICreateProjectFormData {
    projectName: string;
    description: string;
    visibility: ProjectVisibilityType;
    connectionType: string;
}

export interface ICreateProjectCredentialsData {
    projectPassword: string;
    databaseAuth: any;
}

export const ProjectsPage = () => {
    const {projects, createProject} = useProjectContext();

    const [newProjectData, setNewProjectData] = useState<ICreateProjectFormData>({
        projectName: "",
        description: "",
        visibility: ProjectVisibilityType.PRIVATE,
        connectionType: "",
    });

    const [newCredentialsData, setNewCredentialsData] = useState<ICreateProjectCredentialsData>({
        projectPassword: "",
        databaseAuth: {
            type: "psql",
            host: "",
            port: 5432,
            username: "",
            password: "",
            databaseName: "",
        },
    });

    const [isCreateProjectModalOpen, setIsCreateProjectModalOpen] = useState(true);
    const maxSteps = 2;
    const [step, setStep] = useState<number>(1);
    const [members, setMembers] = useState<IMemberRequest[]>([]);

    const openCreateProjectModal = () => {
        setIsCreateProjectModalOpen(true);
    }

    const closeCreateProjectModal = () => {
        setIsCreateProjectModalOpen(false);
    }

    const handleSubmit = () => {

        const cpmr: ICreateProjectMetadataRequest = {
            name: newProjectData.projectName,
            description: newProjectData.description,
            visibility: newProjectData.visibility,
            connectionType: Number(newProjectData.connectionType),
        }

        const cpcr : ICreateProjectCredentialsRequest = {
            projectPassword: newCredentialsData.projectPassword,
            databaseAuth: newCredentialsData.databaseAuth,
        }

        const cpmer: ICreateProjectMembersRequest = {
            members: members.map((member) => member.id),
        }

        const cpr: ICreateProjectRequest = {
            metadata: cpmr,
            credentials: cpcr,
            members: cpmer,
        }

        createProject(cpr).then((res) => {
            if (res) {
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

    const getSubmitButtonText = (): string => {
        if (step === maxSteps) {
            return "Create Project";
        }

        return "Next";
    }

    const primaryButtonAction = (): void => {
        if (step === maxSteps) {
            return handleSubmit();
        }

        return setStep(prev => prev + 1);
    }

    const getCancelButtonText = (): string => {
        if (step === 0) {
            return "Cancel";
        }

        return "Back";
    }

    const cancelButtonAction = (): void => {
        if (step === 0) {
            return closeCreateProjectModal();
        }

        return setStep(prev => prev - 1);
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
                        <ProjectStepper
                            step={step}
                            members={members}
                            setMembers={setMembers}
                            newProjectData={newProjectData}
                            setNewProjectData={setNewProjectData}
                            newCredentialsData={newCredentialsData}
                            setNewCredentialsData={setNewCredentialsData}
                        />
                    }
                    footerType="double"
                    isOpen
                    submitButtonText={getSubmitButtonText()}
                    onSubmit={primaryButtonAction}
                    cancelButtonText={getCancelButtonText()}
                    onCancel={cancelButtonAction}
                    onClose={closeCreateProjectModal}
                />
            )}
        </div>
    )
}

interface IProjectStepperProps {
    step: number;
    members: IMemberRequest[];
    setMembers: React.Dispatch<React.SetStateAction<IMemberRequest[]>>;

    newProjectData: ICreateProjectFormData,
    setNewProjectData: React.Dispatch<React.SetStateAction<ICreateProjectFormData>>;

    newCredentialsData: ICreateProjectCredentialsData,
    setNewCredentialsData: React.Dispatch<React.SetStateAction<ICreateProjectCredentialsData>>;
}

const ProjectStepper = (props: IProjectStepperProps) => {
    return (
        <>

            {props.step === 0 && (
                <CreateProjectForm
                    newProjectData={props.newProjectData}
                    setNewProjectData={props.setNewProjectData}
                />
            )}

            {props.step == 1 && (
                <CreateProjectCredentialsForm connectionType={props.newProjectData.connectionType}
                                              newCredentialsData={props.newCredentialsData}
                                              setNewCredentialsData={props.setNewCredentialsData}/>
            )}

            {props.step == 2 && (
                <AddMembersForm  members={props.members} setMembers={props.setMembers}/>
            )}
        </>
    )
}