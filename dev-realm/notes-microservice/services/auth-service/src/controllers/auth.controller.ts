import { Request, Response } from 'express';
import * as authService from '@/services/auth.service';

async function register(req: Request, res: Response) {
    await authService.register();
}

async function login(req: Request, res: Response) {
    await authService.login();
}

async function refreshTokens(req: Request, res: Response) {
    await authService.refreshTokens();
}

async function logout(req: Request, res: Response) {
    await authService.logout();
}

export { register, login, refreshTokens, logout };
