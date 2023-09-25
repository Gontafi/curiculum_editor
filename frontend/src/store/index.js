import { createStore } from "vuex";
/* eslint-disable */
export default createStore({
  state: {
    components: [],
    filteredComponents: [],
    newComponent: {
      codeKz: '',
      codeRu: '',
      codeEn: '',
      descriptionKz: '',
      descriptionRu: '',
      descriptionEn: '',
      order: 0,
    },
    pagination: {
      page: 1,
      itemsPerPage: 10, // You can adjust the number of items per page
    },
  },
  mutations: {
    SET_COMPONENTS(state, components) {
      state.components = components;
    },
    SET_FILTERED_COMPONENTS(state, filteredComponents) {
      state.filteredComponents = filteredComponents;
    },
    ADD_COMPONENT(state, component) {
      state.components.push(component);
    },
    CLEAR_NEW_COMPONENT(state) {
      state.newComponent = {
        codeKz: '',
        codeRu: '',
        codeEn: '',
        descriptionKz: '',
        descriptionRu: '',
        descriptionEn: '',
        order: 0,
      };
    },
    EDIT_COMPONENT(state, editedComponent) {
        const index = state.components.findIndex((c) => c.id === editedComponent.id);
        if (index !== -1) {
          state.components[index] = editedComponent;
        }
      },
    DELETE_COMPONENT(state, componentId) {
        state.components = state.components.filter((c) => c.id !== componentId);
    },
    // Add mutations for editing, deleting, and searching components here
  },
  actions: {
    async fetchComponents({ commit }) {
      try {
        const response = await fetch('http://localhost:3000/api/professional-component', {
          method: "GET",
          withCredential: true,
          headers: {
            "Access-Control-Allow-Origin": "*",
          }
        });
        if (!response.ok) {
          throw new Error('Failed to fetch components');
        }
        const components = await response.json();
        commit('SET_COMPONENTS', components);
      } catch (error) {
        console.error('Error fetching components:', error);
      }
    },
    async addNewComponent({ commit, state }) {
      try {
        const response = await fetch('http://localhost:3000/api/professional-component', {
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
        commit('ADD_COMPONENT', newComponent);
        commit('CLEAR_NEW_COMPONENT');
      } catch (error) {
        console.error('Error adding a new component:', error);
      }
    },
    async editComponent({ commit }, editedComponent) {
        try {
          const response = await fetch(`http://localhost:3000/api/professional-component/${editedComponent.id}`, {
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
          const response = await fetch(`http://localhost:3000/api/professional-component/${componentId}`, {
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
    totalPages: (state) => {
      return Math.ceil(state.filteredComponents.length / state.pagination.itemsPerPage);
    },
  },
});
