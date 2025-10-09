<template>
  <div class="wrapper">
    <header>
      <div class="logo">Городская Больница</div>
      <div class="user-menu" @click="toggleDropdown">
        <span>{{ userStore.userData?.name }}</span>
        <img :src="userStore.userData?.avatar" alt="Аватар">
        <div class="dropdown-menu" :class="{ show: dropdownOpen }">
          <router-link to="/profile" class="dropdown-item" @click="closeDropdown">Мой профиль</router-link>
          <a href="#" class="dropdown-item" @click="closeDropdown">Сменить пароль</a>
          <a href="#" class="dropdown-item" @click="closeDropdown">Привязать Телеграмм</a>
          <div class="dropdown-divider"></div>
          <a href="#" class="dropdown-item" @click="logout">Выйти</a>
        </div>
      </div>
    </header>

    <div class="main-container">
      <sidebar>
        <ul class="nav-links">
          <li><router-link to="/">Главная</router-link></li>
          <li><router-link to="/tasks" class="active">Задачи</router-link></li>
          <li><a href="#">ЭДО</a></li>
          <li><a href="#">База знаний</a></li>
          <li><a href="#">Порталы</a></li>
          <li><router-link to="/news">Новости</router-link></li>
          <li><router-link to="/support">Поддержка</router-link></li>
        </ul>
      </sidebar>

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
            <button 
              class="view-toggle-btn" 
              :class="{ active: currentView === 'assigned-to-me' }"
              @click="switchView('assigned-to-me')"
            >
              Назначенные мне
            </button>
            <button 
              class="view-toggle-btn" 
              :class="{ active: currentView === 'assigned-by-me' }"
              @click="switchView('assigned-by-me')"
            >
              Назначенные мной
            </button>
          </div>
        </div>

        <div class="tasks-grid">
          <div 
            v-for="task in filteredTasks" 
            :key="task.id"
            class="task-item" 
            :class="[`${task.priority}-priority`, { completed: task.status === 'completed' }]"
            @click="openViewTaskModal(task.id)"
          >
            <div class="task-header">
              <div>
                <div class="task-title">{{ task.title }}</div>
                <div class="task-description">{{ task.description }}</div>
              </div>
              <span class="task-priority" :class="`priority-${task.priority}`">
                {{ priorityText[task.priority] }}
              </span>
            </div>
            <div class="task-meta">
              <span>{{ currentView === 'assigned-to-me' ? 'Создатель: ' + task.creator : 'Исполнитель: ' + task.assignee }}</span>
              <span class="task-status" :class="`status-${task.status}`">
                {{ statusText[task.status] }}
              </span>
            </div>
          </div>
        </div>
      </main>
    </div>

    <!-- Модальное окно создания задачи -->
    <div class="modal" :class="{ show: showCreateModal }">
      <div class="modal-content">
        <div class="modal-header">
          <h2>Создание новой задачи</h2>
          <button class="close-btn" @click="closeCreateTaskModal">&times;</button>
        </div>
        <div class="modal-body">
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
            <textarea 
              id="taskDescription" 
              v-model="newTask.description" 
              placeholder="Подробное описание задачи"
            ></textarea>
          </div>
          <div class="form-group">
            <label for="taskPriority">Приоритет</label>
            <select id="taskPriority" v-model="newTask.priority">
              <option value="low">Низкий</option>
              <option value="medium">Средний</option>
              <option value="high">Высокий</option>
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
    <div class="modal" :class="{ show: showViewModal }">
      <div class="modal-content">
        <div class="modal-header">
          <h2>{{ selectedTask?.title }}</h2>
          <button class="close-btn" @click="closeViewTaskModal">&times;</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label>Создатель</label>
            <div>{{ selectedTask?.creator }}</div>
          </div>
          <div class="form-group">
            <label>Исполнитель</label>
            <select v-model="selectedTaskAssignee">
              <option value="petrov">Петров М.И. (Хирург)</option>
              <option value="sidorova">Сидорова А.В. (Медсестра)</option>
              <option value="kozlov">Козлов Д.С. (Администратор)</option>
              <option value="ivanov">Иванов А.С. (Кардиолог)</option>
            </select>
          </div>
          <div class="form-group">
            <label>Приоритет</label>
            <div>{{ priorityText[selectedTask?.priority] }}</div>
          </div>
          <div class="form-group">
            <label>Статус</label>
            <select v-model="selectedTaskStatus" @change="updateTaskStatus">
              <option value="created">Создана</option>
              <option value="in-progress">В работе</option>
              <option value="completed">Выполнена</option>
            </select>
          </div>
          <div class="form-group">
            <label>Описание</label>
            <div>{{ selectedTask?.description }}</div>
          </div>
          
          <div class="comments-section">
            <h3>Комментарии</h3>
            <div id="taskComments">
              <div v-for="comment in selectedTask?.comments" :key="comment.time" class="comment">
                <div class="comment-header">
                  <span class="comment-author">{{ comment.author }}</span>
                  <span class="comment-time">{{ comment.time }}</span>
                </div>
                <div class="comment-text">{{ comment.text }}</div>
              </div>
            </div>
            <div class="comment-input">
              <input type="text" v-model="newComment" placeholder="Введите комментарий...">
              <button class="btn btn-primary" @click="addComment">Отправить</button>
            </div>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-primary" @click="closeViewTaskModal">Закрыть</button>
        </div>
      </div>
    </div>

    <footer>
      <p>© 2023 Городская Больница. Корпоративная информационная система. Версия 2.1</p>
    </footer>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()

// Состояние
const dropdownOpen = ref(false)
const showCreateModal = ref(false)
const showViewModal = ref(false)
const currentView = ref('assigned-to-me')
const statusFilter = ref('all')
const selectedTaskId = ref(null)
const newComment = ref('')
const selectedTaskAssignee = ref('')
const selectedTaskStatus = ref('')

// Данные
const tasks = ref([
  {
    id: 1,
    title: "Консилиум по пациенту Петров И.В.",
    description: "Обсуждение результатов МРТ и плана лечения. Присутствие обязательно.",
    priority: "high",
    status: "created",
    creator: "Главный врач",
    assignee: "Иванов А.С.",
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
    creator: "Иванов А.С.",
    assignee: "Сидорова А.В.",
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
    assignee: "Иванов А.С.",
    type: "medical",
    comments: [
      { author: "Иванов А.С.", text: "Осмотр завершен, все пациенты стабильны", time: "06.10.2023 09:30" }
    ]
  }
])

const newTask = reactive({
  type: 'medical',
  description: '',
  priority: 'medium'
})

// Тексты для отображения
const statusText = {
  'created': 'Создана',
  'in-progress': 'В работе', 
  'completed': 'Выполнена'
}

const priorityText = {
  'high': 'Высокий',
  'medium': 'Средний',
  'low': 'Низкий'
}

// Вычисляемые свойства
const filteredTasks = computed(() => {
  let filtered = tasks.value
  
  // Фильтрация по виду (назначенные мне/мной)
  if (currentView.value === 'assigned-to-me') {
    filtered = filtered.filter(task => task.assignee === userStore.userData?.name)
  } else {
    filtered = filtered.filter(task => task.creator === userStore.userData?.name)
  }
  
  // Фильтрация по статусу
  if (statusFilter.value === 'active') {
    filtered = filtered.filter(task => task.status !== 'completed')
  } else if (statusFilter.value === 'completed') {
    filtered = filtered.filter(task => task.status === 'completed')
  }
  
  return filtered
})

const selectedTask = computed(() => {
  return tasks.value.find(task => task.id === selectedTaskId.value)
})

// Методы
const toggleDropdown = () => {
  dropdownOpen.value = !dropdownOpen.value
}

const closeDropdown = () => {
  dropdownOpen.value = false
}

const logout = () => {
  userStore.clearUserData()
  router.push('/login')
  closeDropdown()
}

const showCreateTaskModal = () => {
  showCreateModal.value = true
}

const closeCreateTaskModal = () => {
  showCreateModal.value = false
  // Сброс формы
  Object.assign(newTask, {
    type: 'medical',
    description: '',
    priority: 'medium'
  })
}

const createNewTask = () => {
  if (newTask.description) {
    const task = {
      id: tasks.value.length + 1,
      title: getTaskTitle(newTask.type),
      description: newTask.description,
      type: newTask.type,
      priority: newTask.priority,
      status: 'created',
      creator: userStore.userData?.name,
      assignee: 'Петров М.И.',
      comments: []
    }
    
    tasks.value.push(task)
    closeCreateTaskModal()
  }
}

const getTaskTitle = (type) => {
  const titles = {
    'medical': 'Медицинская задача',
    'administrative': 'Административная задача',
    'technical': 'Техническая задача',
    'other': 'Задача'
  }
  return titles[type]
}

const openViewTaskModal = (taskId) => {
  selectedTaskId.value = taskId
  selectedTaskAssignee.value = getAssigneeValue(selectedTask.value.assignee)
  selectedTaskStatus.value = selectedTask.value.status
  showViewModal.value = true
}

const closeViewTaskModal = () => {
  showViewModal.value = false
  selectedTaskId.value = null
  newComment.value = ''
}

const getAssigneeValue = (assignee) => {
  const mapping = {
    'Иванов А.С.': 'ivanov',
    'Петров М.И.': 'petrov', 
    'Сидорова А.В.': 'sidorova',
    'Козлов Д.С.': 'kozlov'
  }
  return mapping[assignee] || 'ivanov'
}

const switchView = (view) => {
  currentView.value = view
}

const filterTasks = () => {
  // Фильтрация происходит в computed свойстве
}

const updateTaskStatus = () => {
  if (selectedTask.value) {
    selectedTask.value.status = selectedTaskStatus.value
  }
}

const addComment = () => {
  if (newComment.value && selectedTask.value) {
    const comment = {
      author: userStore.userData?.name,
      text: newComment.value,
      time: new Date().toLocaleString('ru-RU')
    }
    selectedTask.value.comments.push(comment)
    newComment.value = ''
  }
}

// Обработчики событий
const handleClickOutside = (event) => {
  const userMenu = document.querySelector('.user-menu')
  if (userMenu && !userMenu.contains(event.target)) {
    dropdownOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
/* Стили из вашего tasks.html */
.wrapper {
  display: grid;
  grid-template-rows: auto 1fr auto;
  min-height: 100vh;
}

header {
  background-color: #fff;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  padding: 1rem 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: relative;
}

.logo {
  font-size: 1.5rem;
  font-weight: 300;
  color: #2c5aa0;
}

.user-menu {
  display: flex;
  align-items: center;
  gap: 1rem;
  cursor: pointer;
  position: relative;
}

.user-menu img {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  right: 0;
  background: white;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
  border-radius: 8px;
  padding: 0.5rem 0;
  min-width: 200px;
  display: none;
  z-index: 1000;
}

.dropdown-menu.show {
  display: block;
}

.dropdown-item {
  padding: 0.75rem 1.5rem;
  display: block;
  text-decoration: none;
  color: #333;
  transition: background 0.3s ease;
}

.dropdown-item:hover {
  background: #f8f9fa;
}

.dropdown-divider {
  height: 1px;
  background: #e1e5e9;
  margin: 0.25rem 0;
}

.main-container {
  display: grid;
  grid-template-columns: 250px 1fr;
  gap: 0;
}

sidebar {
  background-color: #2c5aa0;
  color: white;
  padding: 2rem 0;
}

.nav-links {
  list-style: none;
}

.nav-links li a {
  display: block;
  color: #e0e0e0;
  padding: 0.8rem 2rem;
  text-decoration: none;
  transition: all 0.3s ease;
}

.nav-links li a:hover,
.nav-links li a.router-link-active {
  background-color: rgba(255, 255, 255, 0.1);
  color: white;
  border-left: 4px solid #4fc3f7;
}

.content-area {
  padding: 2rem;
  background-color: #fff;
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

footer {
  background-color: #1a1a1a;
  color: #999;
  text-align: center;
  padding: 1.5rem;
  font-size: 0.9rem;
}

@media (max-width: 768px) {
  .main-container {
    grid-template-columns: 1fr;
  }
  sidebar {
    display: none;
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