<template>
  <div class="home">
    <h3> Login Page </h3>
    <b-container fluid>
      <div class="col-12">
        <b-row class="justify-content-sm-center mt-5">

        <b-card class="col-sm-6">
          <b-button v-if="this.$gAuth" pill block variant="outline-danger" @click="loginWithGoogle" >
            Login with Google</b-button>
          <AuthNaver></AuthNaver>
        </b-card>
        </b-row>
      </div>
    </b-container>
  </div>
</template>

<script>
  import AuthNaver from "../AuthNaver";

  export default {
    name: 'test',
    components: {AuthNaver},
    methods: {
      loginWithGoogle() {
        this.$gAuth.getAuthCode()
                .then(authCode => {
                  this.$store.dispatch('loginWithGoogle', {authCode})
                          .then(() => this.backRedirect())
                })
                .catch(() => {
                  alert('인증 코드를 가져오는 데에 에러가 발생했습니다: \n다시 시도해 주세요.')
                })
      },
      backRedirect() {
        let redirectURI;
        redirectURI = decodeURIComponent(getUrlParams().redirect);
        if (redirectURI === 'undefined') {
          this.$router.push('/');
        } else {
          this.$router.push(redirectURI)
        }
      },
    }
  }
  function getUrlParams() {
    var params = {};
    window.location.search.replace(/[?&]+([^=&]+)=([^&]*)/gi, function(str, key, value) { params[key] = value; });
    return params;
  }
</script>
