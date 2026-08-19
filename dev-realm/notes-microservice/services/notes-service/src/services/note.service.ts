import {
    createNote,
    deleteNote,
    getNoteByID,
    getNoteByTitle,
    patchNote,
    updateNote,
} from '@/repositories/note.repository';
import { AppError } from '@shared/src/types';

async function create(args: CreateNoteType) {
    const { title } = args;

    const existing = await getNoteByTitle(title);
    if (existing) {
        throw new AppError('Duplicate title.', 400);
    }

    return createNote(args);
}

async function get(id: string) {
    return getNoteByID(id);
}

async function update(id: string, args: UpdateNoteType) {
    const { title } = args;

    const note = await getNoteByID(id);
    if (!note) {
        throw new AppError('Note not found.', 404);
    }

    const existing = await getNoteByTitle(title);
    if (existing && existing.id !== id) {
        throw new AppError('Duplicate title.', 400);
    }

    return updateNote(id, args);
}

async function patch(id: string, args: PatchNoteType) {
    const { title } = args;

    const note = await getNoteByID(id);
    if (!note) {
        throw new AppError('Note not found.', 404);
    }

    if (title !== undefined) {
        const existing = await getNoteByTitle(title);
        if (existing && existing.id !== id) {
            throw new AppError('Duplicate title.', 400);
        }
    }

    return patchNote(id, args);
}

async function remove(id: string) {
    return deleteNote(id);
}

export { create, get, update, patch, remove };
