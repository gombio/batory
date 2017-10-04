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
    const store = createStore(
      reducers,
      {}, //default state
      applyMiddleware(ReduxThunk)
    );
    // console.log(batoryTheme);
    // console.log(getMuiTheme(batoryTheme));
    return (
      <Provider store={store}>
        <WebSocketProvider>
          <MuiThemeProvider muiTheme={getMuiTheme(batoryTheme)}>
            <div>
              <TopMenu />
              <LeftMenu />
              <div style={style.mainContent}>
                <Calendar />
              </div>
            </div>
          </MuiThemeProvider>
        </WebSocketProvider>
      </Provider>
    );
  }
}

export default App;
