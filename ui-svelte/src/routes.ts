import notFound from "./views/404.svelte";
import Home from "./views/home.svelte";
import Login from "./views/user/login.svelte";
import Register from "./views/user/register.svelte";
import Protected from "./views/protected.svelte";
import { isLoggedIn } from "./services/auth.service.ts";
const routes = [
  {
    name: "/",
    component: Home
  },
  { name: "login", component: Login },
  { name: "register", component: Register },
  {
    name: "protected",
    onlyIf: { guard: isLoggedIn, redirect: "/login" },
    component: Protected
  },
  { name: "*", component: notFound }
];

export { routes };
