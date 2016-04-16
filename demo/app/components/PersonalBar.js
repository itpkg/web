import React, {PropTypes} from 'react'
import { NavDropdown, MenuItem } from 'react-bootstrap'
import { connect } from 'react-redux'
import i18next from 'i18next'
import { browserHistory } from 'react-router'

import { CurrentUser } from '../mixins'
import { refresh, signOut }from '../actions/oauth'
import { ajax } from '../utils'

const Widget = React.createClass({
  mixins: [CurrentUser],
  componentDidMount: function(){
    const {onRefresh} = this.props;
    onRefresh();
  },
  render: function() {
    const {user, oauth, onSignOut} = this.props;
    var title, links;
    if (this.isSignIn()){
      title = i18next.t("users.welcome", {name:user.name});
      links = [
        {href:"/users/dashboard", label:i18next.t("users.dashboard")},
        null,
        {label:i18next.t("users.sign_out"), click:onSignOut},
      ];
    }else{
      title = i18next.t("users.sign_in_or_up");
      links = [
        {href:oauth.google, label:i18next.t("users.sign_in_with_google")}
      ];
    }
    return (
      <NavDropdown title={title} id="personal-bar">
        {links.map(function(l, i){
          return l == null ?
            (<MenuItem key={i} divider />) :
            (l.click ?
              (<MenuItem key={i} onClick={l.click}>{l.label}</MenuItem>) :
              (<MenuItem key={i} href={l.href}>{l.label}</MenuItem>))
        })}
      </NavDropdown>
    );
  }
});

Widget.propTypes = {
    user: PropTypes.object.isRequired,
    oauth: PropTypes.object.isRequired,
    onRefresh: PropTypes.func.isRequired,
    onSignOut: PropTypes.func.isRequired,
};

export default connect(
  state => ({ user: state.currentUser, oauth:state.oauth2 }),
  dispatch => ({
    onSignOut:function(){
      dispatch(signOut());
      browserHistory.push('/');
    },
    onRefresh: function(){
      ajax("get", "/oauth/sign_in", null, function(rst){
        dispatch(refresh(rst));
      });
    }
  })
)(Widget);
