import React from 'react'
import BigCalendar from 'react-big-calendar';
import Dialog from 'material-ui/Dialog';
import FlatButton from 'material-ui/FlatButton';
import SelectField from 'material-ui/SelectField';
import MenuItem from 'material-ui/MenuItem';

import moment from 'moment';

import { connect } from 'react-redux';

import { schedulesFetch, schedulesCreate } from '../../actions';

BigCalendar.momentLocalizer(moment); // or globalizeLocalizer

const INITIAL_STATE = {
  showForm: false,
  project: "youdash", //TODO: Fetch it from project context or some other source
  person: "",
  start: "",
  end: "",
}

class Calendar extends React.Component {
  state = INITIAL_STATE;

  componentDidMount() {
    this.props.schedulesFetch();
  }

  handleChange = (event, index, person) => {
    this.setState({person});
  };

  handleClose() {
    this.setState({showForm: false});
  }

  handleCreate() {
    const { project, person, start, end } = this.state;
    if ("" === person || "" === project || "" === start || "" === end) {
      return;
    }

    this.props.schedulesCreate(project, person, start, end);
    this.setState({...INITIAL_STATE});
  }

  render() {
    const items = [];
    //TODO: get them from some sane source, state and DB
    items.push(<MenuItem value="john@example.com" key="john@example.com" primaryText="john@example.com" />);
    items.push(<MenuItem value="barnaba@example.com" key="barnaba@example.com" primaryText="barnaba@example.com" />);

    return (
      <div>
        <BigCalendar
          selectable
          events={this.props.schedules}
          defaultView='week'
          defaultDate={new Date(2017, 9, 12)}
          onSelectEvent={event => alert(event.title)}
          onSelectSlot={(slotInfo) => {
            this.setState({
              showForm: true,
              start: slotInfo.start,
              end: slotInfo.end,
            });
          }}
        />
        <Dialog
            title="Create schedule"
            modal={true}
            open={this.state.showForm}
            onRequestClose={this.handleClose.bind(this)}
            actions={[
              <div>
                <FlatButton label="Cancel" onClick={this.handleClose.bind(this)} />
                <FlatButton label="Create" onClick={this.handleCreate.bind(this)} />
              </div>
            ]}
          >
          <div>
            <SelectField
              value={this.state.person}
              onChange={this.handleChange}
              maxHeight={200}
            >
              {items}
            </SelectField>
          </div>
        </Dialog>
      </div>
    )
  }
}

const mapStateToProps = state => {
  const { schedules } = state;

  return {
    schedules,
  };
};

export default connect(mapStateToProps, {
  schedulesFetch,
  schedulesCreate,
})(Calendar);
