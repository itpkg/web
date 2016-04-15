import {  REFRESH } from '../../constants'

const initState = {};

export function siteInfo(state = initState, action) {
    //console.log(action);
    switch (action.type) {
        case REFRESH:
            return action.info;
        default:
            return state;
    }
}
