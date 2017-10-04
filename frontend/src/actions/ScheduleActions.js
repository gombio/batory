import {
  SCHEDULE_ADD,
} from './types';

export const scheduleAdd = (schedule) => {
  return {
    type: SCHEDULE_ADD,
    payload: schedule,
  };
};
