import {Input} from "../../input";
import styles from "./style.module.scss";
import React, {useState} from "react";
import {Button, ButtonType} from "../../button";
import {useUserContext} from "../../../contexts/users.context.tsx";

export interface IAddMembersFormProps {
    members: number[];
    setMembers: React.Dispatch<React.SetStateAction<number[]>>;
}

export const AddMembersForm = ({members, setMembers}: IAddMembersFormProps) => {
    const {users} = useUserContext();
    const [searchUser, setSearchUser] = useState<string>("");

    const addMember = (id: number) => {
        setMembers([
            ...members,
            id
        ]);
    }

    return (
        <div className={styles.add_members_form}>
            <div className={styles.search_header}>

                <div className={styles.input_wrapper}>
                    <Input
                        id={"searchUser"}
                        value={searchUser}
                        handleInput={(e: Event) => setSearchUser((e.target as HTMLInputElement).value)}
                        placeholder={"Search User..."}
                    />
                </div>

                <div className={styles.button_wrapper}>
                    <Button text={"Search User"} callback={() => console.log("search user")}></Button>
                </div>
            </div>

            <ul className={styles.user_list}>
                {users.map((user: string, index: number) => (
                    <li className={styles.user_list_items} onClick={() => addMember(index)} key={index}>
                        <span>{user}</span>
                        <div className={styles.button_wrapper}>
                            <Button type={ButtonType.Outline} text={"Add"} callback={() => addMember(index)}/>
                        </div>
                    </li>
                ))}
            </ul>
        </div>
    )
}