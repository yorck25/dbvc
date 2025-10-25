import {createContext, useContext, useEffect, useState, type Dispatch, type FC, type ReactNode} from 'react';
import type {IConnectionType} from "../models/connection.models.ts";
import {NetworkAdapter, setAuthHeader} from "../lib/networkAdapter.tsx";

interface IConnectionTypesContext {
    connectionTypes: IConnectionType[] | undefined;
    setConnectionTypes: Dispatch<IConnectionType[] | undefined>;
}

const ConnectionTypesContext = createContext<IConnectionTypesContext | undefined>(undefined);

export const ConnectionTypesContextProvider: FC<{ children: ReactNode }> = ({children}) => {
    const [connectionTypes, setConnectionTypes] = useState<IConnectionType[]>();

    useEffect(() => {
        fetchConnectionTypes();
    }, []);

    const fetchConnectionTypes = () => {
        const header = setAuthHeader();
        header.append("Content-Type", "application/json");

        const requestOptions: RequestInit = {
            method: NetworkAdapter.GET,
            headers: header,
        }

        fetch('http://localhost:8080/connection-types', requestOptions)
            .then(res => res.json())
            .then((data: IConnectionType[]) => {
                setConnectionTypes(data);
            });
    }

    const contextValue: IConnectionTypesContext = {
        connectionTypes,
        setConnectionTypes,
    };

    return (
        <ConnectionTypesContext.Provider value={contextValue}>
            {children}
        </ConnectionTypesContext.Provider>
    );
};

export default ConnectionTypesContext;

export const useConnectionTypesContext = () => {
    const context = useContext(ConnectionTypesContext);
    if (!context) {
        throw new Error('useItemContext must be used within a ItemContextProvider');
    }
    return context;
};