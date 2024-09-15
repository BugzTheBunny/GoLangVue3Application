<template>
  <HeaderComponent />
  <BodyComponent />
  <FooterComponent />
</template>

<script>
import HeaderComponent from "./components/HeaderComponent.vue"
import BodyComponent from "./components/BodyComponent.vue"
import FooterComponent from "./components/FooterComponent.vue"


export default {
  name: 'App',
  components: {
    HeaderComponent,
    BodyComponent,
    FooterComponent
  }
}
</script>

<style>

</style>