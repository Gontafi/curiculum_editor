<!-- eslint-disable -->
<template>
    <div class="container">
      <h2>Manage Courses</h2>
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th>ID</th>
            <th>Код KZ</th>
            <th>Код RU</th>
            <th>Код EN</th>
            <th>ECTS</th>
            <th>Модуль</th>
            <th>Департамент</th>
            <th>Профессиональный компонент</th>
            <th>Название KZ</th>
            <th>Название RU</th>
            <th>Название EN</th>
            <th>Язык обучения KZ</th>
            <th>Язык обучения RU</th>
            <th>Язык обучения EN</th>
            <th>Форма контроля KZ</th>
            <th>Форма контроля RU</th>
            <th>Форма контроля EN</th>
            <th>Лекционные часы</th>
            <th>Семинарские часы</th>
            <th>Лабораторные часы</th>
            <th>СРО часы</th>
            <th>Действия</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(course, index) in courses" :key="course.id">
              <td>
                <span>{{ course.id }}</span>
              </td>
              <td>
                <span v-if="!course.editing">{{ course.code_kz }}</span>
                <input v-else v-model="editedCourse.code_kz" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.code_ru }}</span>
                <input v-else v-model="editedCourse.code_ru" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.code_en }}</span>
                <input v-else v-model="editedCourse.code_en" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.ects }}</span>
                <input v-else v-model="editedCourse.ects" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.module_id }}</span>
                <input v-else v-model="editedCourse.module_id" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.department_id }}</span>
                <input v-else v-model="editedCourse.department_id" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.professional_component_id }}</span>
                <input v-else v-model="editedCourse.professional_component_id" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.name_kz }}</span>
                <input v-else v-model="editedCourse.name_kz" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.name_ru }}</span>
                <input v-else v-model="editedCourse.name_ru" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.name_en }}</span>
                <input v-else v-model="editedCourse.name_en" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.lang_of_teach_kz }}</span>
                <input v-else v-model="editedCourse.lang_of_teach_kz" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.lang_of_teach_ru }}</span>
                <input v-else v-model="editedCourse.lang_of_teach_ru" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.lang_of_teach_en }}</span>
                <input v-else v-model="editedCourse.lang_of_teach_en" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.control_form_kz }}</span>
                <input v-else v-model="editedCourse.control_form_kz" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.control_form_ru }}</span>
                <input v-else v-model="editedCourse.control_form_ru" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.control_form_en }}</span>
                <input v-else v-model="editedCourse.control_form_en" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.lecture_hour }}</span>
                <input v-else v-model="editedCourse.lecture_hour" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.seminar_hour }}</span>
                <input v-else v-model="editedCourse.seminar_hour" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.lab_hour }}</span>
                <input v-else v-model="editedCourse.lab_hour" />
              </td>
              <td>
                <span v-if="!course.editing">{{ course.sro_hour }}</span>
                <input v-else v-model="editedCourse.sro_hour" />
              </td>
              <td>
                <button @click="editCourse(course)" v-if="!course.editing">Изменить</button>
                <button @click="saveEditedCourse(editedCourse, course)" v-if="course.editing">Сохранить</button>
                <button @click="cancelEdit(course)" v-if="course.editing">Отмена</button>
                <button @click="deleteCourse(course.id)" v-if="!course.editing">Удалить</button>
              </td>
            </tr>
            <tr>
              <td></td>
              <td><input v-model="newCourse.code_kz" /></td>
              <td><input v-model="newCourse.code_ru" /></td>
              <td><input v-model="newCourse.code_en" /></td>
              <td><input v-model="newCourse.ects" /></td>
              <td><input v-model="newCourse.module_id" /></td>
              <td><input v-model="newCourse.department_id" /></td>
              <td><input v-model="newCourse.professional_component_id" /></td>
              <td><input v-model="newCourse.name_kz" /></td>
              <td><input v-model="newCourse.name_ru" /></td>
              <td><input v-model="newCourse.name_en" /></td>
              <td><input v-model="newCourse.lang_of_teach_kz" /></td>
              <td><input v-model="newCourse.lang_of_teach_ru" /></td>
              <td><input v-model="newCourse.lang_of_teach_en" /></td>
              <td><input v-model="newCourse.control_form_kz" /></td>
              <td><input v-model="newCourse.control_form_ru" /></td>
              <td><input v-model="newCourse.control_form_en" /></td>
              <td><input v-model="newCourse.lecture_hour" /></td>
              <td><input v-model="newCourse.seminar_hour" /></td>
              <td><input v-model="newCourse.lab_hour" /></td>
              <td><input v-model="newCourse.sro_hour" /></td>
              <td>
                <button @click="addNewCourse">Добавить</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="pagination">
        <button @click="prevPage">Previous</button>
        <span>Page {{ pagination.page }}</span>
        <button @click="nextPage">Next</button>
      </div>
    </div>
  </template>
<!-- eslint-disable -->
<script>
  import { mapState, mapActions, mapGetters } from 'vuex';

  export default {
    name: "CourseComponent",
    computed: {
      ...mapState("course", ['courses', 'newCourse', 'pagination', 'editedCourse']),
    },
    created() {
      this.fetchCourses();
    },
    methods: {
      ...mapActions("course", ['fetchCourses', 'addNewCourse', 'deleteCourse', 'updateCourse']),
      ...mapGetters("course", ['totalPages']),

      editCourse(course) {
        course.editing = true;
        this.editedCourse.id = String(course.id);
        this.editedCourse.code_kz = course.code_kz;
        this.editedCourse.code_ru = course.code_ru;
        this.editedCourse.code_en = course.code_en;
        this.editedCourse.ects = course.ects;
        this.editedCourse.module_id = course.module_id;
        this.editedCourse.department_id = course.department_id;
        this.editedCourse.professional_component_id = course.professional_component_id;
        this.editedCourse.name_kz = course.name_kz;
        this.editedCourse.name_ru = course.name_ru;
        this.editedCourse.name_en = course.name_en;
        this.editedCourse.lang_of_teach_kz = course.lang_of_teach_kz;
        this.editedCourse.lang_of_teach_ru = course.lang_of_teach_ru;
        this.editedCourse.lang_of_teach_en = course.lang_of_teach_en;
        this.editedCourse.control_form_kz = course.control_form_kz;
        this.editedCourse.control_form_ru = course.control_form_ru;
        this.editedCourse.control_form_en = course.control_form_en;
        this.editedCourse.lecture_hour = course.lecture_hour;
        this.editedCourse.seminar_hour = course.seminar_hour;
        this.editedCourse.lab_hour = course.lab_hour;
        this.editedCourse.sro_hour = course.sro_hour;
      },

      saveEditedCourse(course, currentCourse) {
        this.updateCourse(course);
        this.cancelEdit(currentCourse);
      },

      cancelEdit(course) {
        course.editing = false;
      },

      prevPage() {
        if (this.pagination.page !== 1) {
          this.pagination.page--;
          this.fetchCourses();
        }
      },

      nextPage() {
        this.pagination.page++;
        this.fetchCourses();
      },
    },
  };
</script>
<!-- eslint-disable -->
  <style>
  .container {
    max-width: 1100px;
    margin: 0 auto;
    padding: 20px;
    font-family: Arial, sans-serif;
  }
  
  h2 {
    margin-bottom: 20px;
  }
  
  .table-container {
    margin-bottom: 20px;
    overflow-x: auto;
  }
  
  table {
    width: 100%;
    border-collapse: collapse;
  }
  
  table, th, td {
    border: 1px solid #ccc;
  }
  
  th, td {
    padding: 10px;
    text-align: left;
  }
  
  th {
    background-color: #f2f2f2;
  }
  
  .pagination {
    text-align: center;
  }
  
  .pagination button {
    margin: 0 5px;
    padding: 5px 10px;
    border: 1px solid #ccc;
    background-color: #f2f2f2;
    cursor: pointer;
  }
  
  input[type="text"] {
    width: 100%;
    box-sizing: border-box;
  }
  </style>