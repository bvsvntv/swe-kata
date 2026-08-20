import {
    createProfile,
    deleteProfile,
    getProfile,
    updateProfile,
} from '@/repositories/user.repository';
import { CreateProfileType, UpdateProfileType } from '@/types/user.types';
import { AppError } from '@shared/src/types';

async function get(userID: string) {
    const profile = await getProfile(userID);
    if (!profile) {
        throw new AppError('Profile not found.', 404);
    }

    return await getProfile(userID);
}

async function create(userID: string, args: CreateProfileType) {
    const profile = await getProfile(userID);
    if (profile) {
        throw new AppError('Profile already exists.', 400);
    }

    return await createProfile(userID, args);
}

async function update(userID: string, args: UpdateProfileType) {
    const profile = await getProfile(userID);
    if (!profile) {
        throw new AppError('Profile not found.', 404);
    }

    return await updateProfile(userID, args);
}

async function remove(userID: string) {
    const profile = await getProfile(userID);
    if (!profile) {
        throw new AppError('Profile not found.', 404);
    }

    return await deleteProfile(userID);
}

export { get, create, update, remove };
