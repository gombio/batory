import {
  AUTH_AUTHENTICATED,
} from '../actions/types';

const INITIAL_STATE = false;

export default (state = INITIAL_STATE, action) => {
  switch (action.type) {
    case AUTH_AUTHENTICATED:
      return action.payload;
    default:
      return state;
  }
};
