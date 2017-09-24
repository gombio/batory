import { combineReducers } from 'redux';
import ProjectScopeReducer from './ProjectScopeReducer';
import EventsReducer from './EventsReducer';
// import EmployeeFormReducer from './EmployeeFormReducer';
// import EmployeeReducer from './EmployeeReducer';

export default combineReducers({
  projectScope: ProjectScopeReducer,
  events: EventsReducer,
  // auth: AuthReducer,
  // employeeForm: EmployeeFormReducer,
  // employees: EmployeeReducer,
});
