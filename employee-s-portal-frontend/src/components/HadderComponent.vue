<template>
  <div class="header">
    <img :src="logo" alt="Логотип">
    <div class="logoText">
      <p>{{ orgName }}</p>
      <p class="portalName">{{ portalName }}</p>
    </div>
    <div class="profileMenu" @click.stop="toggleMenu">
      <img class="photoProfile" :src="logo" alt="Профиль">
      <p class="name">{{ formattedUserName }}</p>
      
      <!-- Выпадающее меню -->
      <div v-if="isMenuOpen" class="dropdown-menu" @click.stop>
        <div class="menu-item" @click="goToProfile">
          <span>👤</span> Открыть профиль
        </div>
        <div class="menu-item" @click="bindTelegram">
          <span>📱</span> Привязать Telegram
        </div>
        <div class="menu-item logout" @click="logout">
          <span>🚪</span> Выйти
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '@/stores/user';
import logo from '@/assets/logo.png';

const router = useRouter();
const userStore = useUserStore();
const isMenuOpen = ref(false);

const orgName = 'ГАУЗ ККБСМБ им. Подгорбунского';
const portalName = 'Корпоративный портал';

// Форматируем имя пользователя
const formattedUserName = computed(() => {
  if (!userStore.userData) return 'Гость';
  const { surname, name, patronymic } = userStore.userData;
  return `${surname} ${name?.charAt(0) || ''}.${patronymic?.charAt(0) || ''}.`;
});

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value;
};

const closeMenu = () => {
  isMenuOpen.value = false;
};

const goToProfile = () => {
  if (router.currentRoute.value.path === `/user/${userStore.userData.id}`) {
    // Если уже на этой странице, эмулируем перезагрузку
    window.location.reload();
  } else {
    router.push(`/user/${userStore.userData.id}`);
  }
  closeMenu();
};

const bindTelegram = () => {
  alert('Функция привязки Telegram будет реализована позже');
  closeMenu();
};

const logout = () => {
  userStore.clearUserData();
  router.push('/login');
  closeMenu();
};

// Закрытие меню при клике вне его области
const handleClickOutside = (event) => {
  if (!event.target.closest('.profileMenu')) {
    closeMenu();
  }
};

onMounted(() => {
  document.addEventListener('click', handleClickOutside);
});

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside);
});
</script>

<style scoped>
.header {
  background-color: #FFFFFF;
  display: flex;
  align-items: center;
  height: 120px;
  padding: 10px;
  position: relative;
}

img {
  width: 75px;
  height: 75px;
}

.logoText {
  width: 180px;
  margin-left: 10px;
  font-size: 20px;
  font-weight: bold;
}

.portalName {
  opacity: 50%;
  font-size: 15px;
}

.profileMenu {
  width: 350px;
  margin-top: 20px;
  margin-right: 10px;
  font-size: 20px;
  font-weight: bold;
  position: fixed;
  right: 0px;
  top: 0;
  display: flex;
  cursor: pointer;
  z-index: 1000;
}

.photoProfile {
  margin-top: 0;
  margin-right: 30px;
}

.name {
  font-size: 25px;
  margin-top: 20px;
}

/* Стили для выпадающего меню */
.dropdown-menu {
  position: absolute;
  top: 100%;
  right: 0;
  background-color: white;
  border: 1px solid #ddd;
  border-radius: 5px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  width: 250px;
  z-index: 1001;
  margin-top: 10px;
}

.menu-item {
  padding: 12px 20px;
  cursor: pointer;
  display: flex;
  align-items: center;
  transition: background-color 0.2s;
}

.menu-item:hover {
  background-color: #f5f5f5;
}

.menu-item span {
  margin-right: 10px;
  font-size: 18px;
}

.logout {
  color: #ff4444;
  border-top: 1px solid #eee;
}
</style>