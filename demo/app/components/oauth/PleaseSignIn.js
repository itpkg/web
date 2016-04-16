import React from 'react'
import { Alert } from 'react-bootstrap'
import i18next from 'i18next'

const Widget = React.createClass({
  render: function() {
    return (
      <div className="row">
        <br/>
        <div className="col-md-offset-1 col-md-10">
          <Alert bsStyle="warning">
            <strong>{i18next.t("messages.please_sign_in")}：</strong>{new Date().toLocaleString()}
          </Alert>
        </div>
      </div>
    );
  }
});

export default Widget;
