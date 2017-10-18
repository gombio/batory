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

export const schedulesCreate = (project, person, start, end) => {
  return (dispatch, getState, services) => {
    const { socket } = services;
    socket.emit('schedules.create', {
      project: project,
      person: person,
      start: start.toISOString(),
      end: end.toISOString(),
    });
  }
}
