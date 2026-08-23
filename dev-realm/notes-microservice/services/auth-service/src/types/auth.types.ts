type RegisterUserType = {
    email: string;
    password: string;
    userAgent: string;
    ipAddress: string;
};

type LoginUserType = RegisterUserType;

type StartSessionType = {
    sessionID: string;
    userID: string;
    token: string;
    userAgent?: string;
    ipAddress?: string;
    expiresAt: Date;
};

type UpdateSessionType = {
    sessionID: string;
    token: string;
    expiresAt: Date;
};

export { RegisterUserType, LoginUserType, StartSessionType, UpdateSessionType };
