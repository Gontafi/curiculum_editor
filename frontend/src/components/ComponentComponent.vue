<!-- eslint-disable -->
<template>
  <div class="container">
    <h2>Manage Components</h2>
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>ProfID</th>
            <th>Код KZ</th>
            <th>Код RU</th>
            <th>Код EN</th>
            <th>Описание KZ</th>
            <th>Описание RU</th>
            <th>Описание EN</th>
            <th>Порядок</th>
            <th>Действия</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(component, index) in filteredComponents" :key="component.id">
            <td>
              <span>{{ component.id }}</span>
            </td>
            <td>
              <span v-if="!component.editing">{{ component.prof_id }}</span>
              <input v-else v-model="editedComponent.prof_id" />
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
              <button @click="editComponent(component)" v-if="!component.editing">Изменить</button>
              <button @click="saveEditedComponent(editedComponent, component)" v-if="component.editing">Сохранить</button>
              <button @click="cancelEdit(component)" v-if="component.editing">Отмена</button>
              <button @click="deleteComponent(component.id)" v-if="!component.editing">Удалить</button>
            </td>
          </tr>
          <tr>
            <td></td>
            <td><input v-model="newComponent.prof_id" /></td>
            <td><input v-model="newComponent.code_kz" /></td>
            <td><input v-model="newComponent.code_ru" /></td>
            <td><input v-model="newComponent.code_en" /></td>
            <td><input v-model="newComponent.description_kz" /></td>
            <td><input v-model="newComponent.description_ru" /></td>
            <td><input v-model="newComponent.description_en" /></td>
            <td><input v-model="newComponent.order" /></td>
            <td>
              <button @click="addNewComponent">Добавить</button>
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
  name: "ComponentComponent",
  computed: {
    ...mapState("courseComponent",['filteredComponents','newComponent', "pagination", "editedComponent"]),
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
      this.editedComponent.prof_id = component.prof_id;
      this.editedComponent.code_en = component.code_en;
      this.editedComponent.code_kz = component.code_kz;
      this.editedComponent.code_ru = component.code_ru;
      this.editedComponent.description_en = component.description_en;
      this.editedComponent.description_kz = component.description_kz;
      this.editedComponent.description_ru = component.description_ru;
      this.editedComponent.order = String(component.order);
    },
    saveEditedComponent(component, currentComponent) {
      this.updateComponent(component);
      this.cancelEdit(currentComponent);
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
        this.pagination.page++;
        this.fetchComponents();
    },
  },
};
</script>
<!-- eslint-disable -->
<style>
.container {
  max-width: 1200px;
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