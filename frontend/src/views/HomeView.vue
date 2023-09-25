<script>
import HomeHeaderVue from "@/components/HomeHeader.vue";

export default {
  name: "HomeView",
  components: {
    HomeHeaderVue, // Register the component here
  },
};
</script>

<template>
  <div id="home">
    <HomeHeaderVue></HomeHeaderVue>
    <!-- eslint-disable -->
    <!-- Additional content for the home page -->
    <div class="page-content">
      <h2>Welcome to Our Curriculum Management System</h2>
      <p>
        Our system allows you to easily manage and browse educational
        modules and courses. Whether you're a student or an educator, you'll
        find the tools you need to succeed.
      </p>
      <div>
        <h3>Key Features:</h3>
        <ul class="custom-list">
          <li>Effortless curriculum planning and organization.</li>
          <li>Browse and discover a wide range of educational modules.</li>
          <li>Manage courses, track progress, and access resources.</li>
          <li>Collaborate with fellow educators and students.</li>
          <li>Streamline communication and stay informed about updates.</li>
        </ul>
      </div>
      <h3>Get Started:</h3>
      <p>
        Ready to begin your educational journey? Sign up or log in to get started.
        If you have any questions or need assistance, our support team is here to help.
      </p>

      <!-- Add more content or sections as needed -->
    </div>
  </div>
</template>

<style>
body {
  margin: 0;
}

#home {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.page-content {
  margin: 10px;
  padding: 10px;
  background-color: #e9e9e9; /* Light gray background color for content */
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.page-content h2 {
  font-size: 24px;
  margin-bottom: 10px;
}

.page-content p {
  font-size: 16px;
}

/* Custom style for the list items */
.custom-list {
  list-style-type: none; /* Remove default list bullets */
  padding-left: 0; /* Remove default left padding */
}

.custom-list li {
  margin-bottom: 8px; /* Add spacing between list items */
  text-align: center; /* Center-align the text within each list item */
}

.custom-list li::before {
  content: "\2022"; /* Use a custom bullet character (in this case, a bullet point) */
  color: #3c66aa; /* Set the color of the custom bullet */
  font-weight: bold; /* Optional: Make the bullet bold */
  position: relative; /* Position relative for custom bullet */
  left: 0; /* Adjust the position of the custom bullet to align with the text */
  margin-right: 5px; /* Add some spacing between the bullet and the text */
}
</style>
