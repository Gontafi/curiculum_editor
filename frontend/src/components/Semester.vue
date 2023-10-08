<template>
  <div class="container">
    <h2>Manage Semesters and Courses</h2>
    <table>
      <thead>
      <tr>
        <th>Semester Name</th>
        <th>Course Code</th>
        <th>Course Name</th>
        <th>Actions</th>
      </tr>
      </thead>
      <tbody>
      <!-- Iterate over semesters and courses -->
      <tr v-for="(semester, semesterIndex) in semesters" :key="semesterIndex">
        <td>{{ semester.name }}</td>
        <td><input v-model="newCourse.code" /></td>
        <td><input v-model="newCourse.name" /></td>
        <td>
          <button @click="addCourse(semesterIndex)">Add Course</button>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  name: "SemesterComponent",
  data() {
    return {
      newCourse: {
        code: '',
        name: ''
      }
    };
  },
  computed: {
    // Map your state properties here
    ...mapState("curriculum", ['semesters']),
  },
  methods: {
    // Map your actions here
    ...mapActions("curriculum", ['addCourseToSemester']),
    addCourse(semesterIndex) {
      // You should call an action to add the new course to the semester
      this.addCourseToSemester({ semesterIndex, course: this.newCourse });
      // Clear the input fields
      this.newCourse.code = '';
      this.newCourse.name = '';
    },
  },
};
</script>

<style scoped>
/* Add your component styles here */
</style>
