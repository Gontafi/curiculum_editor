/* eslint-disable */
export default{
  state: {
    filteredModules: [],
    currentModule: {
      editing: false,
    },
    newModule: {
      id: '0',
      code: '',           // Add your module properties here
      name_kz: '',        
      name_ru: '',        
      name_en: '',        
    },
    editedModule: {
      id: '0',
      code: '',           // Add your module properties here
      name_kz: '',        
      name_ru: '',        
      name_en: '',        
    },
    pagination: {
      page: 1,
      itemsPerPage: 10, // You can adjust the number of items per page
    },
  },
  mutations: {
    SET_FILTERED_MODULES(state, filteredModules) {
      state.filteredModules = filteredModules;
    },
    ADD_MODULE(state, module) {
      state.modules.push(module);
    },
    CLEAR_NEW_MODULE(state) {
      state.newModule = {
        id: '0',
        code: '',           // Clear new module properties
        name_kz: '',        
        name_ru: '',        
        name_en: '',        
      };
    },
    EDIT_MODULE(state, editedModule) {
      const index = state.modules.findIndex((m) => m.id === editedModule.id);
      if (index !== -1) {
        state.modules[index] = editedModule;
      }
    },
    DELETE_MODULE(state, moduleId) {
      state.modules = state.modules.filter((m) => m.id !== moduleId);
    },
  },
  actions: {
    async fetchModules({ commit }) {
      try {
        const response = await fetch('http://localhost:8080/api/modules', {
          method: "GET",
        });
        if (!response.ok) {
          throw new Error('Failed to fetch modules');
        }
        const modules = await response.json();
        commit('SET_FILTERED_MODULES', modules);
      } catch (error) {
        console.error('Error fetching modules:', error);
      }
    },
    async addNewModule({ commit, state }) {
      try {
        const response = await fetch('http://localhost:8080/api/modules', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": "*",
          },
          body: JSON.stringify(state.newModule),
        });

        if (!response.ok) {
          throw new Error('Failed to add a new module');
        }

        const newModule = await response.json();
        commit('ADD_MODULE', newModule);
        commit('CLEAR_NEW_MODULE');
      } catch (error) {
        console.error('Error adding a new module:', error);
      }
    },
    async updateModule({ commit }, editedModule) {
      try {
        const response = await fetch(`http://localhost:8080/api/modules/${editedModule.id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": "*",
          },
          body: JSON.stringify(editedModule),
        });

        if (!response.ok) {
          throw new Error('Failed to edit the module');
        }

        commit('EDIT_MODULE', editedModule);
      } catch (error) {
        console.error('Error editing the module:', error);
      }
    },
    async deleteModule({ commit }, moduleId) {
      try {
        const response = await fetch(`http://localhost:8080/api/modules/${moduleId}`, {
          method: 'DELETE',
          headers: {
            "Access-Control-Allow-Origin": "*",
          },
        });

        if (!response.ok) {
          throw new Error('Failed to delete the module');
        }

        commit('DELETE_MODULE', moduleId);
      } catch (error) {
        console.error('Error deleting the module:', error);
      }
    },
  },
  getters: {
    allFilteredModules: (state) => {
      return state.filteredModules;
    },
    totalPages: (state) => {
      return Math.ceil(state.filteredModules.length / state.pagination.itemsPerPage);
    },
  },
};
