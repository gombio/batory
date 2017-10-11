import React from 'react';
import { connect } from 'react-redux';

import FlatButton from 'material-ui/FlatButton';
import TextField from 'material-ui/TextField';
import Dialog from 'material-ui/Dialog';
import Avatar from 'material-ui/Avatar';
import ScheduleIcon from 'material-ui/svg-icons/action/schedule';

import { authenticate } from '../../actions';

class Login extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      login: '',
      password: '',
    };
  }

  handleSignIn() {
    this.props.authenticate({
      login: this.state.login,
      password: this.state.password,
    });
  }

  handleChange(event) {
    this.setState({...this.state, [event.target.id]: event.target.value});
  }

  render() {
    let style = {
      margin: '15px 20px 0 20px',
    }
    if (this.props.connected) {
      style.backgroundColor = '#0097a7';
    }

    return (
      <Dialog
          title="Batory / Sign in"
          modal={true}
          open={true}
          actions={[<FlatButton label="Sign in" onClick={this.handleSignIn.bind(this)} />]}
        >
          <div style={{flexDirection: 'row', display: 'flex', justifyContent: 'flex-start' }}>
            <div>
              <Avatar icon={<ScheduleIcon />} size={ 120 } style={style} />
            </div>
            <div>
              <TextField
                id="login"
                floatingLabelText="Login / E-mail"
                hintText="foo@example.com"
                onChange={this.handleChange.bind(this)}
              /><br />
              <TextField
                id="password"
                floatingLabelText="Password"
                hintText="SupaDupa!#@"
                type="password"
                onChange={this.handleChange.bind(this)}
              /><br />
            </div>
          </div>
      </Dialog>
    );
  }
}

const mapStateToProps = state => {
  const { connected } = state;
  return { connected };
};

export default connect(mapStateToProps, {
  authenticate,
})(Login);
