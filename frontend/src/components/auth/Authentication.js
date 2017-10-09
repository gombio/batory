import React, { Component } from 'react';
import { connect } from 'react-redux';
import { connected, disconnected, scheduleAdd } from '../../actions';

import Login from './Login'

class Authentication extends React.Component {
  render() {
    if (this.props.authenticated) {
      return (this.props.children);
    }

    return (<Login />);
  }
}

const mapStateToProps = state => {
  const { authenticated } = state;

  return {
    authenticated
  };
};

export default connect(mapStateToProps, {
})(Authentication);
