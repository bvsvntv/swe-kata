type StartSessionType = {
    sessionID: string;
    userID: string;
    token: string;
    expiresAt: Date;
};

type UpdateSessionType = {
    sessionID: string;
    token: string;
    expiresAt: Date;
};

export { StartSessionType, UpdateSessionType };
