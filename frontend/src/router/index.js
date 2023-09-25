import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import Login from "../views/LoginView.vue";
import SignUp from "../views/SignUpView.vue";
import DashboardView from "../views/DashboardView.vue";

const routes = [
  {
    path: "/",
    name: "HomeView",
    component: HomeView,
  },
  {
    path: "/login",
    name: "login",
    component: Login,
  },
  {
    path: "/signup",
    name: "sighUp",
    component: SignUp,
  },
  {
    path: "/dashboard",
    name: "Dashboard",
    component: DashboardView, // Use your Dashboard component here
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
