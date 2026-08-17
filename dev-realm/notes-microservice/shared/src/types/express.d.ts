import 'express';

declare global {
    namespace Express {
        interface Request {
            user?: {
                userID: string;
                sessionID: string;
            };
        }
    }
}

export {};
