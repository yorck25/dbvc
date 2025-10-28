import {createContext, useContext, useState, type Dispatch, type FC, type ReactNode} from 'react';
import type {StateUpdater} from "preact/hooks";
import {NetworkAdapter, setAuthHeader} from "../lib/networkAdapter.tsx";
import type {IMemberRequest} from "../models/user.models.ts";

interface IUserContext {
    users: string[];
    setUsers: Dispatch<StateUpdater<string[]>>;

    searchAvailableMembers: (searchValue: string) => Promise<IMemberRequest[]>;
}

const UserContext = createContext<IUserContext | undefined>(undefined);

export const UserContextProvider: FC<{ children: ReactNode }> = ({children}) => {
    const [users, setUsers] = useState<string[]>([]);

    const searchAvailableMembers = async (searchValue: string): Promise<IMemberRequest[]> => {
        const headers = setAuthHeader();
        headers.append('Content-Type', 'application/json');

        const requestOptions: RequestInit = {
            method: NetworkAdapter.GET,
            headers: headers,
        }

        const res: Response = await fetch(`http://localhost:8080/users/search?query=${searchValue}`, requestOptions)
        const result: IMemberRequest[] = await res.json();
        return result;
    }

    const appContextValue: IUserContext = {
        users,
        setUsers,
        searchAvailableMembers
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