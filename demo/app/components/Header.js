import React, {PropTypes} from 'react'
import { Nav, Navbar, NavDropdown, NavItem, MenuItem } from 'react-bootstrap';
import { connect } from 'react-redux'

import PersonalBar from './PersonalBar'

const Widget = React.createClass({
  render: function() {
    const {info} = this.props;
    return (
      <Navbar inverse fixedTop fluid>
          <Navbar.Header>
            <Navbar.Brand>
              <a href = "#">{info.subTitle}</a>
            </Navbar.Brand>
            < Navbar.Toggle />
          </Navbar.Header>
          <Navbar.Collapse>
            <Nav>
              {info.links.map(function(l,i){
                return (<NavItem key={i} href={l.href}>{l.label}</NavItem>);
              })}
              <PersonalBar/>
            </Nav>
            <Nav pullRight>
              <NavItem href="/?locale=en-US" target="_blank"> English </NavItem>
              <NavItem href="/?locale=zh-CN" target="_blank"> 简体中文 </NavItem>
            </Nav>
          </Navbar.Collapse>
        </Navbar>
    );
  }
});

Widget.propTypes = {
    info: PropTypes.object.isRequired,
};

export default connect(
  state => ({ info: state.siteInfo })
)(Widget);
