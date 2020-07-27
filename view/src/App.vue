<template>
  <!--Here we use id Selector to bind css Styling-->
  <default-layout>
    <div id="app">
      <router-view/>
    </div>
  </default-layout>
</template>

<!--Style code-->
<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }
</style>
<script>
    import DefaultLayout from "./components/layout/DefaultLayout";
    export default {
        components: {DefaultLayout}
    }
</script>