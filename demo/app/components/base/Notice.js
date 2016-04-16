import React from 'react'

import { Alert } from 'react-bootstrap';

export const Index = React.createClass({
  render: function() {
    return (
      <Alert bsStyle="warning">
    <strong>Holy guacamole!</strong> Best check yo self, youre not looking too good.
  </Alert>
    );
  }
});

export const Show = React.createClass({
  render: function() {
    return (
      <div>
        notice {this.props.id}
      </div>
    );
  }
});
