import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import Login from "../views/LoginView.vue";
import SignUp from "../views/SignUpView.vue";
import DashboardView from "../views/DashboardView.vue";
import store from "@/store/index.js";
const routes = [
  {
    path: "/",
    name: "HomeView",
    component: HomeView,
    meta: { requiredAuth: false },
  },
  {
    path: "/login",
    name: "login",
    component: Login,
    meta: { requiredAuth: false },
  },
  {
    path: "/signup",
    name: "sighUp",
    component: SignUp,
    meta: { requiredAuth: false },
  },
  {
    path: "/dashboard",
    name: "Dashboard",
    component: DashboardView,
    meta: { requiredAuth: true },
  },
];
/* eslint-disable */
const router = createRouter({
  history: createWebHashHistory(),
  routes,
});
function isAuthenticated(){ 
  return store.getters['authComponent/isAuthenticated']
};

router.beforeEach(async (to, from, next) => {
  if (to.meta.requiredAuth) {
    if (isAuthenticated()){
      return next();
    }
    else {
      return next({ path: "/login" });
    }
  }
  return next();
});
export default router;
