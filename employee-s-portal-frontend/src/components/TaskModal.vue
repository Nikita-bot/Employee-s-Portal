<template>
  <div v-if="task" class="modal-overlay" @click.self="$emit('close')">
    <div class="task-modal">
      <div class="modal-header">
        <h3>Задача #{{ task.id }}</h3>
      </div>
      
      <div class="modal-content">
        <div class="task-info">
          <p><strong>Автор:</strong> {{ task.initiator }}</p>
          <p><strong>Исполнитель:</strong> {{ task.executor }}</p>
          <p><strong>Описание:</strong> {{ task.description }}</p>
          <p><strong>Статус:</strong> {{ task.status }}</p>
          <p><strong>Дата создания:</strong> {{task.create_date }}</p>
          <p><strong>Дата выполнения:</strong> {{ task.execute_date ? task.execute_date : 'Не выполнено' }}</p>
        </div>

        <div class="comments-section">
          <h4>Комментарии:</h4>
          <div v-for="comment in task.comments" :key="comment.id" class="comment">
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
          v-if="isCreator"
          class="delete-btn"
          @click="$emit('delete')"
        >
          Удалить задачу
        </button>
        <button 
          v-if="isExecutor && task.status !== 1"
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
import { ref, defineProps, defineEmits } from 'vue';

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
  }
});

const emit = defineEmits(['close', 'delete', 'complete', 'add-comment']);

const newComment = ref('');

const addComment = () => {
  if (!newComment.value.trim()) return;
  emit('add-comment', newComment.value);
  newComment.value = '';
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

.complete-btn {
  border: none;
  background-color: #5662DE;
  border-radius: 5px;
  color: #FFFFFF;
  width: 160px;
  padding: 8px 15px;
  cursor: pointer;
}

h3 {
  color: #FFFFFF;
}
</style>