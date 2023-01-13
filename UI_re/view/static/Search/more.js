$(function(){
    var content = $(".more_h").html();
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
    $('.more').hide();
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
