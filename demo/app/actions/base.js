import {
    SIGN_IN,
    SIGN_OUT,
    REFRUSH
} from '../constants'

export function signIn(user) {
    return {
        type: SIGN_IN,
        user: user
    }
}

export function signOut() {
    return {
        type: SIGN_OUT
    }
}

export function refresh(info) {
    return {
        type: REFRUSH,
        info: info
    }
}
