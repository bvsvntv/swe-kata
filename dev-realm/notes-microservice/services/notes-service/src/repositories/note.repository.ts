import prisma from '@/lib/prisma';
import {
    CreateNoteType,
    UpdateNoteType,
    PatchNoteType,
} from '@/types/note.types';
import { Note } from 'generated/prisma/client';

async function createNote(args: CreateNoteType): Promise<Note> {
    const { userID, title, content } = args;

    return await prisma.note.create({
        data: {
            userID,
            title,
            content,
        },
    });
}

async function getNoteByID(id: string) {
    return await prisma.note.findUnique({
        where: { id },
        select: {
            id: true,
            title: true,
            content: true,
            createdAt: true,
            updatedAt: true,
        },
    });
}

async function getNoteByTitle(title: string) {
    return await prisma.note.findFirst({
        where: { title },
    });
}

async function updateNote(id: string, args: UpdateNoteType) {
    const { title, content } = args;

    return await prisma.note.update({
        where: { id },
        data: {
            title,
            content,
        },
    });
}

async function patchNote(id: string, args: PatchNoteType) {
    const { title, content } = args;

    return await prisma.note.update({
        where: { id },
        data: {
            title,
            content,
        },
    });
}

async function deleteNote(id: string) {
    await prisma.note.delete({
        where: { id },
    });
}

export {
    createNote,
    getNoteByID,
    getNoteByTitle,
    updateNote,
    patchNote,
    deleteNote,
};
