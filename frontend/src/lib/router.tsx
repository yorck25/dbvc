import {BrowserRouter, Route, Routes} from "react-router-dom";
import {Page404} from "../pages/404";
import {ProjectsPage} from "../pages/projects";

export const Router = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route
                    path="*"
                    element={<Page404/>}
                />

                <Route path="/" element={<ProjectsPage/>}/>
            </Routes>
        </BrowserRouter>
    )
}