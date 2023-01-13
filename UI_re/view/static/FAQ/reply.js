function text(){
    var text = document.getElementById("send-content").value;
    if (text.length>0){
        var tmp = document.getElementById("comments-list")
        tmp.insertAdjacentHTML("beforeend", "<li>\n" +
            "\t\t\t\t<div class=\"comment-main-level\">\n" +
            "\t\t\t\t\t<!-- Avatar -->\n" +
            "\t\t\t\t\t<div class=\"comment-avatar\"><img src=\"https://pic4.zhimg.com/80/v2-8e1108819305c40d20a7790945276a3f_720w.webp\" alt=\"\"></div>\n" +
            "\t\t\t\t\t<!-- Contenedor del Comentario -->\n" +
            "\t\t\t\t\t<div class=\"comment-box\">\n" +
            "\t\t\t\t\t\t<div class=\"comment-head\">\n" +
            "\t\t\t\t\t\t\t<h6 class=\"comment-name\"><a href=\"http://creaticode.com/blog\">Self</a></h6>\n" +
            "\t\t\t\t\t\t\t<span>hace 0 minutos</span>\n" +
            "\t\t\t\t\t\t\t<i class=\"fa fa-xmark\"></i>\n" +
            "\t\t\t\t\t\t\t<i class=\"fa fa-reply\"></i>\n" +
            "\t\t\t\t\t\t\t<i class=\"fa fa-heart\"></i>\n" +
            "\t\t\t\t\t\t</div>\n" +
            "\t\t\t\t\t\t<div class=\"comment-content\">\n" +
            "<p class=\"touch-disappear\">" + text+
            "                            </p>\n" +
            "                            <div class=\"mini-txt undisplay\">\n" +
            "                                <input  type=\"text\" placeholder=\"Write a comment...\">\n" +
            "                                <div class=\"mini-textarea-icons\">\n" +
            "                                    <button class=\"mini-send\" >\n" +
            "                                        <i class=\"fa-solid fa-paper-plane\"></i>\n" +
            "                                    </button>\n" +
            "                                </div>\n" +
            "                            </div>" +
            "\t\t\t\t\t\t</div>\n" +
            "\t\t\t\t\t</div>\n" +
            "\t\t\t\t</div>\n" +
            "\t\t\t</li>")
        $.post("/CommentSave", {"Comment":text});
    }
}

$(".comments-list").on("mouseenter", ".fa-heart", function () {
    $(this).click(function () {
        $(this).toggleClass('like-give');
    })
})
$(".comments-list").on("mouseenter", ".fa-xmark", function () {
    $(this).click(function () {
        $(this).parent().parent().parent().parent().remove();

    })
})
$(".comments-list").on("mouseenter", ".fa-reply", function () {
    $(this).click(function () {
        var obj=$(this).parent().parent()[0].getElementsByClassName("mini-txt")[0];
        obj.classList.remove("undisplay");

    })
})

$(".comments-list").on("mouseenter", ".touch-disappear", function () {
    $(this).click(function () {
        $(".mini-txt").addClass("undisplay");
    })
})
var count=1
$("#comments-list").on("mouseenter", ".mini-send", function () {
    $(this).click(function () {
        console.log(count)
        count += 1
        var txt = $(this).parent().prev()[0].value;
        if (txt.length>0){
            var obj = document.getElementsByClassName("my-list");
            var tmp;
            for(var i=0;i<obj.length;i++){
                console.log(i);
                if (obj[i].contains(this)){
                    tmp = obj[i];
                    break;
                }
            }
            tmp = tmp.getElementsByClassName("reply-list")[0];
            tmp.insertAdjacentHTML("beforeend", "<li>\n" +
                "\t\t\t\t<div class=\"comment-main-level\">\n" +
                "\t\t\t\t\t<div class=\"comment-avatar\"><img src=\"https://pic4.zhimg.com/80/v2-8e1108819305c40d20a7790945276a3f_720w.webp\" alt=\"\"></div>\n" +
                "\t\t\t\t\t<div class=\"comment-box\">\n" +
                "\t\t\t\t\t\t<div class=\"comment-head\">\n" +
                "\t\t\t\t\t\t\t<h6 class=\"comment-name\"><a href=\"http://creaticode.com/blog\">Self</a></h6>\n" +
                "\t\t\t\t\t\t\t<span>hace 0 minutos</span>\n" +
                "\t\t\t\t\t\t\t<i class=\"fa fa-xmark\"></i>\n" +
                "\t\t\t\t\t\t\t<i class=\"fa fa-reply\"></i>\n" +
                "\t\t\t\t\t\t\t<i class=\"fa fa-heart\"></i>\n" +
                "\t\t\t\t\t\t</div>\n" +
                "\t\t\t\t\t\t<div class=\"comment-content\">\n" +
                "<p class=\"touch-disappear\">" + txt +
                "                            </p>\n" +
                "                            <div class=\"mini-txt undisplay\">\n" +
                "                                <input  type=\"text\" placeholder=\"Write a comment...\">\n" +
                "                                <div class=\"mini-textarea-icons\">\n" +
                "                                    <button class=\"mini-send\" >\n" +
                "                                        <i class=\"fa-solid fa-paper-plane\"></i>\n" +
                "                                    </button>\n" +
                "                                </div>\n" +
                "                            </div>" +
                "\t\t\t\t\t\t</div>\n" +
                "\t\t\t\t\t</div>\n" +
                "\t\t\t\t</div>\n" +
                "\t\t\t</li>");
        }
    })
})


