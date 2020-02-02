import Vue from "vue";
import VueRouter from "vue-router";
import Dashboard from "../views/Dashboard.vue";
import Introduction from "../views/Introduction.vue";
import { isLoggedIn } from "@/services/auth.service";
// import
Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "dashboard",
    component: Dashboard,
    meta: {
      auth: true
    }
  },
  {
    path: "/introduction",
    name: "introduction",
    component: Introduction
  },
  {
    path: "/login",
    name: "login",
    component: () =>
      import(/* webpackChunkName: "login" */ "../views/user/Login.vue")
  },
  {
    path: "/register",
    name: "register",
    component: () =>
      import(/* webpackChunkName: "login" */ "../views/user/Register.vue")
  },
  {
    path: "/profile",
    name: "profile",
    meta: {
      auth: true
    },
    component: () =>
      import(/* webpackChunkName: "login" */ "../views/user/Profile.vue")
  },
  {
    path: "/movies",
    name: "movies",
    meta: {
      auth: true
    },
    component: () =>
      import(/* webpackChunkName: "login" */ "../views/movie/Index.vue")
  },
  {
    path: "/collection/:id",
    name: "collection-detail",
    meta: {
      auth: true
    },
    component: () =>
      import(/* webpackChunkName: "login" */ "../views/collection/Detail.vue")
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

router.beforeEach((to, from, next) => {
  if (to.meta.auth && !isLoggedIn()) {
    next({ name: "introduction" });
    return false;
  }

  next();
});

export default router;
