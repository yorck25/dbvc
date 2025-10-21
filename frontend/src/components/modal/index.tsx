import styles from './modal.module.scss';
import {Button} from "../button";
import React from "react";

interface Props {
    title: string;
    message: React.ReactNode;
    isOpen: boolean;
    onClose: () => void;
}

export const Modal = (
    {
        title,
        message,
        isOpen,
        onClose,
    }: Props) => {

    const closePopup = (event?: React.MouseEvent<HTMLElement>) => {
        if (event) {
            event?.stopPropagation();
        }
        onClose();
    };

    return (
        <>
            {isOpen && (
                <div onClick={(e) => closePopup(e)} className={`${styles.overlay} ${isOpen ? styles.show : ''}`}>
                    <div onClick={(e) => e.stopPropagation()} className={`${styles.popup} ${isOpen ? styles.show : ''}`}
                         id="popup">
                        <h2>{title}</h2>
                        <div className={styles.content}>{message}</div>
                        <Button
                            id="close-popup"
                            text="Close Popup"
                            ariaLabel="Close Popup"
                            callback={closePopup}
                        />
                    </div>
                </div>
            )}
        </>
    );
};
