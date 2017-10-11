import {
  // PROJECT_SCOPE_CHANGED,
  SCHEDULE_ADD,
} from '../actions/types';

const INITIAL_STATE = [];

export default (state = INITIAL_STATE, action) => {
  switch (action.type) {
    case SCHEDULE_ADD:
      const schedule = action.payload;
      let e = {
        title: schedule.person,
        start: new Date(schedule.start),
        end: new Date(schedule.end),
      }
      // console.log(schedule);
      const result = [...state]
      result.push(e)

      return result
    default:
      return state;
  }
};
