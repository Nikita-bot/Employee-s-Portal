<template>
  <div class="wrapper">
    <div class="login-container">
      <div class="login-header">
        <h1>Городская Больница</h1>
        <p>Корпоративный портал</p>
      </div>
      
      <div class="login-form">
        <div class="error-message" v-if="errorMessage" @click="errorMessage = ''">
          {{ errorMessage }}
        </div>
        
        <div class="form-group">
          <label for="username">Логин</label>
          <input 
            type="text" 
            id="username" 
            v-model="username"
            placeholder="Введите ваш логин"
            @keypress.enter="attemptLogin"
            :disabled="isLoading"
          >
        </div>
        
        <div class="form-group">
          <label for="password">Пароль</label>
          <input 
            type="password" 
            id="password" 
            v-model="password"
            placeholder="Введите ваш пароль"
            @keypress.enter="attemptLogin"
            :disabled="isLoading"
          >
        </div>
        
        <button class="login-btn" @click="attemptLogin" :disabled="isLoading">
          <span v-if="isLoading">Вход...</span>
          <span v-else>Войти в систему</span>
        </button>
        
        <div class="login-links">
          <a href="#">Забыли пароль?</a> | 
          <a href="#">Техническая поддержка</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()

const username = ref('')
const password = ref('')
const errorMessage = ref('')
const isLoading = ref(false)

const attemptLogin = async () => {
  if (!username.value || !password.value) {
    errorMessage.value = 'Пожалуйста, заполните все поля'
    setTimeout(() => {
      errorMessage.value = ''
    }, 3000)
    return
  }
  
  isLoading.value = true
  errorMessage.value = ''
  
  try {
    const formData = new FormData()
    formData.append('login', username.value)
    formData.append('password', password.value)

    const response = await fetch('/api/v1/auth', {
      method: 'POST',
      body: formData
    })
    
    if (!response.ok) {
      throw new Error('Ошибка авторизации')
    }
    
    
    const authData = await response.json()
    console.log(authData.userId)
    
    // Сохраняем только user_id из ответа
    userStore.setAuthData({
      userId: authData.userId,
    })
    
    // Загружаем полные данные пользователя
    await userStore.fetchUserData()
    
    // Переходим на главную страницу
    router.push('/')
    
  } catch (error) {
    console.error('Login error:', error)
    errorMessage.value = 'Неверный логин или пароль'
    setTimeout(() => {
      errorMessage.value = ''
    }, 3000)
  } finally {
    isLoading.value = false
  }
}
</script>

<style>
/* Стили остаются без изменений */
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}
.wrapper {
    background: linear-gradient(135deg, #8297f3 0%, #3023ee 100%);
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 1rem;
}
.login-container {
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
.form-group input:disabled {
    background-color: #f8f9fa;
    cursor: not-allowed;
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
    background: #ccc;
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
    display: block;
}
</style>