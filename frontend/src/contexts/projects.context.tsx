import {
    createContext,
    useContext,
    useEffect,
    useState,
    type Dispatch,
    type FC,
    type ReactNode,
} from "react";
import type { ICreateProjectRequest, IProjectWithUsers } from "../models/projects.models";
import { NetworkAdapter, setAuthHeader } from "../lib/networkAdapter.tsx";

interface IProjectContext {
    projects: IProjectWithUsers[] | undefined;
    setProjects: Dispatch<any>;
    getProjectById: (projectId: number) => IProjectWithUsers | undefined;
    ensureProjectLoaded: (projectId: number) => Promise<IProjectWithUsers | undefined>;
    createProject: (cpr: ICreateProjectRequest) => Promise<boolean>;
}

const ProjectContext = createContext<IProjectContext | undefined>(undefined);

export const ProjectContextProvider: FC<{ children: ReactNode }> = ({ children }) => {
    const [projects, setProjects] = useState<IProjectWithUsers[]>();

    useEffect(() => {
        fetchProjects();
    }, []);

    const fetchProjects = async () => {
        try {
            const myHeaders = setAuthHeader();
            myHeaders.append("Content-Type", "application/json");

            const requestOptions: RequestInit = {
                method: NetworkAdapter.GET,
                headers: myHeaders,
            };

            const res = await fetch("http://localhost:8080/projects", requestOptions);
            if (!res.ok) throw new Error(res.statusText);

            const data: IProjectWithUsers[] = await res.json();
            setProjects(data);
        } catch (e) {
            console.error("Failed to fetch projects:", e);
        }
    };

    const getProjectById = (projectId: number): IProjectWithUsers | undefined => {
        return projects?.find((p) => p.project.id === projectId);
    };

    const fetchProjectById = async (projectId: number): Promise<IProjectWithUsers | undefined> => {
        try {
            const myHeaders = setAuthHeader();
            myHeaders.append("Content-Type", "application/json");

            const requestOptions: RequestInit = {
                method: NetworkAdapter.GET,
                headers: myHeaders,
            };

            const res = await fetch(`http://localhost:8080/projects/${projectId}`, requestOptions);
            if (!res.ok) throw new Error(res.statusText);

            const data: IProjectWithUsers = await res.json();
            setProjects((prev) => (prev ? [...prev, data] : [data]));
            return data;
        } catch (e) {
            console.error("Failed to fetch project:", e);
            return undefined;
        }
    };

    const ensureProjectLoaded = async (
        projectId: number
    ): Promise<IProjectWithUsers | undefined> => {
        const existing = getProjectById(projectId);
        if (existing) return existing;
        return await fetchProjectById(projectId);
    };

    const createProject = async (cpr: ICreateProjectRequest): Promise<boolean> => {
        try {
            const myHeaders = setAuthHeader();
            myHeaders.append("Content-Type", "application/json");

            const requestOptions: RequestInit = {
                method: NetworkAdapter.POST,
                headers: myHeaders,
                body: JSON.stringify(cpr),
            };

            const res = await fetch("http://localhost:8080/projects", requestOptions);
            if (!res.ok) throw new Error(res.statusText);

            const data: IProjectWithUsers = await res.json();
            setProjects((prev) => (prev ? [...prev, data] : [data]));
            return true;
        } catch (e) {
            console.error("Failed to create project:", e);
            return false;
        }
    };

    const appContextValue: IProjectContext = {
        projects,
        setProjects,
        getProjectById,
        ensureProjectLoaded,
        createProject,
    };

    return (
        <ProjectContext.Provider value={appContextValue}>
            {children}
        </ProjectContext.Provider>
    );
};

export const useProjectContext = () => {
    const context = useContext(ProjectContext);
    if (!context) throw new Error("useProjectContext must be used within a ProjectContextProvider");
    return context;
};
