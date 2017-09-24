import {
  PROJECT_SCOPE_CHANGED,
} from './types';

export const projectScopeChange = (project) => {
  return {
    type: PROJECT_SCOPE_CHANGED,
    payload: project,
  };
};
