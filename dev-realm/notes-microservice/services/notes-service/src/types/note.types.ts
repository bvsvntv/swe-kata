type CreateNoteType = {
    userID: string;
    title: string;
    content: string;
};

type UpdateNoteType = {
    title: string;
    content: string;
};

type PatchNoteType = Partial<UpdateNoteType>;

export { CreateNoteType, UpdateNoteType, PatchNoteType };
