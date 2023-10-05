/* eslint-disable */
import router from "@/router";
export default {
  namespaced: true,
  name:"auth",
  state: {
      logined:false,
      credentials: {
        username: "",
        password: "",
      },
      loading:false,
    },
    
  mutations: {
    setLogin(state, bool) {
      state.logined = bool;
    },
    clearAccessToken(state) {
      state.accessToken = null;
    },
    setLoading(state, bool) {
      state.loading = bool;
    }
  },

  actions: {
    async login({ state, commit }) {
      commit("setLoading", true);
    
      try {
        const response = await Promise.race([
          fetch("http://localhost:8080/api/auth/sign-in/", {
            method: "POST",
            headers: {
              'Content-Type': 'application/json',
              'Access-Control-Allow-Origin': '*',
            },
            body: JSON.stringify(state.credentials),
          }),
          new Promise((_, reject) =>
            setTimeout(() => reject(new Error('Request timed out')), 5000)
          ),
        ]);
    
        if (!response.ok) {
          throw new Error('Failed to login');
        }
        
        const res = await response.json();
        localStorage.setItem("jwt", res.access_token);
        commit("setLogin", true);
        router.push('/dashboard');
      } catch (error) {
        console.error('Error while logging in:', error);
      } finally {
        commit("setLoading", false);
      }
    },
    async signUp({state}) {
      try{
        const response = await fetch("http://localhost:8080/api/auth/sign-up/", {
          method: "POST",
          headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*',
          },
          body: JSON.stringify(state.credentials),
        });
        if (!response.ok){
          throw new Error("Failed to register user");
        }
        router.push('/login');
      } catch(error) {
        console.error("error while trying to registrate", error)
      }
    },
    logout({state}) {
      localStorage.clear();
      console.log("cleared:" + localStorage.getItem("jwt"));
      state.logined = false;
      router.push("/login");
    },
  },

  getters: {
    isAuthenticated: (state) => {
      console.log("trying to access some route");
      console.log(localStorage.getItem("jwt"));
      if (localStorage.getItem("jwt") == null) {
        return false;
      }
      return true;
    }
  },
}