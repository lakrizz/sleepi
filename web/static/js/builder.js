$(document).ready(function () {

    $('.builder-song').each(function(e) {
        console.log(e);
        ele = $(e.target)
        if (ele.hasClass("in-library")) {
            console.log("is in library");
            ele.find(".playlist-remove").hide();
        } 
        else
        {
            console.log("is not in library");
            ele.find(".playlist-remove").show();
        }
    });


    $("#playlist-songs").sortable({
        update: rebuild_playlist,
    }
    );

    $('.builder-song').click(function (e) {
        ele = $(e.target)
        if (ele.hasClass("in-library")) {
            ele.removeClass("in-library");
            target = $('#playlist-songs');
            ele.find(".playlist-remove").show();
            target.append(e.target);
            rebuild_playlist();
        }
    });

    $('.playlist-remove').click(function (e) {
        ele = $(e.target).parent().parent();
        ele.addClass("in-library");
        target = $('#library-songs');
        ele.find(".playlist-remove").hide();
        target.append(ele);
    });

});

function rebuild_playlist() {
    $('#playlist-order').empty();
    $('#playlist-songs div').each(function (i, e) {
        $('#playlist-order').append("<input type='hidden' name='order' value='"+$(e).attr("id")+"' / >");
    });
}