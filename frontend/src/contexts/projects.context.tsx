import {createContext, useContext, useEffect, useState, type Dispatch, type FC, type ReactNode} from 'react';
import type {ICreateProjectRequest, IProject} from '../models/projects.models';
import {NetworkAdapter, setAuthHeader} from "../lib/networkAdapter.tsx";

interface IProjectContext {
    projects: any;
    setProjects: Dispatch<any>;

    createProject: (cpr: ICreateProjectRequest) => Promise<boolean>;
}

const ProjectContext = createContext<IProjectContext | undefined>(undefined);

export const ProjectContextProvider: FC<{ children: ReactNode }> = ({children}) => {
    const [projects, setProjects] = useState<IProject[]>();

    useEffect(() => {
        console.log("projects context");
        fetchProjects();
    }, []);

    const fetchProjects = () => {
        fetch('http://localhost:8080/projects')
            .then(res => res.json())
            .then((data: IProject[]) => {
                setProjects(data);
            });
    }

    const createProject = async (cpr: ICreateProjectRequest): Promise<boolean> => {
        try {
            const myHeaders = setAuthHeader();
            myHeaders.append("Content-Type", "application/json");

            console.log(cpr)

            const requestOptions: RequestInit = {
                method: NetworkAdapter.POST,
                headers: myHeaders,
                body: JSON.stringify(cpr)
            }

            const res: Response = await fetch('http://localhost:8080/projects', requestOptions);
            if(!res.ok) {
                console.error(res.statusText);
                return false;
            }

            const data: IProject = await res.json();
            setProjects((prevProjects) => prevProjects ? [...prevProjects, data] : [data]);
            return true;
        } catch (e) {
            console.error(e);
            return false;
        }

    }

    const appContextValue: IProjectContext = {
        projects,
        setProjects,
        createProject,
    };

    return (
        <ProjectContext.Provider value={appContextValue}>
            {children}
        </ProjectContext.Provider>
    );
};

export default ProjectContext;

export const useProjectContext = () => {
    const context = useContext(ProjectContext);
    if (!context) {
        throw new Error('useItemContext must be used within a ItemContextProvider');
    }
    return context;
};