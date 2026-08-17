import jwt from 'jsonwebtoken';
import { NextFunction, Request, Response } from 'express';
import { verifyAccessToken } from '@/utils/jwt.util';
import { JWTPayloadType } from '@/types/auth.types';
import { AppError } from '@shared/src/types';

function authMiddleware(req: Request, res: Response, next: NextFunction) {
    try {
        const authHeader = req.headers.authorization;
        if (!authHeader) {
            return next(new AppError('Authentication required.', 401));
        }
        if (!authHeader.startsWith('Bearer ')) {
            return next(
                new AppError('Invalid authentication header format.', 401),
            );
        }

        const accessToken = authHeader.split(' ')[1];
        if (!accessToken) {
            return next(new AppError('Access token missing.', 401));
        }

        const payload = verifyAccessToken(accessToken) as JWTPayloadType;
        req.user = {
            userID: payload?.sub,
            sessionID: payload?.sessionID,
        };

        next();
    } catch (error) {
        if (error instanceof jwt.TokenExpiredError) {
            return next(new AppError('Access token expired.', 401));
        }

        if (error instanceof jwt.JsonWebTokenError) {
            return next(new AppError('Invalid access token.', 401));
        }

        return next(error);
    }
}

export { authMiddleware };
