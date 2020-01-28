<template>
  <header>
    <router-link to="/">
      <h1>Collections</h1>
    </router-link>

    <ul>
      <li>
        <router-link to="login">Login</router-link>
      </li>
      <li>
        <router-link to="register">Register</router-link>
      </li>
      <li>
        <a href="#" @click="Logout">Logout</a>
      </li>
    </ul>
  </header>
</template>

<script lang="ts">
import { Logout } from "@/services/auth.service";
import Vue from "vue";
export default Vue.extend({
  methods: {
    Logout() {
      Logout();
    }
  }
});
</script>

<style lang="scss" scoped>
header {
  display: flex;
  width: 100%;
  justify-content: space-between;
  align-content: center;
  padding: 10px 20px;

  h1 {
    color: #000;
    font-weight: 300;
  }

  ul {
    padding: 0;
    list-style: none;
    display: flex;
    margin: 0;
    li {
      margin-right: 15px;
      display: flex;
      flex-direction: column;
      justify-content: center;

      color: #000;
    }
  }
}
</style>
