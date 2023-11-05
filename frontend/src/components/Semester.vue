<template>
  <div>
    <h1>Учебный план</h1>
    <table>
      <thead>
        <tr>
          <th>Семестр</th>
          <th>Дата создания</th>
          <th>Дата обновления</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="semester in semesters" :key="semester.id">
          <td @click="openSemester(semester)">{{ semester.id }}</td>
          <td>{{ semester.createdAt }}</td>
          <td>{{ semester.updatedAt }}</td>
        </tr>
      </tbody>
    </table>
    <button @click="createSemester">Создать семестр</button>
    <div v-if="selectedSemester">
      <h2>Семестр {{ selectedSemester.id }}</h2>
      <button @click="addCourse">Добавить курс</button>
      <!-- Окно добавления курсов -->
      <div v-if="addingCourse">
        <!-- Здесь вставьте форму для добавления курсов -->
        <!-- Пример: -->
        <form @submit="saveCourse">
          <label>Код курса:</label>
          <input v-model="courseCode" required>
          <label>Название курса:</label>
          <input v-model="courseName" required>
          <!-- Добавьте другие поля, необходимые для вашей модели Course -->
          <button type="submit">Сохранить курс</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "SemesterComponent",
  data() {
    return {
      semesters: [], // Здесь храните список семестров из базы данных
      selectedSemester: null,
      addingCourse: false,
      courseCode: "",
      courseName: "",
      // Другие поля, необходимые для вашей модели Course
    };
  },
  methods: {
    openSemester(semester) {
      this.selectedSemester = semester;
    },
    createSemester() {
      // Создайте новый семестр в базе данных и добавьте его в this.semesters
      // Пример: this.semesters.push({ id: 1, createdAt: new Date(), updatedAt: new Date() });
    },
    addCourse() {
      this.addingCourse = true;
    },
    saveCourse() {
      // Сохраните новый курс в базе данных и добавьте его к выбранному семестру
      // Пример: this.selectedSemester.courses.push({ code: this.courseCode, name: this.courseName });
      this.addingCourse = false;
      this.courseCode = "";
      this.courseName = "";
      // Сбросьте другие поля, если необходимо
    },
  },
};
</script>
