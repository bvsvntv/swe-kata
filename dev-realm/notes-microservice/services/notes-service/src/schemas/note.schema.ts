import { z } from 'zod';

const createNoteSchema = z.object({
    title: z.string().min(3),
    content: z.string(),
});

const getNoteParamsSchema = z.object({
    id: z.uuid(),
});

const updateNoteParamsSchema = getNoteParamsSchema;

const updateNoteSchema = createNoteSchema;

const patchNoteSchema = createNoteSchema.partial();

const deleteNoteParamsSchema = updateNoteParamsSchema;

export {
    createNoteSchema,
    getNoteParamsSchema,
    updateNoteParamsSchema,
    updateNoteSchema,
    patchNoteSchema,
    deleteNoteParamsSchema,
};
