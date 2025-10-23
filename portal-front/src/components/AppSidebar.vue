<template>
  <sidebar>
    <ul class="nav-links">
      <li>
        <router-link to="/" :class="{ active: $route.path === '/' }">
          <i class="fas fa-home"></i>
          <span>Главная</span>
        </router-link>
      </li>
      <li>
        <router-link to="/tasks" :class="{ active: $route.path === '/tasks' }">
          <i class="fas fa-tasks"></i>
          <span>Задачи</span>
        </router-link>
      </li>
      <li>
        <a href="#">
          <i class="fas fa-file-alt"></i>
          <span>ЭДО</span>
        </a>
      </li>
      <li>
        <a href="http://documents.kkbsmp.ru/view/list">
          <i class="fas fa-book"></i>
          <span>База знаний</span>
        </a>
      </li>
      <li>
        <a href="https://portals.kkbsmp.ru/">
          <i class="fas fa-external-link-alt"></i>
          <span>Порталы</span>
        </a>
      </li>
      <li>
        <router-link to="/news" :class="{ active: $route.path === '/news' }">
          <i class="fas fa-newspaper"></i>
          <span>Новости</span>
        </router-link>
      </li>
      <li>
        <router-link to="/support" :class="{ active: $route.path === '/support' }">
          <i class="fas fa-headset"></i>
          <span>Поддержка</span>
        </router-link>
      </li>
    </ul>
  </sidebar>
</template>

<script setup>
</script>

<style scoped>
/* Стили из main.css для sidebar */
sidebar {
  background-color: #2c5aa0;
  color: white;
  padding: 1rem 0;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}
.nav-links {
  list-style: none;
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
.nav-links li a.active {
  background-color: rgba(255, 255, 255, 0.1);
  color: white;
  border-left: 4px solid #4fc3f7;
}
</style>