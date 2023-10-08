/* eslint-disable */
import { createStore } from "vuex";
import courseModule from "./module/courseModule";
import courseProfessionalComponent from "./module/professtionalComponent";
import authComponent from "./module/auth";
import courseComponent from "./module/component";
import department from "./module/department"
import course from "./module/course";
import cycles from "./module/cycles";
import prerequisite from "@/store/module/prerequisite";
export default createStore({
  modules: {
    courseModule,
    courseComponent,
    courseProfessionalComponent,
    course,
    department,
    authComponent,
    cycles,
    prerequisite,
  },
});