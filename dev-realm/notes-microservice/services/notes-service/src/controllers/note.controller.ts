import { Request, Response } from 'express';
import * as noteService from '../services/note.service';
import { sendResponse } from '@shared/src/utils/appResponse.util';
import { AppError } from '@shared/src/types';

async function create(req: Request, res: Response) {
    const userID = req.user?.userID;
    const { title, content }: CreateNoteType = req.body;

    if (!userID) {
        throw new AppError('Authentication required.', 401);
    }

    const note = await noteService.create({ userID, title, content });

    return sendResponse(res, 200, {
        success: true,
        message: 'Note has been updated.',
        data: note,
    });
}

async function get(req: Request, res: Response) {
    const { noteID } = req.params;

    const note = await noteService.get(noteID as string);
    return sendResponse(res, 200, {
        success: true,
        message: 'Note has been updated.',
        data: note,
    });
}

async function update(req: Request, res: Response) {
    const { noteID } = req.params;
    const { title, content } = req.body;

    const note = await noteService.update(noteID as string, { title, content });
    return sendResponse(res, 200, {
        success: true,
        message: 'Note has been updated.',
        data: note,
    });
}

async function patch(req: Request, res: Response) {
    const { noteID } = req.params;
    const { title, content } = req.body;

    const note = await noteService.patch(noteID as string, { title, content });
    return sendResponse(res, 200, {
        success: true,
        message: 'Note has been updated.',
        data: note,
    });
}

async function remove(req: Request, res: Response) {
    const { noteID } = req.params;

    await noteService.remove(noteID as string);

    return sendResponse(res, 200, {
        success: true,
        message: 'Note has been deleted.',
    });
}

export { create, get, update, patch, remove };
