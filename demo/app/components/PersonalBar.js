import React, {PropTypes} from 'react'
import { NavDropdown, MenuItem } from 'react-bootstrap'
import { connect } from 'react-redux'
import i18next from 'i18next'

import { CurrentUser } from '../mixins'

const Widget = React.createClass({
  mixins: [CurrentUser],
  render: function() {
    const {user} = this.props;
    var title, links;
    if (this.isSignIn()){
      title = i18next.t("users.welcome", {name:user.name});
      links = [
        {href:"/users/profile", label:i18next.t("users.profile")}
      ];
    }else{
      title = i18next.t("users.sign_in_or_up");
      links = [
        {href:"/users/sign_in", label:i18next.t("users.sign_in_with_google")}
      ];
    }
    return (
      <NavDropdown title={title} id="personal-bar">
        {links.map(function(l, i){
          return l == null ?
            (<MenuItem key={i} divider />) :
            (<MenuItem key={i}>{l.label}</MenuItem>)
        })}
      </NavDropdown>
    );
  }
});

Widget.propTypes = {
    user: PropTypes.object.isRequired,
};

export default connect(
  state => ({ user: state.currentUser })
)(Widget);
