<!-- eslint-disable -->
<template>
    <div>
      <HomeHeaderVue></HomeHeaderVue>
      <SignUp></SignUp>
    </div>
  </template>
<!-- eslint-disable -->  
<script>

import HomeHeaderVue from "@/components/HomeHeader.vue";
import SignUp from "@/components/RegistratingComponent.vue"
export default {
name: "SignupView",
components: {
    HomeHeaderVue,
    SignUp,
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
