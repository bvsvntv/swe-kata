import { Request, Response } from 'express';
import * as authService from '@/services/auth.service';
import { sendResponse } from '@/utils/appResponse.util';

async function register(req: Request, res: Response) {
    const { email, password } = req.body;

    const response = await authService.register(email, password);
    return sendResponse(res, 201, {
        success: true,
        message: 'User has been registered successfully.',
        data: response,
    });
}

async function login(req: Request, res: Response) {
    await authService.login();
}

async function refreshTokens(req: Request, res: Response) {
    await authService.refreshTokens();
}

async function logout(req: Request, res: Response) {
    await authService.logout();
}

export { register, login, refreshTokens, logout };
