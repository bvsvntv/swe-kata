import { NextFunction, Request, Response } from 'express';
import { Logger } from 'winston';

function createRequestLogger(logger: Logger) {
    return function (req: Request, res: Response, next: NextFunction) {
        const start = process.hrtime.bigint();
        const { method, originalUrl, ip } = req;

        logger.info(`Request: ${method} at ${originalUrl} from ${ip}`);

        res.on('finish', () => {
            const durationMs =
                Number(process.hrtime.bigint() - start) / 1_000_000;

            logger.info(
                `Response: ${method} at ${originalUrl} ${res.statusCode} - ${durationMs.toFixed(2)}ms`,
            );
        });

        next();
    };
}

export { createRequestLogger };
