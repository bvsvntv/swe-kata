import { Request, Response } from 'express';
import * as noteService from '../services/note.service';
import { sendResponse } from '@shared/src/utils/appResponse.util';
import { AppError } from '@shared/src/types';

async function createNoteController(req: Request, res: Response) {
    const userID = req.user?.userID;
    const { title, content }: CreateNoteType = req.body;

    if (!userID) {
        throw new AppError('Authentication required.', 401);
    }

    const note = await noteService.create({ userID, title, content });

    return sendResponse(res, 201, {
        success: true,
        message: 'Note has been created.',
        data: note,
    });
}

async function getNoteController(req: Request, res: Response) {
    const { id } = req.params;

    const note = await noteService.get(id as string);
    return sendResponse(res, 200, {
        success: true,
        message: 'Note has been fetched.',
        data: note,
    });
}

async function updateNoteController(req: Request, res: Response) {
    const { id } = req.params;
    const { title, content } = req.body;

    const note = await noteService.update(id as string, { title, content });
    return sendResponse(res, 200, {
        success: true,
        message: 'Note has been updated.',
        data: note,
    });
}

async function patchNoteController(req: Request, res: Response) {
    const { id } = req.params;
    const { title, content } = req.body;

    const note = await noteService.patch(id as string, { title, content });
    return sendResponse(res, 200, {
        success: true,
        message: 'Note has been updated.',
        data: note,
    });
}

async function deleteNoteController(req: Request, res: Response) {
    const { id } = req.params;

    await noteService.remove(id as string);

    return sendResponse(res, 200, {
        success: true,
        message: 'Note has been deleted.',
    });
}

export {
    createNoteController,
    getNoteController,
    updateNoteController,
    patchNoteController,
    deleteNoteController,
};
