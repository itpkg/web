import {
  SIGN_IN,
  SIGN_OUT,
  REFRESH
} from '../constants';

const initialState = {
  info: {}
}

export default function update(state = {
  info: {
    subTitle: ''
  }
}, action) {
  switch (action.type) {
    // case SIGN_IN:
    //   return Object.assign({}, ...state, user:action.user)
    // case SIGN_OUT:
    //   return Object.assign({}, ...state, user:null)
    // case REFRESH:
    //   return Object.assign({}, ...state, info:action.info)
    default: return state;
  }
}