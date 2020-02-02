<template>
  <ul class="movie__list" v-if="movies && movies.length">
    <li class="movie-list__item" v-for="movie of movies" :key="movie.id">
      <figure>
        <img :src="movie.poster_path" v-if="movie.poster_path" />
        <img src="https://via.placeholder.com/250x360" v-else />
        <figcaption>
          <h4>{{ movie.Title }}</h4>
          <p>{{ movie.overview }}</p>
        </figcaption>
      </figure>
    </li>
  </ul>
  <div v-else-if="movies && !movies.length">No movies found.</div>
</template>

<script>
import { Component } from "vue-property-decorator";
import Vue from "vue";

@Component({
  props: ["movies"]
})
export default class MovieList {}
</script>

<style lang="scss" scoped>
.movie-list {
  &__list {
    list-style: none;
    padding: 0;
  }
  &__item figure {
    display: flex;
    margin: 0;
    margin-bottom: 20px;
    img {
      width: 250px;
      margin-right: 15px;
    }
  }
}
</style>
