<template>
  <header>
    <router-link to="/">
      <h1>Collections</h1>
    </router-link>

    <ul>
      <template v-if="!isLoggedIn">
        <li>
          <router-link to="/login">Login</router-link>
        </li>
        <li>
          <router-link to="/register">Register</router-link>
        </li>
      </template>

      <template v-if="isLoggedIn">
        <li>
          <router-link to="/movies">Movies</router-link>
        </li>
        <li>
          <router-link to="/profile">Profile</router-link>
        </li>
        <li>
          <a href="#" @click="Logout">Logout</a>
        </li>
      </template>
    </ul>
  </header>
</template>

<script lang="ts">
import { Logout, isLoggedIn } from "@/services/auth.service";
import { Component } from "vue-property-decorator";
import Vue from "vue";

@Component
export default class Header extends Vue {
  isLoggedIn = isLoggedIn();
  Logout() {
    Logout();
  }
}
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
