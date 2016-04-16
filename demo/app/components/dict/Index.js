import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import i18next from 'i18next'
import { Alert } from 'react-bootstrap'

import SignIn  from '../oauth/PleaseSignIn'
import { CurrentUser } from '../../mixins'

const Widget = React.createClass({
  mixins: [CurrentUser],
  render: function() {
    return  this.isSignIn() ?
        (<div className="row">
          <br/>
          <div className="col-md-offset-1 col-md-7">

          </div>
          <div className="col-md-3">
            <h4></h4>
            <hr/>
          </div>
        </div>) :
        (<SignIn />);
  }
});

Widget.propTypes = {
    user: PropTypes.object.isRequired
};

export default connect(
    state => ({user:state.currentUser})
)(Widget);
