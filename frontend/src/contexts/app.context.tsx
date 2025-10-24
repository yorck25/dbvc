import {
    createContext,
    useContext,
    useEffect,
    useState,
    type Dispatch,
    type FC,
    type ReactNode,
} from "react";
import type {IRegisterRequest, IUser} from "../models/user.models.ts";

import {HTTPMethods} from "../lib/HTTPMethods.tsx";

interface IAppContext {
    user: IUser | undefined;
    setUser: Dispatch<IUser | undefined>;

    isLoggedIn: boolean;
    setIsLoggedIn: Dispatch<boolean>;

    token: string | undefined;
    setToken: Dispatch<string | undefined>;

    registerRequest: (rr: IRegisterRequest) => Promise<boolean>;
}

const AppContext = createContext<IAppContext | undefined>(undefined);

export const AppContextProvider: FC<{ children: ReactNode }> = ({
                                                                    children
                                                                }) => {
    const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
    const [user, setUser] = useState<IUser | undefined>();
    const [token, setToken] = useState<string | undefined>();

    useEffect(() => {
        console.log("get user from local storeage");
    }, []);

    const registerRequest = async (rr: IRegisterRequest): Promise<boolean> => {
        try {
            const myHeaders = new Headers();
            myHeaders.append("Content-Type", "application/json");

            const requestOptions: RequestInit = {
                method: HTTPMethods.POST,
                headers: myHeaders,
                redirect: 'follow',
                body: JSON.stringify(rr),
            };

            const response = await fetch("http://localhost:8080/auth/register", requestOptions);

            if (!response.ok) {
                console.error(new Error("Network response was not ok"));
                return false;
            }

            const result: { token: string, user: IUser } = await response.json();

            setUser(result.user);
            setToken(result.token);
            setIsLoggedIn(true);
            saveTokenInStorage(result.token);

            return true;

        } catch (error) {
            console.error("Fetch error:", error);
            return false;
        }
    };


    const saveTokenInStorage = (token: string) => {
        localStorage.setItem("authToken", token);
    }

    const appContextValue: IAppContext = {
        user,
        setUser,
        isLoggedIn,
        setIsLoggedIn,
        token,
        setToken,
        registerRequest,
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
        throw new Error("useItemContext must be used within a ItemContextProvider");
    }
    return context;
};
