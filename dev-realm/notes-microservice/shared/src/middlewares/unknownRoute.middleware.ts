import { Request, Response, NextFunction } from 'express';
import { AppError } from '../types';

export function unknownRouteHandler(
    req: Request,
    res: Response,
    next: NextFunction,
) {
    next(new AppError(`Can't find requested resource in the server!`, 404));
}
