/* eslint-disable */
export default {
    namespaced: true,
    state: {
      departments: [],
      currentDepartment: {
        editing: false,
      },
      newDepartment: {
        id: '0',
        description_kz: '',
        description_ru: '',
        description_en: '',
      },
      editedDepartment: {
        id: '0',
        description_kz: '',
        description_ru: '',
        description_en: '',
      },
      pagination: {
        page: 1,
        itemsPerPage: 10,
      },
    },
    mutations: {
      SET_DEPARTMENTS(state, departments) {
        state.departments = departments;
      },
      ADD_DEPARTMENT(state, department) {
        state.departments.push(department);
      },
      CLEAR_NEW_DEPARTMENT(state) {
        state.newDepartment = {
          id: '0',
          description_kz: '',
          description_ru: '',
          description_en: '',
        };
      },
      EDIT_DEPARTMENT(state, editedDepartment) {
        const index = state.departments.findIndex((d) => String(d.id) === editedDepartment.id);
        if (index !== -1) {
          state.departments[index] = editedDepartment;
        }
      },
      DELETE_DEPARTMENT(state, departmentId) {
        state.departments = state.departments.filter((d) => d.id !== departmentId);
      },
    },
    actions: {
      async fetchDepartments({ commit, state }) {
        try {
          const response = await fetch(
            `http://localhost:8080/api/department?${state.pagination.page}&perPage=${state.pagination.itemsPerPage}`,
            {
              method: 'GET',
            }
          );
          if (!response.ok) {
            throw new Error('Failed to fetch departments');
          }
          const departments = await response.json();
          commit('SET_DEPARTMENTS', departments);
        } catch (error) {
          console.error('Error fetching departments:', error);
        }
      },
      async addNewDepartment({ commit, state }) {
        try {
          const response = await fetch('http://localhost:8080/api/department', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              'Access-Control-Allow-Origin': '*',
            },
            body: JSON.stringify(state.newDepartment),
          });
  
          if (!response.ok) {
            throw new Error('Failed to add a new department');
          }
  
          const newDepartment = await response.json();
          state.newDepartment.id = newDepartment.id;
          commit('ADD_DEPARTMENT', state.newDepartment);
          commit('CLEAR_NEW_DEPARTMENT');
        } catch (error) {
          console.error('Error adding a new department:', error);
        }
      },
      async updateDepartment({ commit }, editedDepartment) {
        try {
          const response = await fetch(`http://localhost:8080/api/department/${editedDepartment.id}`, {
            method: 'PUT',
            headers: {
              'Content-Type': 'application/json',
              'Access-Control-Allow-Origin': '*',
            },
            body: JSON.stringify(editedDepartment),
          });
  
          if (!response.ok) {
            throw new Error('Failed to edit the department');
          }
  
          commit('EDIT_DEPARTMENT', editedDepartment);
        } catch (error) {
          console.error('Error editing the department:', error);
        }
      },
      async deleteDepartment({ commit }, departmentId) {
        try {
          const response = await fetch(`http://localhost:8080/api/department/${departmentId}`, {
            method: 'DELETE',
            headers: {
              'Access-Control-Allow-Origin': '*',
            },
          });
  
          if (!response.ok) {
            throw new Error('Failed to delete the department');
          }
  
          commit('DELETE_DEPARTMENT', departmentId);
        } catch (error) {
          console.error('Error deleting the department:', error);
        }
      },
    },
    getters: {
      allDepartments: (state) => {
        return state.departments;
      },
      totalPages: (state) => {
        return Math.ceil(state.departments.length / state.pagination.itemsPerPage);
      },
    },
  };
  