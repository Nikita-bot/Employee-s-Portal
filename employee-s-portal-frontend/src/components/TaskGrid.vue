<template>
  <div class="mainContent">
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

    <div class="task-grid">
      
      <div class="grid-header">ID</div>
      <div class="grid-header">
        {{ activeTab === 'assigned' ? 'Автор' : 'Исполнитель' }}
      </div>
      <div class="grid-header">Описание</div>
      <div class="grid-header">Статус</div>
      <div class="grid-header">Дата создания</div>
      <div class="grid-header">Дата назначения</div>

      
      <template v-for="task in displayedTasks" :key="task.id">
        <div class="grid-item">{{ task.id }}</div>
        <div class="grid-item">
          {{ activeTab === 'assigned' ? task.author : task.executor }}
        </div>
        <div class="grid-item">{{ task.description }}</div>
        <div class="grid-item">{{ task.status }}</div>
        <div class="grid-item">{{ task.createdDate }}</div>
        <div class="grid-item">{{ task.assignedDate }}</div>
      </template>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';

const activeTab = ref('assigned');

// Пример данных
const tasksData = {
  assigned: [
    {
      id: 32156,
      author: 'Петров П.П.', // Автор задачи (кто создал)
      executor: 'Сулима Р.И.', // Исполнитель (кому назначена)
      description: 'Добавить пользователя в МИС',
      status: 'Выполнено',
      createdDate: '29.04.2025',
      assignedDate: '29.04.2025'
    },
    {
      id: 32157,
      author: 'Сидоров С.С.',
      executor: 'Иванов И.И.',
      description: 'Обновить регламент',
      status: 'В работе',
      createdDate: '01.05.2025',
      assignedDate: '02.05.2025'
    }
  ],
  created: [
    {
      id: 32158,
      author: 'Иванов И.И.', // Это я (автор)
      executor: 'Петров П.П.', // Кому я назначил
      description: 'Проверить отчет',
      status: 'Новая',
      createdDate: '03.05.2025',
      assignedDate: '03.05.2025'
    }
  ]
};

// Вычисляемое свойство для отображаемых задач
const displayedTasks = computed(() => {
  return activeTab.value === 'assigned' ? tasksData.assigned : tasksData.created;
});
</script>

<style scoped>

    .mainContent {
        margin-top: 40px;
        margin-left: 40px;
        width: 100%;
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
</style>