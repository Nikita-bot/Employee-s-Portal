<template>
  <div class="modal" :class="{ show: showModal }" @click="closeModal">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h2>Смена пароля</h2>
        <button class="close-btn" @click="closeModal">&times;</button>
      </div>
      <div class="modal-body">
        <form @submit.prevent="changePassword">
          <div class="form-group">
            <label for="newPassword">Новый пароль</label>
            <div class="password-toggle">
              <input 
                type="password" 
                id="newPassword" 
                v-model="newPassword"
                placeholder="Введите новый пароль"
                @input="validatePassword"
              >
              <button type="button" class="toggle-password" @click="togglePasswordVisibility('newPassword')">
                👁️
              </button>
            </div>
          </div>

          <div class="form-group">
            <label for="confirmPassword">Подтверждение пароля</label>
            <div class="password-toggle">
              <input 
                type="password" 
                id="confirmPassword" 
                v-model="confirmPassword"
                placeholder="Повторите новый пароль"
                @input="validatePassword"
              >
              <button type="button" class="toggle-password" @click="togglePasswordVisibility('confirmPassword')">
                👁️
              </button>
            </div>
            <div class="error-message" :style="{ display: showPasswordError ? 'block' : 'none' }">
              Пароли не совпадают
            </div>
          </div>

          <div class="success-message" :style="{ display: passwordSuccess ? 'block' : 'none' }">
            Пароль успешно изменен!
          </div>

          <div class="error-message" :style="{ display: passwordError ? 'block' : 'none' }">
            {{ passwordError }}
          </div>
        </form>
      </div>
      <div class="modal-footer">
        <button class="btn btn-outline" @click="closeModal">Отмена</button>
        <button class="btn btn-primary" :disabled="!canSave || isLoading" @click="changePassword">
          <span v-if="isLoading">Сохранение...</span>
          <span v-else>Сохранить</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useUserStore } from '@/stores/user.js'

const userStore = useUserStore()

const showModal = ref(false)
const newPassword = ref('')
const confirmPassword = ref('')
const showPasswordError = ref(false)
const passwordSuccess = ref(false)
const passwordError = ref('')
const isLoading = ref(false)

const canSave = computed(() => {
  return newPassword.value === confirmPassword.value && 
         newPassword.value !== '' && 
         confirmPassword.value !== ''
})

const show = () => {
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  resetForm()
}

const resetForm = () => {
  newPassword.value = ''
  confirmPassword.value = ''
  showPasswordError.value = false
  passwordSuccess.value = false
  passwordError.value = ''
  isLoading.value = false
}

const togglePasswordVisibility = (fieldId) => {
  const field = document.getElementById(fieldId)
  if (field.type === 'password') {
    field.type = 'text'
  } else {
    field.type = 'password'
  }
}

const validatePassword = () => {
  showPasswordError.value = confirmPassword.value !== '' && 
                           newPassword.value !== confirmPassword.value
  passwordError.value = ''
}

const changePassword = async () => {
  if (!canSave.value) return

  isLoading.value = true
  passwordError.value = ''
  const formData = new FormData();
  formData.append('id', userStore.userId);
  formData.append('pass', newPassword.value);


  try {
    const response = await fetch('/api/v1/user/pass', {
      method: 'PATCH',
      body: formData
    })

    if (response.ok) {
      passwordSuccess.value = true
      passwordError.value = ''
      
      setTimeout(() => {
        closeModal()
        // Показываем глобальное уведомление об успехе
        if (typeof window.showNotification === 'function') {
          window.showNotification('Пароль успешно изменен', 'success')
        }
      }, 1500)
    } else {
      const errorText = await response.text()
      throw new Error(errorText || 'Ошибка при смене пароля')
    }
  } catch (error) {
    console.error('Ошибка смены пароля:', error)
    passwordError.value = error.message || 'Произошла ошибка при смене пароля'
    passwordSuccess.value = false
  } finally {
    isLoading.value = false
  }
}

// Экспортируем методы для использования в других компонентах
defineExpose({
  show
})
</script>

<style scoped>
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
  max-width: 450px;
  max-height: 90vh;
  overflow-y: auto;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
}
.modal-header {
  padding: 1.5rem;
  border-bottom: 1px solid #e1e5e9;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #f8f9fa;
  border-radius: 8px 8px 0 0;
}
.modal-header h2 {
  color: #2c5aa0;
  font-weight: 500;
  font-size: 1.25rem;
}
.close-btn {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: #666;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 4px;
}
.close-btn:hover {
  background: #e9ecef;
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
.form-group input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 1rem;
  transition: all 0.3s ease;
}
.form-group input:focus {
  outline: none;
  border-color: #2c5aa0;
  box-shadow: 0 0 0 3px rgba(44, 90, 160, 0.1);
}
.password-toggle {
  position: relative;
}
.toggle-password {
  position: absolute;
  right: 10px;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  color: #666;
  cursor: pointer;
  padding: 0.25rem;
  border-radius: 4px;
}
.toggle-password:hover {
  background: #f0f0f0;
}
.error-message {
  color: #e74c3c;
  font-size: 0.85rem;
  margin-top: 0.5rem;
  display: none;
}
.success-message {
  color: #27ae60;
  font-size: 0.85rem;
  margin-top: 0.5rem;
  display: none;
}
.modal-footer {
  padding: 1rem 1.5rem;
  border-top: 1px solid #e1e5e9;
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
}
.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  transition: all 0.3s ease;
  min-width: 100px;
}
.btn-primary {
  background: #2c5aa0;
  color: white;
}
.btn-primary:hover:not(:disabled) {
  background: #1e3d6f;
}
.btn-primary:disabled {
  background: #ccc;
  cursor: not-allowed;
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
</style>