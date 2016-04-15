import React from 'react'
import { createStore, combineReducers, applyMiddleware } from 'redux'
import { Provider } from 'react-redux'
import { Router, Route, browserHistory } from 'react-router'
import { syncHistoryWithStore, routerReducer } from 'react-router-redux'

import reducers from '../reducers'
import Base from './base/Router'
import NoMatch from './NoMatch'

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
          {Base}
          <Route path="*" component={NoMatch}/>
        </Router>
      </Provider>,
    );
  }
});

export default Widget;
