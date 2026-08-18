import { Request, Response, NextFunction } from 'express';
import { AppError } from '../types';

export function errorHandler(
    error: AppError,
    req: Request,
    res: Response,
    next: NextFunction,
) {
    const statusCode = error.statusCode || 500;
    const message = error.message || 'Internal server error.';

    res.status(statusCode).json({ error: message });

    next();
}
