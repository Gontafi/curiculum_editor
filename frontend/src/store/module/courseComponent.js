/* eslint-disable */
export default {
    state: {
        filteredComponents: [],
        currentComponent: {
          editing: false,
        },
        newComponent: {
          id: '0',
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
          state.components.push(component);
        },
        CLEAR_NEW_COMPONENT(state) {
          state.newComponent = {
            id: '0',
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
            const index = state.components.findIndex((c) => c.id === editedComponent.id);
            if (index !== -1) {
              state.components[index] = editedComponent;
            }
          },
        DELETE_COMPONENT(state, componentId) {
            state.components = state.components.filter((c) => c.id !== componentId);
        },
      },
      actions: {
        async fetchComponents({ commit }) {
          try {
            const response = await fetch('http://localhost:8080/api/professional-component', {
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
            const response = await fetch('http://localhost:8080/api/professional-component', {
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
        async updateComponent({ commit }, editedComponent) {
            try {
              const response = await fetch(`http://localhost:8080/api/professional-component/${editedComponent.id}`, {
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
              const response = await fetch(`http://localhost:8080/api/professional-component/${componentId}`, {
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
      },
};