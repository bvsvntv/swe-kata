import { Request, Response } from 'express';
import * as authService from '@/services/auth.service';
import { sendResponse } from '@shared/src/utils/appResponse.util';
import { AppError } from '@shared/src/types';

async function registerController(req: Request, res: Response) {
    const { email, password } = req.body;

    const response = await authService.register(email, password);
    return sendResponse(res, 201, {
        success: true,
        message: 'User has been registered successfully.',
        data: response,
    });
}

async function loginController(req: Request, res: Response) {
    const { email, password } = req.body;

    const response = await authService.login(email, password);
    return sendResponse(res, 200, {
        success: true,
        message: 'User has been logged in successfully.',
        data: response,
    });
}

async function getProfileController(req: Request, res: Response) {
    const userID = req.user?.userID;
    if (!userID) {
        throw new AppError('Authentication required.', 401);
    }

    const user = await authService.getProfile(userID);
    if (!user) {
        throw new AppError('User not found.', 404);
    }

    return sendResponse(res, 200, {
        success: true,
        message: 'User has been fetched.',
        data: user,
    });
}

async function refreshTokensController(req: Request, res: Response) {
    const { refreshToken } = req.body;
    if (!refreshToken) {
        throw new AppError('Missing refresh token.', 400);
    }

    const response = await authService.refreshTokens(refreshToken);
    return sendResponse(res, 200, {
        success: true,
        message: 'Session has been refreshed successfully.',
        data: response,
    });
}

async function logoutController(req: Request, res: Response) {
    const userID = req.user?.userID;
    if (!userID) {
        throw new AppError('Authentication required.', 401);
    }

    const user = await authService.getProfile(userID);
    if (!user) {
        throw new AppError('User not found.', 404);
    }

    await authService.logout(userID);

    return sendResponse(res, 200, {
        success: true,
        message: 'User has been logged out successfully.',
    });
}

export {
    registerController,
    loginController,
    getProfileController,
    refreshTokensController,
    logoutController,
};
