import {
  CONNECTED,
} from '../actions/types';

const INITIAL_STATE = false;

export default (state = INITIAL_STATE, action) => {
  switch (action.type) {
    case CONNECTED:
      return action.payload;
    default:
      return state;
  }
};
