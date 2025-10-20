import {createContext, useContext, useEffect, useState, type Dispatch, type FC, type ReactNode} from 'react';

interface IProjectContext {
    projects: any;
    setProjects: Dispatch<any>;
}

const ProjectContext = createContext<IProjectContext | undefined>(undefined);

export const ProjectContextProvider: FC<{ children: ReactNode }> = ({children}) => {
    const [projects, setProjects] = useState<any>();

    useEffect(() => {
        console.log("");
    }, []);

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