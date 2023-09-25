<!-- eslint-disable -->
<template>
    <div class="container">
      <h2>Manage Components</h2>
      <div class="table-container">
        <table>
          <thead>
            <tr>
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
                <span v-if="!component.editing">{{ component.codeKz }}</span>
                <input v-else v-model="component.editedCodeKz" />
              </td>
              <td>
                <span v-if="!component.editing">{{ component.codeRu }}</span>
                <input v-else v-model="component.editedCodeRu" />
              </td>
              <td>
                <span v-if="!component.editing">{{ component.codeEn }}</span>
                <input v-else v-model="component.editedCodeEn" />
              </td>
              <td>
                <span v-if="!component.editing">{{ component.descriptionKz }}</span>
                <input v-else v-model="component.editedDescriptionKz" />
              </td>
              <td>
                <span v-if="!component.editing">{{ component.descriptionRu }}</span>
                <input v-else v-model="component.editedDescriptionRu" />
              </td>
              <td>
                <span v-if="!component.editing">{{ component.descriptionEn }}</span>
                <input v-else v-model="component.editedDescriptionEn" />
              </td>
              <td>
                <span v-if="!component.editing">{{ component.order }}</span>
                <input v-else v-model="component.editedOrder" />
              </td>
              <td>
                <button @click="editComponent(component)" v-if="!component.editing">Edit</button>
                <button @click="saveEditedComponent(component)" v-if="component.editing">Submit</button>
                <button @click="cancelEdit(component)" v-if="component.editing">Cancel</button>
                <button @click="deleteComponent(component)" v-if="!component.editing">Delete</button>
              </td>
            </tr>
            <tr>
              <td><input v-model="newComponent.codeKz" /></td>
              <td><input v-model="newComponent.codeRu" /></td>
              <td><input v-model="newComponent.codeEn" /></td>
              <td><input v-model="newComponent.descriptionKz" /></td>
              <td><input v-model="newComponent.descriptionRu" /></td>
              <td><input v-model="newComponent.descriptionEn" /></td>
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
        <span>Page {{ pagination.page }} of {{ totalPages }}</span>
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
      ...mapState(['components', 'filteredComponents', 'newComponent', "pagination"]),
    },
    created() {
      this.fetchComponents();
    },
    methods: {
      ...mapActions(['fetchComponents', 'addNewComponent', 'deleteComponent']),
      ...mapGetters(['totalPages']),
      editComponent(component) {
        component.editing = true;
      },
      saveEditedComponent(component) {
        // Make an API request to update the component here
        // After successfully updating, set component.editing = false;
      },
      cancelEdit(component) {
        component.editing = false;
      },
      prevPage() {
        if (this.pagination.page > 1) {
          this.pagination.page--;
        }
      },
      nextPage() {
        if (this.pagination.page < this.totalPages) {
          this.pagination.page++;
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
