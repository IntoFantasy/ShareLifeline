var colorPalette = ['000000', 'FF9966', '6699FF', '99FF66', 'CC0000', '00CC00', '0000CC', '333333', '0066FF', 'FFFFFF'];
var forePalette = $('.fore-palette');
var backPalette = $('.back-palette');
var data_img;
var help;
var max_height=600;
var max_width=600;

for (var i = 0; i < colorPalette.length; i++) {
  forePalette.append('<a href="#" data-command="forecolor" data-value="' + '#' + colorPalette[i] + '" style="background-color:' + '#' + colorPalette[i] + ';" class="palette-item"></a>');
  backPalette.append('<a href="#" data-command="backcolor" data-value="' + '#' + colorPalette[i] + '" style="background-color:' + '#' + colorPalette[i] + ';" class="palette-item"></a>');
}
var img_upload = document.getElementById("img_upload");
img_upload.addEventListener('change', readFile, false);

$('.toolbar a').click(function(e) {
  var command = $(this).data('command');
  if (command == 'h1' || command == 'h2' || command == 'p') {
    document.execCommand('formatBlock', false, command);
  }
  if (command == 'forecolor' || command == 'backcolor') {
    document.execCommand($(this).data('command'), false, $(this).data('value'));
  }
    if (command == 'createlink') {
  url = prompt('Enter the link here: ','http:\/\/'); document.execCommand($(this).data('command'), false, url);
  }
    if (command == 'insertimage'){
      help = $(this).data('command');
      console.log(this);
    }
  else document.execCommand($(this).data('command'), false, null);
});

function readFile() {

  var file = this.files[0];//这里是抓取到上传的对象。

  if(!/image\/\w+/.test(file.type)) {

    alert("请确保文件为图像类型");

    return false;

  }

  var reader = new FileReader();

  reader.readAsDataURL(file);
  reader.onload = function() {
    console.log(this.result);
    data_img = this.result;
    document.execCommand(help, false, data_img);


  }

}