<!-- eslint-disable -->
<template>
    <div>
      <div class="sign-up-back">
        <div class="sign-up-page">
          <h2>SignUp</h2>
          <form @submit.prevent="this.$store.dispatch('authComponent/signUp')">
            <div class="form-group">
              <label for="username">Username:</label>
              <input type="text" id="username" v-model="credentials.username" required>
            </div>
            <div class="form-group">
              <label for="password">Password:</label>
              <input type="password" id="password" v-model="credentials.password" required>
            </div>
            <button type="submit">Login</button>
          </form>
        </div>
      </div>
    </div>
  </template>
<!-- eslint-disable -->
  <script>
  export default {
    name: "SignUp",
    data () {
        return {
            credentials: this.$store.state.authComponent.credentials,
            accessToken: this.$store.state.authComponent.accessToken,
        }
    },
  };
  </script>
<!-- eslint-disable -->
  <style scoped>
  .sign-up-back {
      padding: 30px;
  }
  .sign-up-page {
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
