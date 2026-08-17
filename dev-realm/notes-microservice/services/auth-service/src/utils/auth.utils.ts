import crypto from 'crypto';

function hashValue(value: string): string {
    return crypto.createHash('sha256').update(value).digest('hex');
}

function generateSessionID(): string {
    return crypto.randomUUID();
}

export { hashValue, generateSessionID };
