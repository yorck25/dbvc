import { Routes, Route } from "react-router-dom";
import { Page404 } from "../pages/404";
import { ProjectsPage } from "../pages/projects";

export const Router = () => {
    return (
        <Routes>
            <Route path={"/"} element={<></>} />
            <Route path="/projects" element={<ProjectsPage />} />
            <Route path="*" element={<Page404 />} />
        </Routes>
    );
};