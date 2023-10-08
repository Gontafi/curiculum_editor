/* eslint-disable */
export default {
    namespaced: true,
    state: {
        prerequisites: [], // Array to store course prerequisites
        newPrerequisite: {
            course_id: null, // ID of the course
            prerequisite_id: null, // ID of the prerequisite course
        },
        selectedPrerequisites: []
    },
    mutations: {
        SET_PREREQUISITES(state, prerequisites) {
            state.prerequisites = prerequisites;
        },
        ADD_PREREQUISITE(state, prerequisite) {
            state.prerequisites.push(prerequisite);
            state.selectedPrerequisites.push(prerequisite)
        },
        DELETE_PREREQUISITE(state, prerequisiteId) {
            state.prerequisites = state.prerequisites.filter((p) => p.id !== prerequisiteId);
        },
        EDIT_PREREQUISITE(state, editedPrerequisite) {
            const index = state.prerequisites.findIndex((c) => String(c.id) === editedPrerequisite.id);
            if (index !== -1) {
                state.prerequisites[index] = editedPrerequisite;
            }
        },
        SET_SELECTED_COURSE(state, courseId) {
            state.selectedCourseId = courseId;
        },
    },
    actions: {
        async fetchPrerequisites({ commit }) {
            try {
                // Fetch prerequisites from your API endpoint
                const response = await fetch('http://localhost:8080/api/prerequisite', {
                    method: 'GET',
                });

                if (!response.ok) {
                    throw new Error('Failed to fetch prerequisites');
                }

                const prerequisites = await response.json();
                commit('SET_PREREQUISITES', prerequisites);
            } catch (error) {
                console.error('Error fetching prerequisites:', error);
            }
        },
        async addPrerequisite({ commit, state }, newPrerequisite) {
            try {
                // Send a POST request to add a new prerequisite
                const response = await fetch('http://localhost:8080/api/prerequisite', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Access-Control-Allow-Origin': '*',
                    },
                    body: JSON.stringify(newPrerequisite),
                });

                if (!response.ok) {
                    throw new Error('Failed to add a new prerequisite');
                }

                const addedPrerequisite = await response.json();
                commit('ADD_PREREQUISITE', addedPrerequisite);
                // Clear the newPrerequisite object
                state.newPrerequisite = {
                    course_id: null,
                    prerequisite_id: null,
                };
            } catch (error) {
                console.error('Error adding a new prerequisite:', error);
            }
        },
        async deletePrerequisite({ commit }, prerequisiteId) {
            try {
                // Send a DELETE request to remove a prerequisite
                const response = await fetch(`http://localhost:8080/api/prerequisite/${prerequisiteId}`, {
                    method: 'DELETE',
                    headers: {
                        'Access-Control-Allow-Origin': '*',
                    },
                });

                if (!response.ok) {
                    throw new Error('Failed to delete the prerequisite');
                }

                commit('DELETE_PREREQUISITE', prerequisiteId);
            } catch (error) {
                console.error('Error deleting the prerequisite:', error);
            }
        },
        async editPrerequisite({ commit }, prerequisite) {
            try {
                // Send a PUT request to edit a prerequisite
                const response = await fetch(`http://localhost:8080/api/prerequisite/${prerequisite.id}`, {
                    method: 'PUT',
                    headers: {
                        'Access-Control-Allow-Origin': '*',
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(prerequisite),
                });

                if (!response.ok) {
                    throw new Error('Failed to edit the prerequisite');
                }

                commit('EDIT_PREREQUISITE', prerequisite);
            } catch (error) {
                console.error('Error editing the prerequisite:', error);
            }
        },
        setSelectedCourse({ commit }, courseId) {
            commit('SET_SELECTED_COURSE', courseId);
        },
    },
    getters: {
        allPrerequisites: (state) => {
            return state.prerequisites;
        },
        selectedCoursePrerequisites: (state) => {
            return state.prerequisites.filter((p) => p.courseId === state.selectedCourseId);
        },
    },
};
