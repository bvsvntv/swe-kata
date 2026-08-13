import prisma from '@/lib/prisma';
import { User } from 'generated/prisma/client';

export async function findUserByEmail(email: string) {
    return await prisma.user.findUnique({
        where: { email },
    });
}

export async function createUser(
    email: string,
    password: string,
): Promise<User> {
    return await prisma.user.create({
        data: {
            email,
            password,
        },
    });
}
