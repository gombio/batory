import { combineReducers } from 'redux';
import ProjectScopeReducer from './ProjectScopeReducer';
import SchedulesReducer from './SchedulesReducer';
import ConnectionReducer from './ConnectionReducer';
// import EmployeeFormReducer from './EmployeeFormReducer';
// import EmployeeReducer from './EmployeeReducer';

export default combineReducers({
  connected: ConnectionReducer,
  projectScope: ProjectScopeReducer,
  schedules: SchedulesReducer,
  // auth: AuthReducer,
  // employeeForm: EmployeeFormReducer,
  // employees: EmployeeReducer,
});
