<template>
  <div class="modal" :class="{ show: showModal }" @click="closeModal">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h2>
          <i class="fab fa-telegram telegram-icon"></i>
          Привязка Telegram
        </h2>
        <button class="close-btn" @click="closeModal">&times;</button>
      </div>
      <div class="modal-body">
        <div class="instructions">
          <h3>Как привязать Telegram:</h3>
          
          <div class="instruction-step">
            <div class="step-number">1</div>
            <div class="step-text">
              Откройте Telegram и найдите бота <strong>@HospitalPortalBot</strong>
            </div>
          </div>
          
          <div class="instruction-step">
            <div class="step-number">2</div>
            <div class="step-text">
              Отправьте боту команду <code>/start</code> для начала работы
            </div>
          </div>
          
          <div class="instruction-step">
            <div class="step-number">3</div>
            <div class="step-text">
              Введите код подтверждения, который вам отправил бот
            </div>
          </div>
        </div>

        <div class="form-group">
          <label for="verificationCode">Код подтверждения</label>
          <input 
            type="text" 
            id="verificationCode" 
            v-model="verificationCode"
            class="code-input" 
            placeholder="ABC123" 
            maxlength="6"
            @input="validateVerificationCode"
          >
          <div class="error-message" :style="{ display: showCodeError ? 'block' : 'none' }">
            {{ codeErrorText }}
          </div>
          <div class="success-message" :style="{ display: showCodeSuccess ? 'block' : 'none' }">
            Код подтвержден!
          </div>
        </div>

        <div class="status-connected" :style="{ display: isConnected ? 'block' : 'none' }">
          ✅ Telegram успешно привязан! Теперь вы будете получать уведомления о задачах и новостях.
        </div>
      </div>
      <div class="modal-footer">
        <button class="btn btn-outline" @click="closeModal">Отмена</button>
        <button class="btn btn-primary" :disabled="!canConnect" @click="connectTelegram" v-if="!isConnected">
          Привязать
        </button>
        <button class="btn btn-danger" @click="disconnectTelegram" v-else>
          Отвязать
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const showModal = ref(false)
const verificationCode = ref('')
const showCodeError = ref(false)
const showCodeSuccess = ref(false)
const codeErrorText = ref('Неверный код подтверждения')
const isConnected = ref(false)

const canConnect = computed(() => {
  return /^[A-Z0-9]{6}$/.test(verificationCode.value.toUpperCase())
})

const show = () => {
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  resetForm()
}

const resetForm = () => {
  verificationCode.value = ''
  showCodeError.value = false
  showCodeSuccess.value = false
  isConnected.value = false
}

const validateVerificationCode = () => {
  const code = verificationCode.value.toUpperCase()
  const isValid = /^[A-Z0-9]{6}$/.test(code)
  
  if (code.length === 6 && !isValid) {
    showCodeError.value = true
    showCodeSuccess.value = false
  } else if (isValid) {
    showCodeError.value = false
    showCodeSuccess.value = true
  } else {
    showCodeError.value = false
    showCodeSuccess.value = false
  }
}

const connectTelegram = () => {
  if (!canConnect.value) return

  // Демо-логика
  if (verificationCode.value.toUpperCase() === 'ABC123') {
    isConnected.value = true
  } else {
    showCodeError.value = true
    codeErrorText.value = 'Неверный код подтверждения'
  }
}

const disconnectTelegram = () => {
  // В реальном приложении AJAX запрос
  console.log('Отвязка Telegram')
  resetForm()
  alert('Telegram успешно отвязан от вашего аккаунта')
}

defineExpose({
  show
})
</script>

<style scoped>
/* Стили из main.css для Telegram модалки */
.instructions {
  background: #f8f9fa;
  padding: 1.5rem;
  border-radius: 8px;
  margin-bottom: 1.5rem;
  border-left: 4px solid #0088cc;
}
.instructions h3 {
  color: #0088cc;
  margin-bottom: 1rem;
  font-size: 1.1rem;
}
.instruction-step {
  display: flex;
  align-items: flex-start;
  gap: 0.75rem;
  margin-bottom: 1rem;
}
.step-number {
  background: #0088cc;
  color: white;
  width: 24px;
  height: 24px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.8rem;
  font-weight: bold;
  flex-shrink: 0;
}
.step-text {
  flex: 1;
  color: #666;
  line-height: 1.5;
}
.telegram-icon {
  color: #0088cc;
  margin-right: 0.5rem;
}
.status-connected {
  background: #e8f5e8;
  color: #27ae60;
  padding: 1rem;
  border-radius: 6px;
  margin-bottom: 1rem;
  display: none;
}
.connection-info {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 1rem;
  background: #f8f9fa;
  border-radius: 6px;
  margin-bottom: 1rem;
}
.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
}
.user-details h4 {
  margin: 0;
  color: #333;
  font-weight: 500;
}
.user-details p {
  margin: 0;
  color: #666;
  font-size: 0.9rem;
}
.hidden {
  display: none !important;
}
.btn-danger {
  background: #e74c3c;
  color: white;
}
.btn-danger:hover {
  background: #c0392b;
}
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