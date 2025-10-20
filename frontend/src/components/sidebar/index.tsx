import style from './style.module.scss';

export const Sidebar = () => {
    return (
        <nav className={style.sidebar}>
            <ul>
                <li><a>Database</a></li>
            </ul>
        </nav>
    )
}