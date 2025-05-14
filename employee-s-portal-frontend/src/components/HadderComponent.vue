<template>
  <div class="header">
    <img :src="logo" alt="Логотип">
    <div class="logoText">
      <p>{{ orgName }}</p>
      <p class="portalName">{{ portalName }}</p>
    </div>
    <div class="profileMenu" @click="handleProfileClick">
      <img class="photoProfile" :src="logo" alt="Профиль">
      <p class="name">{{ userName }}</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import logo from '@/assets/logo.png';

const router = useRouter();

const orgName = ref('ГАУЗ ККБСМБ им. Подгорбунского');
const portalName = ref('Корпоративный портал');
const userName = ref('Иванов И.И.');

const handleProfileClick = () => {
  router.push('/user');
};
</script>

<style scoped>

.header {
  background-color: #FFFFFF;
  display: flex;
  align-items: center;
  height: 120px;
  padding: 10px;
}

img {
  width: 75px;
  height: 75px;
}

.logoText{
    width: 180px;
    margin-left: 10px;
    font-size: 20px;
    font-weight: bold;
}

.portalName{
    opacity: 50%;
    font-size: 15px;
}

.profileMenu{
    width: 300px;
    margin-top: 20px;
    margin-right: 10px;
    font-size: 20px;
    font-weight: bold;
    position: fixed;
    right: 0px;
    top: 0;
    display: flex;
    cursor: pointer;
}


.photoProfile{
    margin-top: 0;
    margin-right: 30px;
}

.name{
    font-size: 25px;
    margin-top: 20px;
} 

</style>