<template>
  <div class="wrapper">
    <header>
      <div class="logo">Городская Больница</div>
      <div class="user-menu" @click="toggleDropdown">
        <span>{{ userStore.userData?.name }}</span>
        <img :src="userStore.userData?.avatar" alt="Аватар">
        <div class="dropdown-menu" :class="{ show: dropdownOpen }">
          <router-link to="/profile" class="dropdown-item" @click="closeDropdown">Мой профиль</router-link>
          <a href="#" class="dropdown-item" @click="closeDropdown">Сменить пароль</a>
          <a href="#" class="dropdown-item" @click="closeDropdown">Привязать Телеграмм</a>
          <div class="dropdown-divider"></div>
          <a href="#" class="dropdown-item" @click="logout">Выйти</a>
        </div>
      </div>
    </header>

    <div class="main-container" :class="{ 'menu-collapsed': menuCollapsed }">
      <sidebar>
        <button class="menu-toggle" @click="toggleMenu">
          <i class="fas" :class="menuCollapsed ? 'fa-chevron-right' : 'fa-chevron-left'"></i>
        </button>
        <ul class="nav-links">
          <li>
            <router-link to="/" class="nav-link">
              <i class="fas fa-home"></i>
              <span>Главная</span>
            </router-link>
          </li>
          <li>
            <router-link to="/tasks" class="nav-link">
              <i class="fas fa-tasks"></i>
              <span>Задачи</span>
            </router-link>
          </li>
          <li>
            <a href="#" class="nav-link">
              <i class="fas fa-file-alt"></i>
              <span>ЭДО</span>
            </a>
          </li>
          <li>
            <a href="#" class="nav-link">
              <i class="fas fa-book"></i>
              <span>База знаний</span>
            </a>
          </li>
          <li>
            <a href="#" class="nav-link">
              <i class="fas fa-external-link-alt"></i>
              <span>Порталы</span>
            </a>
          </li>
          <li>
            <a href="#" class="nav-link">
              <i class="fas fa-newspaper"></i>
              <span>Новости</span>
            </a>
          </li>
          <li>
            <a href="#" class="nav-link">
              <i class="fas fa-headset"></i>
              <span>Поддержка</span>
            </a>
          </li>
        </ul>
      </sidebar>

      <main class="content-area">
        <section class="welcome-banner">
          <h1>Добро пожаловать в корпоративный портал</h1>
          <p>Сегодня: {{ currentDate }}</p>
          <div class="user-role">{{ userStore.userData?.role }} | {{ userStore.userData?.department }}</div>
        </section>

        <section class="quick-access-grid">
          <div class="quick-card" @click="$router.push('/tasks')">
            <h3>Мои задачи</h3>
            <p>Текущие и назначенные задания</p>
          </div>
          <div class="quick-card" @click="openFeature('Электронный документооборот')">
            <h3>Электронный документооборот</h3>
            <p>Работа с документами и подписями</p>
          </div>
          <div class="quick-card" @click="$router.push('/profile')">
            <h3>Мой профиль</h3>
            <p>ЭЦП, аналитика, календарь</p>
          </div>
          <div class="quick-card" @click="openFeature('База знаний')">
            <h3>База знаний</h3>
            <p>Протоколы и стандарты лечения</p>
          </div>
          <div class="quick-card" @click="openFeature('Внешние порталы')">
            <h3>Внешние порталы</h3>
            <p>Доступ к смежным системам</p>
          </div>
          <div class="quick-card" @click="openFeature('Техподдержка')">
            <h3>Техподдержка</h3>
            <p>Помощь с техническими вопросами</p>
          </div>
        </section>

        <section class="news-section">
          <div class="section-header">
            <h2>Последние новости</h2>
            <a href="#" class="view-all">Все новости</a>
          </div>
          <div class="news-grid">
            <!-- Новости будут здесь -->
          </div>
        </section>
      </main>
    </div>

    <footer>
      <p>© 2023 Городская Больница. Корпоративная информационная система. Версия 2.1</p>
    </footer>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()
const dropdownOpen = ref(false)
const menuCollapsed = ref(false)

const currentDate = computed(() => {
  const now = new Date()
  const options = { year: 'numeric', month: 'long', day: 'numeric' }
  return now.toLocaleDateString('ru-RU', options)
})

const toggleDropdown = () => {
  dropdownOpen.value = !dropdownOpen.value
}

const closeDropdown = () => {
  dropdownOpen.value = false
}

const toggleMenu = () => {
  menuCollapsed.value = !menuCollapsed.value
}

const logout = () => {
  userStore.clearUserData()
  router.push('/login')
  closeDropdown()
}

const openFeature = (featureName) => {
  alert(`Функция "${featureName}" в разработке`)
}

// Закрытие dropdown при клике вне его
const handleClickOutside = (event) => {
  const userMenu = document.querySelector('.user-menu')
  if (userMenu && !userMenu.contains(event.target)) {
    dropdownOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
/* Стили из вашего HTML файла */
.wrapper {
  display: grid;
  grid-template-rows: auto 1fr auto;
  min-height: 100vh;
}

header {
  background-color: #fff;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  padding: 1rem 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: relative;
}

.logo {
  font-size: 1.5rem;
  font-weight: 300;
  color: #2c5aa0;
}

.user-menu {
  display: flex;
  align-items: center;
  gap: 1rem;
  cursor: pointer;
  position: relative;
}

.user-menu img {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  right: 0;
  background: white;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
  border-radius: 8px;
  padding: 0.5rem 0;
  min-width: 200px;
  display: none;
  z-index: 1000;
}

.dropdown-menu.show {
  display: block;
}

.dropdown-item {
  padding: 0.75rem 1.5rem;
  display: block;
  text-decoration: none;
  color: #333;
  transition: background 0.3s ease;
}

.dropdown-item:hover {
  background: #f8f9fa;
}

.dropdown-divider {
  height: 1px;
  background: #e1e5e9;
  margin: 0.25rem 0;
}

.main-container {
  display: grid;
  grid-template-columns: 250px 1fr;
  gap: 0;
  transition: all 0.3s ease;
}

.main-container.menu-collapsed {
  grid-template-columns: 70px 1fr;
}

sidebar {
  background-color: #2c5aa0;
  color: white;
  padding: 1rem 0;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.menu-toggle {
  position: absolute;
  top: 1rem;
  right: 1rem;
  background: rgba(255, 255, 255, 0.2);
  border: none;
  color: white;
  width: 30px;
  height: 30px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s ease;
}

.menu-toggle:hover {
  background: rgba(255, 255, 255, 0.3);
}

.nav-links {
  list-style: none;
  margin-top: 2rem;
}

.nav-links li a {
  display: flex;
  align-items: center;
  color: #e0e0e0;
  padding: 0.8rem 1.5rem;
  text-decoration: none;
  transition: all 0.3s ease;
  white-space: nowrap;
}

.nav-links li a i {
  margin-right: 1rem;
  font-size: 1.2rem;
  width: 20px;
  text-align: center;
}

.nav-links li a:hover,
.nav-links li a.router-link-active {
  background-color: rgba(255, 255, 255, 0.1);
  color: white;
  border-left: 4px solid #4fc3f7;
}

.menu-collapsed .nav-links li a span {
  display: none;
}

.menu-collapsed .nav-links li a i {
  margin-right: 0;
}

.content-area {
  padding: 2rem;
  background-color: #fff;
}

.welcome-banner {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 1.5rem;
  border-radius: 8px;
  margin-bottom: 2rem;
}

.welcome-banner h1 {
  font-weight: 300;
  margin-bottom: 0.5rem;
  font-size: 1.5rem;
}

.user-role {
  background: rgba(255, 255, 255, 0.2);
  padding: 0.25rem 0.75rem;
  border-radius: 20px;
  font-size: 0.9rem;
  display: inline-block;
  margin-top: 0.5rem;
}

.quick-access-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
}

.quick-card {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  padding: 1.25rem;
  text-align: center;
  transition: all 0.3s ease;
  cursor: pointer;
  border: 1px solid #e1e5e9;
  min-height: 120px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.quick-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.quick-card h3 {
  color: #2c5aa0;
  margin-bottom: 0.5rem;
  font-weight: 500;
  font-size: 1rem;
}

.quick-card p {
  color: #666;
  font-size: 0.85rem;
  line-height: 1.4;
}

.news-section {
  background: #f8f9fa;
  border-radius: 8px;
  padding: 1.5rem;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
  padding-bottom: 0.75rem;
  border-bottom: 2px solid #e1e5e9;
}

.section-header h2 {
  color: #2c5aa0;
  font-weight: 500;
  font-size: 1.25rem;
}

.news-grid {
  display: grid;
  gap: 1rem;
}

.view-all {
  color: #2c5aa0;
  text-decoration: none;
  font-weight: 500;
}

.view-all:hover {
  text-decoration: underline;
}

footer {
  background-color: #1a1a1a;
  color: #999;
  text-align: center;
  padding: 1.5rem;
  font-size: 0.9rem;
}

@media (max-width: 768px) {
  .main-container {
    grid-template-columns: 1fr;
  }
  sidebar {
    display: none;
  }
  .quick-access-grid {
    grid-template-columns: 1fr 1fr;
  }
  .welcome-banner h1 {
    font-size: 1.25rem;
  }
}
</style>