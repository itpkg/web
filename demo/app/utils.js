import $ from 'jquery';

export function isProduction(){
  return process.env.NODE_ENV === 'production';
}

export function ajax(method, url, data, done, fail){
  if(fail == undefined){
    fail = function(e){
      alert(e.responseText);
    }
  }
  $.ajax({
    xhrFields: {
      withCredentials: true
    },
    type: method,
    url: API+url,
    data:data
  }).done(done).fail(fail);
}
