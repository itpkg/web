import React, { PropTypes } from 'react'
import { connect } from 'react-redux'

import Header from './Header'
import Footer from './Footer'
import {refresh} from '../actions/base'
import {ajax} from '../utils'

const Widget = React.createClass({
  componentDidMount: function(){
    const {onRefresh} = this.props;
    onRefresh();
  },
  render: function() {
    return (
      <div>
        <Header/>
        <div className="container-fluid">
            {this.props.children}
            <hr/>
            <Footer/>
        </div>
      </div>
    );
  }
});

Widget.propTypes = {
    onRefresh: PropTypes.func.isRequired
};

export default connect(
  state=>({info:state.siteInfo}),
  dispatch => ({
    onRefresh: function(){
      ajax("get", "/site/info", null, function(rst){
        dispatch(refresh(rst));
      });
    }
  })
)(Widget);
