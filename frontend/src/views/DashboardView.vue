<!-- eslint-disable -->
<template>
    <div>
      <v-app>
        <!-- Navigation Drawer -->
        <v-navigation-drawer v-bind:width="230" app>
          <v-list dense>
            <v-list-item v-for="(item, index) in items" :key="index" @click="changeTab(index)">
              <v-list-item-action>
                <v-icon>{{ item.icon }}</v-icon>
              </v-list-item-action>
              <v-list-item-subtitle>
                <v-list-item-title class="item-text">{{ item.text }}</v-list-item-title>
              </v-list-item-subtitle>
            </v-list-item>
          </v-list>
          <v-list-item @click="logout">
              <v-list-item-action>
                <v-icon>mdi-logout</v-icon>
              </v-list-item-action>
              <v-list-item-subtitle>
                <v-list-item-title>Выйти с аккаунта</v-list-item-title>
              </v-list-item-subtitle>
            </v-list-item>
        </v-navigation-drawer>
  
        <!-- Main Content -->
        <v-app-bar app color="#3c66aa">
          <v-toolbar-title>{{ currentTab.text }}</v-toolbar-title>
        </v-app-bar>
  
        <v-main>
          <!-- Content for each tab goes here -->
          <v-container>
            <v-row v-if="currentTab.name === 'ProfessionalComponent'"> <!-- Add similar blocks for other models -->
                <ProfessionalComponent></ProfessionalComponent>
            </v-row>
            <v-row v-if="currentTab.name === 'ComponentComponent'"> <!-- Add similar blocks for other models -->
                <ComponentComponent></ComponentComponent>
            </v-row>
            <v-row v-if="currentTab.name === 'Department'"> <!-- Add similar blocks for other models -->
                <Departments></Departments>
            </v-row>
            <v-row v-if="currentTab.name === 'Course'"> <!-- Add similar blocks for other models -->
                <Course></Course>
            </v-row>
            <v-row v-if="currentTab.name === 'Cycle'"> <!-- Add similar blocks for other models -->
                <Cycles></Cycles>
            </v-row>
            <v-row v-if="currentTab.name === 'Module'"> <!-- Add similar blocks for other models -->
                <ModuleComponent></ModuleComponent>
            </v-row>
            <v-row v-if="currentTab.name === 'Prerequisite'"> <!-- Add similar blocks for other models -->
              <PrerequisiteComponent></PrerequisiteComponent>
            </v-row>
            <v-row v-if="currentTab.name === 'Semester'"> <!-- Add similar blocks for other models -->
              <semester></semester>
            </v-row>
          </v-container>
        </v-main>
        <v-dialog v-model="logoutDialog" max-width="400">
        <v-card>
          <v-card-title class="headline">Выйти с аккаунта</v-card-title>
          <v-card-text>
            Вы точно хотите выйти?
          </v-card-text>
          <v-card-actions>
            <v-btn color="primary" @click="confirmLogout">Да</v-btn>
            <v-btn color="grey" text @click="cancelLogout">Отмена</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      </v-app>
    </div>
  </template>
<!-- eslint-disable -->
<script>
import ComponentComponent from "@/components/ComponentComponent.vue"
import ProfessionalComponent from "@/components/ProfessionalComponent.vue";
import Course from "@/components/Course.vue";
import Departments from "@/components/Departments.vue";
import ModuleComponent from "@/components/ModuleComponent.vue";
import Cycles from "@/components/Cycles.vue";
import PrerequisiteComponent from "@/components/PrerequisiteComponent.vue";
import semester from "@/components/Semester.vue";
import store from "@/store/index.js";
  export default {
    name: "DashboardView",
    components: {
      ProfessionalComponent,
    ComponentComponent,
    ModuleComponent,
    Departments,
    Cycles,
    Course,
      PrerequisiteComponent,
      semester,
},
    data() {
      return {
        items: [
          { text: "Профессиональные Компоненты", name: "ProfessionalComponent", icon: "mdi-view-module" },
          { text: "Компоненты", name: "ComponentComponent", icon: "mdi-view-module" },
          { text: "Модули", name: "Module", icon: "mdi-library" },
          { text: "Циклы", name: "Cycle", icon: "mdi-refresh" },
          { text: "Кафедры", name: "Department", icon: "mdi-bank" },
          { text: "Курсы, Дисципилины", name: "Course", icon: "mdi-school" },
          { text: "Пререквезиты", name: "Prerequisite", icon: "mdi-library"},
          { text: "Семестр", name: "Semester", icon: "mdi-library" },
          { text: "Учебный план", name: "Curiculum", icon: "mdi-library" },
        ],
        currentTab: null,
        logoutDialog: false,
      };
    },
    created() {
      // Set the initial tab when the component is created
      this.currentTab = this.items[0];
    },
    methods: {
      changeTab(index) {
        this.currentTab = this.items[index];
      },
      logout() {
      this.logoutDialog = true;
    },

    // Add a method to confirm log out
    confirmLogout() {
      return this.$store.dispatch('authComponent/logout')
    },

    // Add a method to cancel log out
    cancelLogout() {
      this.logoutDialog = false;
    },
    },
  };
  </script>
<!-- eslint-disable -->
<style>
.item-title {
  background-color: #2c3e50;
}
.item-text {
    color:black;
}
</style>
