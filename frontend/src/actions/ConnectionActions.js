import {
  CONNECTED,
} from './types';

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
