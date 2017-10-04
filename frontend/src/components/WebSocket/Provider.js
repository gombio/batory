import React, { Component } from 'react';
import { connect } from 'react-redux';
import { connected, disconnected, scheduleAdd } from '../../actions';

import Socket from '../../Socket';

class Provider extends React.Component {
  componentDidMount() {
    let ws = new WebSocket('ws://localhost:8081/ws');
    let socket = this.socket = new Socket(ws);
    socket.on('connect', this.onConnect.bind(this));
    socket.on('disconnect', this.onDisonnect.bind(this));
    socket.on('schedule.add', this.onScheduleAdd.bind(this));
  }

  onConnect() {
    this.props.connected(); //send action that we're now connected
    this.socket.emit('schedules.list');
  }

  onDisonnect() {
    this.props.disconnected(); //send action that we're now connected
  }

  onScheduleAdd(schedule) {
    this.props.scheduleAdd(schedule);
  }

  render() {
    return (this.props.children);
  }
}

const mapStateToProps = state => {
  const { projectScope, events } = state;

  return {
    projectScope,
    events,
  };
};

export default connect(mapStateToProps, {
  connected,
  disconnected,
  scheduleAdd,
})(Provider);
