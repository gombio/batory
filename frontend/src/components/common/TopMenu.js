import React from 'react';
import { connect } from 'react-redux';

import { projectScopeChange } from '../../actions';

// ToolBar
import {Toolbar, ToolbarGroup } from 'material-ui/Toolbar';

// TopBar
import AppBar from 'material-ui/AppBar';

// Menu
import DropDownMenu from 'material-ui/DropDownMenu';
import MenuItem from 'material-ui/MenuItem';

class TopMenu extends React.Component {
  handleChange(event, index, value) {
    this.props.projectScopeChange(value);
  };

  render() {
    return (
      <AppBar>
        <Toolbar style={{backgroundColor: 'transparent'}}>
          <ToolbarGroup firstChild={true}>
            <DropDownMenu value={this.props.projectScope} onChange={this.handleChange.bind(this)}>
              <MenuItem value={null} primaryText="All projects" />
              <MenuItem value={"projectFoo"} primaryText="Project Foo" />
              <MenuItem value={"projectBar"} primaryText="Project Bar" />
              <MenuItem value={"projectBaz"} primaryText="Project Baz" />
            </DropDownMenu>
          </ToolbarGroup>
        </Toolbar>
      </AppBar>
    );
  }
}

const mapStateToProps = state => {
  const { projectScope } = state;

  return {
    projectScope,
  };
};

export default connect(mapStateToProps, {
  projectScopeChange,
})(TopMenu);
