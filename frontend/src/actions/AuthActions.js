import {
  AUTH_AUTHENTICATED,
} from './types';

export const authenticate = (user) => {
  return (dispatch, getState, services) => {
    const { socket } = services;
    socket.on('auth.login.success', () => dispatch(authenticated()));
    // socket.on('auth.login.success', dispatch(authenticated()));
    socket.emit('auth.login', user);
  }
};

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
