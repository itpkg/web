import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import { Form, FormGroup, ControlLabel, FormControl, Button } from 'react-bootstrap'

import i18next from 'i18next'

import {ajax} from '../../utils'

const Widget = React.createClass({
  getInitialState: function() {
    return {result:null, keyword:null};
  },
  handleChange: function(e) {
    this.setState({keyword: e.target.value});
  },
  handleSubmit: function(e) {
    e.preventDefault();
    ajax(
      'post',
      '/dict',
      {
        keyword:this.state.keyword
      },
      function(rst){
        this.setState({result:rst});
      }.bind(this),
      null,
      true
    );
  },
  render: function() {
    const rst = this.state.result;
    return (
      <fieldset>
        <Form inline>
          <FormGroup controlId="keyword">
            <ControlLabel>{i18next.t("form.keyword")}：</ControlLabel>
            <FormControl type="text" onChange={this.handleChange}/>
          </FormGroup>
          &nbsp;
          <Button bsStyle="primary" type="submit" onClick={this.handleSubmit}>{i18next.t("buttons.search")}</Button>
        </Form>
        <br/>
        { rst ? <pre>{rst}</pre> : <br/>}
    </fieldset>
    )
  }
});

export default Widget;
