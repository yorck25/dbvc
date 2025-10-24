import "./app.css";
import {AppContextProvider} from "./contexts/app.context";
import {ProjectContextProvider} from "./contexts/projects.context";
import {Router} from "./lib/router";
import {BrowserRouter} from "react-router-dom";

export function App() {
    return (
        <AppContextProvider>
            <BrowserRouter>
                <div className={"main"}>
                    <ProjectContextProvider>
                        <div className={"main_wrapper"}>
                            <Router/>
                        </div>
                    </ProjectContextProvider>
                </div>
            </BrowserRouter>
        </AppContextProvider>
    );
}
