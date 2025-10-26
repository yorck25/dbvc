import {createContext, useContext, useEffect, useState, type Dispatch, type FC, type ReactNode} from 'react';
import type {StateUpdater} from "preact/hooks";

interface IUserContext {
    users: string[];
    setUsers: Dispatch<StateUpdater<string[]>>;
}

const UserContext = createContext<IUserContext | undefined>(undefined);

export const UserContextProvider: FC<{ children: ReactNode }> = ({children}) => {
    const [users, setUsers] = useState<string[]>(["Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Heidi"]);

    useEffect(() => {
    }, []);

    const appContextValue: IUserContext = {
        users,
        setUsers,
    };

    return (
        <UserContext.Provider value={appContextValue}>
            {children}
        </UserContext.Provider>
    );
};

export default UserContext;

export const useUserContext = () => {
    const context = useContext(UserContext);
    if (!context) {
        throw new Error('useItemContext must be used within a ItemContextProvider');
    }
    return context;
};