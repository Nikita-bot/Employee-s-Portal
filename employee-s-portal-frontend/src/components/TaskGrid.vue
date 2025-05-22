<template>
  <div class="mainContent">
    <div class="menuTask">
      <div class="chooseBlock">
        <div 
          class="blockEl"
          :class="{ active: activeTab === 'assigned' }"
          @click="activeTab = 'assigned'"
        >
          Назначенные мне
        </div>
        <div class="line"></div>
        <div 
          class="blockEl"
          :class="{ active: activeTab === 'created' }"
          @click="activeTab = 'created'"
        >
          Созданные мной
        </div>
      </div>

      <div class="create-task-container">
        <button class="create-btn" @click="createTask">
          + Создать новую задачу
        </button>
      </div>
    </div>
    
    <div class="task-grid">
      <!-- Заголовки таблицы -->
      <div 
        class="grid-header sortable"
        @click="sortBy('id')"
      >
        ID
        <span v-if="sortField === 'id'" class="sort-icon">
          {{ sortDirection === 'asc' ? '↑' : '↓' }}
        </span>
      </div>
      <div 
        class="grid-header sortable"
        @click="sortBy(activeTab === 'assigned' ? 'author' : 'executor')"
      >
        {{ activeTab === 'assigned' ? 'Автор' : 'Исполнитель' }}
        <span 
          v-if="sortField === (activeTab === 'assigned' ? 'author' : 'executor')" 
          class="sort-icon"
        >
          {{ sortDirection === 'asc' ? '↑' : '↓' }}
        </span>
      </div>
      <div 
        class="grid-header sortable"
        @click="sortBy('description')"
      >
        Описание
        <span v-if="sortField === 'description'" class="sort-icon">
          {{ sortDirection === 'asc' ? '↑' : '↓' }}
        </span>
      </div>
      <div 
        class="grid-header sortable"
        @click="sortBy('status')"
      >
        Статус
        <span v-if="sortField === 'status'" class="sort-icon">
          {{ sortDirection === 'asc' ? '↑' : '↓' }}
        </span>
      </div>
      <div 
        class="grid-header sortable"
        @click="sortBy('createdDate')"
      >
        Дата создания
        <span v-if="sortField === 'createdDate'" class="sort-icon">
          {{ sortDirection === 'asc' ? '↑' : '↓' }}
        </span>
      </div>
      <div 
        class="grid-header sortable"
        @click="sortBy('assignedDate')"
      >
        Дата выполнения
        <span v-if="sortField === 'assignedDate'" class="sort-icon">
          {{ sortDirection === 'asc' ? '↑' : '↓' }}
        </span>
      </div>

      <!-- Строки с задачами -->
      <template v-for="task in displayedTasks" :key="task.id">
        <div 
          class="grid-item clickable"
          @click="openTaskModal(task)"
        >
          {{ task.id }}
        </div>
        <div 
          class="grid-item clickable"
          @click="openTaskModal(task)"
        >
          {{ activeTab === 'assigned' ? task.author : task.executor }}
        </div>
        <div 
          class="grid-item clickable"
          @click="openTaskModal(task)"
        >
          {{ task.description }}
        </div>
        <div 
          class="grid-item clickable"
          @click="openTaskModal(task)"
        >
          {{ task.status }}
        </div>
        <div 
          class="grid-item clickable"
          @click="openTaskModal(task)"
        >
          {{ task.createdDate }}
        </div>
        <div 
          class="grid-item clickable"
          @click="openTaskModal(task)"
        >
          {{ task.assignedDate }}
        </div>
      </template>
    </div>

    <!-- Модальное окно задачи -->
    <div v-if="selectedTask" class="modal-overlay" @click.self="closeModal">
      <div class="task-modal">
        <div class="modal-header">
          <h3>Задача #{{ selectedTask.id }}</h3>

        </div>
        
        <div class="modal-content">
          <div class="task-info">
            <p><strong>Автор:</strong> {{ selectedTask.author }}</p>
            <p><strong>Исполнитель:</strong> {{ selectedTask.executor }}</p>
            <p><strong>Описание:</strong> {{ selectedTask.description }}</p>
            <p><strong>Статус:</strong> {{ selectedTask.status }}</p>
            <p><strong>Дата создания:</strong> {{ selectedTask.createdDate }}</p>
            <p><strong>Дата назначения:</strong> {{ selectedTask.assignedDate }}</p>
          </div>

          <div class="comments-section">
            <h4>Комментарии:</h4>
            <div v-for="comment in selectedTask.comments" :key="comment.id" class="comment">
              <p><strong>{{ comment.author }}:</strong> {{ comment.text }}</p>
              <small>{{ comment.date }}</small>
            </div>
            
            <div class="new-comment">
              <textarea 
                v-model="newComment"
                placeholder="Напишите комментарий..."
                rows="3"
              ></textarea>
              <button @click="addComment">Отправить</button>
            </div>
          </div>
        </div>

        <div class="modal-footer">
          <button 
            v-if="activeTab === 'created'"
            class="delete-btn"
            @click="deleteTask"
          >
            Удалить задачу
          </button>
          <button 
            v-if="activeTab === 'assigned' && selectedTask.status !== 'Выполнено'"
            class="complete-btn"
            @click="completeTask"
          >
            Выполнить задачу
          </button>
          <button class="close-btn" @click="closeModal">Закрыть</button>
        </div>
      </div>
    </div>

    <!-- Модальное окно создания задачи -->
    <div v-if="isCreateModalOpen" class="modal-overlay" @click.self="closeCreateModal">
      <div class="task-modal">
        <div class="modal-header">
          <h3>Создание новой задачи</h3>
        </div>
        
        <div class="modal-content">
          <div class="task-form">
            <div class="form-group">
              <label>Тип задачи:</label>
              <select v-model="newTask.type" class="form-input">
                <option v-for="type in taskTypes" :key="type.value" :value="type.value">
                  {{ type.label }}
                </option>
              </select>
            </div>
            
            <div class="form-group">
              <label>Исполнитель:</label>
              <select v-model="newTask.executor" class="form-input">
                <option v-for="user in users" :key="user.id" :value="user.name">
                  {{ user.name }} ({{ user.position }})
                </option>
              </select>
            </div>
            
            <div class="form-group">
              <label>Комментарий:</label>
              <textarea 
                v-model="newTaskComment"
                class="form-input"
                rows="3"
                placeholder="Введите комментарий к задаче..."
              ></textarea>
            </div>
          </div>
        </div>

        <div class="modal-footer">
          <button class="save-btn" @click="saveNewTask">
            Сохранить задачу
          </button>
          <button class="close-btn" @click="closeCreateModal">
            Отмена
          </button>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useUserStore } from '@/stores/user';

const userStore = useUserStore();

// Состояния
const activeTab = ref('assigned');
const selectedTask = ref(null);
const newComment = ref('');
const sortField = ref('');
const sortDirection = ref('asc');
const isCreateModalOpen = ref(false);
const isLoading = ref(false);

// Данные для создания задачи
const newTask = ref({
  task_id: 1,
  executor: '',
  description: '',
  status: 1,
  create_date: new Date().toISOString().split('T')[0],
  execute_date: ''
});

const newTaskComment = ref('');
const taskTypes = ref([
  { value: 1, label: 'Проверить отчет' },
  { value: 2, label: 'Добавить пользователя' },
  { value: 3, label: 'Обновить данные' },
  { value: 4, label: 'Другое' }
]);

const users = ref([]);
const tasks = ref({
  assigned: [],
  created: []
});

/* Новые функции для работы с API */

const fetchTasks = async () => {
  isLoading.value = true;
  try {
    const response = await fetch(`/api/v1/tasks/user/${userStore.userData.id}`);
    if (!response.ok) throw new Error('Ошибка загрузки задач');
    
    const data = await response.json();
    tasks.value.assigned = data.executor || [];
    tasks.value.created = data.initiator || [];
    
    // Загрузка пользователей (можно вынести в отдельный метод)
    const usersResponse = await fetch('/api/v1/users');
    if (usersResponse.ok) {
      users.value = await usersResponse.json();
    }
  } catch (error) {
    console.error('Ошибка загрузки задач:', error);
    alert('Не удалось загрузить задачи');
  } finally {
    isLoading.value = false;
  }
};

const fetchTaskById = async (id) => {
  try {
    const response = await fetch(`/api/v1/tasks/${id}`);
    if (!response.ok) throw new Error('Ошибка загрузки задачи');
    const data = await response.json();
    return data.task;
  } catch (error) {
    console.error('Ошибка загрузки задачи:', error);
    alert('Не удалось загрузить задачу');
    return null;
  }
};

const createNewTask = async () => {
  if (!newTask.value.executor || !newTask.value.description) {
    alert('Пожалуйста, заполните все обязательные поля');
    return;
  }

  try {
    const response = await fetch('/api/v1/tasks', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        ...newTask.value,
        initiator: userStore.userData.id,
        description: newTask.value.description + (newTaskComment.value ? `\nКомментарий: ${newTaskComment.value}` : '')
      })
    });

    if (!response.ok) throw new Error('Ошибка создания задачи');
    
    alert('Задача успешно создана');
    closeCreateModal();
    fetchTasks();
  } catch (error) {
    console.error('Ошибка создания задачи:', error);
    alert('Не удалось создать задачу');
  }
};

/* Неизмененные функции из вашего исходного кода */

const displayedTasks = computed(() => {
  return activeTab.value === 'assigned' ? tasks.value.assigned : tasks.value.created;
});

const sortBy = (field) => {
  if (sortField.value === field) {
    sortDirection.value = sortDirection.value === 'asc' ? 'desc' : 'asc';
  } else {
    sortField.value = field;
    sortDirection.value = 'asc';
  }
  
  const taskList = activeTab.value === 'assigned' ? tasks.value.assigned : tasks.value.created;
  taskList.sort((a, b) => {
    const valueA = a[field];
    const valueB = b[field];
    
    if (valueA < valueB) return sortDirection.value === 'asc' ? -1 : 1;
    if (valueA > valueB) return sortDirection.value === 'asc' ? 1 : -1;
    return 0;
  });
};

const openTaskModal = async (task) => {
  isLoading.value = true;
  try {
    const fullTask = await fetchTaskById(task.id);
    if (fullTask) {
      selectedTask.value = {
        ...fullTask,
        author: users.value.find(u => u.id === fullTask.initiator)?.name || 'Неизвестно',
        executor: users.value.find(u => u.id === fullTask.executor)?.name || 'Неизвестно',
        status: getStatusText(fullTask.status),
        createdDate: formatDate(fullTask.create_date),
        assignedDate: formatDate(fullTask.execute_date),
        comments: []
      };
    }
  } finally {
    isLoading.value = false;
  }
};

const closeModal = () => {
  selectedTask.value = null;
  newComment.value = '';
};

const addComment = () => {
  if (!newComment.value.trim()) return;
  
  const comment = {
    id: Date.now(),
    author: 'Я (текущий пользователь)',
    text: newComment.value,
    date: new Date().toLocaleString()
  };
  
  const taskList = activeTab.value === 'assigned' ? tasks.value.assigned : tasks.value.created;
  const taskIndex = taskList.findIndex(t => t.id === selectedTask.value.id);
  
  if (taskIndex !== -1) {
    taskList[taskIndex].comments.push(comment);
    newComment.value = '';
  }
};

const deleteTask = () => {
  if (confirm('Вы уверены, что хотите удалить эту задачу?')) {
    const taskList = activeTab.value === 'assigned' ? tasks.value.assigned : tasks.value.created;
    const taskIndex = taskList.findIndex(t => t.id === selectedTask.value.id);
    
    if (taskIndex !== -1) {
      taskList.splice(taskIndex, 1);
      closeModal();
    }
  }
};

const createTask = () => {
  newTask.value = {
    task_id: 1,
    executor: '',
    description: '',
    status: 1,
    create_date: new Date().toISOString().split('T')[0],
    execute_date: ''
  };
  newTaskComment.value = '';
  isCreateModalOpen.value = true;
};

const closeCreateModal = () => {
  isCreateModalOpen.value = false;
};

const completeTask = () => {
  const taskList = tasks.value.assigned;
  const taskIndex = taskList.findIndex(t => t.id === selectedTask.value.id);
  
  if (taskIndex !== -1) {
    taskList[taskIndex].status = 3; // 3 = Выполнено
    taskList[taskIndex].execute_date = new Date().toISOString().split('T')[0];
    closeModal();
  }
};

/* Вспомогательные функции */

const formatDate = (dateString) => {
  if (!dateString) return '';
  return new Date(dateString).toLocaleDateString('ru-RU');
};

const getStatusText = (statusCode) => {
  const statuses = {
    1: 'В работе',
    2: 'Выполнено',
  };
  return statuses[statusCode] || 'Неизвестно';
};

// Инициализация
onMounted(() => {
  fetchTasks();
});
</script>

<style scoped>
/* Существующие стили */
.mainContent {
  margin-top: 40px;
  margin-left: 40px;
  width: 1500px;
}

.chooseBlock {
  display: flex;
  justify-content: space-evenly;
  align-items: center;
  width: 300px;
  height: 50px;
  background-color: #FFFFFF;
  border-radius: 15px;
}

.blockEl {
  border-radius: 15px;
  height: 30px;
  padding: 3px 5px 0px 5px;
  color: #5662DE;
  cursor: pointer;
  transition: background-color 0.2s;
}

.blockEl:hover {
  background-color: #e0e0e0;
}

.blockEl.active {
  background-color: #5662DE;
  color: #FFFFFF;
}

.line {
  width: 1.5px;
  background-color: #5662DE;
  height: 40px;
  margin: 0 5px;
}

.task-grid {
  width: 90%;
  margin-top: 50px;
  display: grid;
  grid-template-columns: 10% 15% 30% 15% 15% 15%;
  gap: 1px; 
  border-radius: 15px;
  overflow: hidden; 
  font-size: 20px;
}

.grid-header {
  background-color: #5662DE;
  color: white;
  padding: 20px 10px;
  text-align: center;
  font-size: 20px;
}

.grid-item {
  background-color: white;
  padding: 20px 10px;
  text-align: center;
  word-break: break-word;
  margin-top: 20px;
}

.clickable {
  cursor: pointer;
  transition: background-color 0.2s;
}

.clickable:hover {
  background-color: #f0f0f0;
}

/* Стили модального окна */
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

.task-info {
  margin-bottom: 20px;
}

.task-info p {
  margin: 10px 0;
}

.comments-section {
  border-top: 1px solid #eee;
  padding-top: 20px;
}

.comment {
  margin-bottom: 15px;
  padding-bottom: 15px;
  border-bottom: 1px solid #f5f5f5;
}

.comment small {
  color: #888;
}

.new-comment {
  margin-top: 20px;
}

.new-comment textarea {
  width: 100%;
  padding: 10px;
  border-radius: 5px;
  border: 1px solid #ddd;
  resize: vertical;
}

.new-comment button {
  margin-top: 10px;
  padding: 8px 15px;
  background-color: #5662DE;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.modal-footer {
  padding: 15px 20px;
  background-color: #f9f9f9;
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.close-btn {
  padding: 8px 15px;
  background-color: #ddd;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.delete-btn {
  padding: 8px 15px;
  background-color: #ff4444;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

h3{
  color:#FFFFFF;
}

.create-task-container{
  background-color: #5662DE;
  height: 50px;
  width: 300px;
  border-radius: 15px;
  text-align: center;
}

.create-btn{
  background-color: #5662DE;
  border: none;
  border-radius: 15px;
  color: #FFFFFF;
  font-size: 17px;
  width: 100%;
  height: 100%;
  cursor: pointer;
  transition: background-color 0.2s;
}

.create-btn:hover{
  background-color: #3944bd;
}

.menuTask{
  width: 90%;
  display: flex;
  justify-content: space-between;
  margin-right: 130px;
}

.complete-btn{
  border: none;
  background-color: #5662DE;
  border-radius: 5px;
  color: #FFFFFF;
  width: 150px;
}


.sortable {
  cursor: pointer;
  transition: background-color 0.2s;
  padding: 5px;
  border-radius: 5px;
}

.sortable:hover {
  background-color: #3f4ac0;
}

.sortable span {
  margin-left: 5px;
  font-weight: bold;
}
/* Добавляем новые стили для формы */
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

/* ... остальные существующие стили ... */
</style>