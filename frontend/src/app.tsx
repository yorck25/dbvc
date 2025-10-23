import "./app.css";
import {Sidebar} from "./components/sidebar";
import { ProjectContextProvider } from "./contexts/projects.context";
import {Router} from "./lib/router";
import {BrowserRouter} from "react-router-dom";

export function App() {
    return (
        <BrowserRouter>
            <div className={"main"}>
                <Sidebar/>

                <ProjectContextProvider>
                    <div className={"main_wrapper"}>
                        <Router/>
                    </div>
                </ProjectContextProvider>
            </div>
        </BrowserRouter>
    );
}
