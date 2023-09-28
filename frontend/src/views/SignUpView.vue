<!-- eslint-disable -->
<template>
    <div>
      <HomeHeaderVue></HomeHeaderVue>
      <SignUp></SignUp>
    </div>
  </template>
<!-- eslint-disable -->  
<script>

import axios from "axios";
import HomeHeaderVue from "@/components/HomeHeader.vue";
import SignUp from "@/components/RegistratingComponent.vue"
export default {
name: "SignupView",
components: {
    HomeHeaderVue,
    SignUp,
},
data() {
    return {
    username: "",
    password: "",
    };
},
methods: {
    signup() {
        // Send a POST request to your backend API for user registration
        // You can use Axios or another HTTP library to make the request
        const userData = {
            username: this.username,
            password: this.password,
        };

        axios
            .post("/api/auth/sign-up", userData)
            .then((response) => {
            // Registration successful, you can handle the response here
            const accessToken = response.data.access_token;
            // Store the JWT access token in your Vuex store or local storage
            // For example, you can use Vuex to store the token
            this.$store.commit("setAccessToken", accessToken);

            // Redirect to the login page or another page as needed
            this.$router.push("/login");
            })
            .catch((error) => {
            // Handle registration error (e.g., duplicate username)
            console.error("Registration failed:", error);
            // Display an error message to the user if needed
        });
    },
},
};
</script>
<!-- eslint-disable -->  
<style scoped>
  .login-back {
    padding: 30px;
  }
  .login-page {
    max-width: 400px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  }
  
  .form-group {
    padding: 5px;
    margin-bottom: 15px;
  }
  
  label {
    display: block;
    font-weight: bold;
  }
  
  input[type="text"],
  input[type="email"],
  input[type="password"] {
    width: 100%;
    padding: 5px;
    padding-top: 15px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  
  button {
    display: block;
    width: 100%;
    padding: 10px;
    background-color: #007bff;
    color: #fff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
</style>
