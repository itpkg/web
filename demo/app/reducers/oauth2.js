import {  OAUTH2 } from '../constants'

const initState = {};

export function oauth2(state = initState, action) {
    switch (action.type) {
        case OAUTH2:          
            return action.info;
        default:
            return state;
    }
}
