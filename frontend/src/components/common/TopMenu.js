import React from 'react';

// ToolBar
import {Toolbar, ToolbarGroup } from 'material-ui/Toolbar';

// TopBar
import AppBar from 'material-ui/AppBar';

// Menu
import DropDownMenu from 'material-ui/DropDownMenu';
import MenuItem from 'material-ui/MenuItem';

const style = {
  appBar: {
    // position: 'fixed',
  },
}

class TopMenu extends React.Component {
  state = {
    value: 1
  }

  handleChange(event, index, value) {
    console.log(event, index, value);
    this.setState({value});
  };

  render() {
    return (
      <AppBar style={style.appBar}>
        <Toolbar style={{backgroundColor: 'transparent'}}>
          <ToolbarGroup firstChild={true}>
            <DropDownMenu value={this.state.value} onChange={this.handleChange.bind(this)}>
              <MenuItem value={1} primaryText="All projects" />
              <MenuItem value={2} primaryText="Project Foo" />
              <MenuItem value={3} primaryText="Project Bar" />
              <MenuItem value={4} primaryText="Project Baz" />
            </DropDownMenu>
          </ToolbarGroup>
        </Toolbar>
      </AppBar>
    );
  }
}

export default TopMenu;
