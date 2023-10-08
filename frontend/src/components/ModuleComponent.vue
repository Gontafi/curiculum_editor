<!-- eslint-disable -->
<template>
    <div class="container">
      <h2>Manage Modules</h2>
      <div class="table-container">
        <!-- Table for displaying modules -->
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Code</th>
              <th>Name KZ</th>
              <th>Name RU</th>
              <th>Name EN</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(module, index) in filteredModules" :key="module.id">
              <td>
                <span>{{ module.id }}</span>
              </td>
              <td>
                <span v-if="!module.editing">{{ module.code }}</span>
                <input v-else v-model="editedModule.code" />
              </td>
              <td>
                <span v-if="!module.editing">{{ module.name_kz }}</span>
                <input v-else v-model="editedModule.name_kz" />
              </td>
              <td>
                <span v-if="!module.editing">{{ module.name_ru }}</span>
                <input v-else v-model="editedModule.name_ru" />
              </td>
              <td>
                <span v-if="!module.editing">{{ module.name_en }}</span>
                <input v-else v-model="editedModule.name_en" />
              </td>
              <td>
                <button @click="editModule(module)" v-if="!module.editing">Edit</button>
                <button @click="saveEditedModule(editedModule, module)" v-if="module.editing">Submit</button>
                <button @click="cancelEdit(module)" v-if="module.editing">Cancel</button>
                <button @click="deleteModule(module.id)" v-if="!module.editing">Delete</button>
              </td>
            </tr>
            <!-- Add new module row -->
            <tr>
              <td></td>
              <td><input v-model="newModule.code"/></td>
              <td><input v-model="newModule.name_kz" /></td>
              <td><input v-model="newModule.name_ru" /></td>
              <td><input v-model="newModule.name_en" /></td>
              <td>
                <button @click="addNewModule">Add</button>
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
    name: "ModuleComponent",
    computed: {
      ...mapState('courseModule', ['filteredModules', 'newModule', 'pagination', 'editedModule']),
    },
    created() {
      this.fetchModules();
    },
    methods: {
      ...mapActions('courseModule', ['fetchModules', 'addNewModule', 'deleteModule', 'updateModule']),
      ...mapGetters('courseModule', ['totalPages']),
      editModule(module) {
        module.editing = true;
        this.editedModule.id = String(module.id);
        this.editedModule.code = module.code;
        this.editedModule.name_kz = module.name_kz;
        this.editedModule.name_ru = module.name_ru;
        this.editedModule.name_en = module.name_en;
      },
      saveEditedModule(module, currentmodule) {
        this.updateModule(module);
        this.cancelEdit(currentmodule);
      },
      cancelEdit(module) {
        module.editing = false;
      },
      prevPage() {
        this.pagination.page--;
        this.fetchModules();
      },
      nextPage() {
        this.pagination.page++;
        this.fetchModules();
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
    padding: 10px;
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