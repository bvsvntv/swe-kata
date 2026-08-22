import { env } from '@/config/env.config';
import { createWinstonLogger } from '@shared/src/utils/log.util';
import { Logger } from 'winston';

const logger: Logger = createWinstonLogger({
    level: env.LOG_LEVEL,
    env: env.SERVER_ENV,
    directory: env.LOG_DIRECTORY,
});

export { logger };
