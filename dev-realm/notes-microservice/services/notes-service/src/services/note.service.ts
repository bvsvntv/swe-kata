import {
    createNote,
    deleteNote,
    getNote,
    patchNote,
    updateNote,
} from '@/repositories/note.repository';

async function create() {
    console.log('create @ note service');
    return createNote();
}

async function get() {
    console.log('get @ note service');
    return getNote();
}

async function update() {
    console.log('update @ note service');
    return updateNote();
}

async function patch() {
    console.log('pathch @ note service');
    return patchNote();
}

async function remove() {
    console.log('remove @ note service');
    return deleteNote();
}

export { create, get, update, patch, remove };
