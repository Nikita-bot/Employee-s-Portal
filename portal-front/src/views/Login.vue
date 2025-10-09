<template>
  <div class="login-container">
    <div class="login-box">
      <div class="login-header">
        <h1>Городская Больница</h1>
        <p>Корпоративный портал</p>
      </div>
      
      <form @submit.prevent="handleLogin" class="login-form">
        <div v-if="errorMessage" class="error-message">
          {{ errorMessage }}
        </div>
        
        <div class="form-group">
          <label for="username">Логин</label>
          <input 
            v-model="form.username" 
            type="text" 
            id="username" 
            placeholder="Введите ваш логин"
            required
          >
        </div>
        
        <div class="form-group">
          <label for="password">Пароль</label>
          <input 
            v-model="form.password" 
            type="password" 
            id="password" 
            placeholder="Введите ваш пароль"
            required
          >
        </div>
        
        <button type="submit" class="login-btn" :disabled="loading">
          {{ loading ? 'Вход...' : 'Войти в систему' }}
        </button>
        
        <div class="login-links">
          <a href="#">Забыли пароль?</a> | 
          <a href="#">Техническая поддержка</a>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()
const loading = ref(false)
const errorMessage = ref('')

const form = reactive({
  username: '',
  password: ''
})

const handleLogin = async () => {
  loading.value = true
  errorMessage.value = ''

  try {
    // Имитация запроса к API
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // Демо-логика (в реальном приложении здесь будет запрос к вашему Go API)
    if (form.username && form.password) {
      const userData = {
        id: 1,
        name: 'Иванов Александр Сергеевич',
        role: 'Врач-кардиолог',
        department: 'Кардиологическое отделение',
        avatar: 'https://i.imgur.com/3QZ9a7s.png'
      }
      
      userStore.setUserData(userData)
      router.push('/')
    } else {
      errorMessage.value = 'Неверный логин или пароль'
    }
  } catch (error) {
    errorMessage.value = 'Ошибка при входе в систему'
  } finally {
    loading.value = false
  }
}

// Обработка нажатия Enter
const handleKeyPress = (event) => {
  if (event.key === 'Enter') {
    handleLogin()
  }
}
</script>

<style scoped>
.login-container {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 1rem;
}

.login-box {
  background: white;
  border-radius: 12px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 420px;
  overflow: hidden;
}

.login-header {
  background: #2c5aa0;
  color: white;
  padding: 2rem;
  text-align: center;
}

.login-header h1 {
  font-weight: 300;
  font-size: 1.5rem;
  margin-bottom: 0.5rem;
}

.login-form {
  padding: 2rem;
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
  padding: 0.75rem 1rem;
  border: 2px solid #e1e5e9;
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.3s ease;
}

.form-group input:focus {
  outline: none;
  border-color: #2c5aa0;
}

.login-btn {
  width: 100%;
  background: #2c5aa0;
  color: white;
  border: none;
  padding: 0.75rem;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.3s ease;
}

.login-btn:hover:not(:disabled) {
  background: #1e3d6f;
}

.login-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.login-links {
  margin-top: 1.5rem;
  text-align: center;
}

.login-links a {
  color: #2c5aa0;
  text-decoration: none;
  font-size: 0.9rem;
}

.login-links a:hover {
  text-decoration: underline;
}

.error-message {
  background: #ffe6e6;
  color: #d63031;
  padding: 0.75rem;
  border-radius: 6px;
  margin-bottom: 1rem;
  text-align: center;
}
</style>