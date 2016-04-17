import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import { Input, ButtonInput } from 'react-bootstrap'

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
        <form className="form-horizontal" onSubmit={this.handleSubmit}>
          <Input type="text" onChange={this.handleChange} labelClassName="col-md-2" wrapperClassName="col-md-9" label={i18next.t("form.keyword")}/>
          <ButtonInput type="submit" wrapperClassName="col-md-offset-2 col-md-10" value={i18next.t("buttons.search")} />
        </form>
        <br/>
        { rst ? <pre className="col-md-offset-1 col-md-10">{rst}</pre> : <br/>}
      </fieldset>
    )
  }
});

export default Widget;
