/* eslint-disable */
export default {
    namespaced: true,
    state: {
      cycles: [],
      currentCycle: {
        editing: false,
      },
      newCycle: {
        id: '0',
        module_id: '0',
        code_kz: '',
        code_ru: '',
        code_en: '',
      },
      editedCycle: {
        id: '0',
        module_id: '0',
        code_kz: '',
        code_ru: '',
        code_en: '',
      },
      pagination: {
        page: 1,
        itemsPerPage: 10,
      },
    },
    mutations: {
      SET_CYCLES(state, cycles) {
        state.cycles = cycles;
      },
      ADD_CYCLE(state, cycle) {
        state.cycles.push(cycle);
      },
      CLEAR_NEW_CYCLE(state) {
        state.newCycle = {
          id: '0',
          module_id: '0',
          code_kz: '',
          code_ru: '',
          code_en: '',
        };
      },
      EDIT_CYCLE(state, editedCycle) {
        const index = state.cycles.findIndex((c) => String(c.id) === editedCycle.id);
        if (index !== -1) {
          state.cycles[index] = editedCycle;
        }
      },
      DELETE_CYCLE(state, cycleId) {
        state.cycles = state.cycles.filter((c) => c.id !== cycleId);
      },
    },
    actions: {
      async fetchCycles({ commit, state }) {
        try {
          const response = await fetch(
            `http://localhost:8080/api/cycle?${state.pagination.page}&perPage=${state.pagination.itemsPerPage}`,
            {
              method: 'GET',
            }
          );
          if (!response.ok) {
            throw new Error('Failed to fetch cycles');
          }
          const cycles = await response.json();
          commit('SET_CYCLES', cycles);
        } catch (error) {
          console.error('Error fetching cycles:', error);
        }
      },
      async addNewCycle({ commit, state }) {
        try {
          const response = await fetch('http://localhost:8080/api/cycle', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              'Access-Control-Allow-Origin': '*',
            },
            body: JSON.stringify(state.newCycle),
          });
  
          if (!response.ok) {
            throw new Error('Failed to add a new cycle');
          }
  
          const newCycle = await response.json();
          state.newCycle.id = newCycle.id;
          commit('ADD_CYCLE', state.newCycle);
          commit('CLEAR_NEW_CYCLE');
        } catch (error) {
          console.error('Error adding a new cycle:', error);
        }
      },
      async updateCycle({ commit }, editedCycle) {
        try {
          const response = await fetch(`http://localhost:8080/api/cycle/${editedCycle.id}`, {
            method: 'PUT',
            headers: {
              'Content-Type': 'application/json',
              'Access-Control-Allow-Origin': '*',
            },
            body: JSON.stringify(editedCycle),
          });
  
          if (!response.ok) {
            throw new Error('Failed to edit the cycle');
          }
  
          commit('EDIT_CYCLE', editedCycle);
        } catch (error) {
          console.error('Error editing the cycle:', error);
        }
      },
      async deleteCycle({ commit }, cycleId) {
        try {
          const response = await fetch(`http://localhost:8080/api/cycle/${cycleId}`, {
            method: 'DELETE',
            headers: {
              'Access-Control-Allow-Origin': '*',
            },
          });
  
          if (!response.ok) {
            throw new Error('Failed to delete the cycle');
          }
  
          commit('DELETE_CYCLE', cycleId);
        } catch (error) {
          console.error('Error deleting the cycle:', error);
        }
      },
    },
    getters: {
      allCycles: (state) => {
        return state.cycles;
      },
      totalPages: (state) => {
        return Math.ceil(state.cycles.length / state.pagination.itemsPerPage);
      },
    },
  };
  