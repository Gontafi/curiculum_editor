/* eslint-disable */
import { createStore } from "vuex";
import courseModule from "./module/courseModule";
import courseComponent from "./module/courseComponent";
import auth from "./module/auth";
export default createStore({
  modules: {
    courseModule,
    courseComponent,
    auth,
  },
});