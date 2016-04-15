import $ from 'jquery';

export function Get(url, data, done, fail){
  $.ajax({
    url: API+url,
    xhrFields: {
      withCredentials: true
    },
    data:data
  }).done(done).fail(fail);
}
