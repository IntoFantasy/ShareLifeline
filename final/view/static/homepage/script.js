/**
 * Variables
 */
const signupButton = document.getElementById('signup-button'),
    loginButton = document.getElementById('login-button'),
    userForms = document.getElementById('user_options-forms')

/**
 * Add event listener to the "Sign Up" button
 */
signupButton.addEventListener('click', () => {
  userForms.classList.remove('bounceRight')
  userForms.classList.add('bounceLeft')
}, false)

/**
 * Add event listener to the "Login" button
 */
loginButton.addEventListener('click', () => {
  userForms.classList.remove('bounceLeft')
  userForms.classList.add('bounceRight')
}, false)

function login() {
  // window.location.href='/homepage';
  // var arr = {
  //   "emailLogin":"emailLogin",
  //   "passwordLogin":"passwordLogin"
  // }
  $.ajax({
    type: "POST",
    dataType: "json",
    url: "http://127.0.0.1:9300/login" ,
    data:{
      emailLogin: $("[name=emailLogin]").val(),
      passwordLogin: $("[name=passwordLogin]").val(),
      nameSignUp: $("[name=nameSignUp]").val(),
      emailSignUp: $("[name=emailSignUp]").val(),
      passwordSignUp: $("[name=passwordSignUp]").val()
    },
    success: function(result){
      console.log(result);
      if(result.Code === 200){
        // var email=$("[name=emailLogin]").val();
        // var pwd=$("[name=passwordLogin]").val()
        // var today=new Date();
        // var tmp="email="+email+" password="+pwd+" date="+today.toGMTString();
        // window.document.cookie=tmp;
        // localStorage.setItem("email", email);
        // localStorage.setItem("password", pwd);
        // localStorage.setItem("date", today);
        // console.log(sessionStorage.getItem("password"));
        // console.log(sessionStorage.getItem("date"));
        window.location.href='http://127.0.0.1:9700/Question';
        // window.open('/homepage');
        // console.log(document.cookie)
        alert("success");
      }else if (result.Code === 300){
        alert("fail, because the email has been signed up");
      }else if (result.Code === 400){
        alert("fail, please don't fill in both sign up and login");
      } else{
        alert("fail");
      }
    }
  });
}

function signUp(){
  $.ajax({
      type: "GET",
      dataType: "json",
      url: "http://127.0.0.1:9300/Login",
      data: {
        nameSignUp: $("[name=nameSignUp]").val(),
        emailSignUp: $("[name=emailSignUp]").val(),
        passwordSignUp: $("[name=passwordSignUp]").val()
      },
      success: function(result){
        console.log(result);
        if (result.Code == 200){
          window.location.href='/homepage';
          alert("success");
        }else{
          alert("fail");
        }
      }
    }
  );
}