import {
    createProfile,
    deleteProfile,
    getProfile,
    updateProfile,
} from '@/repositories/user.repository';
import { CreateProfileType, UpdateProfileType } from '@/types/user.types';
import { AppError } from '@shared/src/types';
import { logger } from '@/lib/logger';

async function get(userID: string) {
    const profile = await getProfile(userID);
    if (!profile) {
        logger.warn('Domain: profile fetch failed - profile not found', {
            event: 'profile_fetch_failed',
            userId: userID,
        });
        throw new AppError('Profile not found.', 404);
    }

    logger.info('Domain: profile fetched', {
        event: 'profile_fetched',
        userId: userID,
    });

    return await getProfile(userID);
}

async function create(userID: string, args: CreateProfileType) {
    const profile = await getProfile(userID);
    if (profile) {
        logger.warn('Domain: profile creation failed - already exists', {
            event: 'profile_create_failed',
            reason: 'already_exists',
            userId: userID,
        });
        throw new AppError('Profile already exists.', 400);
    }

    const created = await createProfile(userID, args);

    logger.info('Domain: profile created', {
        event: 'profile_created',
        userId: created.userID,
        profileId: created.id,
    });

    return created;
}

async function update(userID: string, args: UpdateProfileType) {
    const profile = await getProfile(userID);
    if (!profile) {
        logger.warn('Domain: profile update failed - profile not found', {
            event: 'profile_update_failed',
            reason: 'not_found',
            userId: userID,
        });
        throw new AppError('Profile not found.', 404);
    }

    const updated = await updateProfile(userID, args);

    logger.info('Domain: profile updated', {
        event: 'profile_updated',
        userId: userID,
    });

    return updated;
}

async function remove(userID: string) {
    const profile = await getProfile(userID);
    if (!profile) {
        logger.warn('Domain: profile deletion failed - profile not found', {
            event: 'profile_delete_failed',
            reason: 'not_found',
            userId: userID,
        });
        throw new AppError('Profile not found.', 404);
    }

    await deleteProfile(userID);

    logger.info('Domain: profile deleted', {
        event: 'profile_deleted',
        userId: userID,
    });
}

export { get, create, update, remove };
