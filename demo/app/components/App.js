import React from 'react'
import {
  Link
} from 'react-router'
import {
  connect
} from 'react-redux'
import $ from 'jquery'

import Header from './Header'
import {
  refresh
} from '../actions/base'
import call from '../ajax';

const Widget = React.createClass({
      componentDidMount: function() {
        const {
          refresh
        } = this.props;

        call(
          "GET",
          null,
          '/site/info',
          function(rst) {
            refresh(rst);
          }.bind(this)
        );
      },
      render: function() {
        const {
          children,
          info
        } = this.props;
        return ( < div >
          < Header / >
          < div className = "container-fluid" > {
            info.Lang
          } {
            children
          } < /div> < /div>)
        }
      });

    export default connect(
      state => ({
        info: state.info,
        user: state.user
      }), {
        refresh
      }
    )(Widget);