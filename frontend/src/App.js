import React, { Component } from 'react';
import './App.css';
import BigCalendar from 'react-big-calendar';
import moment from 'moment';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme';
import lightBaseTheme from 'material-ui/styles/baseThemes/lightBaseTheme';
import AppBar from 'material-ui/AppBar';

//XXX: DUMMY DATA
import events from './dummy/events';

import RaisedButton from 'material-ui/RaisedButton';

BigCalendar.momentLocalizer(moment); // or globalizeLocalizer

class App extends Component {
  render() {
    console.log(events);
    return (
      <MuiThemeProvider muiTheme={getMuiTheme(lightBaseTheme)}>
        <div>
          <AppBar title="Batory v.0.0.1" />
          <BigCalendar
            selectable
            events={events}
            defaultView='week'
            defaultDate={new Date(2017, 9, 12)}
            onSelectEvent={event => alert(event.title)}
            onSelectSlot={(slotInfo) => alert(
              `selected slot: \n\nstart ${slotInfo.start.toLocaleString()} ` +
              `\nend: ${slotInfo.end.toLocaleString()}`
            )}
          />
        </div>
      </MuiThemeProvider>
    );
  }
}

export default App;
