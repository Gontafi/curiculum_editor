<template>
    <div>
        
    </div>
  </template>
  
  <script>
  export default {
    name: "CuriculumComponent",
    data() {
      return {
        semesters: [], // Здесь храните список семестров из базы данных
        selectedSemester: null,
        addingCourse: false,
        courseCode: "",
        courseName: "",
        // Другие поля, необходимые для вашей модели Course
      };
    },
  };
  </script>
  