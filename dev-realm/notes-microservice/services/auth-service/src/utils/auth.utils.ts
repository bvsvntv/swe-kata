import crypto from 'crypto';

function hashValue(value: string): string {
    return crypto.createHash('sha256').update(value).digest('hex');
}

export { hashValue };
