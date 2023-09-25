const state = {
    accessToken: null,
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
    // Your actions for login, logout, and token refresh can go here
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