<template>
  <div class="profile container">
    <div class="profile__collections">
      <h5>My Collections</h5>
      <ul>
        <li v-for="collection in collections" :key="collection.ID">
          <router-link
            :to="{ name: 'collection-detail', params: { id: collection.ID } }"
          >
            {{ collection.title }}
          </router-link>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import { Component } from "vue-property-decorator";
import Vue from "vue";
import { GetMyCollections } from "@/services/collection.service";
@Component
export default class Profile extends Vue {
  collections = [];
  mounted() {
    this.getMyCollections();
  }
  async getMyCollections() {
    try {
      this.collections = await GetMyCollections();
    } catch (error) {
      console.error(error);
    }
  }
}
</script>

<style lang="scss" scoped></style>
