import React from 'react'
import { createStore, combineReducers, applyMiddleware } from 'redux'
import { Provider } from 'react-redux'
import { Router, Route, browserHistory } from 'react-router'
import { syncHistoryWithStore,routerReducer } from 'react-router-redux'

import reducers from '../reducers'
import NoMatch from './NoMatch'
import Layout from './Layout'

import Base from './base/Router'
import Oauth from './oauth/Router'
import Dict from './dict/Router'


const store = createStore(
  combineReducers({
    ...reducers,
    routing: routerReducer
  })
)

const history = syncHistoryWithStore(browserHistory, store)

const Widget = React.createClass({
  render: function() {
    return (
      <Provider store={store}>
        <Router history={history}>
          <Route path='/' component={Layout}>
            {Base}
            {Oauth}
            {Dict}
            <Route path="*" component={NoMatch}/>
          </Route>
        </Router>
      </Provider>,
    );
  }
});

export default Widget;
