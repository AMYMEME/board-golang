<template>
  <div class="home">
    <h3> 잠시만 기다려주세요 </h3>
    <div id="naverIdLogin" class="mt-4"></div>
  </div>
</template>
<script>
  export default {
    name: 'NaverCallBack',
    data() {
      return {
        initiate: (store) => {
          const naverLogin = new naver.LoginWithNaverId(
                  {
                    clientId: `_HXiSkrZSfB9vq94TsZC`,
                    callbackUrl: `http://localhost:8080/auth/naver`,
                    callbackHandle: true,
                    loginButton: {color: "green", type: 3, height: 50}
                  }
          );
          naverLogin.init();

          window.addEventListener('load', function () {
            naverLogin.getLoginStatus(function (status) {
              if (status) {
                var name = naverLogin.user.getName();

                if( name === undefined || name == null) {
                  alert("이름은 필수정보입니다. 정보제공을 동의해주세요.");
                  naverLogin.reprompt();
                }
              } else {
                alert("AccssToken이 올바르지 않아 callback 처리에 실패하였습니다.");
                window.location.replace('/login')
              }
            });
          });

          naverLogin.getLoginStatus((status) => {
            if (status) {
              var name = naverLogin.user.getName();
              if (name===undefined || name === null) {
                return;
              }
              var result = {}
              result['name'] = name;
              result['uniqId'] = naverLogin.user.getId();
              store.dispatch('loginWithNaver', result)
                      .then(() => {
                        window.location.replace('/')
                      })
            } else {
              console.log("AccessToken이 올바르지 않습니다.");
            }
          })
        }
      }
    },
    mounted() {
      const script = document.createElement('script')
      script.src = 'https://static.nid.naver.com/js/naveridlogin_js_sdk_2.0.0.js'
      script.onload = () => {
        this.initiate(this.$store)
      }
      document.body.appendChild(script)
    },
  }
</script>
