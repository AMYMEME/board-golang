<template>
  <div class="home">
    <h3> 잠시만 기다려주세요 </h3>
  </div>
</template>
<script>
  export default {
    name: 'AuthNaver',
    data() {
      return {
        params: getUrlParams()
      }
    },
    methods: {

    },
    mounted() {
      if (window.location.hash) {
        removeHashbang();
      }
    },
    created() {
      this.$axios.post('http://localhost:8090/auth/naver', {
        access_token: this.params.access_token,
        state: this.params.state,
        token_type: this.params.token_type,
        expires_in: this.params.expires_in
      })
              .then(()=> {
                window.location.replace('/')
              })
              .catch((err) => {
                this.error = JSON.stringify(err)
                alert('에러가 발생했습니다 : '+this.error)
              })
    },
  }
  function removeHashbang() {
    window.location.replace(window.location.href.replace('#','?'))
  }
  function getUrlParams() {
    var params = {};
    window.location.search.replace(/[?&]+([^=&]+)=([^&]*)/gi, function(str, key, value) { params[key] = value; });
    return params;
  }
</script>
