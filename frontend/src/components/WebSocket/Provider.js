import React from 'react';
import { connect } from 'react-redux';
import { serverConnect } from '../../actions';

class Provider extends React.Component {
  componentDidMount() {
    this.props.serverConnect();
  }

  render() {
    return (this.props.children);
  }
}

export default connect(null, {
  serverConnect,
})(Provider);
