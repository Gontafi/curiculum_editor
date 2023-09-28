/* eslint-disable */
export default {
    namespaced: true,
    state: {
      semesterCourses: [],
      currentSemesterCourse: {
        editing: false,
      },
      newSemesterCourse: {
        id: '0',
        semester_id: '0',
        course_id: '0',
      },
      editedSemesterCourse: {
        id: '0',
        semester_id: '0',
        course_id: '0',
      },
      pagination: {
        page: 1,
        itemsPerPage: 10,
      },
    },
    mutations: {
      SET_SEMESTER_COURSES(state, semesterCourses) {
        state.semesterCourses = semesterCourses;
      },
      ADD_SEMESTER_COURSE(state, semesterCourse) {
        state.semesterCourses.push(semesterCourse);
      },
      CLEAR_NEW_SEMESTER_COURSE(state) {
        state.newSemesterCourse = {
          id: '0',
          semester_id: '0',
          course_id: '0',
        };
      },
      EDIT_SEMESTER_COURSE(state, editedSemesterCourse) {
        const index = state.semesterCourses.findIndex((sc) => String(sc.id) === editedSemesterCourse.id);
        if (index !== -1) {
          state.semesterCourses[index] = editedSemesterCourse;
        }
      },
      DELETE_SEMESTER_COURSE(state, semesterCourseId) {
        state.semesterCourses = state.semesterCourses.filter((sc) => sc.id !== semesterCourseId);
      },
    },
    actions: {
      async fetchSemesterCourses({ commit, state }) {
        try {
          const response = await fetch(
            `http://localhost:8080/api/semester-course?${state.pagination.page}&perPage=${state.pagination.itemsPerPage}`,
            {
              method: 'GET',
            }
          );
          if (!response.ok) {
            throw new Error('Failed to fetch semester courses');
          }
          const semesterCourses = await response.json();
          commit('SET_SEMESTER_COURSES', semesterCourses);
        } catch (error) {
          console.error('Error fetching semester courses:', error);
        }
      },
      async addNewSemesterCourse({ commit, state }) {
        try {
          const response = await fetch('http://localhost:8080/api/semester-course', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              'Access-Control-Allow-Origin': '*',
            },
            body: JSON.stringify(state.newSemesterCourse),
          });
  
          if (!response.ok) {
            throw new Error('Failed to add a new semester course');
          }
  
          const newSemesterCourse = await response.json();
          state.newSemesterCourse.id = newSemesterCourse.id;
          commit('ADD_SEMESTER_COURSE', state.newSemesterCourse);
          commit('CLEAR_NEW_SEMESTER_COURSE');
        } catch (error) {
          console.error('Error adding a new semester course:', error);
        }
      },
      async updateSemesterCourse({ commit }, editedSemesterCourse) {
        try {
          const response = await fetch(`http://localhost:8080/api/semester-course/${editedSemesterCourse.id}`, {
            method: 'PUT',
            headers: {
              'Content-Type': 'application/json',
              'Access-Control-Allow-Origin': '*',
            },
            body: JSON.stringify(editedSemesterCourse),
          });
  
          if (!response.ok) {
            throw new Error('Failed to edit the semester course');
          }
  
          commit('EDIT_SEMESTER_COURSE', editedSemesterCourse);
        } catch (error) {
          console.error('Error editing the semester course:', error);
        }
      },
      async deleteSemesterCourse({ commit }, semesterCourseId) {
        try {
          const response = await fetch(`http://localhost:8080/api/semester-course/${semesterCourseId}`, {
            method: 'DELETE',
            headers: {
              'Access-Control-Allow-Origin': '*',
            },
          });
  
          if (!response.ok) {
            throw new Error('Failed to delete the semester course');
          }
  
          commit('DELETE_SEMESTER_COURSE', semesterCourseId);
        } catch (error) {
          console.error('Error deleting the semester course:', error);
        }
      },
    },
    getters: {
      allSemesterCourses: (state) => {
        return state.semesterCourses;
      },
      totalPages: (state) => {
        return Math.ceil(state.semesterCourses.length / state.pagination.itemsPerPage);
      },
    },
  };
  