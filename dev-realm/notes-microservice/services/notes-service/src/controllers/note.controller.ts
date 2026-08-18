import { Request, Response } from 'express';
import * as noteService from '../services/note.service';

async function create(req: Request, res: Response) {
    await noteService.create();
}

async function get(req: Request, res: Response) {
    await noteService.get();
}

async function update(req: Request, res: Response) {
    await noteService.update();
}

async function patch(req: Request, res: Response) {
    await noteService.patch();
}

async function remove(req: Request, res: Response) {
    await noteService.remove();
}

export { create, get, update, patch, remove };
