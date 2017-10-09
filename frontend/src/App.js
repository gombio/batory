import React, { Component } from 'react';
import { createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import ReduxThunk from 'redux-thunk';
import './App.css';

import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import TopMenu from './components/common/TopMenu'
import LeftMenu from './components/common/LeftMenu'
import Calendar from './components/schedule/Calendar'
// import Homepage from './components/homepage/Homepage'
import Authentication from './components/auth/Authentication'

import reducers from './reducers';
import batoryTheme from './themes/batory'

import WebSocketProvider from './components/WebSocket/Provider'

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
    //XXX: This may be totally stupid!
    // const ws = new WebSocket('ws://localhost:8081/ws');
    const store = createStore(
      reducers,
      {}, //default state
      // applyMiddleware(ReduxThunk.withExtraArgument({ ws }))
      applyMiddleware(ReduxThunk)
    );
    // console.log(batoryTheme);
    // console.log(getMuiTheme(batoryTheme));
    return (
      <Provider store={store}>
        <WebSocketProvider>
          <MuiThemeProvider muiTheme={getMuiTheme(batoryTheme)}>
            <Authentication>
              <div>
                <TopMenu />
                <LeftMenu />
                <div style={style.mainContent}>
                  <Calendar />
                </div>
              </div>
            </Authentication>
          </MuiThemeProvider>
        </WebSocketProvider>
      </Provider>
    );
  }
}

export default App;
