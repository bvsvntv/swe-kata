import 'express';

declare global {
    namespace Express {
        interface Request {
            user?: {
                userID: string;
            };
        }
    }
}

export {};
