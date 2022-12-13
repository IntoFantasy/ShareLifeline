function search(){
    let search_content = document.getElementById("search_content").value;
    // console.log(search_content);
    let url = "/searchResult?searchContent=" + search_content;
    console.log(url);
    window.location.href=url;
    // window.location.href()

    // $.ajax({
    //     type: "POST",
    //     dataType: "json",
    //     url: "http://127.0.0.1:9300/searchResult" ,
    //     data:{
    //         SearchContent: search_content,
    //     },
    //     success: function(result){
    //         console.log(result);
    //         if(result.Code === 200){
    //
    //             window.location.href='/homepage';
    //             // window.open('/homepage');
    //             // console.log(document.cookie)
    //             // alert("success");
    //
    //         }else{
    //             alert("fail");
    //         }
    //     }
    // });
}