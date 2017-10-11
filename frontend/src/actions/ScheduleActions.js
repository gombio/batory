import {
  SCHEDULE_ADD,
} from './types';

export const schedulesFetch = () => {
  return (dispatch, getState, services) => {
    const { socket } = services;
    socket.on('schedule.add', (schedule) => dispatch(scheduleAdd(schedule)));
    socket.emit('schedules.list');
  }
};

export const scheduleAdd = (schedule) => {
  return {
    type: SCHEDULE_ADD,
    payload: schedule,
  };
};
