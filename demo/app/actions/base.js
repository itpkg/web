import { REFRESH } from '../constants'

export function refresh(info) {
    return {
        type: REFRESH,
        info: info
    }
}
