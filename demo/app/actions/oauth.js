import { SIGN_IN, SIGN_OUT, OAUTH2 } from '../constants'


export function signIn(token) {
    return {
        type: SIGN_IN,
        token:token
    }
}

export function signOut() {
    return {
        type: SIGN_OUT
    }
}

export function refresh(info) {
    return {
        type: OAUTH2,
        info: info
    }
}
