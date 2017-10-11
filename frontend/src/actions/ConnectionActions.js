import Socket from '../Socket';

import {
  CONNECTED,
} from './types';

//XXX: This is WebSocket entrypoint!
export const serverConnect = () => { //XXX: note, connect is a reserved keyword by react-redux
  return (dispatch, getState, services) => {
    services.socket = new Socket(new WebSocket('ws://localhost:8081/ws'));
    services.socket.on('connect', () => dispatch(connected()));
    services.socket.on('disconnect', () => dispatch(disconnected()));
  }
};

export const connected = () => {
  return {
    type: CONNECTED,
    payload: true,
  };
};

export const disconnected = () => {
  return {
    type: CONNECTED,
    payload: false,
  };
};
