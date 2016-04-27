import { DICT_NOTE_ADD, DICT_NOTE_REF, DICT_NOTE_DEL } from '../constants'

export function refDictNotes(items) {
    return {
        type: DICT_NOTE_DEL,
        notes: items
    }
}

export function addDictNote(item) {
    return {
        type: DICT_NOTE_ADD,
        note: item
    }
}

export function delDictNotes(id) {
    return {
        type: DICT_NOTE_DEL,
        id:id
    }
}
