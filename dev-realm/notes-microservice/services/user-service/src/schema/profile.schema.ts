import { z } from 'zod';

const createProfileSchema = z.object({
    firstName: z
        .string()
        .min(1, 'First name is required.')
        .max(32, 'First name must be 32 characters or less.'),
    lastName: z
        .string()
        .min(1, 'First name is required.')
        .max(32, 'First name must be 32 characters or less.'),
    bio: z.string().max(500, 'Bio must be 500 characters or less.').optional(),
    avatarURL: z.url('Please provide valid avatar URL.').optional(),
});

const updateProfileSchema = createProfileSchema.partial();

export { createProfileSchema, updateProfileSchema };
