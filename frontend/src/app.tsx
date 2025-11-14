import "./app.css";
import {AppContextProvider} from "./contexts/app.context";
import {ProjectContextProvider} from "./contexts/projects.context";
import {Router} from "./lib/router";
import {BrowserRouter} from "react-router-dom";
import {ConfigContextProvider} from "./contexts/connection-types.context.tsx";
import {UserContextProvider} from "./contexts/users.context.tsx";

export function App() {
    return (
        <AppContextProvider>
            <BrowserRouter>
                <div className={"main"}>
                    <ConfigContextProvider>
                        <UserContextProvider>
                            <ProjectContextProvider>
                                <div className={"main_wrapper"}>
                                    <Router/>
                                </div>
                            </ProjectContextProvider>
                        </UserContextProvider>
                    </ConfigContextProvider>
                </div>
            </BrowserRouter>
        </AppContextProvider>
    );
}
