<!-- eslint-disable -->
<template>
    <div>
        <HomeHeaderVue></HomeHeaderVue>
        <div class="login-back">
            <div class="login-page">
                <h2>Login</h2>
                <form @submit.prevent="login">
                    <div class="form-group">
                        <label for="username">Username:</label>
                        <input type="text" id="username" v-model="username" required>
                    </div>
                    <div class="form-group">
                        <label for="password">Password:</label>
                        <input type="password" id="password" v-model="password" required>
                    </div>
                    <button type="submit">Login</button>
                </form>
            </div>
        </div>
    </div>
</template>
<!-- eslint-disable -->
<script>
import axios from "axios";
import HomeHeaderVue from "@/components/HomeHeader.vue";

export default {
  name: "LoginView",
  components: {
    HomeHeaderVue,
  },
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    login() {
      // Prepare the user data to send in the request
      const userData = {
        username: this.username,
        password: this.password,
      };

      // Send a POST request to your backend API for authentication
      axios
        .post("/api/auth/sign-in", userData) // Replace with your API endpoint
        .then((response) => {
          // Authentication successful, you can handle the response here
          const accessToken = response.data.access_token;
          // Store the JWT access token in your Vuex store or local storage
          // For example, you can use Vuex to store the token
          this.$store.commit("setAccessToken", accessToken);

          // Redirect the user to the dashboard or another page
          this.$router.push("/dashboard"); // Replace with the desired route
        })
        .catch((error) => {
          // Handle authentication error
          console.error("Login failed:", error);
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
