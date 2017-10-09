import { combineReducers } from 'redux';
import ProjectScopeReducer from './ProjectScopeReducer';
import SchedulesReducer from './SchedulesReducer';
import ConnectionReducer from './ConnectionReducer';
import AuthenticationReducer from './AuthenticationReducer';


export default combineReducers({
  connected: ConnectionReducer,
  authenticated: AuthenticationReducer,
  projectScope: ProjectScopeReducer,
  schedules: SchedulesReducer,
  // auth: AuthReducer,
  // employeeForm: EmployeeFormReducer,
  // employees: EmployeeReducer,
});
