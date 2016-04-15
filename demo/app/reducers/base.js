import {
    SIGN_IN,
    SIGN_OUT,
    REFRUSH
} from '../constants'

export default function(state = {}, action) {
    switch (action) {
        case SIGN_IN:
            return {
                user: action.user
            };
        case SIGN_OUT:
            return {
                user: null
            };
        case REFRUSH:
            return {
                site: action.info
            };
        default:
            return state;
    }
}
