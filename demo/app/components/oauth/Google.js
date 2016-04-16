import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import i18next from 'i18next'
import { Alert } from 'react-bootstrap'
import parse from 'url-parse'
import { browserHistory } from 'react-router'

import {signIn} from '../../actions/oauth'
import {ajax} from '../../utils'

const Widget = React.createClass({
  componentDidMount() {
        const {onSignIn} = this.props;
        onSignIn();
  },
  render: function() {
    return (
      <div className="row">
        <br/>
        <div className="col-md-offset-1 col-md-10">
          <Alert bsStyle="warning">
            <strong>{i18next.t("messages.please_waiting")}</strong>{new Date().toLocaleString()}
          </Alert>
        </div>
      </div>
    );
  }
});

Widget.propTypes = {
    onSignIn: PropTypes.func.isRequired
};

export default connect(
    state => ({}),
    dispatch => ({
        onSignIn: function () {
            ajax(
                'post',
                '/oauth/sign_in',
                {
                    type: 'google',
                    code: parse(location.href, true).query.code
                },
                function (tkn) {
                    dispatch(signIn(tkn));
                    browserHistory.push('/');
                },
              )
        }
    })
)(Widget);
