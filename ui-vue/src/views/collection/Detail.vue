<template>
  <div class="container" v-if="collection">
    {{ collection.title }}
    <movie-list :movies="collection.Movies" />
  </div>
</template>

<script>
import { Component } from "vue-property-decorator";
import Vue from "vue";
import { GetCollectionById } from "@/services/collection.service";
import MovieList from "@/components/movie/List";
@Component({
  components: { MovieList }
})
export default class CollectionDetail extends Vue {
  collection = null;
  mounted() {
    this.getCollection();
  }
  async getCollection() {
    try {
      this.collection = await GetCollectionById(this.$route.params.id);
      console.log(this.collection);
    } catch (error) {
      console.error(error);
    }
  }
}
</script>

<style lang="scss" scoped></style>
