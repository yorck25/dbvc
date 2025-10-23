import {Link} from "react-router-dom";
import style from "./style.module.scss";

export const Sidebar = () => {
    return (
        <nav className={style.sidebar}>
            <ul>
                <li>
                    <Link to="/projects">Projects</Link>
                </li>
            </ul>
        </nav>
    )
}