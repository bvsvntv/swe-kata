import { AppError } from '@shared/src/types';
import { Request, Response } from 'express';
import * as userProfileService from '../services/user.service';
import { CreateProfileType } from '@/types/user.types';
import { sendResponse } from '@shared/src/utils/appResponse.util';

async function createProfileController(req: Request, res: Response) {
    const userID = req.user?.userID;
    const { firstName, lastName, bio, avatarURL }: CreateProfileType = req.body;

    if (!userID) {
        throw new AppError('Authentication required.', 401);
    }

    const profile = await userProfileService.create(userID, {
        firstName,
        lastName,
        bio,
        avatarURL,
    });

    return sendResponse(res, 201, {
        success: true,
        message: 'Profile has been created.',
        data: profile,
    });
}

async function getProfileController(req: Request, res: Response) {
    const userID = req.user?.userID;

    if (!userID) {
        throw new AppError('Authentication required.', 401);
    }

    const profile = await userProfileService.get(userID);

    return sendResponse(res, 200, {
        success: true,
        message: 'Profile has been fetched.',
        data: profile,
    });
}

async function updateProfileController(req: Request, res: Response) {
    const userID = req.user?.userID;
    const { firstName, lastName, bio, avatarURL }: CreateProfileType = req.body;

    if (!userID) {
        throw new AppError('Authentication required.', 401);
    }

    const profile = await userProfileService.update(userID, {
        firstName,
        lastName,
        bio,
        avatarURL,
    });

    return sendResponse(res, 200, {
        success: true,
        message: 'Profile has been updated.',
        data: profile,
    });
}

async function deleteProfileController(req: Request, res: Response) {
    const userID = req.user?.userID;

    if (!userID) {
        throw new AppError('Authentication required.', 401);
    }

    await userProfileService.remove(userID);

    return sendResponse(res, 200, {
        success: true,
        message: 'Profile has been deleted.',
    });
}

export {
    createProfileController,
    getProfileController,
    updateProfileController,
    deleteProfileController,
};
