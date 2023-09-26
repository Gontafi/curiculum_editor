<!-- eslint-disable -->
<template>
    <div class="container">
      <h2>Manage Components</h2>
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Code KZ</th>
              <th>Code RU</th>
              <th>Code EN</th>
              <th>Description KZ</th>
              <th>Description RU</th>
              <th>Description EN</th>
              <th>Order</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(component, index) in filteredComponents" :key="component.id">
              <td>
                <span>{{ component.id }}</span>
              </td>
              <td>
                <span v-if="!component.editing">{{ component.code_kz }}</span>
                <input v-else v-model="editedComponent.code_kz" />
              </td>
              <td>
                <span v-if="!component.editing">{{ component.code_ru }}</span>
                <input v-else v-model="editedComponent.code_ru" />
              </td>
              <td>
                <span v-if="!component.editing">{{ component.code_en }}</span>
                <input v-else v-model="editedComponent.code_en" />
              </td>
              <td>
                <span v-if="!component.editing">{{ component.description_kz }}</span>
                <input v-else v-model="editedComponent.description_kz" />
              </td>
              <td>
                <span v-if="!component.editing">{{ component.description_ru }}</span>
                <input v-else v-model="editedComponent.description_ru" />
              </td>
              <td>
                <span v-if="!component.editing">{{ component.description_en }}</span>
                <input v-else v-model="editedComponent.description_en" />
              </td>
              <td>
                <span v-if="!component.editing">{{ component.order }}</span>
                <input v-else v-model="editedComponent.order" />
              </td>
              <td>
                <button @click="editComponent(component)" v-if="!component.editing">Edit</button>
                <button @click="saveEditedComponent(editedComponent, component)" v-if="component.editing">Submit</button>
                <button @click="cancelEdit(component)" v-if="component.editing">Cancel</button>
                <button @click="deleteComponent(component.id)" v-if="!component.editing">Delete</button>
              </td>
            </tr>
            <tr>
              <td></td>
              <td><input v-model="newComponent.code_kz" /></td>
              <td><input v-model="newComponent.code_ru" /></td>
              <td><input v-model="newComponent.code_en" /></td>
              <td><input v-model="newComponent.description_kz" /></td>
              <td><input v-model="newComponent.description_ru" /></td>
              <td><input v-model="newComponent.description_en" /></td>
              <td><input v-model="newComponent.order" /></td>
              <td>
                <button @click="addNewComponent">Add</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="pagination">
        <button @click="prevPage" :disabled="pagination.page === 1">Previous</button>
        <span>Page {{ pagination.page }} of {{ totalPages() }}</span>
        <button @click="nextPage" :disabled="pagination.page === totalPages">Next</button>
      </div>
    </div>
  </template>
<!-- eslint-disable -->
  <script>
  import { mapState, mapActions, mapGetters } from 'vuex';
  export default {
    name: "ComponentComponent",
    computed: {
      ...mapState("courseComponent",['filteredComponents', 'newComponent', "pagination", "editedComponent"]),
    },
    created() {
      this.fetchComponents();
    },
    methods: {
       ...mapActions("courseComponent", ['fetchComponents', 'addNewComponent', 'deleteComponent', 'updateComponent']),
       ...mapGetters("courseComponent", ['totalPages', "allFilteredComponents"]),
      editComponent(component) {
        component.editing = true;
        this.editedComponent.id = String(component.id);
        this.editedComponent.code_en = component.code_en;
        this.editedComponent.code_kz = component.code_kz;
        this.editedComponent.code_ru = component.code_ru;
        this.editedComponent.description_en = component.description_en;
        this.editedComponent.description_kz = component.description_kz;
        this.editedComponent.description_ru = component.description_ru;
        this.editedComponent.order = String(component.order);
      },
      saveEditedComponent(component, currentCompoent) {
        this.updateComponent(component);
        this.cancelEdit(currentCompoent);
      },
      cancelEdit(component) {
        component.editing = false;
      },
      prevPage() {
        if (this.pagination.page > 1) {
          this.pagination.page--;
          this.fetchComponents();
        }
      },
      nextPage() {
        if (this.pagination.page < this.totalPages) {
          this.pagination.page++;
          this.fetchComponents();
        }
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
