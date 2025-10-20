export interface IProjectRequest {
    id: number;
    ownerId: number;
    name: string;
    createdAt: Date;
    updatedAt: Date;
    active: boolean;
    connectionType: number;
}

export interface IProjects {
    id: number;
    ownerId: number;
    name: string;
    createdAt: Date;
    updatedAt: Date;
    active: boolean;
    connectionType: number;
}