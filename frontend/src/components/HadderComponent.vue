<template>
  <div class="header">
    <div class="logoBlock">
      <img :src="logo" alt="Логотип">
      <div class="logoText">
        <p>{{ orgName }}</p>
        <p class="portalName">{{ portalName }}</p>
      </div>
    </div>
    

    <!-- Новые вкладки -->
    <div class="tabs">
      <div 
        v-for="tab in tabs" 
        :key="tab.name"
        class="tab"
        :class="{ active: activeTab === tab.name }"
        @click="selectTab(tab)"
      >
        {{ tab.title }}
      </div>
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
        <div class="menu-item" @click="changePassword">
          <span>🔐</span> Сменить пароль
        </div>
        <div class="menu-item logout" @click="logout">
          <span>🚪</span> Выйти
        </div>
      </div>
    </div>

  <!-- Модальное окно смены пароля -->
    <div class="modal-overlay" v-if="isModalOpen" @click.self="closeModal">
      <div class="modal-content">
        <span class="close" @click="closeModal">&times;</span>
        <h2>Смена пароля</h2>
        
        <div class="form-group">
          <label for="newPassword">Введите новый пароль:</label>
          <input 
            type="password" 
            id="newPassword" 
            v-model="newPassword" 
            @input="validatePasswords"
          >
        </div>
        
        <div class="form-group">
          <label for="confirmPassword">Повторите пароль:</label>
          <input 
            type="password" 
            id="confirmPassword" 
            v-model="confirmPassword" 
            @input="validatePasswords"
            :class="{ 'error': !passwordsMatch && confirmPassword }"
          >
          <p class="error-message" v-if="!passwordsMatch && confirmPassword">
            Пароли не совпадают
          </p>
        </div>
        
        <button 
          class="change-btn" 
          @click="submitPasswordChange" 
          :disabled="!passwordsMatch || !newPassword"
        >
          Сменить пароль
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useUserStore } from '@/stores/user';
import logo from '@/assets/logo.png';

const router = useRouter();
const userStore = useUserStore();

if (!userStore.userData.id) {
  router.push('/login')
}
const route = useRoute();

const isMenuOpen = ref(false);
const activeTab = ref(null);
const isModalOpen = ref(false);

const newPassword = ref('');
const confirmPassword = ref('');
const passwordsMatch = ref(true);

const orgName = 'ГАУЗ ККБСМБ им. Подгорбунского';
const portalName = 'Корпоративный портал';

const tabs = [
  { name: 'chat', title: 'Чат', route: '/chat' },
  { name: 'news', title: 'Новости', route: '/news' },
  { name: 'support', title: 'Поддержка', route: '/support' },
];

const selectTab = (tab) => {
  activeTab.value = tab.name;
  router.push(tab.route);
};

const changePassword = () => {
  isModalOpen.value = true;
  newPassword.value = '';
  confirmPassword.value = '';
  passwordsMatch.value = true;
};

const closeModal = () => {
  isModalOpen.value = false;
};

const validatePasswords = () => {
  passwordsMatch.value = newPassword.value === confirmPassword.value;
};

const submitPasswordChange = async () => {
  if (!passwordsMatch.value || !newPassword.value) return;

  try {
    const newPassStruct ={
      "id":userStore.userData.id,
      "pass":newPassword.value,
    }
    
    console.log(newPassStruct);

    const response = await fetch('/api/v1/user/pass', {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(newPassStruct)
    });

    if (response.ok) {
      alert('Пароль успешно изменён');
      closeModal();
    } else {
      alert('Ошибка при изменении пароля: ' + 'Неизвестная ошибка');
    }
  } catch (error) {
    console.error('Ошибка при изменении пароля:', error);
    alert('Произошла ошибка при изменении пароля');
  }
};

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

const handleClickOutside = (event) => {
  if (!event.target.closest('.profileMenu')) {
    closeMenu();
  }
};


watch(
  () => route.path,
  (newPath) => {
    const currentTab = tabs.find(tab => newPath.startsWith(tab.route));
    activeTab.value = currentTab ? currentTab.name : null;
  },
  { immediate: true } // Вызывается сразу при загрузке
);

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
  justify-content: space-between;
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

.logoBlock{
  display: flex;
}

.portalName {
  opacity: 50%;
  font-size: 15px;
}


.tabs {
  display: flex;
  margin-left: 50px;
  gap: 110px;
  align-items: center;
}

.tab {
  padding: 10px 20px;
  font-weight: bold;
  cursor: pointer;
  color: #000000;
  transition: all 0.2s;
  border-bottom: 3px solid transparent;
  font-size: 22px;
}

.tab:hover {
  background-color: #f0f0f0;
}

.tab.active {

  border-bottom: 3px solid #5662DE;
}

.profileMenu {
  width: 350px;
  font-size: 20px;
  font-weight: bold;
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
  z-index: 1002;
}

.modal-content {
  background-color: white;
  padding: 25px;
  border-radius: 8px;
  width: 400px;
  max-width: 90%;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
  position: relative;
}

.modal-content h2 {
  margin-top: 0;
  color: #5662DE;
  text-align: center;
}

.close {
  position: absolute;
  top: 15px;
  right: 15px;
  font-size: 24px;
  cursor: pointer;
  color: #aaa;
}

.close:hover {
  color: #666;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: bold;
}

.form-group input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 16px;
}

.form-group input.error {
  border-color: #ff4444;
}

.error-message {
  color: #ff4444;
  font-size: 14px;
  margin-top: 5px;
}

.change-btn {
  background-color: #5662DE;
  color: white;
  border: none;
  padding: 12px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
  width: 100%;
  transition: background-color 0.2s;
}

.change-btn:hover {
  background-color: #454fa7;
}

.change-btn:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}
</style>