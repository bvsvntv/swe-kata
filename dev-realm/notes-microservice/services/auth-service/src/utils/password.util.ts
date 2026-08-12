import 'dotenv/config';
import bcrypt from 'bcrypt';

export async function hashPassword(password: string): Promise<string> {
    return await bcrypt.hash(
        password,
        Number(process.env.HASH_SALT_ROUNDS) || 10,
    );
}
