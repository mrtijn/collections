<template>
  <div class="login">
    <div class="login__container">
      <h2>Login</h2>
      <form @submit.prevent="submitForm">
        <input
          type="text"
          id="username"
          required
          class="input"
          v-model="username"
          placeholder="Username"
        />
        <input
          type="password"
          class="input"
          placeholder="password"
          required
          v-model="password"
        />
        <button class="btn btn__primary btn--arrow">Login</button>
      </form>

      <div v-if="errorMessage" class="error">{{ errorMessage }}</div>
    </div>
  </div>
</template>
<script lang="ts">
import { Login, isLoggedIn } from "@/services/auth.service";
import Vue from "vue";
import { Component } from "vue-property-decorator";

@Component
export default class LoginView extends Vue {
  username = "";
  password = "";
  errorMessage = "";
  created() {
    if (isLoggedIn()) {
      this.$router.push({ name: "dashboard" });
    }
  }
  async submitForm() {
    try {
      await Login(this.username, this.password);
      this.errorMessage = "";
      console.log(this);
      this.$router.push("/protected");
    } catch (error) {
      console.log(error);
      this.errorMessage = error.message;
    }
  }
}
</script>

<style lang="scss">
.login {
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
    z-index: 1;

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
