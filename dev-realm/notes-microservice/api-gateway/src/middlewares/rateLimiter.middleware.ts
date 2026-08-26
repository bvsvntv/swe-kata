import rateLimit from 'express-rate-limit';

export const rateLimiter = rateLimit({
    windowMs: 15 * 60 * 1000, // 15 minutes
    max: 100, // Limit each IP to 100 requests per windowMs
    handler: (_, res) => {
        res.status(429).json({
            success: false,
            message: 'Too many requests. Please try again later',
        });
    },
});
