<template>
  <HeaderComponent />
  <div>
    <router-view/>
  </div>
  <FooterComponent />
</template>

<script>
import HeaderComponent from "./components/HeaderComponent.vue"
import FooterComponent from "./components/FooterComponent.vue"
import { store } from './components/store.js'

const getCookie = (name) => {
  return document.cookie.split("; ").reduce((r, v) => {
    const parts = v.split("=");
    return parts[0] === name ? decodeURIComponent(parts[1]) : r;
  },"");
}

export default {
  name: 'App',
  components: {
    HeaderComponent,
    FooterComponent
  },
  data() {
    return{
      store
    }
  },
  beforeMount() {
    // check for cookie
    let data = getCookie("_site_data");
    if (data !== ""){
      let cookieData = JSON.parse(data);

      // update store
      store.token = cookieData.token.token;
      store.user = {
        id: cookieData.id,
        first_name: cookieData.first_name,
        last_name: cookieData.last_name,
        email: cookieData.email,
      }
    }
  }
}
</script>

<style>

</style>