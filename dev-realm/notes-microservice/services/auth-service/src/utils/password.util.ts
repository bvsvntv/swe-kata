import { env } from '@/config/env.config';
import bcrypt from 'bcrypt';

async function hashPassword(password: string): Promise<string> {
    return await bcrypt.hash(password, env.HASH_SALT_ROUNDS);
}

async function checkPassword(
    password: string,
    passwordHash: string,
): Promise<boolean> {
    return await bcrypt.compare(password, passwordHash);
}

export { hashPassword, checkPassword };
