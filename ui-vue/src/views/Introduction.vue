<template>
  <div class="introduction">
    <div class="introduction__container">
      <h2>Welcome</h2>
      <router-link class="btn btn__primary " to="/login">
        Login
      </router-link>
      <router-link class="btn btn__primary " to="/register">
        Register
      </router-link>
    </div>
  </div>
</template>
<script lang="ts">
import { isLoggedIn } from "@/services/auth.service";
import Vue from "vue";
import { Component } from "vue-property-decorator";

@Component
export default class IntroductionView extends Vue {
  created() {
    if (isLoggedIn()) {
      this.$router.push({ name: "dashboard" });
    }
  }
}
</script>

<style lang="scss">
.introduction {
  display: flex;
  flex-direction: row;
  background-color: lighten(#000, 5%);
  width: 100%;
  align-items: space-between;
  justify-content: center;
  flex: 1;
  position: relative;

  &:before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    height: 100%;
    width: 100%;
    background-image: url("https://source.unsplash.com/collection/1416491/1600x900");
    background-size: cover;
    opacity: 0.4;
    z-index: 0;
  }
  &__container {
    background-color: #fff;
    min-width: 400px;
    padding: 15px;
    align-self: center;
    display: flex;
    flex-direction: column;
    z-index: 1;

    a {
      margin-bottom: 5px;
    }

    h2 {
      border-bottom: 1px solid #dadada;
      margin: 0 0 15px 0;
      padding-bottom: 15px;
    }

    form {
      display: flex;
      flex-direction: column;
    }
  }
}
</style>
