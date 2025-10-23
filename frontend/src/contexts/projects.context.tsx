import {createContext, useContext, useEffect, useState, type Dispatch, type FC, type ReactNode} from 'react';
import type { Project } from '../models/projects.models';

interface IProjectContext {
    projects: any;
    setProjects: Dispatch<any>;
}

const ProjectContext = createContext<IProjectContext | undefined>(undefined);

export const ProjectContextProvider: FC<{ children: ReactNode }> = ({children}) => {
    const [projects, setProjects] = useState<Project[]>();

    useEffect(() => {
        console.log("projects context");
        fetchProjects();
    }, []);

    const fetchProjects = () => {
        fetch('http://localhost:8080/projects')
            .then(res => res.json())
            .then((data: Project[]) => {
                setProjects(data);
            });
    }

    const appContextValue: IProjectContext = {
        projects,
        setProjects,
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