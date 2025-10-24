export interface IRegisterRequest {
    firstName: string;
    lastName: string;
    email: string;
    username: string;
    password: string;
    termsAccepted: boolean;
}

export interface IUser {
    id: number;
    firstName: string;
    lastName: string;
    email: string;
    username: string;
    createdAt: string;
    updatedAt: string;
    active: boolean;
}