import { env } from '@/config/env.config';
import bcrypt from 'bcrypt';

export async function hashPassword(password: string): Promise<string> {
    return await bcrypt.hash(password, env.HASH_SALT_ROUNDS);
}
