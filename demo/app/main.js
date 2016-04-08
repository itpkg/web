require('bootstrap/dist/css/bootstrap.css');
require('bootstrap/dist/css/bootstrap-theme.css');
require('react-select/dist/react-select.css');
require('./main.css');

import i18next from 'i18next';
import XHR from 'i18next-xhr-backend';
import Cache from 'i18next-localstorage-cache';
import LngDetector from 'i18next-browser-languagedetector';

import React from 'react'
import ReactDOM from 'react-dom'
import {createStore, combineReducers, applyMiddleware} from 'redux'
import {Provider} from 'react-redux'
import {Router, Route, browserHistory} from 'react-router'
import {syncHistoryWithStore, routerReducer} from 'react-router-redux'

import * as reducers from './reducers'
import App from './components/App'

i18next
    .use(XHR)
    .use(Cache)
    .use(LngDetector)
    .init({
        detection: {
            order: ['querystring', 'cookie', 'localStorage', 'navigator'],
            lookupQuerystring: 'locale',
            lookupCookie: 'locale',
            lookupLocalStorage: 'locale',
            caches: ['localStorage', 'cookie'],
            cookieMinutes: 60 * 24 * 365 * 10,
        },
        backend: {
            loadPath: API_HOST + '/locales/{{lng}}.json',
            crossDomain: process.env.NODE_ENV != "production",
        },
        cache: {
            enable: process.env.NODE_ENV == "production",
            prefix: "locales_",
            expirationTime: 7 * 24 * 60 * 60 * 1000
        }
    });

const store = createStore(
    combineReducers({
        ...reducers,
        routing: routerReducer
    })
)

const history = syncHistoryWithStore(browserHistory, store)

ReactDOM.render(
  (<Provider store = { store } >
    <Router history = { history }>
      <Route path = "/" component = { App } >
      </Route>
    </Router>
  </Provider>),
    document.getElementById('root')
)
