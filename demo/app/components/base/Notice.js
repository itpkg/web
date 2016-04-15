import React from 'react'

export const Index = React.createClass({
  render: function() {
    return (
      <div>
        notices
      </div>
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
