<template>
  <header class="app-header">
    <div class="logo">Городская Больница</div>
    <div class="user-menu" @click="toggleDropdown">
      <span class="user-name">{{ userStore.user.user.surname  + ' ' + userStore.user.user.name[0] +'.'+ userStore.user.user.patronymic[0] +'.'}}</span>
      <img :src="userStore.user.user.avatar || '/default-avatar.png'" alt="Аватар" class="user-avatar">
      <div class="dropdown-menu" :class="{ show: showDropdown }">
        <router-link to="/profile" class="dropdown-item">
          <i class="fas fa-user"></i>
          Мой профиль
        </router-link>
        <button class="dropdown-item" @click="showChangePasswordModal">
          <i class="fas fa-key"></i>
          Сменить пароль
        </button>
        <button class="dropdown-item" @click="showTelegramModal">
          <i class="fab fa-telegram"></i>
          Привязать Телеграмм
        </button>
        <div class="dropdown-divider"></div>
        <button class="dropdown-item logout" @click="logout">
          <i class="fas fa-sign-out-alt"></i>
          Выйти
        </button>
      </div>
    </div>
    
    <!-- Модалки -->
    <PasswordModal ref="passwordModal" />
    <TelegramModal ref="telegramModal" />
  </header>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import PasswordModal from '@/components/PasswordModal.vue'
import TelegramModal from '@/components/TelegramModal.vue'

const router = useRouter()
const userStore = useUserStore()
const showDropdown = ref(false)

const passwordModal = ref(null)
const telegramModal = ref(null)

const toggleDropdown = () => {
  showDropdown.value = !showDropdown.value
}

const showChangePasswordModal = () => {
  passwordModal.value?.show()
  showDropdown.value = false
}

const showTelegramModal = () => {
  telegramModal.value?.show()
  showDropdown.value = false
}

const logout = () => {
  userStore.logout()
  router.push('/login')
}

// Закрытие dropdown при клике вне
document.addEventListener('click', (event) => {
  const userMenu = document.querySelector('.user-menu')
  if (userMenu && !userMenu.contains(event.target)) {
    showDropdown.value = false
  }
})
</script>

<style scoped>
.app-header {
  background-color: #fff;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  padding: 1rem 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: sticky;
  top: 0;
  z-index: 100;
  height: 70px;
}

.logo {
  font-size: 1.5rem;
  font-weight: 300;
  color: #2c5aa0;
  letter-spacing: 0.5px;
}

.user-menu {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  cursor: pointer;
  position: relative;
  padding: 0.5rem 0.75rem;
  border-radius: 8px;
  transition: background-color 0.3s ease;
}

.user-menu:hover {
  background-color: #f8f9fa;
}

.user-name {
  font-weight: 500;
  color: #333;
  font-size: 0.95rem;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #e1e5e9;
  transition: border-color 0.3s ease;
}

.user-menu:hover .user-avatar {
  border-color: #2c5aa0;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  right: 0;
  background: white;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
  border-radius: 8px;
  padding: 0.5rem 0;
  min-width: 220px;
  display: none;
  z-index: 1000;
  margin-top: 0.5rem;
  border: 1px solid #e1e5e9;
}

.dropdown-menu.show {
  display: block;
  animation: dropdownFadeIn 0.2s ease;
}

@keyframes dropdownFadeIn {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.dropdown-item {
  padding: 0.75rem 1.5rem;
  display: flex;
  align-items: center;
  gap: 0.75rem;
  text-decoration: none;
  font-size: 0.9rem;
  transition: all 0.3s ease;
  color: #333;
  cursor: pointer;
  border: none;
  background: none;
  width: 100%;
  text-align: left;
  border-radius: 0;
}

.dropdown-item:hover {
  background: #f8f9fa;
  color: #2c5aa0;
}

.dropdown-item i {
  width: 16px;
  text-align: center;
  font-size: 0.9rem;
}

.dropdown-item.logout {
  color: #e74c3c;
}

.dropdown-item.logout:hover {
  background: #ffe6e6;
  color: #c0392b;
}

.dropdown-divider {
  height: 1px;
  background: #e1e5e9;
  margin: 0.25rem 0;
}

/* Иконки Font Awesome */
.fas, .fab {
  font-family: 'Font Awesome 6 Free';
  font-weight: 900;
}

.fab {
  font-family: 'Font Awesome 6 Brands';
}

/* Адаптивность */
@media (max-width: 768px) {
  .app-header {
    padding: 1rem;
  }
  
  .logo {
    font-size: 1.25rem;
  }
  
  .user-name {
    display: none;
  }
  
  .dropdown-menu {
    min-width: 200px;
    right: -10px;
  }
}

@media (max-width: 480px) {
  .app-header {
    padding: 0.75rem;
  }
  
  .logo {
    font-size: 1.1rem;
  }
  
  .user-avatar {
    width: 35px;
    height: 35px;
  }
}
</style>