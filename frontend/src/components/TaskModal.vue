<template>
  <div v-if="task" class="modal-overlay" @click.self="$emit('close')">
    <div class="task-modal">
      <div class="modal-header">
        <h3>Задача #{{ task.id }}</h3>
      </div>
      
      <div class="tabs">
        <button 
          class="tab-button" 
          :class="{ 'active': activeTab === 'info' }"
          @click="activeTab = 'info'"
        >
          Информация
        </button>
        <button 
          class="tab-button" 
          :class="{ 'active': activeTab === 'journal' }"
          @click="activeTab = 'journal'"
        >
          Журнал
        </button>
      </div>
      
      <div class="modal-content">
        <!-- Вкладка с информацией о задаче -->
        <div v-if="activeTab === 'info'" class="task-info">
          <p><strong>Автор:</strong> {{ formatFullName(task.initiator.surname, task.initiator.name, task.initiator.patronymic) }}</p>
          <p><strong>Исполнитель:</strong> 
            <span v-if="!isChangingExecutor">{{ formatFullName(task.executor.surname, task.executor.name, task.executor.patronymic) }}</span>
            <select v-else v-model="selectedExecutor" class="executor-select">
              <option v-for="user in departmentUsers" :key="user.id" :value="user.id">
                {{ formatFullName(user.surname, user.name, user.patronymic) }}
              </option>
            </select>
            <button 
              v-if="isExecutor && !isChangingExecutor && task.status !== 'В работе'" 
              class="change-executor-btn"
              @click="startChangeExecutor"
            >
              Сменить исполнителя
            </button>
          </p>
          <p><strong>Приоритет:</strong> 
            <span class="priority-badge" :class="getPriorityClass(task.priority)">
              {{ getPriorityText(task.priority) }}
            </span>
          </p>
          <p><strong>Описание:</strong> {{ task.description }}</p>
          <p><strong>Статус:</strong> {{ task.status }}</p>
          <p><strong>Дата создания:</strong> {{ task.create_date }}</p>
          <p><strong>Дата выполнения:</strong> {{ task.execute_date ? task.execute_date : 'Не выполнено' }}</p>

          <div class="comments-section">
            <h4>Комментарии:</h4>
            <div v-for="comment in task.comments" :key="comment.id" class="comment">
              <p><strong>{{ formatFullName(comment.author.surname, comment.author.name, comment.author.patronymic) }}:</strong> {{ comment.comment }}</p>
              <small>{{ comment.creation_date }}</small>
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

        <!-- Вкладка с журналом изменений -->
        <div v-if="activeTab === 'journal'" class="journal-section">
          <div v-if="journalLoading" class="loading">Загрузка журнала...</div>
          <div v-else>
            <table class="journal-table">
              <thead>
                <tr>
                  <th>Дата</th>
                  <th>Действие</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(entry, index) in journalEntries" :key="index">
                  <td>{{ entry.creation_date }}</td>
                  <td>{{ entry.action }}</td>
                </tr>
              </tbody>
            </table>
            <div v-if="journalEntries.length === 0" class="no-entries">
              В журнале нет записей
            </div>
          </div>
        </div>
      </div>

      <div class="modal-footer">
        <div v-if="isChangingExecutor && activeTab === 'info'" class="executor-actions">
          <button class="save-btn" @click="saveNewExecutor">Сохранить</button>
          <button class="cancel-btn" @click="cancelChangeExecutor">Отмена</button>
        </div>
        <button 
          v-if="isCreator && activeTab === 'info'"
          class="delete-btn"
          @click="$emit('delete')"
        >
          Удалить задачу
        </button>
        <button 
          v-if="isExecutor && task.status === 'Создана' && activeTab === 'info'"
          class="complete-btn"
          @click="$emit('take')"  
        >
          Взять в работу
        </button>
        <button 
          v-if="isExecutor && task.status === 'В работе' && activeTab === 'info'"
          class="complete-btn"
          @click="$emit('complete')"
        >
          Выполнить задачу
        </button>
        <button class="close-btn" @click="$emit('close')">Закрыть</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, defineProps, defineEmits, onMounted, watch } from 'vue';

const props = defineProps({
  task: {
    type: Object,
    required: true
  },
  isCreator: {
    type: Boolean,
    default: false
  },
  isExecutor: {
    type: Boolean,
    default: false
  },
  currentUserId: {
    type: Number,
    required: true
  }
});

const emit = defineEmits(['close', 'delete', 'complete', 'take', 'add-comment', 'refresh-tasks']);

const newComment = ref('');
const departmentUsers = ref([]);
const isChangingExecutor = ref(false);
const selectedExecutor = ref(null);
const activeTab = ref('info');
const journalEntries = ref([]);
const journalLoading = ref(false);

const formatFullName = (surname, name, patronymic) => {
  if (!surname) return '';
  
  let formatted = surname;
  if (name) {
    formatted += ` ${name.charAt(0)}.`;
    if (patronymic) {
      formatted += `${patronymic.charAt(0)}.`;
    }
  }
  return formatted;
};

// Функция для получения текстового представления приоритета
const getPriorityText = (priority) => {
  const priorityMap = {
    1: 'Низкий',
    2: 'Нормальный',
    3: 'Высокий'
  };
  return priorityMap[priority] || 'Не указан';
};

// Функция для получения класса CSS в зависимости от приоритета
const getPriorityClass = (priority) => {
  const priorityClassMap = {
    1: 'priority-low',
    2: 'priority-normal',
    3: 'priority-high'
  };
  return priorityClassMap[priority] || '';
};

const fetchJournal = async () => {
  try {
    journalLoading.value = true;
    const response = await fetch(`/api/v1/journal/${props.task.id}`);
    
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    
    const data = await response.json();
    journalEntries.value = data.journal || [];
  } catch (error) {
    console.error('Ошибка при получении журнала:', error);
  } finally {
    journalLoading.value = false;
  }
};

const changeExecutor = async () => {
  try {
    const response = await fetch(`/api/v1/tasks/executor/${props.task.id}`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        executor_id: selectedExecutor.value
      })
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    
    const newExecutor = departmentUsers.value.find(user => user.id === selectedExecutor.value);
    
    props.task.executor = newExecutor;
    isChangingExecutor.value = false;
    emit('close');
    emit('refresh-tasks')
    
  } catch (error) {
    console.error('Ошибка при изменении исполнителя:', error);
  }
};

const addComment = () => {
  if (!newComment.value.trim()) return;
  emit('add-comment', newComment.value);
  newComment.value = '';
};

const fetchDepartmentUsers = async () => {
  try {
    console.log(props.task.id)
    const response = await fetch(`/api/v1/depatments/user/${props.task.id}`);
    
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    
    const data = await response.json();
    
    departmentUsers.value = data.users;
  } catch (error) {
    console.error('Ошибка при получении пользователей отдела:', error);
  }
};

const startChangeExecutor = async () => {
  await fetchDepartmentUsers();
  selectedExecutor.value = props.task.executor.id;
  isChangingExecutor.value = true;
};

const cancelChangeExecutor = () => {
  isChangingExecutor.value = false;
};

const saveNewExecutor = async () => {
  if (selectedExecutor.value && selectedExecutor.value !== props.task.executor.id) {
    await changeExecutor()
  }
  isChangingExecutor.value = false;
};

// Загружаем пользователей отдела при открытии модального окна, если задача назначена текущему пользователю
onMounted(() => {
  if (props.isExecutor) {
    fetchDepartmentUsers();
  }
});

// Следим за изменением задачи
watch(() => props.task, () => {
  if (props.isExecutor) {
    fetchDepartmentUsers();
  }
});

// Загружаем журнал при переключении на вкладку
watch(() => activeTab.value, (newTab) => {
  if (newTab === 'journal') {
    fetchJournal();
  }
});
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

.tabs {
  display: flex;
  border-bottom: 1px solid #ddd;
}

.tab-button {
  padding: 10px 20px;
  margin: 2px 0px;
  background: none;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  font-size: 16px;
  transition: all 0.3s;
}

.tab-button:hover {
  background-color: #f5f5f5;
}

.tab-button.active {
  background-color: #5662DE;
  color: white;
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
  display: flex;
  align-items: center;
}

.executor-select {
  padding: 5px;
  border-radius: 4px;
  border: 1px solid #ddd;
  margin-left: 10px;
}

.change-executor-btn {
  margin-left: 10px;
  padding: 4px 8px;
  background-color: #5662DE;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.8em;
}

/* Стили для отображения приоритета */
.priority-badge {
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 0.85em;
  font-weight: 500;
}

.priority-low {
  background-color: #e8f5e9;
  color: #2e7d32;
}

.priority-normal {
  background-color: #e3f2fd;
  color: #1565c0;
}

.priority-high {
  background-color: #ffebee;
  color: #c62828;
  font-weight: bold;
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

/* Стили для журнала */
.journal-section {
  margin-top: 10px;
}

.journal-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 15px;
}

.journal-table th, .journal-table td {
  padding: 10px;
  text-align: left;
  border-bottom: 1px solid #eee;
}

.journal-table th {
  background-color: #f5f5f5;
  font-weight: 500;
}

.loading {
  padding: 20px;
  text-align: center;
  color: #666;
}

.no-entries {
  padding: 20px;
  text-align: center;
  color: #888;
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

.complete-btn {
  border: none;
  background-color: #5662DE;
  border-radius: 5px;
  color: #FFFFFF;
  width: 160px;
  padding: 8px 15px;
  cursor: pointer;
}

.executor-actions {
  display: flex;
  gap: 10px;
  margin-right: auto;
}

.save-btn {
  padding: 8px 15px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.cancel-btn {
  padding: 8px 15px;
  background-color: #f44336;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

h3 {
  color: #FFFFFF;
}

strong{
  margin-right: 5px;
}
</style>