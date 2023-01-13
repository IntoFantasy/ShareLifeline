
function change(){
    window.location.href='/homepage';
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "http://127.0.0.1:9300/modify",
        data: {
            name: $("[name=name]").val(),
            password: $("[name=password]").val(),
            email: $("[name=email]").val(),
            address: $("[name=address]").val(),
            birthday: $("[name=birthday]").val(),
            introduction: $("[name=introduction]").val(),
            phone: $("[name=phone]").val(),
            sex: $("[name=sex]").val()
        },
        success: function(result){
            if (result.Code==200){
                window.location.href='/homepage';
                alert("success");
            }else{
                alert("fail");
            }
        }
    });
}