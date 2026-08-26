import { env } from '@/config/env.config';
import { ServiceConfigType } from '@/types';

const serviceConfigs: ServiceConfigType[] = [
    {
        name: 'auth-service',
        path: '/api/v1/auth/',
        url: env.AUTH_SERVICE_URL,
        pathRewrite: { '^/': '/api/v1/auth/' },
        timeout: 5000,
    },
    {
        name: 'user-service',
        path: '/api/v1/user/',
        url: env.USER_SERVICE_URL,
        pathRewrite: { '^/': '/api/v1/user/' },
    },
    {
        name: 'notes-service',
        path: '/api/v1/notes/',
        url: env.NOTES_SERVICE_URL,
        pathRewrite: { '^/': '/api/v1/notes/' },
    },
];
