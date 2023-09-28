/* eslint-disable */
export default {
    namespaced: true,
    name:"courseProfessionalComponent",
    state: {
        filteredProfessionalComponents: [],
        newProfessionalComponent: {
          id: '0',
          code_kz: '',           
          code_ru: '',           
          code_en: '',           
          description_kz: '',    
          description_ru: '',    
          description_en: '',    
          order: '0',            
        },
        editedProfessionalComponent: {
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
        SET_FILTERED_PROFESSIONAL_COMPONENTS(state, filteredProfessionalComponents) {
          state.filteredProfessionalComponents = filteredProfessionalComponents;
        },
        ADD_PROFESSIONAL_COMPONENT(state, component) {
          state.filteredProfessionalComponents.push(component);
        },
        CLEAR_NEW_PROFESSIONAL_COMPONENT(state) {
          state.newProfessionalComponent = {
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
        EDIT_PROFESSIONAL_COMPONENT(state, editedProfessionalComponent) {
            const index = state.filteredProfessionalComponents.findIndex((c) => String(c.id) === editedProfessionalComponent.id);
            if (index !== -1) {
              state.filteredProfessionalComponents[index] = editedProfessionalComponent;
            }
          },
        DELETE_PROFESSIONAL_COMPONENT(state, componentId) {
            state.filteredProfessionalComponents = state.filteredProfessionalComponents.filter((c) => c.id !== componentId);
        },
      },
      actions: {
        async fetchProfessionalComponents({ commit, state }) {
          try {
            const response = await fetch(`http://localhost:8080/api/professional-component?page=${state.pagination.page}&perPage=${state.pagination.itemsPerPage}`, {
              method: "GET",
            });
            if (!response.ok) {
              throw new Error('Failed to fetch components');
            }
            const components = await response.json();
            commit('SET_FILTERED_PROFESSIONAL_COMPONENTS', components);
          } catch (error) {
            console.error('Error fetching components:', error);
          }
        },
        async addnewProfessionalComponent({ commit, state }) {
          try {
            const response = await fetch('http://localhost:8080/api/professional-component', {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
                "Access-Control-Allow-Origin": "*",
              },
              body: JSON.stringify(state.newProfessionalComponent),
            });
    
            if (!response.ok) {
              throw new Error('Failed to add a new component');
            }
    
            const newProfessionalComponent = await response.json();
            state.newProfessionalComponent.id = newProfessionalComponent.id;
            commit('ADD_PROFESSIONAL_COMPONENT', state.newProfessionalComponent);
            commit('CLEAR_NEW_PROFESSIONAL_COMPONENT');
          } catch (error) {
            console.error('Error adding a new component:', error);
          }
        },
        async updateProfessionalComponent({ commit }, editedProfessionalComponent) {
            try {
              const response = await fetch(`http://localhost:8080/api/professional-component/${editedProfessionalComponent.id}`, {
                method: 'PUT',
                headers: {
                  'Content-Type': 'application/json',
                  "Access-Control-Allow-Origin": "*",
                },
                body: JSON.stringify(editedProfessionalComponent),
    
              });
        
              if (!response.ok) {
                throw new Error('Failed to edit the component');
              }
        
              commit('EDIT_PROFESSIONAL_COMPONENT', editedProfessionalComponent);
            } catch (error) {
              console.error('Error editing the component:', error);
            }
          },
          async deleteProfessionalComponent({ commit }, componentId) {
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
        
              commit('DELETE_PROFESSIONAL_COMPONENT', componentId);
            } catch (error) {
              console.error('Error deleting the component:', error);
            }
          },
      },
      getters: {
        allfilteredProfessionalComponents: (state) => {
          return state.filteredProfessionalComponents;
        },
        totalPages: (state) => {
          return Math.ceil(state.filteredProfessionalComponents.length / state.pagination.itemsPerPage);
        },
      },
};

