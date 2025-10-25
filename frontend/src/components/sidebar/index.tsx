import { Link, useNavigate } from "react-router-dom";
import style from "./style.module.scss";
import { useAppContext } from "../../contexts/app.context";
import {LogoutIcon} from "../icons";

export const Sidebar = () => {
  const { logout } = useAppContext();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate("/");
  };

  return (
    <div className={style.sidebar}>
      <div className={style.logo_container}>
        <div className={style.search_bar}>
          <input type="text"></input>
        </div>

        <div className={style.divider} />
      </div>

      <div className={style.nav_list}>
        <nav>
          <ul>
            <li>
              <Link to="/projects">projects</Link>
            </li>
          </ul>
        </nav>

        <nav>
          <ul>
            <li onClick={() => handleLogout()}>
              <LogoutIcon />
              <Link to="/">Logout</Link>
            </li>
          </ul>
        </nav>
      </div>
    </div>
  );
};
