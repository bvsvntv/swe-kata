import prisma from '@/lib/prisma';
import { CreateProfileType, UpdateProfileType } from '@/types/user.types';
import { UserProfile } from 'generated/prisma/client';

async function createProfile(
    userID: string,
    args: CreateProfileType,
): Promise<UserProfile> {
    const { firstName, lastName, bio, avatarURL } = args;

    return await prisma.userProfile.create({
        data: {
            userID,
            firstName,
            lastName,
            bio,
            avatarURL,
        },
    });
}

async function getProfile(userID: string): Promise<UserProfile | null> {
    return await prisma.userProfile.findUnique({
        where: { userID },
        select: {
            id: true,
            userID: true,
            firstName: true,
            lastName: true,
            bio: true,
            avatarURL: true,
            createdAt: true,
            updatedAt: true,
        },
    });
}

async function updateProfile(
    userID: string,
    args: UpdateProfileType,
): Promise<UserProfile> {
    const { firstName, lastName, bio, avatarURL } = args;

    return await prisma.userProfile.update({
        where: { userID },
        data: {
            firstName,
            lastName,
            bio,
            avatarURL,
        },
    });
}

async function deleteProfile(userID: string) {
    await prisma.userProfile.delete({
        where: { userID },
    });
}

export { createProfile, getProfile, updateProfile, deleteProfile };
