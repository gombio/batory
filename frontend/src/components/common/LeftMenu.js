import React from 'react';
import { connect } from 'react-redux';

//TopBar
import Drawer from 'material-ui/Drawer';
import AppBar from 'material-ui/AppBar';
// Menu
import Menu from 'material-ui/Menu';
import MenuItem from 'material-ui/MenuItem';
import Divider from 'material-ui/Divider';
//Icons
import HomeIcon from 'material-ui/svg-icons/action/home';
import ScheduleIcon from 'material-ui/svg-icons/action/schedule';
import PeopleIcon from 'material-ui/svg-icons/social/people';
import ErrorIcon from 'material-ui/svg-icons/alert/error';
import ExitIcon from 'material-ui/svg-icons/action/exit-to-app';
import SettingsIcon from 'material-ui/svg-icons/action/settings';
import HealingIcon from 'material-ui/svg-icons/image/healing';
import LightBulb from 'material-ui/svg-icons/action/lightbulb-outline';

class LeftMenu extends React.Component {
  isConnected() {
    let color = (this.props.connected) ? 'ffff00' : 'dddddd';

    return (<div>Batory v.0.0.1 <LightBulb color={color} /></div>);
  }

  render() {
    return (
      <Drawer open={true}>
        <AppBar title={this.isConnected()} showMenuIconButton={false} titleStyle={{color: 'white'}} />
        <Menu>
          <MenuItem primaryText="Home" leftIcon={<HomeIcon />} />
          <Divider />
          <MenuItem primaryText="People" leftIcon={<PeopleIcon />} />
          <MenuItem primaryText="On-Call schedule" leftIcon={<ScheduleIcon />} />
          <MenuItem primaryText="Health checks" leftIcon={<HealingIcon />} />
          <MenuItem primaryText="Issues" leftIcon={<ErrorIcon />} />
          <Divider />
          <MenuItem primaryText="Settings" leftIcon={<SettingsIcon />} />
          <Divider />
          <MenuItem primaryText="Exit" leftIcon={<ExitIcon />} />
        </Menu>
      </Drawer>
    );
  }
}

const mapStateToProps = state => {
  const { connected } = state;

  return {
    connected,
  };
};

export default connect(mapStateToProps, {
  //
})(LeftMenu);
