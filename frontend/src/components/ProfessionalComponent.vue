<!-- eslint-disable -->
<template>
  <div class="container">
    <h2>Управление Профессиональными компонентами</h2>
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th>ID</th>
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
          <tr v-for="(component, index) in filteredProfessionalComponents" :key="component.id">
            <td>
              <span>{{ component.id }}</span>
            </td>
            <td>
              <span v-if="!component.editing">{{ component.code_kz }}</span>
              <input v-else v-model="editedProfessionalComponent.code_kz" />
            </td>
            <td>
              <span v-if="!component.editing">{{ component.code_ru }}</span>
              <input v-else v-model="editedProfessionalComponent.code_ru" />
            </td>
            <td>
              <span v-if="!component.editing">{{ component.code_en }}</span>
              <input v-else v-model="editedProfessionalComponent.code_en" />
            </td>
            <td>
              <span v-if="!component.editing">{{ component.description_kz }}</span>
              <input v-else v-model="editedProfessionalComponent.description_kz" />
            </td>
            <td>
              <span v-if="!component.editing">{{ component.description_ru }}</span>
              <input v-else v-model="editedProfessionalComponent.description_ru" />
            </td>
            <td>
              <span v-if="!component.editing">{{ component.description_en }}</span>
              <input v-else v-model="editedProfessionalComponent.description_en" />
            </td>
            <td>
              <span v-if="!component.editing">{{ component.order }}</span>
              <input v-else v-model="editedProfessionalComponent.order" />
            </td>
            <td>
              <button @click="editComponent(component)" v-if="!component.editing">Изменить</button>
              <button @click="saveeditedProfessionalComponent(editedProfessionalComponent, component)" v-if="component.editing">Сохранить</button>
              <button @click="cancelEdit(component)" v-if="component.editing">Отмена</button>
              <button @click="deleteProfessionalComponent(component.id)" v-if="!component.editing">Удалить</button>
            </td>
          </tr>
          <tr>
            <td></td>
            <td><input v-model="newProfessionalComponent.code_kz" /></td>
            <td><input v-model="newProfessionalComponent.code_ru" /></td>
            <td><input v-model="newProfessionalComponent.code_en" /></td>
            <td><input v-model="newProfessionalComponent.description_kz" /></td>
            <td><input v-model="newProfessionalComponent.description_ru" /></td>
            <td><input v-model="newProfessionalComponent.description_en" /></td>
            <td><input v-model="newProfessionalComponent.order" /></td>
            <td>
              <button @click="addnewProfessionalComponent">Добавить</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="pagination">
      <button @click="prevPage">Previous</button>
      <span>Page {{ pagination.page }}</span>
      <button @click="nextPage">Next</button>
    </div>
  </div>
</template>
<!-- eslint-disable -->
<script>
import { mapState, mapActions, mapGetters } from 'vuex';

export default {
  name: 'ComponentComponent',
  computed: {
    ...mapState('courseProfessionalComponent', ['filteredProfessionalComponents', 'newProfessionalComponent', 'pagination', 'editedProfessionalComponent']),
  },
  created() {
    this.fetchProfessionalComponents();
  },
  methods: {
    ...mapActions('courseProfessionalComponent', ['fetchProfessionalComponents', 'addnewProfessionalComponent', 'deleteProfessionalComponent', 'updateProfessionalComponent']),
    ...mapGetters('courseProfessionalComponent', ['totalPages', 'allfilteredProfessionalComponents']),
    editComponent(component) {
      component.editing = true;
      this.editedProfessionalComponent.id = String(component.id);
      this.editedProfessionalComponent.code_en = component.code_en;
      this.editedProfessionalComponent.code_kz = component.code_kz;
      this.editedProfessionalComponent.code_ru = component.code_ru;
      this.editedProfessionalComponent.description_en = component.description_en;
      this.editedProfessionalComponent.description_kz = component.description_kz;
      this.editedProfessionalComponent.description_ru = component.description_ru;
      this.editedProfessionalComponent.order = String(component.order);
    },
    saveeditedProfessionalComponent(component, currentComponent) {
      this.updateProfessionalComponent(component);
      this.cancelEdit(currentComponent);
    },
    cancelEdit(component) {
      component.editing = false;
    },
    prevPage() {
      this.pagination.page--;
      console.log(this.pagination.page);
      this.fetchProfessionalComponents();
    },
    nextPage() {
      console.log(this.pagination.page);
      this.pagination.page++;
      this.fetchProfessionalComponents();
    },
  },
};
</script>
<!-- eslint-disable -->
<style scoped>
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
