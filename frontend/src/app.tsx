import "./app.css";
import {AppContextProvider} from "./contexts/app.context";
import {ProjectContextProvider} from "./contexts/projects.context";
import {Router} from "./lib/router";
import {BrowserRouter} from "react-router-dom";
import {ConnectionTypesContextProvider} from "./contexts/connection-types.context.tsx";
import {UserContextProvider} from "./contexts/users.context.tsx";

export function App() {
    return (
        <AppContextProvider>
            <BrowserRouter>
                <div className={"main"}>
                    <ConnectionTypesContextProvider>
                        <UserContextProvider>
                            <ProjectContextProvider>
                                <div className={"main_wrapper"}>
                                    <Router/>
                                </div>
                            </ProjectContextProvider>
                        </UserContextProvider>
                    </ConnectionTypesContextProvider>
                </div>
            </BrowserRouter>
        </AppContextProvider>
    );
}
