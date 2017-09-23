import React, { Component } from 'react';
import './App.css';
import BigCalendar from 'react-big-calendar';
import moment from 'moment';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme';
// import lightBaseTheme from 'material-ui/styles/baseThemes/lightBaseTheme';

import TopMenu from './components/common/TopMenu'
import LeftMenu from './components/common/LeftMenu'
// import Homepage from './components/homepage/Homepage'

//XXX: DUMMY DATA
import events from './dummy/events';

BigCalendar.momentLocalizer(moment); // or globalizeLocalizer

const style = {
  mainContent: {
    paddingTop: 24,
    paddingLeft: 275,
    paddingRight: 24,
    position: 'relative',
    display: 'flex',
    flex: 1,
    flexDirection: 'column',
    minHeight: 0,
    height: '100vh',
  }
}

class App extends Component {
  render() {
    return (
      // <MuiThemeProvider muiTheme={getMuiTheme(lightBaseTheme)}>
      <MuiThemeProvider muiTheme={getMuiTheme(darkBaseTheme)}>
        <div>
          <TopMenu />
          <LeftMenu />
          <div style={style.mainContent}>
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
        </div>
      </MuiThemeProvider>
    );
  }
}

export default App;
