import {Routes, Route, Outlet, useNavigate} from "react-router-dom";
import {Page404} from "../pages/404";
import {ProjectsPage} from "../pages/projects";
import {RegisterPage} from "../pages/register";
import {LoginPage} from "../pages/login";
import {Sidebar} from "../components/sidebar";
import {ProjectContextProvider} from "../contexts/projects.context";
import {useEffect} from "react";
import {useAppContext} from "../contexts/app.context.tsx";

export const Router = () => {
    const navigate = useNavigate();
    const {checkUserLogin, isLoggedIn} = useAppContext();

    useEffect(() => {
        if (!checkUserLogin()) {
            navigate("/login");
        }
    }, [isLoggedIn])

    return (
        <Routes>
            <Route element={<MainLayout/>}>
                <Route path={"/"} element={<></>} />
                <Route path="/projects" element={<ProjectsPage/>}/>
            </Route>

            <Route element={<AuthLayout/>}>
                <Route path="/register" element={<RegisterPage/>}/>
                <Route path="/login" element={<LoginPage/>}/>
            </Route>

            <Route path="*" element={<Page404/>}/>
        </Routes>
    );
};

export const MainLayout = () => (
    <div className="main">
        <Sidebar/>
        <ProjectContextProvider>
            <div className="main_wrapper">
                <Outlet/>
            </div>
        </ProjectContextProvider>
    </div>
);

export const AuthLayout = () => (
    <div className="auth_wrapper">
        <Outlet/>
    </div>
);
