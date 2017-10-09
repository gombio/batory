import {
  AUTH_AUTHENTICATED,
} from './types';

export const authenticated = () => {
  return {
    type: AUTH_AUTHENTICATED,
    payload: true,
  };
};

export const unauthenticated = () => {
  return {
    type: AUTH_AUTHENTICATED,
    payload: false,
  };
};
