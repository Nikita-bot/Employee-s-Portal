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
  const issueDescription = ref('');
  const supportTasks = ref([]);

  const fetchSupportTasks = async () => {
    try {
      const response = await fetch('/api/v1/tasks/support');
      if (!response.ok) throw new Error('Ошибка получения задач');
      supportTasks.value = (await response.json()).task_list || [];
    } catch (error) {
      console.error('Ошибка при получении задач:', error);
      alert('Не удалось загрузить задачи поддержки');
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
      
      if (!response.ok) throw new Error('Ошибка создания задачи');
      
      await fetchSupportTasks();
      alert('Ваш запрос отправлен в техническую поддержку');
      
      selectedIssue.value = '';
      issueDescription.value = '';
    } catch (error) {
      console.error('Ошибка создания задачи:', error);
      alert('Не удалось отправить запрос в поддержку');
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
}

.support-header {
  margin: 30px 0;
  text-align: center;
  width: 100%;
}

.support-header h1 {
  color: #5662DE;
  font-size: 28px;
}

.support-form-wrapper {
  width: 100%;
  display: flex;
  justify-content: center;
}

.support-form-container {
  background-color: white;
  border-radius: 10px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  width: 60vw; 
  box-sizing: border-box; 
}

.support-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
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
  padding: 10px 12px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 14px;
  transition: border-color 0.2s;
}

.form-select {
  height: 40px;
}

.form-textarea {
  min-height: 120px;
  resize: vertical;
}

.form-select:focus, .form-textarea:focus {
  outline: none;
  border-color: #5662DE;
}

.submit-btn {
  background-color: #5662DE;
  color: white;
  border: none;
  padding: 10px 16px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 16px;
  align-self: flex-end;
  transition: background-color 0.2s;
}

.submit-btn:hover {
  background-color: #454fc4;
}
</style>