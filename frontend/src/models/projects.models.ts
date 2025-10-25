export interface ICreateProjectRequest {
    name: string;
    description: string;
    visibility: string;
    connectionType: number;
}

export interface IProject {
    id: number;
    ownerId: number;
    name: string;
    description: string;
    createdAt: string;
    updatedAt: string;
    active: boolean;
    visibility: string;
    connectionType: number;
}