import {createContext, useContext, useEffect, useState, type Dispatch, type FC, type ReactNode} from 'react';

interface IAppContext {
    user: string;
    setuser: Dispatch<string>;
}

const AppContext = createContext<IAppContext | undefined>(undefined);

export const AppContextProvider: FC<{ children: ReactNode }> = ({children}) => {
    const [user, setuser] = useState<string>("");


    useEffect(() => {
        console.log("get user from local storeage")
    }, []);


    const appContextValue: IAppContext = {
        user,
        setuser,
    };

    return (
        <AppContext.Provider value={appContextValue}>
            {children}
        </AppContext.Provider>
    );
};

export default AppContext;

export const useAppContext = () => {
    const context = useContext(AppContext);
    if (!context) {
        throw new Error('useItemContext must be used within a ItemContextProvider');
    }
    return context;
};