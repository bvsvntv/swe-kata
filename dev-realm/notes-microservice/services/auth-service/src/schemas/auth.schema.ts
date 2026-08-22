import { z } from 'zod';

const registerSchema = z.object({
    email: z.email(),
    password: z.string().min(8).max(16),
});

const loginSchema = z.object({
    email: z.email(),
    password: z.string().min(8).max(16),
});

export { registerSchema, loginSchema };
