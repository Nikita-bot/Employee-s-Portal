<template>
  <div class="mainBlock">
    <!-- Весь блок будет динамическим -->
    <div class="infoBlock">
      <img class="photoProfileMain" :src="logo" alt="Фото профиля">
      <div class="info">
        <!-- Статичные метки -->
        <div class="infoName">
          <p class="elInfo">ФИО:</p>
          <p class="elInfo">Должность:</p>
          <p class="elInfo">Отделение:</p>
          <p class="elInfo">Работает с:</p>
        </div>
        <!-- Динамические данные -->
        <div class="infoData">
          <p class="elData">{{ userData.fullName }}</p>
          <p class="elData">{{ userData.position }}</p>
          <p class="elData">{{ userData.department }}</p>
          <p class="elData">{{ userData.employmentDate }}</p>
        </div>
      </div>
    </div>

    <div class="contactsBlock">
      <div class="contacts">
        <div class="contactsInfo">
          <!-- Статичные метки -->
          <div class="elInfo">Почта:</div>
          <div class="elInfo">Телефон:</div>
        </div>
        <div class="contactsData">
          <!-- Динамические данные -->
          <div class="elData">{{ userData.contacts.email }}</div>
          <div class="elData">{{ userData.contacts.phone }}</div>
        </div>
      </div>
      <div class="line"></div>
      <div class="contacts">
        <div class="contactsInfo">
          <!-- Статичные метки -->
          <div class="elInfo">Телеграмм:</div>
          <div class="elInfo">Руководитель:</div>
        </div>
        <div class="contactsData">
          <!-- Динамические данные -->
          <div class="elData" ><a v-if="userData.contacts.telegram"  :href="`https://t.me/${userData.contacts.telegram}`" target="_blank">{{ userData.contacts.telegram }}</a><span v-else>Не указан</span></div>
          <div class="elData">{{ userData.contacts.manager }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { useUserStore } from '@/stores/user';
import logo from '@/assets/logo.png';

const userStore = useUserStore();

// Вычисляемые свойства для форматирования данных
const userData = computed(() => {
  if (!userStore.userData) return {
    fullName: '',
    position: '',
    department: '',
    employmentDate: '',
    contacts: {
      email: '',
      phone: '',
      telegram: '',
      manager: ''
    }
  };

  return {
    fullName: `${userStore.userData.surname} ${userStore.userData.name} ${userStore.userData.patronymic}`,
    position: userStore.userData.position,
    department: userStore.userData.department?.name || '',
    employmentDate: formatDate(userStore.userData.employment_date),
    contacts: {
      email: userStore.userData.email,
      phone: userStore.userData.phone,
      telegram: userStore.userData.tg_link,
      manager: 'Петров П.П.' // Здесь можно добавить логику для получения руководителя
    }
  };
});

// Функция для форматирования даты
const formatDate = (dateString) => {
  if (!dateString) return '';
  const date = new Date(dateString);
  return date.toLocaleDateString('ru-RU');
};

</script>

<style scoped>
        .mainBlock{
            background-color: #FFFFFF;
            width: 1000px;
            height: 70%;
            position: absolute;
            border-radius: 15px;
            top: 7%;
            left: 25%;
            padding: 15px;
        }

        .infoBlock{
            display: flex;
            margin-top: 50px;
            margin-left: 20px;
            font-size: 25px;
        }

        .photoProfileMain{
            width: 300px;
            height: 300px;
        }

        .info{
            display: flex;
            margin-left: 50px;
        }
        .infoName{
            margin-right: 20px;
            text-align: right;
        }

        .elInfo{
            margin-top: 15px;
            color: #5662DE;
            
        }

        .elInfo:not(:first-child){
            margin-top: 30px;
        }

        .elData{
            margin-top: 15px;
        }

        .elData:not(:first-child){
            margin-top: 30px;
        }

        .contactsBlock{
            display: flex;
            font-size: 25px;
            margin-top: 50px;
            margin-left: 20px;
        }

        .contacts{
            display: flex;
            margin-right: 50px;
            margin-left: 50px;
        }

        .contacts:not(:first-child){
          margin-right: 0;
        }

        .contactsInfo{
            margin-right: 20px;
        }

        .line{
            width: 2px;
            height: 90px;
            background-color: #5662DE;
            margin-top: 30px;
        }
</style>