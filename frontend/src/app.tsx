import "./app.css";
import {AppContextProvider} from "./contexts/app.context";
import {ProjectContextProvider} from "./contexts/projects.context";
import {Router} from "./lib/router";
import {BrowserRouter} from "react-router-dom";
import {ConnectionTypesContextProvider} from "./contexts/connection-types.context.tsx";

export function App() {
    return (
        <AppContextProvider>
            <BrowserRouter>
                <div className={"main"}>
                    <ConnectionTypesContextProvider>
                        <ProjectContextProvider>
                            <div className={"main_wrapper"}>
                                <Router/>
                            </div>
                        </ProjectContextProvider>
                    </ConnectionTypesContextProvider>
                </div>
            </BrowserRouter>
        </AppContextProvider>
    );
}
