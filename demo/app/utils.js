import $ from 'jquery';

import {TOKEN} from './constants'

export function isProduction(){
  return process.env.NODE_ENV === 'production';
}

export function ajax(method, url, data, done, fail, auth){
  if(fail == undefined){
    fail = function(e){
      alert(e.responseText);
    }
  }
  var args = {
    xhrFields: {
      withCredentials: true
    },
    type: method,
    url: API+url,
    data:data
  };

  if(auth){
    args.beforeSend= function (xhr) {
      xhr.setRequestHeader('Authorization', 'bearer '+sessionStorage.getItem(TOKEN));
    };
  }

  $.ajax(args).done(done).fail(fail);
}
