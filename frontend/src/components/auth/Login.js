import React from 'react';
import Paper from 'material-ui/Paper';
import FlatButton from 'material-ui/FlatButton';
import TextField from 'material-ui/TextField';
import Dialog from 'material-ui/Dialog';
import Avatar from 'material-ui/Avatar';
import ScheduleIcon from 'material-ui/svg-icons/action/schedule';
import FontIcon from 'material-ui/FontIcon';

class Login extends React.Component {
  render() {
    return (
      <Dialog
          title="Batory / Sign in"
          modal={true}
          open={true}
          actions={[<FlatButton label="Sign in" />]}
        >
          <div style={{flexDirection: 'row', display: 'flex', justifyContent: 'flex-start' }}>
            <div>
              <Avatar icon={<ScheduleIcon />} size={ 120 } style={{ margin: '15px 20px 0 20px', backgroundColor: '#0097a7' }} />
            </div>
            <div>
              <TextField
                floatingLabelText="Login / E-mail"
                hintText="foo@example.com"
              /><br />
              <TextField
                floatingLabelText="Password"
                hintText="SupaDupa!#@"
              /><br />
            </div>
          </div>
      </Dialog>
    );
  }
}

export default Login;
