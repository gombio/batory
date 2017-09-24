import {
  PROJECT_SCOPE_CHANGED,
} from '../actions/types';

import _ from 'lodash';

//XXX: DUMMY DATA
import events from '../dummy/events';

const INITIAL_STATE = [];

export default (state = INITIAL_STATE, action) => {
  switch (action.type) {
    case PROJECT_SCOPE_CHANGED:
      const project = action.payload;
      let result = INITIAL_STATE;
      if (null == project) {
        result = []
        _.map(events, (projectEvents, projectName) => {
          projectEvents.map((e) => {
            result.push(e);
          });

          return {}
        });

        return result;
      }

      result = (undefined == events[project]) ? INITIAL_STATE : events[project];

      return result;
    default:
      return state;
  }

  return state;
};
