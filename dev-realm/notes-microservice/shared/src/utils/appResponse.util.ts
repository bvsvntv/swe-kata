import { Response } from 'express';
import { ApiResponse } from '../types';

export function sendResponse<T>(
    res: Response,
    statusCode: number,
    payload: ApiResponse<T>,
) {
    return res.status(statusCode).json(payload);
}
