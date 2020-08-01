<template>
  <div class="home">
    <h3> Login Page </h3>
    <b-container fluid>
      <div class="col-12">
        <b-row class="justify-content-sm-center mt-5">

        <b-card class="col-sm-6">
          <b-button v-if="this.$gAuth" pill block variant="outline-danger" @click="loginWithGoogle" >
            Login with Google</b-button>
          <div id="naverIdLogin" class="mt-4"></div>
        </b-card>
        </b-row>
      </div>
    </b-container>
  </div>
</template>

<script>
  export default {
    name: 'test',
    data() {
      return {
        initiate: () => {
          const naverLogin = new naver.LoginWithNaverId(
                  {
                    clientId: `_HXiSkrZSfB9vq94TsZC`,
                    callbackUrl: `http://localhost:8080/auth/naver`,
                    isPopup: false,
                    callbackHandle: true,
                    loginButton: {color: "green", type: 3, height: 50}
                  }
          );
          naverLogin.init();

          window.addEventListener('load', function () {
            naverLogin.getLoginStatus(function (status) {
              if (status) {
                var name = naverLogin.user.getNickName();

                if (name === undefined || name === null) {
                  alert("이름은 필수정보입니다. 정보제공을 동의해주세요.");
                  naverLogin.reprompt();
                }
              } else {
                console.log("callback 처리에 실패하였습니다.");
              }
            });
          });
        }
      }
    },
    methods: {
      loginWithGoogle() {
        this.$gAuth.getAuthCode()
                .then(authCode => {
                  this.$store.dispatch('loginWithGoogle', {authCode})
                          .then(() => this.redirect())
                })
                .catch(() => {
                  alert('인증 코드를 가져오는 데에 에러가 발생했습니다: \n다시 시도해 주세요.')
                })
      },
      redirect() {
        let redirectURI;
        redirectURI = decodeURIComponent(getUrlParams().redirect);
        if (redirectURI === 'undefined') {
          this.$router.push('/');
        } else {
          this.$router.push(redirectURI)
        }
      },
    },
    mounted() {
      const script = document.createElement('script')
      script.src = 'https://static.nid.naver.com/js/naveridlogin_js_sdk_2.0.0.js'
      script.onload = () => this.initiate(this)
      document.body.appendChild(script)
    },
  }
  function getUrlParams() {
    var params = {};
    window.location.search.replace(/[?&]+([^=&]+)=([^&]*)/gi, function(str, key, value) { params[key] = value; });
    return params;
  }
</script>
