import dotenv from 'dotenv';
import { z } from 'zod';

dotenv.config({
    path: './.env',
});

const envSchema = z.object({
    SERVER_PORT: z.coerce.number(),
    SERVER_ENV: z.enum(['development', 'stage', 'production']),
    DATABASE_URL: z.string(),
    ACCESS_TOKEN_SECRET: z.string(),
    LOG_LEVEL: z.string(),
    LOG_DIRECTORY: z.string(),
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
