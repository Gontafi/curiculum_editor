/* eslint-disable */
import { createStore } from "vuex";
import courseModule from "./module/courseModule";
import courseComponent from "./module/courseComponent";
export default createStore({
  modules: {
    courseComponent,
    courseModule,
  },
});
console.log(this.modules);