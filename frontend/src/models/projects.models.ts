export interface IProjectRequest {
    id: number;
    ownerId: number;
    name: string;
    createdAt: Date;
    updatedAt: Date;
    active: boolean;
    connectionType: number;
}

export interface Project {
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