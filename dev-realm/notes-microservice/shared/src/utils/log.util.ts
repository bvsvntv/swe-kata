import path from 'node:path';
import fs from 'node:fs';
import { createLogger, transports, format, Logger } from 'winston';
const { combine, timestamp, printf, errors } = format;
import LokiTransport from 'winston-loki';

type LogConfig = {
    serviceName: string;
    env: string;
    level: string;
    directory?: string;
    lokiTransportHost?: string;
};

export function createWinstonLogger(config: LogConfig): Logger {
    const logDir = config.directory || 'logs';
    // Create the log directory if it doesn't exist
    if (!fs.existsSync(logDir)) {
        fs.mkdirSync(logDir);
    }

    return createLogger({
        defaultMeta: { service: config.serviceName },
        format: combine(
            timestamp({ format: 'YYYY-MM-DD HH:mm:ss' }),
            format.json(),
            printf(({ level, message, timestamp, stack, service }) => {
                const text = `[${timestamp}] [${level.toUpperCase()}] [${service}]: ${message}`;
                return stack ? text + '\n' + stack : text;
            }),
            errors({ stack: true }),
        ),
        transports: [
            new transports.Console(),
            new transports.File({
                filename: path.join(logDir, 'server.log'),
                level: config.level ?? 'info',
            }),
            new transports.File({
                filename: path.join(logDir, 'server-error.log'),
                level: 'error', // Log only errors to this file
            }),
            new LokiTransport({
                host: config.lokiTransportHost ?? 'http://localhost:13100',
                labels: {
                    service: config.serviceName,
                    env: config.env,
                },
            }),
        ],
        exceptionHandlers: [
            new transports.File({
                filename: path.join(logDir, 'server-error.log'),
            }),
        ],
        rejectionHandlers: [
            new transports.File({
                filename: path.join(logDir, 'server-error.log'),
            }),
        ],
    });
}
