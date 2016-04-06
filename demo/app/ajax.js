import $ from "jquery"

export default function(method, data, url, done, fail){
  $.ajax({
    method:method,
    data:data,
    xhrFields: {withCredentials: true},
    url:API_HOST+url,
  }).done(done).fail(fail);
}
