import styles from './style.module.scss';
import {Button} from "../button";
import {useEffect} from 'preact/hooks';

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
                        <div className={styles.header}>
                            <div className={styles.top_row}>
                                <h2>{title}</h2>
                                <button
                                    className={styles.closeButton}
                                    onClick={(e) => closePopup(e)}
                                    aria-label="Close modal"
                                >
                                    <svg
                                        xmlns="http://www.w3.org/2000/svg"
                                        width="20"
                                        height="20"
                                        viewBox="0 0 24 24"
                                        fill="none"
                                        stroke="currentColor"
                                        strokeWidth="2"
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                    >
                                        <line x1="18" y1="6" x2="6" y2="18"/>
                                        <line x1="6" y1="6" x2="18" y2="18"/>
                                    </svg>
                                </button>
                            </div>

                            <p>Das ist ein super hint... :)</p>
                        </div>

                        <div className={styles.content}>{content}</div>

                        <div className={styles.footer}>
                            {footerType === 'single' ? (
                                <div className={styles.button_wrapper}>
                                    <Button
                                        id="primary-button"
                                        text={primaryButtonText}
                                        ariaLabel={primaryButtonText}
                                        callback={handlePrimaryClick}
                                    />
                                </div>
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