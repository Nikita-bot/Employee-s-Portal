<template>
  <div class="tasks-wrapper">
    <AppHeader />
    
    <div class="main-container">
      <AppSidebar />
      
      <main class="content-area">
        <div class="page-header">
          <div>
            <h1>Задачи</h1>
            <div class="task-count">Всего задач: {{ filteredTasks.length }}</div>
          </div>
        </div>

        <div class="task-controls">
          <button class="btn btn-primary" @click="showCreateTaskModal">Создать задачу</button>
          <select class="filter-select" v-model="statusFilter" @change="filterTasks">
            <option value="all">Все задачи</option>
            <option value="active">Не выполненные</option>
            <option value="completed">Выполненные</option>
          </select>
          <div class="view-toggle">
            <button class="view-toggle-btn" :class="{ active: currentView === 'assigned-to-me' }" 
                    @click="switchView('assigned-to-me')">Назначенные мне</button>
            <button class="view-toggle-btn" :class="{ active: currentView === 'assigned-by-me' }" 
                    @click="switchView('assigned-by-me')">Назначенные мной</button>
          </div>
        </div>

        <div class="tasks-grid">
          <div v-for="task in filteredTasks" :key="task.id" 
               :class="['task-item', `${task.priority}-priority`, { completed: task.status === 'completed' }]"
               @click="openViewTaskModal(task.id)">
            <div class="task-header">
              <div>
                <div class="task-title">{{ task.title }}</div>
                <div class="task-description">{{ task.description }}</div>
              </div>
              <span :class="['task-priority', `priority-${task.priority}`]">
                {{ getPriorityText(task.priority) }}
              </span>
            </div>
            <div class="task-meta">
              <span>{{ currentView === 'assigned-to-me' ? 'Создатель: ' + task.creator : 'Исполнитель: ' + task.assignee }}</span>
              <span :class="['task-status', `status-${task.status}`]">
                {{ getStatusText(task.status) }}
              </span>
            </div>
          </div>
        </div>
      </main>
    </div>

    <!-- Модальное окно создания задачи -->
    <div class="modal" :class="{ show: showCreateModal }" @click="closeCreateTaskModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h2>Создание новой задачи</h2>
          <button class="close-btn" @click="closeCreateTaskModal">&times;</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label for="taskTitle">Название задачи</label>
            <input type="text" id="taskTitle" v-model="newTask.title" placeholder="Введите название задачи">
          </div>
          <div class="form-group">
            <label for="taskType">Тип задачи</label>
            <select id="taskType" v-model="newTask.type">
              <option value="medical">Медицинская задача</option>
              <option value="administrative">Административная</option>
              <option value="technical">Техническая</option>
              <option value="other">Другая</option>
            </select>
          </div>
          <div class="form-group">
            <label for="taskDescription">Описание</label>
            <textarea id="taskDescription" v-model="newTask.description" 
                      placeholder="Подробное описание задачи"></textarea>
          </div>
          <div class="form-group">
            <label for="taskPriority">Приоритет</label>
            <select id="taskPriority" v-model="newTask.priority">
              <option value="low">Низкий</option>
              <option value="medium">Средний</option>
              <option value="high">Высокий</option>
            </select>
          </div>
          <div class="form-group">
            <label for="taskAssignee">Исполнитель</label>
            <select id="taskAssignee" v-model="newTask.assignee">
              <option value="petrov">Петров М.И. (Хирург)</option>
              <option value="sidorova">Сидорова А.В. (Медсестра)</option>
              <option value="kozlov">Козлов Д.С. (Администратор)</option>
              <option value="ivanov">Иванов А.С. (Кардиолог)</option>
            </select>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-outline" @click="closeCreateTaskModal">Отмена</button>
          <button class="btn btn-primary" @click="createNewTask">Создать задачу</button>
        </div>
      </div>
    </div>

    <!-- Модальное окно просмотра задачи -->
    <div class="modal" :class="{ show: showViewModal }" @click="closeViewTaskModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h2>{{ selectedTask.title }}</h2>
          <button class="close-btn" @click="closeViewTaskModal">&times;</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label>Создатель</label>
            <div>{{ selectedTask.creator }}</div>
          </div>
          <div class="form-group">
            <label>Исполнитель</label>
            <select v-model="selectedTask.assignee" @change="updateTaskAssignee">
              <option value="petrov">Петров М.И. (Хирург)</option>
              <option value="sidorova">Сидорова А.В. (Медсестра)</option>
              <option value="kozlov">Козлов Д.С. (Администратор)</option>
              <option value="ivanov">Иванов А.С. (Кардиолог)</option>
            </select>
          </div>
          <div class="form-group">
            <label>Приоритет</label>
            <div>{{ getPriorityText(selectedTask.priority) }}</div>
          </div>
          <div class="form-group">
            <label>Статус</label>
            <select v-model="selectedTask.status" @change="updateTaskStatus">
              <option value="created">Создана</option>
              <option value="in-progress">В работе</option>
              <option value="completed">Выполнена</option>
            </select>
          </div>
          <div class="form-group">
            <label>Описание</label>
            <div>{{ selectedTask.description }}</div>
          </div>
          
          <div class="comments-section">
            <h3>Комментарии</h3>
            <div class="comment" v-for="comment in selectedTask.comments" :key="comment.time">
              <div class="comment-header">
                <span class="comment-author">{{ comment.author }}</span>
                <span class="comment-time">{{ comment.time }}</span>
              </div>
              <div class="comment-text">{{ comment.text }}</div>
            </div>
            <div class="comment-input">
              <input type="text" v-model="newComment" placeholder="Введите комментарий..." @keypress.enter="addComment">
              <button class="btn btn-primary" @click="addComment">Отправить</button>
            </div>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-primary" @click="closeViewTaskModal">Закрыть</button>
        </div>
      </div>
    </div>

    <AppFooter />
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import AppHeader from '@/components/AppHeader.vue'
import AppSidebar from '@/components/AppSidebar.vue'
import AppFooter from '@/components/AppFooter.vue'

const showCreateModal = ref(false)
const showViewModal = ref(false)
const currentView = ref('assigned-to-me')
const statusFilter = ref('all')
const newComment = ref('')

const tasks = ref([
  {
    id: 1,
    title: "Консилиум по пациенту Петров И.В.",
    description: "Обсуждение результатов МРТ и плана лечения. Присутствие обязательно.",
    priority: "high",
    status: "created",
    creator: "Главный врач",
    assignee: "ivanov",
    type: "medical",
    comments: [
      { author: "Главный врач", text: "Подготовьте заключение до встречи", time: "05.10.2023 10:00" }
    ]
  },
  {
    id: 2,
    title: "Заполнение медицинской документации",
    description: "Завершить ведение истории болезней за октябрь",
    priority: "medium",
    status: "in-progress",
    creator: "ivanov",
    assignee: "sidorova",
    type: "administrative",
    comments: []
  },
  {
    id: 3,
    title: "Плановый осмотр пациентов",
    description: "Ежедневный обход пациентов кардиологического отделения",
    priority: "low",
    status: "completed",
    creator: "Администрация",
    assignee: "ivanov",
    type: "medical",
    comments: [
      { author: "Иванов А.С.", text: "Осмотр завершен, все пациенты стабильны", time: "06.10.2023 09:30" }
    ]
  }
])

const selectedTask = reactive({
  id: 0,
  title: '',
  description: '',
  priority: '',
  status: '',
  creator: '',
  assignee: '',
  type: '',
  comments: []
})

const newTask = reactive({
  title: '',
  description: '',
  type: 'medical',
  priority: 'medium',
  assignee: 'petrov'
})

const filteredTasks = computed(() => {
  let filtered = tasks.value

  // Фильтрация по виду (назначенные мне/мной)
  if (currentView.value === 'assigned-to-me') {
    filtered = filtered.filter(task => task.assignee === 'ivanov')
  } else {
    filtered = filtered.filter(task => task.creator === 'ivanov')
  }

  // Фильтрация по статусу
  if (statusFilter.value === 'active') {
    filtered = filtered.filter(task => task.status !== 'completed')
  } else if (statusFilter.value === 'completed') {
    filtered = filtered.filter(task => task.status === 'completed')
  }

  return filtered
})

const getAssigneeName = (assigneeValue) => {
  const mapping = {
    'petrov': 'Петров М.И.',
    'sidorova': 'Сидорова А.В.',
    'kozlov': 'Козлов Д.С.',
    'ivanov': 'Иванов А.С.'
  }
  return mapping[assigneeValue] || 'Петров М.И.'
}

const getStatusText = (status) => {
  const texts = {
    'created': 'Создана',
    'in-progress': 'В работе',
    'completed': 'Выполнена'
  }
  return texts[status]
}

const getPriorityText = (priority) => {
  const texts = {
    'high': 'Высокий',
    'medium': 'Средний',
    'low': 'Низкий'
  }
  return texts[priority]
}

const switchView = (view) => {
  currentView.value = view
}

const filterTasks = () => {
  // Фильтрация происходит автоматически через computed свойство
}

const showCreateTaskModal = () => {
  showCreateModal.value = true
}

const closeCreateTaskModal = () => {
  showCreateModal.value = false
  Object.assign(newTask, {
    title: '',
    description: '',
    type: 'medical',
    priority: 'medium',
    assignee: 'petrov'
  })
}

const createNewTask = () => {
  if (newTask.title && newTask.description) {
    const task = {
      id: tasks.value.length + 1,
      title: newTask.title,
      description: newTask.description,
      type: newTask.type,
      priority: newTask.priority,
      status: 'created',
      creator: 'ivanov',
      assignee: newTask.assignee,
      comments: []
    }
    
    tasks.value.push(task)
    closeCreateTaskModal()
  } else {
    alert('Пожалуйста, заполните название и описание задачи')
  }
}

const openViewTaskModal = (taskId) => {
  const task = tasks.value.find(t => t.id === taskId)
  if (task) {
    Object.assign(selectedTask, task)
    showViewModal.value = true
  }
}

const closeViewTaskModal = () => {
  showViewModal.value = false
}

const updateTaskStatus = () => {
  // Обновление происходит реактивно через v-model
}

const updateTaskAssignee = () => {
  // Обновление происходит реактивно через v-model
}

const addComment = () => {
  if (newComment.value.trim()) {
    const comment = {
      author: 'Иванов А.С.',
      text: newComment.value,
      time: new Date().toLocaleString('ru-RU')
    }
    selectedTask.comments.push(comment)
    newComment.value = ''
  }
}
</script>

<style scoped>
.tasks-wrapper {
  display: grid;
  grid-template-rows: auto 1fr auto;
  min-height: 100vh;
  background-color: #f8f9fa;
  color: #333;
  line-height: 1.6;
}

.main-container {
  display: grid;
  grid-template-columns: 250px 1fr;
  gap: 0;
}

.content-area {
  padding: 2rem;
  background-color: #fff;
  min-height: calc(100vh - 140px);
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 2rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid #eee;
}

.page-header h1 {
  font-weight: 300;
  color: #2c5aa0;
}

.task-count {
  color: #666;
  font-size: 0.9rem;
  margin-top: 0.5rem;
}

.task-controls {
  display: flex;
  gap: 1rem;
  margin-bottom: 2rem;
  align-items: center;
}

.btn {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
  transition: all 0.3s ease;
}

.btn-primary {
  background: #2c5aa0;
  color: white;
}

.btn-primary:hover {
  background: #1e3d6f;
}

.view-toggle {
  display: flex;
  background: #f8f9fa;
  border-radius: 6px;
  padding: 0.25rem;
  margin-left: auto;
}

.view-toggle-btn {
  padding: 0.5rem 1rem;
  border: none;
  background: none;
  cursor: pointer;
  border-radius: 4px;
  transition: all 0.3s ease;
}

.view-toggle-btn.active {
  background: #2c5aa0;
  color: white;
}

.filter-select {
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  background: white;
  min-width: 150px;
}

.tasks-grid {
  display: grid;
  gap: 1rem;
}

.task-item {
  background: #f8f9fa;
  border-left: 4px solid #2c5aa0;
  border-radius: 6px;
  padding: 1.5rem;
  transition: all 0.3s ease;
  cursor: pointer;
}

.task-item:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transform: translateY(-2px);
}

.task-item.high-priority {
  border-left-color: #e74c3c;
}

.task-item.medium-priority {
  border-left-color: #f39c12;
}

.task-item.completed {
  border-left-color: #27ae60;
  opacity: 0.7;
}

.task-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 0.5rem;
}

.task-title {
  font-weight: 500;
  color: #2c5aa0;
  margin-bottom: 0.5rem;
  font-size: 1.1rem;
}

.task-priority {
  padding: 0.25rem 0.75rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 500;
}

.priority-high {
  background: #e74c3c;
  color: white;
}

.priority-medium {
  background: #f39c12;
  color: white;
}

.priority-low {
  background: #2c5aa0;
  color: white;
}

.task-description {
  color: #666;
  margin-bottom: 1rem;
  line-height: 1.5;
}

.task-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.9rem;
  color: #888;
}

.task-status {
  padding: 0.25rem 0.75rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 500;
}

.status-created {
  background: #e3f2fd;
  color: #2c5aa0;
}

.status-in-progress {
  background: #fff3e0;
  color: #f39c12;
}

.status-completed {
  background: #e8f5e8;
  color: #27ae60;
}

/* Modal Styles */
.modal {
  display: none;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  z-index: 1000;
  align-items: center;
  justify-content: center;
}

.modal.show {
  display: flex;
}

.modal-content {
  background: white;
  border-radius: 8px;
  width: 90%;
  max-width: 600px;
  max-height: 90vh;
  overflow-y: auto;
}

.modal-header {
  padding: 1.5rem;
  border-bottom: 1px solid #e1e5e9;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-header h2 {
  color: #2c5aa0;
  font-weight: 500;
}

.close-btn {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: #666;
}

.modal-body {
  padding: 1.5rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: #555;
  font-weight: 500;
}

.form-group input,
.form-group select,
.form-group textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 1rem;
}

.form-group textarea {
  resize: vertical;
  min-height: 100px;
}

.modal-footer {
  padding: 1rem 1.5rem;
  border-top: 1px solid #e1e5e9;
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
}

.btn-outline {
  background: transparent;
  border: 1px solid #2c5aa0;
  color: #2c5aa0;
}

.comments-section {
  margin-top: 2rem;
  padding-top: 1rem;
  border-top: 1px solid #e1e5e9;
}

.comment {
  padding: 1rem 0;
  border-bottom: 1px solid #f0f0f0;
}

.comment:last-child {
  border-bottom: none;
}

.comment-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.5rem;
}

.comment-author {
  font-weight: 500;
  color: #2c5aa0;
}

.comment-time {
  color: #999;
  font-size: 0.8rem;
}

.comment-text {
  color: #333;
  line-height: 1.5;
}

.comment-input {
  display: flex;
  gap: 0.5rem;
  margin-top: 1rem;
}

.comment-input input {
  flex: 1;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
}

@media (max-width: 768px) {
  .main-container {
    grid-template-columns: 1fr;
  }
  
  .task-controls {
    flex-direction: column;
    align-items: stretch;
  }
  
  .view-toggle {
    margin-left: 0;
  }
  
  .task-header {
    flex-direction: column;
    gap: 0.5rem;
  }
  
  .task-meta {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.5rem;
  }
}
</style>