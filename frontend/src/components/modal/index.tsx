import styles from './style.module.scss';
import {Button} from "../button";
import { useEffect } from 'preact/hooks';

type FooterType = 'single' | 'double';

interface Props {
    title: string;
    content: React.ReactNode;
    isOpen: boolean;
    onClose: () => void;
    footerType?: FooterType;
    primaryButtonText?: string;
    onPrimaryClick?: () => void;
    cancelButtonText?: string;
    submitButtonText?: string;
    onCancel?: () => void;
    onSubmit?: () => void;
}

export const Modal = (
    {
        title,
        content,
        isOpen,
        onClose,
        footerType = 'single',
        primaryButtonText = 'Close',
        onPrimaryClick,
        cancelButtonText = 'Cancel',
        submitButtonText = 'Submit',
        onCancel,
        onSubmit,
    }: Props) => {

    const closePopup = (event?: React.MouseEvent<HTMLElement>) => {
        if (event) {
            event.stopPropagation();
        }
        onClose();
    };

    const handlePrimaryClick = () => {
        if (onPrimaryClick) {
            onPrimaryClick();
        } else {
            closePopup();
        }
    };

    const handleCancel = () => {
        if (onCancel) {
            onCancel();
        } else {
            closePopup();
        }
    };

    const handleSubmit = () => {
        if (onSubmit) {
            onSubmit();
        }
    };

    useEffect(() => {
        const handleEscape = (event: KeyboardEvent) => {
            if (event.key === 'Escape' && isOpen) {
                closePopup();
            }
        };

        if (isOpen) {
            document.addEventListener('keydown', handleEscape);
        }

        return () => {
            document.removeEventListener('keydown', handleEscape);
        };
    }, [isOpen]);

    return (
        <>
            {isOpen && (
                <div onClick={(e) => closePopup(e)} className={`${styles.overlay} ${isOpen ? styles.show : ''}`}>
                    <div onClick={(e) => e.stopPropagation()} className={`${styles.popup} ${isOpen ? styles.show : ''}`}
                         id="popup">
                        <h2>{title}</h2>
                        <div className={styles.content}>{content}</div>

                        <div className={styles.footer}>
                            {footerType === 'single' ? (
                                <Button
                                    id="primary-button"
                                    text={primaryButtonText}
                                    ariaLabel={primaryButtonText}
                                    callback={handlePrimaryClick}
                                />
                            ) : (
                                <div className={styles.buttonGroup}>
                                    <Button
                                        id="cancel-button"
                                        text={cancelButtonText}
                                        ariaLabel={cancelButtonText}
                                        callback={handleCancel}
                                    />
                                    <Button
                                        id="submit-button"
                                        text={submitButtonText}
                                        ariaLabel={submitButtonText}
                                        callback={handleSubmit}
                                    />
                                </div>
                            )}
                        </div>
                    </div>
                </div>
            )}
        </>
    );
};