<template>
  <div id="app">
    <page-header></page-header>
    <register></register>
    <login></login>

    <div id="content" class="container">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import PageHeader from './components/PageHeader.vue';
import Register from './components/account/Register.vue';
import Login from './components/account/Login.vue';

export default {
  name: 'App',
  components: {
    PageHeader,
    Register,
    Login,
  },
};
</script>


<style lang="scss">
@import url('https://fonts.googleapis.com/css?family=Quicksand');

.is-disable {
  pointer-events: none;
  opacity: 0.65;
}

@font-face {
    font-family: 'oddzy';
    src: url('./assets/fonts/oddzy.eot');
    src: url('./assets/fonts/oddzy.eot?#iefix') format('embedded-opentype'),
         url('./assets/fonts/oddzy.woff') format('woff'),
         url('./assets/fonts/oddzy.ttf') format('truetype'),
         url('./assets/fonts/oddzy.svg#oddzy') format('svg');
    font-weight: normal;
    font-style: normal;
}
[class*='icon-']:before{
  display: inline-block;
  font-family: 'oddzy';
  font-style: normal;
  font-weight: normal;
  line-height: 1;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale
}
.icon-greyhound:before{content:'\0041';}
.icon-harness:before{content:'\0042';}
.icon-thoroughbred:before{content:'\0043';}

</style>
