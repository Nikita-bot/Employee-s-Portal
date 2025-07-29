<template>
  <div class="modal-overlay" v-if="isOpen" @click.self="closeModal">
    <div class="news-modal">
      <div class="modal-header">
        <h2>Создание новости</h2>
        <button class="close-btn" @click="closeModal">×</button>
      </div>
      
      <div class="modal-content">
        <div class="form-group">
          <label for="news-title">Заголовок</label>
          <input 
            id="news-title" 
            v-model="newNews.title" 
            class="form-input" 
            placeholder="Введите заголовок новости"
          >
        </div>
        
        <div class="form-group">
          <label for="news-content">Содержание</label>
          <textarea 
            id="news-content" 
            v-model="newNews.content" 
            class="form-textarea" 
            placeholder="Введите текст новости..."
            rows="6"
          ></textarea>
        </div>
        
        <div class="form-actions">
          <button class="cancel-btn" @click="closeModal">Отмена</button>
          <button class="submit-btn" @click="createNews">Опубликовать</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, defineProps, defineEmits } from 'vue';
import { useUserStore } from '@/stores/user';

const userStore = useUserStore();
const props = defineProps({
  isOpen: Boolean
});

const emits = defineEmits(['close', 'created']);

const newNews = ref({
  title: '',
  content: '',
  author: userStore.userData.id,
  date: new Date().toLocaleString()
});

const createNews = async () => {
  if (!newNews.value.title.trim()) {
    alert('Введите заголовок новости');
    return;
  }
  
  if (!newNews.value.content.trim()) {
    alert('Введите содержание новости');
    return;
  }

  try {
    const response = await fetch('/api/v1/news', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(newNews.value)
    });

    if (!response.ok) throw new Error('Ошибка создания новости');
    
    emits('created');
    closeModal();
    alert('Новость успешно создана');
  } catch (error) {
    console.error('Ошибка при создании новости:', error);
    alert('Не удалось создать новость');
  }
};

const closeModal = () => {
  newNews.value = {
    title: '',
    content: '',
    author: userStore.userData.id,
    date: new Date().toLocaleString()
  };
  emits('close');
};
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

.news-modal {
  background-color: white;
  border-radius: 10px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
  width: 600px;
  max-width: 90vw;
  max-height: 90vh;
  overflow-y: auto;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  border-bottom: 1px solid #eee;
}

.modal-header h2 {
  color: #5662DE;
  font-size: 22px;
  margin: 0;
}

.close-btn {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: #888;
  padding: 0 10px;
}

.close-btn:hover {
  color: #333;
}

.modal-content {
  padding: 20px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  color: #333;
  font-weight: 500;
}

.form-input, .form-textarea {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 14px;
  transition: border-color 0.2s;
}

.form-input:focus, .form-textarea:focus {
  outline: none;
  border-color: #5662DE;
}

.form-textarea {
  resize: vertical;
  min-height: 120px;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 20px;
}

.cancel-btn {
  background-color: #f5f5f5;
  color: #333;
  border: none;
  padding: 10px 16px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  transition: background-color 0.2s;
}

.cancel-btn:hover {
  background-color: #e0e0e0;
}

.submit-btn {
  background-color: #5662DE;
  color: white;
  border: none;
  padding: 10px 16px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  transition: background-color 0.2s;
}

.submit-btn:hover {
  background-color: #454fc4;
}
</style>