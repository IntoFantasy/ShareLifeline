function save(){
    var editor = $('#editor').prop("outerHTML");
    console.log(editor);
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "http://127.0.0.1:9300/save" ,
        data:{
            editor: editor
        },
        success: function(result){
            console.log(result);
            if(result.Code === 200){
                alert("success");

            }else{
                alert("fail");
            }
        }
    });
}

function release(){
    var editor = $('#editor').prop("outerHTML");
    console.log(editor);
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "http://127.0.0.1:9300/save" ,
        data:{
            editor: editor
        },
        success: function(result){
            console.log(result);
            if(result.Code === 200){
                alert("success");

            }else{
                alert("fail");
            }
        }
    });
}