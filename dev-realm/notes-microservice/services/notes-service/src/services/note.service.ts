import {
    createNote,
    deleteNote,
    getNoteByID,
    getNoteByTitle,
    patchNote,
    updateNote,
} from '@/repositories/note.repository';
import { AppError } from '@shared/src/types';
import { logger } from '@/lib/logger';

async function create(args: CreateNoteType) {
    const { title } = args;

    const existing = await getNoteByTitle(title);
    if (existing) {
        logger.warn('Domain: note creation failed - duplicate title', {
            event: 'note_create_failed',
            reason: 'duplicate_title',
            title,
            userId: args.userID,
        });
        throw new AppError('Duplicate title.', 400);
    }

    const note = await createNote(args);

    logger.info('Domain: note created', {
        event: 'note_created',
        noteId: note.id,
        title: note.title,
        userId: note.userID,
    });

    return note;
}

async function get(id: string) {
    const note = await getNoteByID(id);
    if (!note) {
        logger.warn('Domain: note fetch failed - note not found', {
            event: 'note_fetch_failed',
            noteId: id,
        });
        throw new AppError('Note not found.', 404);
    }

    logger.info('Domain: note fetched', {
        event: 'note_fetched',
        noteId: note.id,
    });

    return note;
}

async function update(id: string, args: UpdateNoteType) {
    const { title } = args;

    const note = await getNoteByID(id);
    if (!note) {
        logger.warn('Domain: note update failed - note not found', {
            event: 'note_update_failed',
            reason: 'not_found',
            noteId: id,
        });
        throw new AppError('Note not found.', 404);
    }

    const existing = await getNoteByTitle(title);
    if (existing && existing.id !== id) {
        logger.warn('Domain: note update failed - duplicate title', {
            event: 'note_update_failed',
            reason: 'duplicate_title',
            noteId: id,
            title,
        });
        throw new AppError('Duplicate title.', 400);
    }

    const updated = await updateNote(id, args);

    logger.info('Domain: note updated', {
        event: 'note_updated',
        noteId: updated.id,
        userId: updated.userID,
    });

    return updated;
}

async function patch(id: string, args: PatchNoteType) {
    const { title } = args;

    const note = await getNoteByID(id);
    if (!note) {
        logger.warn('Domain: note patch failed - note not found', {
            event: 'note_patch_failed',
            reason: 'not_found',
            noteId: id,
        });
        throw new AppError('Note not found.', 404);
    }

    if (title !== undefined) {
        const existing = await getNoteByTitle(title);
        if (existing && existing.id !== id) {
            logger.warn('Domain: note patch failed - duplicate title', {
                event: 'note_patch_failed',
                reason: 'duplicate_title',
                noteId: id,
                title,
            });
            throw new AppError('Duplicate title.', 400);
        }
    }

    const patched = await patchNote(id, args);

    logger.info('Domain: note patched', {
        event: 'note_patched',
        noteId: patched.id,
        userId: patched.userID,
    });

    return patched;
}

async function remove(id: string) {
    const note = await getNoteByID(id);
    if (!note) {
        logger.warn('Domain: note deletion failed - note not found', {
            event: 'note_delete_failed',
            reason: 'not_found',
            noteId: id,
        });
        throw new AppError('Note not found.', 404);
    }

    await deleteNote(id);

    logger.info('Domain: note deleted', {
        event: 'note_deleted',
        noteId: note.id,
    });
}

export { create, get, update, patch, remove };
