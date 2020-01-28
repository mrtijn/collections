<template>
  <div class="register">
    <form class="container" @submit.prevent="register">
      <h2>Register</h2>
      <div class="form-group">
        <label for="">Email* </label>
        <input type="text" v-model="form.email" required />
      </div>
      <div class="form-group">
        <label for="">Password*</label>
        <input type="password" v-model="form.password" required />
      </div>
      <div class="form-group">
        <label for="">First name</label>
        <input type="text" v-model="form.first_name" />
      </div>
      <div class="form-group">
        <label for="">Last name</label>
        <input type="text" v-model="form.last_name" />
      </div>
      <button class="btn btn__primary">Register</button>
    </form>
  </div>
</template>

<script>
import Vue from "vue";
import { Component } from "vue-property-decorator";
import { isLoggedIn, Register } from "@/services/auth.service";

@Component
export default class RegisterView extends Vue {
  form = {};
  created() {
    console.log("gO!");
    if (isLoggedIn()) {
      this.$router.push({ name: "dashboard" });
    }
  }

  async register() {
    try {
      Register(this.form);
    } catch (error) {
      console.error(error);
    }
  }
}
</script>

<style lang="scss" scoped>
.register {
  .container {
    display: flex;
    flex-direction: column;

    input {
      width: 100%;
    }
  }
}
</style>
