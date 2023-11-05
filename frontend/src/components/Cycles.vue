<!-- eslint-disable -->
<template>
    <div class="container">
      <h2>Manage Cycles</h2>
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Модуль ID</th>
              <th>Код имя KZ</th>
              <th>Код имя RU</th>
              <th>Код имя EN</th>
              <th>Действия</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(cycle, index) in cycles" :key="cycle.id">
              <td>
                <span>{{ cycle.id }}</span>
              </td>
              <td>
                <span v-if="!cycle.editing">{{ cycle.module_id }}</span>
                <input v-else v-model="editedCycle.module_id" />
              </td>
              <td>
                <span v-if="!cycle.editing">{{ cycle.code_kz }}</span>
                <input v-else v-model="editedCycle.code_kz" />
              </td>
              <td>
                <span v-if="!cycle.editing">{{ cycle.code_ru }}</span>
                <input v-else v-model="editedCycle.code_ru" />
              </td>
              <td>
                <span v-if="!cycle.editing">{{ cycle.code_en }}</span>
                <input v-else v-model="editedCycle.code_en" />
              </td>
              <td>
                <button @click="editCycle(cycle)" v-if="!cycle.editing">Изменить</button>
                <button @click="saveEditedCycle(editedCycle, cycle)" v-if="cycle.editing">Сохранить</button>
                <button @click="cancelEdit(cycle)" v-if="cycle.editing">Отмена</button>
                <button @click="deleteCycle(cycle.id)" v-if="!cycle.editing">Удалить</button>
              </td>
            </tr>
            <tr>
              <td></td>
              <td><input v-model="newCycle.module_id" /></td>
              <td><input v-model="newCycle.code_kz" /></td>
              <td><input v-model="newCycle.code_ru" /></td>
              <td><input v-model="newCycle.code_en" /></td>
              <td>
                <button @click="addNewCycle">Добавить</button>
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
    name: "CycleComponent",
    computed: {
      ...mapState("cycles", ['cycles', 'newCycle', 'pagination', 'editedCycle']),
    },
    created() {
      this.fetchCycles();
    },
    methods: {
      ...mapActions("cycles", ['fetchCycles', 'addNewCycle', 'deleteCycle', 'updateCycle']),
      ...mapGetters("cycles", ['totalPages']),
  
      editCycle(cycle) {
        cycle.editing = true;
        this.editedCycle.id = String(cycle.id);
        this.editedCycle.module_id = cycle.module_id;
        this.editedCycle.code_kz = cycle.code_kz;
        this.editedCycle.code_ru = cycle.code_ru;
        this.editedCycle.code_en = cycle.code_en;
        // Set other editedCycle properties as needed
      },
  
      saveEditedCycle(cycle, currentCycle) {
        this.updateCycle(cycle);
        this.cancelEdit(currentCycle);
      },
  
      cancelEdit(cycle) {
        cycle.editing = false;
      },
  
      prevPage() {
          this.pagination.page--;
          this.fetchCycles();
      },
  
      nextPage() {
          this.pagination.page++;
          this.fetchCycles();
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
