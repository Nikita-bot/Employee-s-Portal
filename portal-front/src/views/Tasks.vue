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
            <button class="view-toggle-btn" :class="{ active: currentView === 'executor' }" 
                    @click="switchView('executor')">Назначенные мне</button>
            <button class="view-toggle-btn" :class="{ active: currentView === 'initiator' }" 
                    @click="switchView('initiator')">Назначенные мной</button>
          </div>
        </div>

        <div class="tasks-grid">
          <div v-for="task in filteredTasks" :key="task.id" 
               :class="['task-item', `${getPriorityClass(task.priority)}-priority`, { completed: task.status === 2 }]"
               @click="openViewTaskModal(task.id)">
            <div class="task-header">
              <div>
                <div class="task-title">{{ task.task_name || 'Без названия' }}</div>
                <div class="task-description">{{ task.description }}</div>
              </div>
              <span :class="['task-priority', `priority-${getPriorityClass(task.priority)}`]">
                {{ getPriorityText(task.priority) }}
              </span>
            </div>
            <div class="task-meta">
              <span>{{ currentView === 'executor' ? 'Создатель: ' + getFullName(task.initiator) : 'Исполнитель: ' + getFullName(task.executor) }}</span>
              <span :class="['task-status', `status-${getStatusClass(task.status)}`]">
                {{ getStatusText(task.status) }}
              </span>
            </div>
            <div class="task-date">
              Создана: {{ formatDate(task.create_date) }}
              <span v-if="task.execute_date"> | Выполнена: {{ formatDate(task.execute_date) }}</span>
            </div>
          </div>
        </div>

        <div class="loading-spinner" v-if="isLoadingTasks">
          Загрузка задач...
        </div>

        <div class="no-tasks" v-if="!isLoadingTasks && filteredTasks.length === 0">
          Задачи не найдены
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
            <label for="taskType">Тип задачи *</label>
            <select id="taskType" v-model="newTask.task_id" required @change="onTaskTypeChange">
              <option value="">Выберите тип задачи</option>
              <option v-for="taskType in taskTypes" :key="taskType.id" :value="taskType.id">
                {{ taskType.name }}
              </option>
            </select>
          </div>
          <div class="form-group" v-if="selectedTaskType?.type !== 'mass'">
            <label for="taskExecutor">Исполнитель *</label>
            <select id="taskExecutor" v-model="newTask.executor" required>
              <option value="">Выберите исполнителя</option>
              <option value="0">Назначить автоматически</option>
              <option v-for="user in availableUsers" :key="user.id" :value="user.id">
                {{ user.surname }} {{ user.name }} {{ user.patronymic }} ({{ user.employee.position }})
              </option>
            </select>
          </div>
          
          <div class="form-group" v-else>
            <label>Исполнитель</label>
            <div class="mass-task-info">
              Для массовых задач исполнитель назначается автоматически
            </div>
          </div>

          <div class="form-group">
            <label for="taskDescription">Описание *</label>
            <textarea id="taskDescription" v-model="newTask.description" 
                      placeholder="Подробное описание задачи" required></textarea>
          </div>
          <div class="form-group">
           <label for="taskPriority">Приоритет *</label>
           <select id="taskPriority" v-model="newTask.priority" required>
             <option value="1">Низкий</option>
             <option value="2">Средний</option>
             <option value="3">Высокий</option>
           </select>
         </div>
         <div class="error-message" v-if="createTaskError">
           {{ createTaskError }}
         </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-outline" @click="closeCreateTaskModal">Отмена</button>
          <button class="btn btn-primary" @click="createNewTask" :disabled="isCreatingTask">
           {{ isCreatingTask ? 'Создание...' : 'Создать задачу' }}
         </button>
        </div>
      </div>
    </div>

    <!-- Модальное окно просмотра задачи -->
    <div class="modal" :class="{ show: showViewModal }" @click="closeViewTaskModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h2>{{ selectedTask.task_name || 'Без названия' }}</h2>
          <button class="close-btn" @click="closeViewTaskModal">&times;</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label>Создатель</label>
            <div>{{ getFullName(selectedTask.initiator) }}</div>
          </div>
          <div class="form-group">
            <label>Исполнитель</label>
            <div v-if="!canChangeExecutor">{{ getFullName(selectedTask.executor) }}</div>
            <select v-else v-model="selectedTask.executor_id" @change="updateTaskExecutor">
              <option value="">Выберите исполнителя</option>
              <option v-for="user in availableUsers" :key="user.id" :value="user.id">
                {{ user.surname }} {{ user.name }} {{ user.patronymic }} ({{ user.employee.position }})
              </option>
            </select>
          </div>
          <div class="form-group">
            <label>Приоритет</label>
            <div>{{ getPriorityText(selectedTask.priority) }}</div>
          </div>
          <div class="form-group">
            <label>Статус</label>
            <div v-if="!canChangeStatus">{{ getStatusText(selectedTask.status) }}</div>
            <select v-else v-model="selectedTask.status" @change="updateTaskStatus">
              <option value="0">Создана</option>
              <option value="1">В работе</option>
              <option value="2">Выполнена</option>
            </select>
          </div>
          <div class="form-group">
            <label>Описание</label>
            <div>{{ selectedTask.description }}</div>
          </div>
          <div class="form-group">
            <label>Дата создания</label>
            <div>{{ formatDate(selectedTask.create_date) }}</div>
          </div>
          <div class="form-group" v-if="selectedTask.execute_date">
            <label>Дата выполнения</label>
            <div>{{ formatDate(selectedTask.execute_date) }}</div>
          </div>
          <div class="form-group">
            <label>Комментарии</label>
            <div class="comments-section">
              <div v-if="taskComments.length === 0" class="no-comments">
                Нет комментариев
              </div>
              <div v-else class="comments-list">
                <div v-for="comment in taskComments" :key="comment.id" class="comment-item">
                  <div class="comment-header">
                    <span class="comment-author">{{ getFullName(comment.author) }}</span>
                    <span class="comment-date">{{ formatDateTime(comment.creation_date) }}</span>
                  </div>
                  <div class="comment-text">{{ comment.comment }}</div>
                </div>
              </div>
              
              <!-- Форма добавления комментария -->
              <div class="add-comment">
                <textarea v-model="newComment" placeholder="Добавить комментарий..." 
                          class="comment-textarea"></textarea>
                <button class="btn btn-primary btn-sm" @click="addComment" 
                        :disabled="!newComment.trim() || isAddingComment">
                  {{ isAddingComment ? 'Добавление...' : 'Добавить' }}
                </button>
              </div>
            </div>
          </div>
          <div class="task-actions" v-if="canChangeStatus || canChangeExecutor">
            <button class="btn btn-primary" @click="markAsInProgress" v-if="selectedTask.status === 0 && canChangeStatus">
              Взять в работу
            </button>
            <button class="btn btn-success" @click="markAsCompleted" v-if="selectedTask.status === 1 && canChangeStatus">
              Завершить задачу
            </button>
          </div>

          <div class="error-message" v-if="viewTaskError">
            {{ viewTaskError }}
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-danger" @click="deleteTask" v-if="canDeleteTask" :disabled="isDeletingTask">
            {{ isDeletingTask ? 'Удаление...' : 'Удалить задачу' }}
          </button>
          <button class="btn btn-primary" @click="closeViewTaskModal">Закрыть</button>
        </div>
      </div>
    </div>

    <AppFooter />
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue'
import { useUserStore } from '@/stores/user.js'
import AppHeader from '@/components/AppHeader.vue'
import AppSidebar from '@/components/AppSidebar.vue'
import AppFooter from '@/components/AppFooter.vue'

const userStore = useUserStore()

// Состояния UI
const showCreateModal = ref(false)
const showViewModal = ref(false)
const currentView = ref('executor')
const statusFilter = ref('all')

// Состояния загрузки
const isLoadingTasks = ref(false)
const isCreatingTask = ref(false)
const isLoadingTaskDetails = ref(false)
const isDeletingTask = ref(false)

// Состояния ошибок
const tasksError = ref('')
const createTaskError = ref('')
const viewTaskError = ref('')
const taskTypesError = ref('')
const availableUsersError = ref('')

// Данные
const tasks = ref([])
const taskTypes = ref([])
const availableUsers = ref([])

// Контроллеры для отмены запросов
const abortControllers = {
  tasks: null,
  taskTypes: null,
  availableUsers: null,
  taskDetails: null,
  createTask: null,
  updateExecutor: null,
  updateStatus: null,
  deleteTask: null
}

// Комментарии
const taskComments = ref([])
const newComment = ref('')
const isAddingComment = ref(false)
const isLoadingComments = ref(false)
const commentsError = ref('')

const selectedTask = reactive({
  id: 0,
  task_id: '',
  description: '',
  priority: 0,
  status: 0,
  initiator: {},
  executor: {},
  executor_id: 0,
  create_date: '',
  execute_date: ''
})

const newTask = reactive({
  task_id: '',
  description: '',
  priority: 2,
  executor: '',
  initiator: userStore.userId,
  branch_id: userStore.user?.user?.employee?.branch_id,
  create_date: new Date().toISOString().split('T')[0],
  status: 0
})

const loadTaskComments = async (taskId) => {
  commentsError.value = ''
  isLoadingComments.value = true

  try {
    const response = await fetch(`/api/v1/comments/${taskId}`)
    
    if (!response.ok) {
      throw new Error(`Failed to fetch comments: ${response.status}`)
    }

    const data = await response.json()
    taskComments.value = data.comments || []
  } catch (error) {
    console.error('Error loading comments:', error)
    commentsError.value = 'Ошибка при загрузке комментариев'
  } finally {
    isLoadingComments.value = false
  }
}

const addComment = async () => {
  if (!newComment.value.trim()) return

  isAddingComment.value = true
  commentsError.value = ''

  const commentData = {
    user_task_id: selectedTask.id,
    author_id: parseInt(userStore.userId),
    comment: newComment.value.trim(),
    creation_date: new Date().toISOString()
  }

  try {
    const response = await fetch('/api/v1/comments', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(commentData)
    })

    if (response.ok) {
      newComment.value = ''
      await loadTaskComments(selectedTask.id) // Перезагружаем комментарии
      if (typeof window.showNotification === 'function') {
        window.showNotification('Комментарий добавлен!', 'success')
      }
    } else {
      const errorText = await response.text()
      throw new Error(errorText || 'Ошибка при добавлении комментария')
    }
  } catch (error) {
    console.error('Error adding comment:', error)
    commentsError.value = error.message || 'Ошибка при добавлении комментария'
  } finally {
    isAddingComment.value = false
  }
}

const formatDateTime = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleString('ru-RU')
}

// Вычисляемые свойства
const filteredTasks = computed(() => {
  let filtered = tasks.value

  // Фильтрация по виду (назначенные мне/мной)
  if (currentView.value === 'executor') {
    filtered = filtered.filter(task => task.taskType === 'executor')
  } else {
    filtered = filtered.filter(task => task.taskType === 'initiator')
  }

  // Фильтрация по статусу
  if (statusFilter.value === 'active') {
    filtered = filtered.filter(task => task.status !== 2)
  } else if (statusFilter.value === 'completed') {
    // Для выполненных задач показываем их в обоих списках
    filtered = filtered.filter(task => task.status === 2)
  }

  return filtered
})

const selectedTaskType = computed(() => {
  return taskTypes.value.find(type => type.id === parseInt(newTask.task_id))
})

const canChangeExecutor = computed(() => {
  return (selectedTask.initiator?.id == userStore.userId || 
          selectedTask.executor?.id == userStore.userId) && 
          selectedTask.status != 2
})

const canChangeStatus = computed(() => {
  // Статус может быть изменен только исполнителем и только если задача не выполнена
  return selectedTask.executor?.id == userStore.userId && selectedTask.status !== 2
})

const canDeleteTask = computed(() => {
  
  return selectedTask.initiator?.id == userStore.userId && selectedTask.status === 0
})

// Вспомогательные функции
const getFullName = (user) => {
  if (!user) return 'Не указан'
  return `${user.surname || ''} ${user.name || ''} ${user.patronymic || ''}`.trim() || 'Не указан'
}

const getStatusClass = (status) => {
  const classes = {
    0: 'created',
    1: 'in-progress',
    2: 'completed'
  }
  return classes[status] || 'created'
}

const getPriorityClass = (priority) => {
  const classes = {
    1: 'low',
    2: 'medium',
    3: 'high'
  }
  return classes[priority] || 'medium'
}

const getStatusText = (status) => {
  const texts = {
    0: 'Создана',
    1: 'В работе',
    2: 'Выполнена'
  }
  return texts[status] || 'Создана'
}


const getPriorityText = (priority) => {
  const texts = {
    1: 'Низкий',
    2: 'Средний',
    3: 'Высокий'
  }
  return texts[priority] || 'Средний'
}

const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleDateString('ru-RU')
}

// Функции для работы с API
const loadTasks = async () => {
  if (!userStore.userId) return

  tasksError.value = ''
  isLoadingTasks.value = true

  try {
    // Отменяем предыдущий запрос, если он существует
    if (abortControllers.tasks) {
      abortControllers.tasks.abort()
    }

    // Создаем новый контроллер для этого запроса
    const controller = new AbortController()
    abortControllers.tasks = controller

    const response = await fetch(`/api/v1/tasks/user/${userStore.userId}`, {
      signal: controller.signal
    })

    if (!response.ok) {
      throw new Error(`Failed to fetch tasks: ${response.status} ${response.statusText}`)
    }

    const data = await response.json()
    
    // Объединяем задачи, созданные пользователем и назначенные ему
    const initiatorTasks = data.initiator || []
    const executorTasks = data.executor || []
    
    // Добавляем тип задачи для различия
    initiatorTasks.forEach(task => task.taskType = 'initiator')
    executorTasks.forEach(task => task.taskType = 'executor')
    
    tasks.value = [...initiatorTasks, ...executorTasks]
  } catch (error) {
    if (error.name !== 'AbortError') {
      console.error('Error loading tasks:', error)
      tasksError.value = error.message || 'Произошла ошибка при загрузке задач'
    }
  } finally {
    isLoadingTasks.value = false
  }
}

const loadTaskTypes = async () => {
  taskTypesError.value = ''

  try {
    if (abortControllers.taskTypes) {
      abortControllers.taskTypes.abort()
    }

    const controller = new AbortController()
    abortControllers.taskTypes = controller

    const response = await fetch('/api/v1/taskList', {
      signal: controller.signal
    })

    if (!response.ok) {
      throw new Error(`Failed to fetch task types: ${response.status} ${response.statusText}`)
    }

    const data = await response.json()
    taskTypes.value = data.task_list || []
  } catch (error) {
    if (error.name !== 'AbortError') {
      console.error('Error loading task types:', error)
      taskTypesError.value = error.message || 'Произошла ошибка при загрузке типов задач'
    }
  }
}

const loadAvailableUsers = async (taskId = null) => {
  availableUsersError.value = ''

  try {
    if (abortControllers.availableUsers) {
      abortControllers.availableUsers.abort()
    }

    const controller = new AbortController()
    abortControllers.availableUsers = controller

    let url = '/api/v1/taskList/users'
    if (taskId) {
      url += `/${taskId}`
    }

    const response = await fetch(url, {
      signal: controller.signal
    })

    if (!response.ok) {
      throw new Error(`Failed to fetch available users: ${response.status} ${response.statusText}`)
    }

    const data = await response.json()
    availableUsers.value = data.user_list || []
  } catch (error) {
    if (error.name !== 'AbortError') {
      console.error('Error loading available users:', error)
      availableUsersError.value = error.message || 'Произошла ошибка при загрузке списка пользователей'
    }
  }
}

const onTaskTypeChange = async () => {
  if (newTask.task_id) {
    await loadAvailableUsers(newTask.task_id)
  } else {
    availableUsers.value = []
  }
}

const switchView = (view) => {
  currentView.value = view
}

const filterTasks = () => {
  // Фильтрация происходит автоматически через computed свойство
}

const showCreateTaskModal = async () => {
  await loadTaskTypes()
  showCreateModal.value = true
}

const closeCreateTaskModal = () => {
  showCreateModal.value = false
  createTaskError.value = ''
  Object.assign(newTask, {
    task_id: '',
    description: '',
    priority: 2,
    executor: '',
    initiator: userStore.userId,
    branch_id: userStore.user?.user?.employee?.branch_id,
    create_date: new Date().toISOString().split('T')[0],
    status: 0
  })
  availableUsers.value = []
}

const createNewTask = async () => {
  // Проверка обязательных полей
  if (!newTask.task_id || !newTask.description ) {
    createTaskError.value = 'Пожалуйста, заполните все обязательные поля'
    return
  }

  createTaskError.value = ''
  isCreatingTask.value = true
  console.log(newTask.executor)

  const taskData = {
      task_id: parseInt(newTask.task_id),
      executor: selectedTaskType.value?.type === 'mass' ? 0 : parseInt(newTask.executor),
      initiator: parseInt(userStore.userId),
      description: newTask.description,
      status: 0, // Создана
      priority: parseInt(newTask.priority), // Высокий при срочной, иначе средний
      branch_id: userStore.user?.user?.employee?.branch_id,
      create_date: new Date().toISOString().split('T')[0]
    } 

  try {
    if (abortControllers.createTask) {
      abortControllers.createTask.abort()
    }

    const controller = new AbortController()
    abortControllers.createTask = controller

    console.log(newTask) 
    

    const response = await fetch('/api/v1/tasks', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(taskData)
    })

    if (response.ok) {
      await loadTasks()
      closeCreateTaskModal()
      if (typeof window.showNotification === 'function') {
        window.showNotification('Задача успешно создана!', 'success')
      }
    } else {
      const errorText = await response.text()
      throw new Error(errorText || 'Ошибка при создании задачи')
    }
  } catch (error) {
    if (error.name !== 'AbortError') {
      console.error('Error creating task:', error)
      createTaskError.value = error.message || 'Произошла ошибка при создании задачи'
    }
  } finally {
    isCreatingTask.value = false
  }
}

const openViewTaskModal = async (taskId) => {
  viewTaskError.value = ''
  isLoadingTaskDetails.value = true

  try {
    if (abortControllers.taskDetails) {
      abortControllers.taskDetails.abort()
    }

    const controller = new AbortController()
    abortControllers.taskDetails = controller

    const response = await fetch(`/api/v1/tasks/${taskId}`, {
      signal: controller.signal
    })

    if (!response.ok) {
      throw new Error(`Failed to fetch task details: ${response.status} ${response.statusText}`)
    }
    const data = await response.json()
    Object.assign(selectedTask, data.task)
    selectedTask.executor_id = selectedTask.executor?.id || 0
    
    // Загружаем комментарии и доступных пользователей параллельно
    await Promise.all([
      loadTaskComments(taskId),
      loadAvailableUsers(selectedTask.task_id)
    ])
    
    showViewModal.value = true
    showViewModal.value = true
  } catch (error) {
    if (error.name !== 'AbortError') {
      console.error('Error loading task details:', error)
      viewTaskError.value = error.message || 'Ошибка при загрузке данных задачи'
    }
  } finally {
    isLoadingTaskDetails.value = false
  }
}

const closeViewTaskModal = () => {
  showViewModal.value = false
  viewTaskError.value = ''
  showViewModal.value = false
  viewTaskError.value = ''
  commentsError.value = ''
  newComment.value = ''
  taskComments.value = []
  // Сброс значений selectedTask
  Object.assign(selectedTask, {
    id: 0,
    task_id: '',
    task_name: '',
    description: '',
    priority: 0,
    status: 0,
    initiator: {},
    executor: {},
    executor_id: 0,
    create_date: '',
    execute_date: ''
  })
}

const updateTaskExecutor = async () => {
  if (!selectedTask.executor_id) return

  viewTaskError.value = ''

  try {
    if (abortControllers.updateExecutor) {
      abortControllers.updateExecutor.abort()
    }

    const controller = new AbortController()
    abortControllers.updateExecutor = controller

    const response = await fetch(`/api/v1/tasks/executor/${selectedTask.id}`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        executor_id: selectedTask.executor_id
      }),
      signal: controller.signal
    })

    if (response.ok) {
      await loadTasks()
      // Обновляем selectedTask после изменения исполнителя
      const updatedTask = tasks.value.find(task => task.id === selectedTask.id)
      if (updatedTask) {
        Object.assign(selectedTask, updatedTask)
      }
      if (typeof window.showNotification === 'function') {
        window.showNotification('Исполнитель изменен!', 'success')
      }
    } else {
      const errorText = await response.text()
      throw new Error(errorText || 'Ошибка при изменении исполнителя')
    }
  } catch (error) {
    if (error.name !== 'AbortError') {
      console.error('Error updating executor:', error)
      viewTaskError.value = error.message || 'Ошибка при изменении исполнителя'
    }
  }
}

const updateTaskStatus = async () => {
  viewTaskError.value = ''

  try {
    if (abortControllers.updateStatus) {
      abortControllers.updateStatus.abort()
    }

    const controller = new AbortController()
    abortControllers.updateStatus = controller

    const response = await fetch(`/api/v1/tasks/${selectedTask.id}`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        status: parseInt(selectedTask.status)
      }),
      signal: controller.signal
    })

    if (response.ok) {
      await loadTasks()
      // Обновляем selectedTask после изменения статуса
      const updatedTask = tasks.value.find(task => task.id === selectedTask.id)
      if (updatedTask) {
        Object.assign(selectedTask, updatedTask)
      }
      if (typeof window.showNotification === 'function') {
        window.showNotification('Статус задачи обновлен!', 'success')
      }
    } else {
      const errorText = await response.text()
      throw new Error(errorText || 'Ошибка при обновлении статуса')
    }
  } catch (error) {
    if (error.name !== 'AbortError') {
      console.error('Error updating status:', error)
      viewTaskError.value = error.message || 'Ошибка при обновлении статуса'
    }
  }
}

const markAsInProgress = async () => {
  selectedTask.status = 1
  await updateTaskStatus()
}

const markAsCompleted = async () => {
  selectedTask.status = 2
  await updateTaskStatus()
}

const deleteTask = async () => {
  if (!confirm('Вы уверены, что хотите удалить эту задачу?')) {
    return
  }

  viewTaskError.value = ''
  isDeletingTask.value = true

  try {
    if (abortControllers.deleteTask) {
      abortControllers.deleteTask.abort()
    }

    const controller = new AbortController()
    abortControllers.deleteTask = controller

    const response = await fetch(`/api/v1/tasks/${selectedTask.id}`, {
      method: 'DELETE',
      signal: controller.signal
    })

    if (response.ok) {
      await loadTasks()
      closeViewTaskModal()
      if (typeof window.showNotification === 'function') {
        window.showNotification('Задача удалена!', 'success')
      }
    } else {
      const errorText = await response.text()
      throw new Error(errorText || 'Ошибка при удалении задачи')
    }
  } catch (error) {
    if (error.name !== 'AbortError') {
      console.error('Error deleting task:', error)
      viewTaskError.value = error.message || 'Ошибка при удалении задачи'
    }
  } finally {
    isDeletingTask.value = false
  }
}

// Жизненный цикл компонента
onMounted(() => {
  if (userStore.isAuthenticated) {
    loadTasks()
    loadTaskTypes()
  }
})

onUnmounted(() => {
  // Отменяем все активные запросы при размонтировании компонента
  Object.keys(abortControllers).forEach(key => {
    if (abortControllers[key]) {
      abortControllers[key].abort()
    }
  })
})
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

.task-item.low-priority {
  border-left-color: #2c5aa0;
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

.task-date {
  font-size: 0.8rem;
  color: #888;
  margin-top: 0.5rem;
}

.loading-spinner {
  text-align: center;
  padding: 2rem;
  color: #666;
}

.no-tasks {
  text-align: center;
  padding: 3rem;
  color: #999;
  font-style: italic;
}

.error-message {
  background: #fde8e8;
  color: #e74c3c;
  padding: 1rem;
  border-radius: 6px;
  margin-top: 1rem;
}

.btn-success {
  background: #27ae60;
  color: white;
}

.btn-success:hover {
  background: #219653;
}

.btn-danger {
  background: #e74c3c;
  color: white;
}

.btn-danger:hover {
  background: #c0392b;
}

.task-actions {
  display: flex;
  gap: 1rem;
  margin-top: 1rem;
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

.btn-outline:hover {
  background: #2c5aa0;
  color: white;
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

.comments-section {
  border: 1px solid #e1e5e9;
  border-radius: 6px;
  padding: 1rem;
  max-height: 300px;
  overflow-y: auto;
}

.comments-list {
  margin-bottom: 1rem;
}

.comment-item {
  padding: 0.75rem;
  border-bottom: 1px solid #f0f0f0;
}

.comment-item:last-child {
  border-bottom: none;
}

.comment-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.5rem;
  font-size: 0.85rem;
}

.comment-author {
  font-weight: 500;
  color: #2c5aa0;
}

.comment-date {
  color: #888;
}

.comment-text {
  color: #333;
  line-height: 1.4;
}

.no-comments {
  text-align: center;
  color: #888;
  font-style: italic;
  padding: 1rem;
}

.add-comment {
  margin-top: 1rem;
  padding-top: 1rem;
  border-top: 1px solid #e1e5e9;
}

.comment-textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  resize: vertical;
  min-height: 80px;
  margin-bottom: 0.5rem;
}

.btn-sm {
  padding: 0.4rem 0.8rem;
  font-size: 0.85rem;
}
</style>