<template>
  <div class="support-page">
    <div class="support-header">
      <h1>Техническая поддержка</h1>
    </div>
    <div class="support-form-wrapper">
      <div class="support-form-container">
        <div class="support-form">
          <div class="form-group">
            <label for="issue-type">Тип проблемы</label>
            <select id="issue-type" v-model="selectedIssue" class="form-select">
              <option disabled value="">Выберите тип проблемы</option>
              <option v-for="issue in supportTasks" :key="issue.id" :value="issue.id">
                {{ issue.name }}
              </option>
            </select>
          </div>
          
          <div class="form-group">
            <label for="priority">Приоритет</label>
            <select id="priority" v-model="selectedPriority" class="form-select">
              <option v-for="priority in priorityOptions" :key="priority.value" :value="priority.value">
                {{ priority.text }}
              </option>
            </select>
          </div>
          
          <div class="form-group">
            <label for="issue-description">Описание проблемы</label>
            <textarea 
              id="issue-description" 
              v-model="issueDescription" 
              class="form-textarea" 
              placeholder="Опишите вашу проблему подробно..."
            ></textarea>
          </div>
          
          <button class="submit-btn" @click="submitIssue">
            Отправить запрос
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
  import { ref, onMounted } from 'vue';
  import { useUserStore } from '@/stores/user';

  const userStore = useUserStore();

  const selectedIssue = ref('');
  const selectedPriority = ref(2); // По умолчанию нормальный приоритет
  const issueDescription = ref('');
  const supportTasks = ref([]);
  
  const priorityOptions = [
    { value: 1, text: 'Низкий' },
    { value: 2, text: 'Нормальный' },
    { value: 3, text: 'Высокий' }
  ];

  const fetchSupportTasks = async () => {
    try {
      const response = await fetch('/api/v1/taskList');
      if (!response.ok) throw new Error('Ошибка получения задач');
      
      const data = await response.json();
      supportTasks.value = (data.task_list || []).filter(task => task.type === "support");
      
    } catch (error) {
      console.error('Ошибка при получении задач поддержки:', error);
      alert('Не удалось загрузить типы проблем');
    }
  };

  const createSupportTask = async () => {
    if (!selectedIssue.value) {
      alert('Пожалуйста, выберите тип проблемы');
      return;
    }
    
    if (!issueDescription.value.trim()) {
      alert('Пожалуйста, опишите вашу проблему');
      return;
    }
    
    try {
      const taskToCreate = {
        task_id: selectedIssue.value,
        initiator: userStore.userData.id,
        description: issueDescription.value,
        priority: selectedPriority.value,
        status: 0, 
        create_date: new Date().toISOString().split('T')[0] 
      };
      
      const response = await fetch('/api/v1/tasks', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(taskToCreate)
      });
      
      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(errorData.message || `Ошибка HTTP: ${response.status}`);
      }
      
      alert('Ваш запрос отправлен в техническую поддержку');
      
      // Сброс формы
      selectedIssue.value = '';
      issueDescription.value = '';
      selectedPriority.value = 2; // Сброс к нормальному приоритету
      
    } catch (error) {
      console.error('Ошибка создания задачи:', error);
      alert(`Не удалось отправить запрос в поддержку: ${error.message}`);
    }
  };

  const submitIssue = () => {
    createSupportTask();
  };

  onMounted(() => {
    fetchSupportTasks();
  });
</script>

<style scoped>
.support-page {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f7fb;
  padding: 20px;
}

.support-header {
  margin: 30px 0;
  text-align: center;
  width: 100%;
}

.support-header h1 {
  color: #5662DE;
  font-size: 28px;
  font-weight: 600;
}

.support-form-wrapper {
  width: 100%;
  display: flex;
  justify-content: center;
  flex: 1;
}

.support-form-container {
  background-color: white;
  border-radius: 12px;
  padding: 30px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  width: 100%;
  max-width: 600px;
  box-sizing: border-box; 
  height: 520px;
}

.support-form {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group label {
  color: #333;
  font-size: 16px;
  font-weight: 500;
}

.form-select, .form-textarea {
  padding: 12px 16px;
  border: 1px solid #ddd;
  border-radius: 8px;
  font-size: 14px;
  transition: all 0.2s ease;
  font-family: inherit;
}

.form-select {
  height: 48px;
  background-color: #fff;
}

.form-textarea {
  min-height: 140px;
  resize: vertical;
  line-height: 1.5;
}

.form-select:focus, .form-textarea:focus {
  outline: none;
  border-color: #5662DE;
  box-shadow: 0 0 0 3px rgba(86, 98, 222, 0.1);
}

.form-select:hover, .form-textarea:hover {
  border-color: #bbb;
}

.submit-btn {
  background-color: #5662DE;
  color: white;
  border: none;
  padding: 14px 24px;
  border-radius: 8px;
  cursor: pointer;
  font-size: 16px;
  font-weight: 500;
  align-self: flex-end;
  transition: all 0.2s ease;
  min-width: 160px;
}

.submit-btn:hover {
  background-color: #454fc4;
  transform: translateY(-1px);
  box-shadow: 0 4px 8px rgba(86, 98, 222, 0.2);
}

.submit-btn:active {
  transform: translateY(0);
  box-shadow: 0 2px 4px rgba(86, 98, 222, 0.2);
}

/* Адаптивность для мобильных устройств */
@media (max-width: 768px) {
  .support-page {
    padding: 16px;
  }
  
  .support-header {
    margin: 20px 0;
  }
  
  .support-header h1 {
    font-size: 24px;
  }
  
  .support-form-container {
    padding: 24px 20px;
  }
  
  .support-form {
    gap: 20px;
  }
  
  .submit-btn {
    align-self: stretch;
    width: 100%;
  }
}
</style>