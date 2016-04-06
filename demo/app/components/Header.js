import React from 'react';
import {Nav, Navbar, NavDropdown, NavItem, MenuItem} from 'react-bootstrap';
import { connect } from 'react-redux'

const Widget = React.createClass({
  render:function(){
    const {info} = this.props;
    return (<Navbar inverse fixedTop fluid>
    <Navbar.Header>
      <Navbar.Brand>
        <a href="#">{info}</a>
      </Navbar.Brand>
      <Navbar.Toggle />
    </Navbar.Header>
    <Navbar.Collapse>
      <Nav>
        <NavItem eventKey={1} href="#">Link</NavItem>
        <NavItem eventKey={2} href="#">Link</NavItem>
        <NavDropdown eventKey={3} title="Dropdown" id="basic-nav-dropdown">
          <MenuItem eventKey={3.1}>Action</MenuItem>
          <MenuItem eventKey={3.2}>Another action</MenuItem>
          <MenuItem eventKey={3.3}>Something else here</MenuItem>
          <MenuItem divider />
          <MenuItem eventKey={3.3}>Separated link</MenuItem>
        </NavDropdown>
      </Nav>
      <Nav pullRight>
        <NavItem eventKey={1} href="#">Link Right</NavItem>
        <NavItem eventKey={2} href="#">Link Right</NavItem>
      </Nav>
    </Navbar.Collapse>
  </Navbar>);
  }
});


export default connect(
  state=>({info:state.info, user:state.user})
)(Widget);
