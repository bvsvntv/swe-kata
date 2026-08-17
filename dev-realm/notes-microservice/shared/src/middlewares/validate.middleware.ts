import { Request, Response, NextFunction } from 'express';
import { ZodObject } from 'zod';
import { validateTarget } from '../types';

export function validate(
    schema: ZodObject<any>,
    target: validateTarget = 'body',
) {
    return (req: Request, res: Response, next: NextFunction) => {
        const payload = req[target];
        const result = schema.safeParse(payload);

        if (!result.success) {
            const errors = result.error.issues.map((error) => ({
                field: error.path.join(''),
                message: error.message,
            }));

            res.status(400).json({
                errors: errors
                    .map((error) => `${error.field}: ${error.message}`)
                    .join(', '),
            });
        }

        req[target] = result.data;

        // Call next
        next();
    };
}
