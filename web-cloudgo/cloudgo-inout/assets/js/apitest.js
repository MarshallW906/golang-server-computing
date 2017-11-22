$(document).ready(function() {
    $.ajax({
        url: "/api/test"
    }).then(function(data) {
        $('.rand').append(data.randtoken);
    });
});