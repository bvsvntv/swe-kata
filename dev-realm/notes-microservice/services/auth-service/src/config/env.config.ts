import dotenv from 'dotenv';
import { z } from 'zod';

dotenv.config({
    path: './.env',
});

const envSchema = z.object({
    SERVER_PORT: z.coerce.number(),
    SERVER_ENV: z.enum(['development', 'stage', 'production']),
    DATABASE_URL: z.string(),
    HASH_SALT_ROUNDS: z.coerce.number(),
    ACCESS_TOKEN_SECRET: z.string(),
    ACCESS_TOKEN_EXPIRES_IN: z.string(),
    REFRESH_TOKEN_SECRET: z.string(),
    REFRESH_TOKEN_EXPIRES_IN: z.string(),
});

const parsedEnv = envSchema.safeParse(process.env);

if (!parsedEnv.success) {
    console.error(
        'Invalid environment variables: ',
        z.treeifyError(parsedEnv.error),
    );

    process.exit(1);
}

export const env = parsedEnv.data;
