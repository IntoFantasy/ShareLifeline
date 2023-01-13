$(function(){
    var more = true;
    //限制字符个数
    var hide = function (){
    $(".more_h").each(function(){
    var maxheight=320;
    if($(this).text().length>maxheight){
    $(this).text($(this).text().substring(0,maxheight));
    $(this).html($(this).html()+'...');
    more = true;
    $('.more').html("MORE  <i class=\"fa fa-angle-down\"></i>");
} else {
    $(this).next().next().hide();
}
});
};
    hide();
    $('.more').click(function(){
    if(more){
        // $(this).prev().html(content);
        var new_content = $(this).prev().html();
        $(this).prev().prev().html(new_content);
    // $(".more_h").html(content);
    $(this).html("FOLD  <i class=\"fa fa-angle-up\"></i>");
    more = false;
}else{
    hide();
}
});
})

$(function () {
    $(".like-give").click(function () {
        $(this).toggleClass('give-like-click');
    })
})

$(function () {
    $(".collection-give").click(function () {
        $(this).toggleClass('collection-give-click');
    })
})

$(function () {
    $(".user-recommendation-follow").click(function () {
        if ($(this).hasClass("unfollow")){
            $(this).html("<i class=\"fa-solid fa-plus\"></i><p>Follow</p>");
            var followee = $(this).prev()[0].getElementsByClassName("user-recommendation-name")[0].innerText
            $(this).html("<i class=\"fa-solid fa-x\"></i><p>Unfollow</p>");
            $.get("/FollowPost", {"type": "unfollow", "follower": "user", "followee": followee})
        }
        else {
            var followee = $(this).prev()[0].getElementsByClassName("user-recommendation-name")[0].innerText
            $(this).html("<i class=\"fa-solid fa-x\"></i><p>Unfollow</p>");
            $.get("/FollowPost", {"type": "follow", "follower": "user", "followee": followee})
        }
        $(this).toggleClass('unfollow')

    })
})
