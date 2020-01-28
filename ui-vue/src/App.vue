<template>
  <div id="app">
    <Header />
    <router-view />
  </div>
</template>
<script lang="ts">
import Header from "@/components/Header.vue";
import Vue from "vue";
import { Component } from "vue-property-decorator";

@Component({
  components: { Header }
})
export default class App extends Vue {}
</script>
<style lang="scss"></style>
