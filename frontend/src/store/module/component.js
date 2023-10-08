/* eslint-disable */
export default {
  namespaced: true,
  name:"courseComponent",
  state: {
    filteredComponents: [],
    newComponent: {
      id: '0',
      prof_id: 0,
      code_kz: '',
      code_ru: '',
      code_en: '',
      description_kz: '',
      description_ru: '',
      description_en: '',
      order: '0',
    },
    editedComponent: {
      id: '0',
      prof_id: 0,
      code_kz: '',
      code_ru: '',
      code_en: '',
      description_kz: '',
      description_ru: '',
      description_en: '',
      order: '0',
    },
    pagination: {
      page: 1,
      itemsPerPage: 10, // You can adjust the number of items per page
    },
  },
  mutations: {
    SET_FILTERED_COMPONENTS(state, filteredComponents) {
      state.filteredComponents = filteredComponents;
    },
    ADD_COMPONENT(state, component) {
      state.filteredComponents.push(component);
    },
    CLEAR_NEW_COMPONENT(state) {
      state.newComponent = {
        id: '0',
        prof_id: 0,
        code_kz: '',
        code_ru: '',
        code_en: '',
        description_kz: '',
        description_ru: '',
        description_en: '',
        order: '0',
      };
    },
    EDIT_COMPONENT(state, editedComponent) {
      const index = state.filteredComponents.findIndex((c) => String(c.id) === editedComponent.id);
      if (index !== -1) {
        state.filteredComponents[index] = editedComponent;
      }
    },
    DELETE_COMPONENT(state, componentId) {
      state.filteredComponents = state.filteredComponents.filter((c) => c.id !== componentId);
    },
  },
  actions: {
    async fetchComponents({ commit, state }) {
      try {
        const response = await fetch(`http://localhost:8080/api/component?page=${state.pagination.page}&perPage=${state.pagination.itemsPerPage}`, {
          method: "GET",
        });
        if (!response.ok) {
          throw new Error('Failed to fetch components');
        }
        const components = await response.json();
        commit('SET_FILTERED_COMPONENTS', components);
      } catch (error) {
        console.error('Error fetching components:', error);
      }
    },
    async addNewComponent({ commit, state }) {
      try {
        state.newComponent.prof_id = parseInt(state.newComponent.prof_id)
        const response = await fetch('http://localhost:8080/api/component', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": "*",
          },
          body: JSON.stringify(state.newComponent),
        });

        if (!response.ok) {
          throw new Error('Failed to add a new component');
        }

        const newComponent = await response.json();
        state.newComponent.id = newComponent.id;
        commit('ADD_COMPONENT', state.newComponent);
        commit('CLEAR_NEW_COMPONENT');
      } catch (error) {
        console.error('Error adding a new component:', error);
      }
    },
    async updateComponent({ commit }, editedComponent) {
      try {
        editedComponent.prof_id = parseInt(editedComponent.prof_id)
        const response = await fetch(`http://localhost:8080/api/component/${editedComponent.id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": "*",
          },
          body: JSON.stringify(editedComponent),

        });

        if (!response.ok) {
          throw new Error('Failed to edit the component');
        }

        commit('EDIT_COMPONENT', editedComponent);
      } catch (error) {
        console.error('Error editing the component:', error);
      }
    },
    async deleteComponent({ commit }, componentId) {
      try {
        const response = await fetch(`http://localhost:8080/api/component/${componentId}`, {
          method: 'DELETE',
          headers: {
            "Access-Control-Allow-Origin": "*",
          },
        });

        if (!response.ok) {
          throw new Error('Failed to delete the component');
        }

        commit('DELETE_COMPONENT', componentId);
      } catch (error) {
        console.error('Error deleting the component:', error);
      }
    },
  },
  getters: {
    allFilteredComponents: (state) => {
      return state.filteredComponents;
    },
    totalPages: (state) => {
      return Math.ceil(state.filteredComponents.length / state.pagination.itemsPerPage);
    },
  },
};

