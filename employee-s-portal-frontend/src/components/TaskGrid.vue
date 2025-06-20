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
          {{ activeTab === 'assigned' ? formatFullName(task.initiator.surname, task.initiator.name, task.initiator.patronymic) : formatFullName(task.executor.surname, task.executor.name, task.executor.patronymic) }}
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

    <TaskModal
      v-if="isTaskModalOpen"
      :task="selectedTask"
      :is-creator="activeTab === 'created'"
      :is-executor="activeTab === 'assigned'"
      @close="closeTaskModal"
      @delete="handleDeleteTask"
      @complete="handleCompleteTask"
      @add-comment="handleAddComment"
      @refresh-tasks="fetchTasks"
    />

    <!-- <CreateTaskModal
      v-if="isCreateModalOpen"
      :task-types="taskTypes"
      :department-users="departmentUsers"
      @close="closeCreateModal"
      @save="handleCreateTask"
      @task-type-changed="handleTaskTypeChange"
    /> -->
    <CreateTaskModal
      v-if="isCreateModalOpen"
      :task-types="taskTypes"
      @close="closeCreateModal"
      @save="handleCreateTask"
    />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useUserStore } from '@/stores/user';
import TaskModal from './TaskModal.vue';
import CreateTaskModal from './CreateTaskModal.vue';

const userStore = useUserStore();

// Состояния
const activeTab = ref('assigned');
const selectedTask = ref(null);
const sortField = ref('');
const sortDirection = ref('asc');
const isTaskModalOpen = ref(false);
const isCreateModalOpen = ref(false);
const isLoading = ref(false);

// Данные
const tasks = ref({
  assigned: [],
  created: []
});
const taskTypes = ref([]);
// const departmentUsers = ref([]);

const displayedTasks = computed(() => {
  const taskList = activeTab.value === 'assigned' 
    ? tasks.value.assigned 
    : tasks.value.created;
  
  if (sortField.value) {
    return [...taskList].sort((a, b) => {
      const valueA = a[sortField.value];
      const valueB = b[sortField.value];
      
      if (valueA < valueB) return sortDirection.value === 'asc' ? -1 : 1;
      if (valueA > valueB) return sortDirection.value === 'asc' ? 1 : -1;
      return 0;
    });
  }
  
  return taskList;
});

const sortBy = (field) => {
  if (sortField.value === field) {
    sortDirection.value = sortDirection.value === 'asc' ? 'desc' : 'asc';
  } else {
    sortField.value = field;
    sortDirection.value = 'asc';
  }
};

const fetchTasks = async () => {
  isLoading.value = true;
  try {
    const response = await fetch(`http://localhost:8080/api/v1/tasks/user/${userStore.userData.id}`);
    if (!response.ok) throw new Error('Ошибка загрузки задач');
    
    const data = await response.json();
    tasks.value.assigned = data.executor || [];
    tasks.value.created = data.initiator || [];
    
    // Форматируем даты
    tasks.value.assigned.forEach(task => {
      task.createdDate = task.create_date;
      task.assignedDate = task.execute_date ? task.execute_date : '-';
      task.status = getStatusText(task.status);
    });
    
    tasks.value.created.forEach(task => {
      task.createdDate = task.create_date;
      task.assignedDate = task.execute_date ? task.execute_date : '-';
      task.status = getStatusText(task.status);
    });
    
  } catch (error) {
    console.error('Ошибка загрузки задач:', error);
    alert('Не удалось загрузить задачи');
  } finally {
    isLoading.value = false;
  }
};

const fetchTaskTypes = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/v1/taskList');
    if (!response.ok) throw new Error('Ошибка загрузки типов задач');
    taskTypes.value = (await response.json()).task_list || [];
  } catch (error) {
    console.error('Ошибка загрузки типов задач:', error);
    alert('Не удалось загрузить типы задач');
  }
};

// const fetchDepartmentUsers = async (departmentId) => {
//   try {
//     const response = await fetch(`http://localhost:8080/api/v1/depatments/user/${departmentId}`);
//     if (!response.ok) throw new Error('Ошибка загрузки пользователей отдела');
//     departmentUsers.value = (await response.json()).users || [];
//   } catch (error) {
//     console.error('Ошибка загрузки пользователей отдела:', error);
//     alert('Не удалось загрузить пользователей отдела');
//   }
// };

const openTaskModal = async (task) => {
  isLoading.value = true;
  try {
    const response = await fetch(`http://localhost:8080/api/v1/tasks/${task.id}`);
    if (!response.ok) throw new Error('Ошибка загрузки задачи');
    
    const fullTask = (await response.json()).task;
    selectedTask.value = {
      ...fullTask,
      createdDate: fullTask.create_date,
      assignedDate: fullTask.execute_date ? fullTask.execute_date : '-',
      status: getStatusText(fullTask.status),
      comments: sortCommentsByDate(await fetchComments(fullTask.id), true) // Здесь можно добавить загрузку комментариев
    };
    
    isTaskModalOpen.value = true;
  } catch (error) {
    console.error('Ошибка загрузки задачи:', error);
    alert('Не удалось загрузить задачу');
  } finally {
    isLoading.value = false;
  }
};

function sortCommentsByDate(comments, ascending = false) {
  return comments.sort((a, b) => {
    const parseDate = (dateStr) => {
      const [datePart, timePart] = dateStr.split(', ');
      const [day, month, year] = datePart.split('.');
      const [hours, minutes, seconds] = timePart.split(':');
      return new Date(`${year}-${month}-${day}T${hours}:${minutes}:${seconds}`);
    };

    const dateA = parseDate(a.creation_date);
    const dateB = parseDate(b.creation_date);
    
    return ascending ? dateA - dateB : dateB - dateA;
  });
}

const closeTaskModal = () => {
  isTaskModalOpen.value = false;
  selectedTask.value = null;
};

const handleDeleteTask = async () => {
  if (!selectedTask.value) return;
  
  try {
    const response = await fetch(`http://localhost:8080/api/v1/tasks/${selectedTask.value.id}`, {
      method: 'DELETE'
    });
    
    if (!response.ok) throw new Error('Ошибка удаления задачи');
    
    await fetchTasks();
    closeTaskModal();
    alert('Задача успешно удалена');
  } catch (error) {
    console.error('Ошибка удаления задачи:', error);
    alert('Не удалось удалить задачу');
  }
};

const handleCompleteTask = async () => {
  if (!selectedTask.value) return;
  
  try {
    const response = await fetch(`http://localhost:8080/api/v1/tasks/${selectedTask.value.id}`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
      }
    });
    
    if (!response.ok) throw new Error('Ошибка завершения задачи');
    
    await fetchTasks();
    closeTaskModal();
    alert('Задача успешно завершена');
  } catch (error) {
    console.error('Ошибка завершения задачи:', error);
    alert('Не удалось завершить задачу');
  }
};

const fetchComments = async (id) =>{
  try {
    const response = await fetch(`http://localhost:8080/api/v1/comments/${id}`);
    if (!response.ok) throw new Error('Ошибка загрузки комментариев задач');
    const data = await response.json();
    return data.comments || [];
  } catch (error) {
    console.error('Ошибка загрузки типов задач:', error);
    alert('Не удалось загрузить типы задач');
    return [];
  }
}

const handleAddComment = async (commentText) => {
  if (!selectedTask.value || !commentText.trim()) return;
  
  try {
    const newCommentPost = {
      author: userStore.userData.id,
      text: commentText,
      date: new Date().toLocaleString()
    };

    const newComment = {
      author: {
        id: userStore.userData.id,
        name: userStore.userData.name,
        surname: userStore.userData.surname,
        patronymic: userStore.userData.patronymic
      },
      comment: commentText,
      creation_date: new Date().toLocaleString()
    }

    
    const response = await fetch('http://localhost:8080/api/v1/comments', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        user_task_id: selectedTask.value.id, // предполагается, что id задачи хранится в selectedTask.value.id
        author_id: newCommentPost.author,
        comment: newCommentPost.text,
        creation_date: newCommentPost.date
      })
    });

    if (!response.ok) {
      throw new Error('Ошибка при отправке комментария на сервер');
    }

    if (!selectedTask.value.comments) {
      selectedTask.value.comments = [];
    }

    selectedTask.value.comments.push(newComment);
    
  } catch (error) {
    console.error('Ошибка добавления комментария:', error);
    alert('Не удалось добавить комментарий');
  }
};

const createTask = () => {
  isCreateModalOpen.value = true;
};

const closeCreateModal = () => {
  isCreateModalOpen.value = false;
};

// const handleTaskTypeChange = (taskId) => {
//   const selectedType = taskTypes.value.find(t => t.id === taskId);
//   if (selectedType) {
//     fetchDepartmentUsers(selectedType.department_id);
//   }
// };

const handleCreateTask = async ({ task }) => {
  try {
    const taskToCreate = {
      ...task,
      initiator: userStore.userData.id,
    };
    
    const response = await fetch('http://localhost:8080/api/v1/tasks', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(taskToCreate)
    });
    
    if (!response.ok) throw new Error('Ошибка создания задачи');
    
    await fetchTasks();
    closeCreateModal();
    alert('Задача успешно создана');
  } catch (error) {
    console.error('Ошибка создания задачи:', error);
    alert('Не удалось создать задачу');
  }
};

// Вспомогательные функции

const getStatusText = (statusCode) => {
  const statuses = {
    0: 'В работе',
    1: 'Выполнено',
  };
  return statuses[statusCode] || 'Неизвестно';
};

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

onMounted(() => {
  fetchTasks();
  fetchTaskTypes();
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
  grid-template-columns: 7% 15% 30% 15% 15% 18%;
  gap: 1px; 
  border-radius: 15px;
  overflow: hidden;
  font-size: 20px;
  
  /* Вертикальная прокрутка */
  max-height: 80vh; /* или любая другая подходящая высота */
  overflow-y: auto;
  
  /* Стилизация скроллбара (невидимый) */
  scrollbar-width: none; /* Для Firefox */
  -ms-overflow-style: none; /* Для IE и Edge */
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


</style>