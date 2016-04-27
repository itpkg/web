import {  DICT_NOTE_REF, DICT_NOTE_ADD, DICT_NOTE_DEL } from '../constants'

const initState = {
  cur : null,
};

export function dictNotes(state = initState, action) {
    switch (action.type) {
        case DICT_NOTE_ADD:
            return action.info;
        default:
            return state;
    }
}
