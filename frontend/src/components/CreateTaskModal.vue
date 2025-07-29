<template>
  <div class="modal-overlay" @click.self="$emit('close')">
    <div class="task-modal">
      <div class="modal-header">
        <h3>Создание новой задачи</h3>
      </div>
      
      <div class="modal-content">
        <div class="task-form">
          <div class="form-group">
            <label>Тип задачи:</label>
            <select 
              v-model="newTask.task_id" 
              class="form-input"
            >
              <option value="" disabled>Выберите тип задачи</option>
              <option 
                v-for="type in taskTypes" 
                :key="type.id" 
                :value="type.id"
              >
                {{ type.name }}
              </option>
            </select>
          </div>
<!--           
          <div class="form-group">
            <label>Исполнитель:</label>
            <select 
              v-model="newTask.executor" 
              class="form-input"
              :disabled="!newTask.task_id"
            >
              <option value="" disabled>Выберите исполнителя</option>
              <option 
                v-for="user in departmentUsers" 
                :key="user.id" 
                :value="user.id"
              >
                {{ user.surname }} {{ user.name }} {{ user.patronymic }} ({{ user.position }})
              </option>
            </select>
          </div> -->
          
          <div class="form-group">
            <label>Описание задачи:</label>
            <textarea 
              v-model="newTask.description"
              class="form-input"
              rows="3"
              placeholder="Введите описание задачи..."
            ></textarea>
          </div>

        </div>
      </div>

      <div class="modal-footer">
        <button class="save-btn" @click="saveTask">
          Сохранить задачу
        </button>
        <button class="close-btn" @click="$emit('close')">
          Отмена
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const props = defineProps({
  taskTypes: {
    type: Array,
    default: () => []
  },
  // departmentUsers: {
  //   type: Array,
  //   default: () => []
  // }
});

const emit = defineEmits(['close', 'save', 'task-type-changed']);

const newTask = ref({
  task_id: null,
  description: '',
  status: 0,
  create_date: String(new Date().toISOString().split('T')[0]),
});

const saveTask = () => {
  if (!newTask.value.description) {
    alert('Пожалуйста, заполните все обязательные поля');
    return;
  }
  
  emit('save', {
    task: newTask.value,
  });
};
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.task-modal {
  background-color: white;
  border-radius: 15px;
  width: 700px;
  max-height: 80vh;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.modal-header {
  padding: 20px;
  background-color: #5662DE;
  color: white;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-content {
  padding: 20px;
  overflow-y: auto;
  flex-grow: 1;
}

.task-form {
  padding: 15px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
  color: #5662DE;
}

.form-input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
  font-size: 16px;
}

select.form-input {
  height: 40px;
}

.modal-footer {
  padding: 15px 20px;
  background-color: #f9f9f9;
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.save-btn {
  padding: 8px 15px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.save-btn:hover {
  background-color: #3e8e41;
}

.close-btn {
  padding: 8px 15px;
  background-color: #ddd;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

h3{
  color: white;
}
</style>