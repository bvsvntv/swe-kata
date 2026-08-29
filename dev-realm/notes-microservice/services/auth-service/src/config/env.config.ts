import dotenv from 'dotenv';
import { z } from 'zod';
import packageJson from '../../package.json';

dotenv.config({
    path: './.env',
});

const envSchema = z.object({
    SERVICE_NAME: z.string().default(packageJson.name),
    SERVER_PORT: z.coerce.number(),
    SERVER_ENV: z.enum(['development', 'stage', 'production']),
    DATABASE_URL: z.string(),
    HASH_SALT_ROUNDS: z.coerce.number(),
    ACCESS_TOKEN_SECRET: z.string(),
    ACCESS_TOKEN_EXPIRES_IN: z.string(),
    REFRESH_TOKEN_SECRET: z.string(),
    REFRESH_TOKEN_EXPIRES_IN: z.string(),
    LOG_LEVEL: z.string(),
    LOG_DIRECTORY: z.string(),
    REDIS_URL: z.string(),
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
