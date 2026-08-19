import { z } from 'zod';

const createNoteSchema = z.object({
    title: z.string().min(3),
    content: z.string(),
});

const updateNoteParamsSchema = z.object({
    noteID: z.uuid(),
});

const updateNoteSchema = createNoteSchema;

const patchNoteSchema = createNoteSchema.partial();

const deleteNoteParamsSchema = updateNoteParamsSchema;

export {
    createNoteSchema,
    updateNoteParamsSchema,
    updateNoteSchema,
    patchNoteSchema,
    deleteNoteParamsSchema,
};
