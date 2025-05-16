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
import { ref, computed } from 'vue';

const activeTab = ref('assigned');
const selectedTask = ref(null);
const newComment = ref('');

const sortField = ref('');
const sortDirection = ref('asc');
const isCreateModalOpen = ref(false);

const newTask = ref({
  type: '',
  executor: ''
});
const newTaskComment = ref('');

const taskTypes = ref([
  { value: 'review', label: 'Проверить отчет' },
  { value: 'add_user', label: 'Добавить пользователя' },
  { value: 'update', label: 'Обновить данные' },
  { value: 'other', label: 'Другое' }
]);

// Список пользователей (можно заменить на данные с сервера)
const users = ref([
  { id: 1, name: 'Петров П.П.', position: 'Руководитель' },
  { id: 2, name: 'Сулима Р.И.', position: 'Разработчик' },
  { id: 3, name: 'Иванов И.И.', position: 'Аналитик' }
]);
// Пример данных с комментариями
const tasksData = ref({
  assigned: [
    {
      id: 32156,
      author: 'Петров П.П.',
      executor: 'Сулима Р.И.',
      description: 'Добавить пользователя в МИС',
      status: 'Выполнено',
      createdDate: '29.04.2025',
      assignedDate: '29.04.2025',
      comments: [
        {
          id: 1,
          author: 'Сулима Р.И.',
          text: 'Задача выполнена',
          date: '29.04.2025 15:30'
        }
      ]
    },
    {
      id: 32154,
      author: 'Петров П.F.',
      executor: 'Сулима Р.И.',
      description: 'Добавить пользователя в МИС',
      status: 'В работе',
      createdDate: '29.04.2025',
      assignedDate: '',
      comments: [
        {
          id: 1,
          author: 'Сулима Р.И.',
          text: 'Задача Невыполнима !!!!',
          date: '29.04.2025 15:30'
        }
      ]
    }
  ],
  created: [
    {
      id: 32158,
      author: 'Иванов И.И.',
      executor: 'Петров П.П.',
      description: 'Проверить отчет',
      status: 'Новая',
      createdDate: '03.05.2025',
      assignedDate: '03.05.2025',
      comments: []
    }
  ]
});

const sortBy = (field) => {
  if (sortField.value === field) {
    // Если уже сортируем по этому полю, меняем направление
    sortDirection.value = sortDirection.value === 'asc' ? 'desc' : 'asc';
  } else {
    // Если новое поле, сортируем по возрастанию
    sortField.value = field;
    sortDirection.value = 'asc';
  }
  
  // Применяем сортировку
  const taskList = activeTab.value === 'assigned' ? tasksData.value.assigned : tasksData.value.created;
  taskList.sort((a, b) => {
    let valueA = a[field];
    let valueB = b[field];
    
    // Для дат преобразуем в timestamp
    if (field.includes('Date') && valueA && valueB) {
      valueA = new Date(valueA.split('.').reverse().join('-')).getTime();
      valueB = new Date(valueB.split('.').reverse().join('-')).getTime();
    }
    
    // Сравниваем значения
    if (valueA < valueB) return sortDirection.value === 'asc' ? -1 : 1;
    if (valueA > valueB) return sortDirection.value === 'asc' ? 1 : -1;
    return 0;
  });
};

const displayedTasks = computed(() => {
  return activeTab.value === 'assigned' ? tasksData.value.assigned : tasksData.value.created;
});

const openTaskModal = (task) => {
  selectedTask.value = { ...task };
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
  
  // Находим задачу в данных и добавляем комментарий
  const taskList = activeTab.value === 'assigned' ? tasksData.value.assigned : tasksData.value.created;
  const taskIndex = taskList.findIndex(t => t.id === selectedTask.value.id);
  
  if (taskIndex !== -1) {
    taskList[taskIndex].comments.push(comment);
    // selectedTask.value.comments.push(comment);
    newComment.value = '';
  }
};

const deleteTask = () => {
  if (confirm('Вы уверены, что хотите удалить эту задачу?')) {
    const taskList = activeTab.value === 'assigned' ? tasksData.value.assigned : tasksData.value.created;
    const taskIndex = taskList.findIndex(t => t.id === selectedTask.value.id);
    
    if (taskIndex !== -1) {
      taskList.splice(taskIndex, 1);
      closeModal();
    }
  }
};


const createTask = () => {
  // Сбрасываем форму
  newTask.value = {
    type: '',
    executor: ''
  };
  newTaskComment.value = '';
  
  // Открываем модальное окно
  isCreateModalOpen.value = true;
};

const closeCreateModal = () => {
  isCreateModalOpen.value = false;
};

const saveNewTask = () => {
  if (!newTask.value.type || !newTask.value.executor) {
    alert('Пожалуйста, заполните все обязательные поля');
    return;
  }
  
  // Создаем новую задачу
  const taskDescription = taskTypes.value.find(t => t.value === newTask.value.type)?.label || 'Новая задача';
  
  const newTaskObj = {
    id: Date.now(),
    author: 'Я (текущий пользователь)',
    executor: newTask.value.executor,
    description: taskDescription,
    status: 'Новая',
    createdDate: new Date().toLocaleDateString(),
    assignedDate: '',
    comments: []
  };
  
  // Добавляем комментарий, если он есть
  if (newTaskComment.value.trim()) {
    newTaskObj.comments.push({
      id: Date.now(),
      author: 'Я (текущий пользователь)',
      text: newTaskComment.value,
      date: new Date().toLocaleString()
    });
  }
  
  // Добавляем задачу в список
  tasksData.value.created.unshift(newTaskObj);
  
  // Закрываем модальное окно
  closeCreateModal();
};

const completeTask = () => {
  const taskList = tasksData.value.assigned;
  const taskIndex = taskList.findIndex(t => t.id === selectedTask.value.id);
  
  if (taskIndex !== -1) {
    taskList[taskIndex].status = 'Выполнено';
    taskList[taskIndex].assignedDate = new Date().toLocaleDateString();
    closeModal();
  }
};
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