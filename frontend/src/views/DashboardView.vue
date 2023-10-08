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
                <v-list-item-title>Log Out</v-list-item-title>
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
          </v-container>
        </v-main>
        <v-dialog v-model="logoutDialog" max-width="400">
        <v-card>
          <v-card-title class="headline">Confirm Log Out</v-card-title>
          <v-card-text>
            Are you sure you want to log out?
          </v-card-text>
          <v-card-actions>
            <v-btn color="primary" @click="confirmLogout">Yes</v-btn>
            <v-btn color="grey" text @click="cancelLogout">Cancel</v-btn>
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
},
    data() {
      return {
        items: [
          { text: "ProfessionalComponent", name: "ProfessionalComponent", icon: "mdi-view-module" },
          { text: "Components", name: "ComponentComponent", icon: "mdi-view-module" },
          { text: "Modules", name: "Module", icon: "mdi-library" },
          { text: "Cycles", name: "Cycle", icon: "mdi-refresh" },
          { text: "Departments", name: "Department", icon: "mdi-bank" },
          { text: "Courses", name: "Course", icon: "mdi-school" },
          { text: "Prerequisite", name: "Prerequisite", icon: "mdi-library"},
          { text: "Semester", name: "Semester", icon: "mdi-library" },
          { text: "Curiculum", name: "Curiculum", icon: "mdi-library" },
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
