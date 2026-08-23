import { Request, Response } from 'express';
import * as authService from '@/services/auth.service';
import { sendResponse } from '@shared/src/utils/appResponse.util';
import { AppError } from '@shared/src/types';
import { clearCookies, setCookies } from '@/utils/auth.utils';
import { REFRESH_TOKEN_COOKIE } from '@/constants';

async function registerController(req: Request, res: Response) {
    const { email, password } = req.body;
    const userAgent = req.headers['user-agent'] || 'unknown';
    const ipAddress = req.ip || 'unknown';

    const { accessToken, refreshToken } = await authService.register({
        email,
        password,
        userAgent,
        ipAddress,
    });

    setCookies(res, refreshToken);

    return sendResponse(res, 201, {
        success: true,
        message: 'User has been registered successfully.',
        data: { accessToken },
    });
}

async function loginController(req: Request, res: Response) {
    const { email, password } = req.body;
    const userAgent = req.headers['user-agent'] || 'unknown';
    const ipAddress = req.ip || 'unknown';

    const { accessToken, refreshToken } = await authService.login({
        email,
        password,
        userAgent,
        ipAddress,
    });

    setCookies(res, refreshToken);

    return sendResponse(res, 200, {
        success: true,
        message: 'User has been logged in successfully.',
        data: { accessToken },
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
    const refreshToken = req.cookies?.[REFRESH_TOKEN_COOKIE];

    if (!refreshToken) {
        throw new AppError('Missing refresh token.', 401);
    }

    const { accessToken, refreshToken: newRefreshToken } =
        await authService.refreshTokens(refreshToken);

    setCookies(res, newRefreshToken);

    return sendResponse(res, 200, {
        success: true,
        message: 'Session has been refreshed successfully.',
        data: { accessToken },
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

    clearCookies(res);

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
