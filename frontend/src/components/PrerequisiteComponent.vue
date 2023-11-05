<template>
  <div class="container">
    <h2>Управление Пререквезитами</h2>
    <div class="table-container flex-row">
      <!-- Block 1: All Courses -->
      <div class="block">
        <h3>Все Курсы</h3>
        <table>
          <thead>
          <tr>
            <th>Курсы</th>
            <th>Действия</th>
          </tr>
          </thead>
          <tbody>
          <!-- Iterate over all courses -->
          <tr v-for="(course, index) in courses" :key="index">
            <td>{{ course.name_kz + "\n" + course.name_ru + "\n" + course.name_en }}</td>
            <td>
              <button @click="selectCourse(course)">Выбрать</button>
            </td>
          </tr>
          </tbody>
        </table>
        <nav aria-label="Page navigation">
          <ul class="pagination">
            <li class="page-item">
              <a class="page-link" @click="changePage(currentPage - 1)" aria-label="Previous">
                <span aria-hidden="true">&laquo;</span>
              </a>
            </li>
            <li class="page-item">
              <span aria-hidden="false">{{currentPage}}</span>
            </li>
            <li class="page-item">
              <a class="page-link" @click="changePage(currentPage + 1)" aria-label="Next">
                <span aria-hidden="true">&raquo;</span>
              </a>
            </li>
          </ul>
        </nav>
      </div>

      <!-- Block 2: Available Prerequisite Courses -->
      <div class="block">
        <h3>Доступные Пререквезиты Куров</h3>
        <table>
          <thead>
          <tr>
            <th>Курсы</th>
            <th>Действия</th>
          </tr>
          </thead>
          <tbody>
          <!-- Iterate over available prerequisite courses -->
          <tr v-for="(course, index) in availablePrerequisiteCourses" :key="index">
            <td>{{ course.name_kz + "\n" + course.name_ru + "\n" + course.name_en }}</td>
            <td>
              <button @click="addPrerequisited(course)">Добавить как пререквезит</button>
            </td>
          </tr>
          </tbody>
        </table>
      </div>

      <!-- Block 3: Selected Course Prerequisites -->
      <div class="block">
        <h3>Пререквезиты выбранного курса</h3>
        <table>
          <thead>
          <tr>
            <th>Курсы</th>
            <th>Действия</th>
          </tr>
          </thead>
          <tbody>
          <!-- Display prerequisites for the selected course -->
          <tr v-for="(prerequisite, index) in selectedCoursePrerequisites" :key="index">
            <td>{{ prerequisite.course.name_kz + "\n" + prerequisite.course.name_ru + "\n" + prerequisite.course.name_en }}</td>
            <td>
              <button @click="removePrerequisite(prerequisite)">Remove Prerequisite</button>
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import {mapState, mapActions, mapGetters} from 'vuex';

export default {
  name: "PrerequisiteComponent",
  data() {
    return {
      selectedCourse: null, // To store the selected course
      currentPage: 1, // Current page for pagination
      pageSize: 10, // Number of items per page
    };
  },
  created() {
    // Fetch courses and prerequisites when the component is created
    this.fetchCourses();
    this.fetchPrerequisites();
  },
  computed: {
    // Map your state properties here
    ...mapState("course", ['courses', 'pagination']),
    ...mapState("prerequisite", ['prerequisites']),
    // Calculate available prerequisite courses (courses not selected as prerequisites for the current course)
    availablePrerequisiteCourses() {
      if (this.selectedCourse) {
        return this.courses.filter(course => {
          return !this.prerequisites.some(prerequisite => prerequisite.courseId === this.selectedCourse.id && prerequisite.prerequisiteId === course.id);
        });
      }
      return [];
    },

    // Calculate selected course prerequisites
    selectedCoursePrerequisites() {
      if (this.selectedCourse) {
        return this.prerequisites
            .filter(prerequisite => prerequisite.course_id === this.selectedCourse.id)
            // .slice((this.currentPage - 1) * this.pageSize, this.currentPage * this.pageSize);
      }
      return [];
    },

  },
  methods: {
    // Map your actions here
    ...mapActions("course", ['fetchCourses']),
    ...mapGetters("course", ['totalPages']),
    ...mapActions("prerequisite", ['fetchPrerequisites', 'addPrerequisite', 'deletePrerequisite']),

    selectCourse(course) {
      this.selectedCourse = course;
      this.currentPage = 1; // Reset the current page when a course is selected
    },

    addPrerequisited(to_course) {
      if (this.selectedCourse) {
        // Call an action to add the prerequisite for the selected course
        this.addPrerequisite({
          course_id: this.selectedCourse.id,
          prerequisite_id: to_course.id,
        });
      }
    },

    removePrerequisite(prerequisite) {
      if (this.selectedCourse) {
        // Call an action to remove the prerequisite for the selected course
        this.deletePrerequisite(prerequisite.id);
      }
    },

    changePage(page) {
      if (page !== 0 ){
        this.currentPage = page;
        this.pagination.page = page;
        console.log( this.pagination.page);
        this.fetchCourses();
      }
    },
  },
};
</script>

<style scoped>
/* Add your component styles here */
.table-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  gap: 20px; /* Adjust the gap as needed */
}

.block {
  flex: 1;
  max-width: calc(33.33% - 20px); /* Adjust the width as needed */
}
</style>