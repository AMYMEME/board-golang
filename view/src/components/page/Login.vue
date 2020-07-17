<template>
  <div class="home">
    <h3> Login Page </h3>
    <b-container fluid>
      <div class="col-12">
        <b-row class="justify-content-sm-center mt-5">

        <b-card class="col-sm-6">
          <b-button pill block variant="outline-danger" @click="loginWithGoogle" >
            Login with Google</b-button>
          <b-button pill block variant="outline-success" @click="" >
            Login with Naver</b-button>
        </b-card>
        </b-row>
      </div>
    </b-container>
  </div>
</template>

<script>
  export default {
    name: 'test',
    methods: {
      loginWithGoogle() {
        this.$gAuth.getAuthCode()
                .then(authCode => {
                  //on success
                  return this.$axios.post('http://localhost:8090/auth/google', {
                    code: authCode,
                    redirect_uri: 'postmessage',
                  })
                })
                .then((res) => {
                  console.log(res.data);
                })
                .catch(err => {
                  alert('에러가 발생했습니다: '+err+'\n다시 시도해 주세요.');
                })
      }
    }
  }
  function f() {

  }
</script>
