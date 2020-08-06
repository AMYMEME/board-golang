<template>
  <div class="home">
    <h3> Login Page </h3>
    <b-container fluid>
      <div class="col-12">
        <b-row class="justify-content-sm-center mt-5">
        <b-card class="col-sm-6">
          <AuthGoogle></AuthGoogle>
          <AuthNaver class="mt-4"></AuthNaver>
          <AuthKakao class="mt-4"></AuthKakao>
        </b-card>
        </b-row>
      </div>
    </b-container>
  </div>
</template>

<script>
  import AuthNaver from "../AuthNaver";
  import AuthKakao from "../AuthKakao";
  import AuthGoogle from "../AuthGoogle";

  export default {
    name: 'test',
    components: {AuthGoogle, AuthKakao, AuthNaver},
    methods: {
    },
    mounted() {
      if (this.$store.state.accessToken) {
        alert('로그인이 완료되었습니다')
        location.replace('/');
      }
    }
  }
</script>
