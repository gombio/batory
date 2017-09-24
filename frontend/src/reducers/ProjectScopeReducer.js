import {
  PROJECT_SCOPE_CHANGED,
} from '../actions/types';

const INITIAL_STATE = null;

export default (state = INITIAL_STATE, action) => {
  switch (action.type) {
    case PROJECT_SCOPE_CHANGED:
      return action.payload; //XXX: Am I changing state here properly?
    default:
      return state;
  }

  return state;
};
