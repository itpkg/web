require("bootstrap/dist/css/bootstrap.css");
require("bootstrap/dist/css/bootstrap-theme.css");
require("./main.css");

import i18next from 'i18next';
import XHR from 'i18next-xhr-backend';
import Cache from 'i18next-localstorage-cache';
import LanguageDetector from 'i18next-browser-languagedetector';

import React from 'react'
import ReactDOM from 'react-dom'

import {
    isProduction
} from './utils'
import Root from './components/Root'

i18next
    .use(XHR)
    .use(Cache)
    .use(LanguageDetector)
    .init({
            cache: {
                enable: isProduction(),
                prefix: 'locales_',
                expirationTime: 7 * 24 * 60 * 60 * 1000
            },
            backend: {
                loadPath: API + '/locales/{{lng}}.json',
                crossDomain: true
            },
            detection: {
                order: ['querystring', 'localStorage', 'cookie', 'navigator'],
                lookupQuerystring: 'locale',
                lookupCookie: 'locale',
                lookupLocalStorage: 'locale',

                caches: ['localStorage', 'cookie'],
                cookieMinutes: 365 * 24 * 60
            }
        },
        (err, t) => {
            console.log("Lang: " + i18next.language);
            ReactDOM.render( < Root / > , document.getElementById('root'));
        }
    );
