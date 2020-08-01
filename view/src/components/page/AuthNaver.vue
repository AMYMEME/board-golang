<template>
  <div class="home">
    <h3> 잠시만 기다려주세요 </h3>
  </div>
</template>
<script>
  export default {
    name: 'AuthNaver',
    mounted() {
      if (window.location.hash) {
        removeHashbang();
      }
    },
    created() {
      var params = getUrlParams()
      this.$store.dispatch('loginWithNaver', {params})
              .then(() => this.$router.push('/'))
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
