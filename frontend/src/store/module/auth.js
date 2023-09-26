/* eslint-disable */
const state = {
    accessToken: null,
    input: {
      username: "",
      password: "",
    }
  };
  
  const mutations = {
    setAccessToken(state, token) {
      state.accessToken = token;
    },
    clearAccessToken(state) {
      state.accessToken = null;
    },
  };
  
  const actions = {
    async login({state, commit}) {
      try {
        const response = await fetch("http://localhost:8080/login", {
            method: "POST",
            body: JSON.stringify(state.input),
          });
        if (!response.ok) {
          throw new Error('Failed to login');
        }
        commit("setAccessToken", response.token)
      } catch (error) {
        console.error('Error adding while login:', error);
      }
    },
    async signUp({state, commit}) {
      try{
        const response = await fetch("http://localhost:8080/sign-up", {
          method: "POST",
          body: JSON.stringify(state.input),
        });
        if (!response.ok){
          throw new Error("Failed to register user");
        }
        commit()
      } catch(error) {
        console.error("error while trying to registrate", error)
      }
    },
  };
  
  const getters = {
    isAuthenticated: (state) => !!state.accessToken,
  };
  
  export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters,
  };