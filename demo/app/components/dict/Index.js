import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import i18next from 'i18next'
import { Alert } from 'react-bootstrap'

import SignIn  from '../oauth/PleaseSignIn'
import { CurrentUser } from '../../mixins'
import {Edit, Show, Index} from './Notes'
import Search from './Search'
import {ajax} from '../../utils'

const Widget = React.createClass({
  mixins: [CurrentUser],
  getInitialState: function() {
    return {notes:[], note:{}};
  },
  componentDidMount() {
    ajax("get", "/dict/notes", null, function(rst){
      this.setState({notes:rst});
    }.bind(this), null, true)
  },
  render: function() {
    return  this.isSignIn() ?
        (<div className="row">
          <br/>
          <div className="col-md-offset-1 col-md-7">
            <Search/>
          </div>
          <div className="col-md-3">
            <Edit note={this.state.note}/>
            <br/>
            <Index notes={this.state.notes}/>
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
