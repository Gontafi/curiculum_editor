<!-- eslint-disable -->
<template>
    <div class="container">
      <h2>Manage Departments</h2>
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Description KZ</th>
              <th>Description RU</th>
              <th>Description EN</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(department, index) in departments" :key="department.id">
              <td>
                <span>{{ department.id }}</span>
              </td>
              <td>
                <span v-if="!department.editing">{{ department.description_kz }}</span>
                <input v-else v-model="editedDepartment.description_kz" />
              </td>
              <td>
                <span v-if="!department.editing">{{ department.description_ru }}</span>
                <input v-else v-model="editedDepartment.description_ru" />
              </td>
              <td>
                <span v-if="!department.editing">{{ department.description_en }}</span>
                <input v-else v-model="editedDepartment.description_en" />
              </td>
              <td>
                <button @click="editDepartment(department)" v-if="!department.editing">Edit</button>
                <button @click="saveEditedDepartment(editedDepartment, department)" v-if="department.editing">Submit</button>
                <button @click="cancelEdit(department)" v-if="department.editing">Cancel</button>
                <button @click="deleteDepartment(department.id)" v-if="!department.editing">Delete</button>
              </td>
            </tr>
            <tr>
              <td></td>
              <td><input v-model="newDepartment.description_kz" /></td>
              <td><input v-model="newDepartment.description_ru" /></td>
              <td><input v-model="newDepartment.description_en" /></td>
              <td>
                <button @click="addNewDepartment">Add</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="pagination">
        <button @click="prevPage">Previous</button>
        <span>Page {{ pagination.page }} of {{ totalPages() }}</span>
        <button @click="nextPage">Next</button>
      </div>
    </div>
  </template>
<!-- eslint-disable -->
  <script>
  import { mapState, mapActions, mapGetters } from 'vuex';
  
  export default {
    name: "DepartmentComponent",
    computed: {
      ...mapState("department", ['departments', 'newDepartment', 'pagination', 'editedDepartment']),
    },
    created() {
      this.fetchDepartments();
    },
    methods: {
      ...mapActions("department", ['fetchDepartments', 'addNewDepartment', 'deleteDepartment', 'updateDepartment']),
      ...mapGetters("department", ['totalPages']),
      
      editDepartment(department) {
        department.editing = true;
        this.editedDepartment.id = String(department.id);
        this.editedDepartment.description_kz = department.description_kz;
        this.editedDepartment.description_ru = department.description_ru;
        this.editedDepartment.description_en = department.description_en;
      },
      
      saveEditedDepartment(department, currentDepartment) {
        this.updateDepartment(department);
        this.cancelEdit(currentDepartment);
      },
      
      cancelEdit(department) {
        department.editing = false;
      },
      
      prevPage() {
        this.pagination.page--;
        this.fetchDepartments();
      },
      
      nextPage() {
        this.pagination.page++;
        this.fetchDepartments();
      },
    },
  };
  </script>
<!-- eslint-disable -->
  <style>
  .container {
    max-width: 1100px;
    margin: 0 auto;
    padding: 20px;
    font-family: Arial, sans-serif;
  }
  
  h2 {
    margin-bottom: 20px;
  }
  
  .table-container {
    margin-bottom: 20px;
    overflow-x: auto;
  }
  
  table {
    width: 100%;
    border-collapse: collapse;
  }
  
  table, th, td {
    border: 1px solid #ccc;
  }
  
  th, td {
    padding: 5px;
    text-align: left;
  }
  
  th {
    background-color: #f2f2f2;
  }
  
  .pagination {
    text-align: center;
  }
  
  .pagination button {
    margin: 0 5px;
    padding: 5px 10px;
    border: 1px solid #ccc;
    background-color: #f2f2f2;
    cursor: pointer;
  }
  
  input[type="text"] {
    width: 100%;
    box-sizing: border-box;
  }
  </style>
