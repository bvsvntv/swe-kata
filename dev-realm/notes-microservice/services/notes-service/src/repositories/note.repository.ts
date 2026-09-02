import { metricRegistry } from '@/lib/metrics';
import prisma from '@/lib/prisma';
import {
    CreateNoteType,
    UpdateNoteType,
    PatchNoteType,
} from '@/types/note.types';
import { createDBMetrics } from '@shared/src/utils/dbMetric.util';
import { Note } from 'generated/prisma/client';

const db = createDBMetrics(metricRegistry, 'notes');

async function createNote(args: CreateNoteType): Promise<Note> {
    const { userID, title, content } = args;

    return db.measure('note', 'create', () =>
        prisma.note.create({
            data: {
                userID,
                title,
                content,
            },
        }),
    );
}

async function getNoteByID(id: string) {
    return db.measure('note', 'findUnique', () =>
        prisma.note.findUnique({
            where: { id },
            select: {
                id: true,
                title: true,
                content: true,
                createdAt: true,
                updatedAt: true,
            },
        }),
    );
}

async function getNoteByTitle(title: string) {
    return db.measure('note', 'findFirst', () =>
        prisma.note.findFirst({
            where: { title },
        }),
    );
}

async function updateNote(id: string, args: UpdateNoteType) {
    const { title, content } = args;

    return db.measure('note', 'update', () =>
        prisma.note.update({
            where: { id },
            data: {
                title,
                content,
            },
        }),
    );
}

async function patchNote(id: string, args: PatchNoteType) {
    const { title, content } = args;

    return db.measure('note', 'patch', () =>
        prisma.note.update({
            where: { id },
            data: {
                title,
                content,
            },
        }),
    );
}

async function deleteNote(id: string) {
    db.measure('note', 'delete', () =>
        prisma.note.delete({
            where: { id },
        }),
    );
}

export {
    createNote,
    getNoteByID,
    getNoteByTitle,
    updateNote,
    patchNote,
    deleteNote,
};
