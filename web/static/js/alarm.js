$(document).ready(function () {

    $("#volumewarmup").click(function (e) {
        if ($(e.target).is(":checked")) {
            $('#volume-warmup').removeClass('div-disabled');
            $('#volume-warmup').find('*').attr('disabled', false);
        } else {
            $('#volume-warmup').addClass('div-disabled');
            $('#volume-warmup').find('*').attr('disabled', true);
        }
    });
});

