import React from "react";
import styles from './style.module.scss';

export enum ButtonType {
    Default = 'default',
    Outline = 'outline',
    Link = 'link',
    Text = 'text',
}

interface Props {
    id?: string;
    text?: string;
    callback?: (event?: React.MouseEvent<HTMLButtonElement>) => void;
    ariaLabel?: string;
    type?: ButtonType;
    icon?: React.ReactNode;
    large?: boolean;
    dark?: boolean;
    disabled?: boolean;
}

export const Button = (
    {
        id,
        text,
        callback = () => console.log(""),
        ariaLabel,
        type = ButtonType.Default,
        icon,
        large,
        dark = false,
        disabled = false,
    }: Props) => {

    const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        callback?.(event);
    };

    return (
        <button
            id={id}
            disabled={disabled}
            aria-label={ariaLabel}
            onClick={handleClick}
            className={`${styles[type]} ${large ? styles.large : ''} ${dark ? styles.dark : ''}`}
        >
            {icon && icon}
            <span>{text}</span>
        </button>
    );
};
