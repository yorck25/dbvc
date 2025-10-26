import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import TableBody from "@mui/material/TableBody";
import type {IProject} from "../../../models/projects.models.ts";
import {EllipsisVerticalIcon} from "../../icons";
import TableContainer from "@mui/material/TableContainer";
import styles from "./style.module.scss";
import {useConnectionTypesContext} from "../../../contexts/connection-types.context.tsx";
import {useUserContext} from "../../../contexts/users.context.tsx";

export const ProjectsTable = ({projects}: { projects: IProject[] }) => {
    const {users} = useUserContext();
    const {connectionTypes} = useConnectionTypesContext();

    const dummyMembers = users;

    const renderMembers = () => {
        return (
            <div className={styles.members_container}>
                {dummyMembers.slice(0, 5).map((member: string) => (
                    renderMember(member[0])
                ))}

                {renderMember("+" + (dummyMembers.length - 5))}
            </div>
        )
    }

    const renderMember = (key: string) => {
        const colors = ["#eb4034", "#9fd631", "#bf26c9", "#539667"];
        const random = Math.floor(Math.random() * colors.length);

        return (
            <div className={styles.profile_image_container} style={{background: colors[random]}}>
                {key.toUpperCase()}
            </div>
        );
    }

    const getConnectionType = (connectionTypeId: number): string => {
        if (!connectionTypes) {
            return "null";
        }

        const connectionType = connectionTypes.find(ct => ct.id === connectionTypeId);
        if (!connectionType) {
            return "null";
        }

        return connectionType.typeName;
    }

    return (
        <div className={styles.projects_table}>
            <TableContainer component={Paper}>
                <Table sx={{minWidth: 650}} aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell>Name</TableCell>
                            <TableCell align={"left"}>Description</TableCell>
                            <TableCell align={"left"}>Connection Type</TableCell>
                            <TableCell align={"left"}>Members</TableCell>
                            <TableCell align={"left"}></TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {projects?.map((row: IProject) => (
                            <TableRow
                                key={row.id}
                                sx={{'&:last-child td, &:last-child th': {border: 0}}}
                            >
                                <TableCell align={"left"} component="th" scope="row">
                                    {row.name}
                                </TableCell>
                                <TableCell align={"left"}>{row.description}</TableCell>
                                <TableCell align={"left"}>{getConnectionType(row.connectionType)}</TableCell>
                                <TableCell align={"left"}>{renderMembers()}</TableCell>
                                <TableCell align={"left"}><span
                                    className={styles.option_button}>{EllipsisVerticalIcon()}</span></TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </div>
    )
}