<template>
  <div class="login-page">
    <div class="header">
      <img :src="logo" alt="Логотип">
      <div class="logoText">
        <p>{{ orgName }}</p>
        <p class="portalName">{{ portalName }}</p>
      </div>
    </div>
    
    <div class="container">
      <div class="mainBlock">
        <h2>Авторизация</h2>
        <form @submit.prevent="handleLogin">
          <input 
            v-model="loginForm.username"
            class="login" 
            placeholder="Логин" 
            type="text"
          >
          <br/>
          <input 
            v-model="loginForm.password"
            class="password" 
            placeholder="Пароль" 
            type="password"
          >
          <br>
          <input 
            class="bLogin" 
            type="submit" 
            value="ВОЙТИ"
            :disabled="isLoading"
          >
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import logo from '@/assets/logo.png'; 

const router = useRouter();

const orgName = 'ГАУЗ ККБСМБ им. Подгорбунского';
const portalName = 'Корпоративный портал';

const loginForm = ref({
  username: '',
  password: ''
});

const isLoading = ref(false);

const handleLogin = async () => {
  if (!loginForm.value.username || !loginForm.value.password) {
    alert('Пожалуйста, заполните все поля');
    return;
  }

  isLoading.value = true;
  
  try {
    // const response = await authService.login(loginForm.value);
    
    router.push('/tasks'); // Перенаправляем на страницу задач
  } catch (error) {
    console.error('Ошибка авторизации:', error);
    alert('Неверный логин или пароль');
  } finally {
    isLoading.value = false;
  }
};
</script>

<style scoped>
.login-page {
  font-family: Arial, sans-serif;
}

* {
  margin: 0;
  padding: 0;
  color: #565656;
  font-weight: 700;
}

body {
  background-color: #EAEAEA;
}

.header {
  background-color: #FFFFFF;
  display: flex;
  height: 100px;
  padding: 10px;
}

img {
  width: 75px;
  height: 75px;
  margin-top: 10px;
}

.logoText {
  width: 180px;
  margin-top: 20px;
  margin-left: 10px;
  font-size: 20px;
  font-weight: bold;
}

.portalName {
  opacity: 50%;
  font-size: 15px;
}

.container {
  width: 100%;
  height: 100%;
  position: fixed;
  top: 0;
  left: 0;
  overflow: auto;  
  background-color: #EAEAEA;
}

.mainBlock {
  background-color: #FFFFFF;
  width: 750px;
  height: 375px;
  position: absolute;
  border-radius: 15px;
  text-align: center;
  top: 30%;
  left: 50%;
  transform: translateX(-50%);
}

h2 {
  font-size: 33px;
  padding-top: 50px;
}

input {
  margin-top: 25px;
  height: 45px;
  width: 350px;
  border-radius: 15px;
  background-color: #EAEAEA;
  border: none;
  font-weight: bold;
  font-size: 15px;
  padding-left: 15px;
}

.bLogin {
  background-color: #5662DE;
  color: #FFFFFF;
  cursor: pointer;
}

.bLogin:hover {
  background-color: #2d39bb;
}

.bLogin:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}
</style>