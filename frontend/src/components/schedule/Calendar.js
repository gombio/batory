import React from 'react'
import BigCalendar from 'react-big-calendar';
import moment from 'moment';

import { connect } from 'react-redux';
import _ from 'lodash';

BigCalendar.momentLocalizer(moment); // or globalizeLocalizer

class Calendar extends React.Component {
  render() {
    return (
      <BigCalendar
        selectable
        events={this.props.events}
        defaultView='week'
        defaultDate={new Date(2017, 9, 12)}
        onSelectEvent={event => alert(event.title)}
        onSelectSlot={(slotInfo) => alert(
          `selected slot: \n\nstart ${slotInfo.start.toLocaleString()} ` +
          `\nend: ${slotInfo.end.toLocaleString()}`
        )}
      />
    )
  }
}

const mapStateToProps = state => {
  const { events } = state;

  return {
    events,
  };
  // const employees = _.map(state.employees, (val, uid) => {
  //   return {
  //     ...val,
  //     uid: uid,
  //   };
  // });
  //
  // return { employees };
};

export default connect(mapStateToProps, {
  // employeesFetch,
})(Calendar);
