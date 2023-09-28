/* eslint-disable */
export default {
    namespaced: true,
    state: {
      courses: [],
      currentCourse: {
        editing: false,
      },
      newCourse: {
        id: '0',
        code_kz: '',
        code_ru: '',
        code_en: '',
        ects: 0,
        module_id: '0',
        department_id: '0',
        professional_component_id: '0',
        name_kz: '',
        name_ru: '',
        name_en: '',
        lang_of_teach_kz: '',
        lang_of_teach_ru: '',
        lang_of_teach_en: '',
        control_form_kz: '',
        control_form_ru: '',
        control_form_en: '',
        lecture_hour: 0,
        seminar_hour: 0,
        lab_hour: 0,
        sro_hour: 0,
      },
      editedCourse: {
        id: '0',
        code_kz: '',
        code_ru: '',
        code_en: '',
        ects: 0,
        module_id: '0',
        department_id: '0',
        professional_component_id: '0',
        name_kz: '',
        name_ru: '',
        name_en: '',
        lang_of_teach_kz: '',
        lang_of_teach_ru: '',
        lang_of_teach_en: '',
        control_form_kz: '',
        control_form_ru: '',
        control_form_en: '',
        lecture_hour: 0,
        seminar_hour: 0,
        lab_hour: 0,
        sro_hour: 0,
      },
      pagination: {
        page: 1,
        itemsPerPage: 10,
      },
    },
    mutations: {
      SET_COURSES(state, courses) {
        state.courses = courses;
      },
      ADD_COURSE(state, course) {
        state.courses.push(course);
      },
      CLEAR_NEW_COURSE(state) {
        state.newCourse = {
          id: '0',
          code_kz: '',
          code_ru: '',
          code_en: '',
          ects: 0,
          module_id: '0',
          department_id: '0',
          professional_component_id: '0',
          name_kz: '',
          name_ru: '',
          name_en: '',
          lang_of_teach_kz: '',
          lang_of_teach_ru: '',
          lang_of_teach_en: '',
          control_form_kz: '',
          control_form_ru: '',
          control_form_en: '',
          lecture_hour: 0,
          seminar_hour: 0,
          lab_hour: 0,
          sro_hour: 0,
        };
      },
      EDIT_COURSE(state, editedCourse) {
        const index = state.courses.findIndex((c) => String(c.id) === editedCourse.id);
        if (index !== -1) {
          state.courses[index] = editedCourse;
        }
      },
      DELETE_COURSE(state, courseId) {
        state.courses = state.courses.filter((c) => c.id !== courseId);
      },
    },
    actions: {
      async fetchCourses({ commit, state }) {
        try {
          const response = await fetch(
            `http://localhost:8080/api/course?${state.pagination.page}&perPage=${state.pagination.itemsPerPage}`,
            {
              method: 'GET',
            }
          );
          if (!response.ok) {
            throw new Error('Failed to fetch courses');
          }
          const courses = await response.json();
          commit('SET_COURSES', courses);
        } catch (error) {
          console.error('Error fetching courses:', error);
        }
      },
      async addNewCourse({ commit, state }) {
        try {
          const response = await fetch('http://localhost:8080/api/course', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              'Access-Control-Allow-Origin': '*',
            },
            body: JSON.stringify(state.newCourse),
          });
  
          if (!response.ok) {
            throw new Error('Failed to add a new course');
          }
  
          const newCourse = await response.json();
          state.newCourse.id = newCourse.id;
          commit('ADD_COURSE', state.newCourse);
          commit('CLEAR_NEW_COURSE');
        } catch (error) {
          console.error('Error adding a new course:', error);
        }
      },
      async updateCourse({ commit }, editedCourse) {
        try {
          const response = await fetch(`http://localhost:8080/api/course/${editedCourse.id}`, {
            method: 'PUT',
            headers: {
              'Content-Type': 'application/json',
              'Access-Control-Allow-Origin': '*',
            },
            body: JSON.stringify(editedCourse),
          });
  
          if (!response.ok) {
            throw new Error('Failed to edit the course');
          }
  
          commit('EDIT_COURSE', editedCourse);
        } catch (error) {
          console.error('Error editing the course:', error);
        }
      },
      async deleteCourse({ commit }, courseId) {
        try {
          const response = await fetch(`http://localhost:8080/api/course/${courseId}`, {
            method: 'DELETE',
            headers: {
              'Access-Control-Allow-Origin': '*',
            },
          });
  
          if (!response.ok) {
            throw new Error('Failed to delete the course');
          }
  
          commit('DELETE_COURSE', courseId);
        } catch (error) {
          console.error('Error deleting the course:', error);
        }
      },
    },
    getters: {
      allCourses: (state) => {
        return state.courses;
      },
      totalPages: (state) => {
        return Math.ceil(state.courses.length / state.pagination.itemsPerPage);
      },
    },
  };
  