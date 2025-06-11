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
          <p class="elData">{{ userDataLocal.fullName || 'Не указан' }}</p>
          <p class="elData">{{ userDataLocal.position || 'Не указан' }}</p>
          <p class="elData">{{ userDataLocal.department || 'Не указан' }}</p>
          <p class="elData">{{ userDataLocal.employmentDate || 'Не указан' }}</p>
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
          <div class="elData">{{ userDataLocal.contacts?.email || 'Не указан' }}</div>
          <div class="elData">{{ userDataLocal.contacts?.phone || 'Не указан' }}</div>
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
          <div class="elData">
            <a v-if="userDataLocal.contacts?.telegram" :href="`https://t.me/${userDataLocal.contacts.telegram}`" target="_blank">
              {{ userDataLocal.contacts.telegram }}
            </a>
            <span v-else>Не указан</span>
          </div>
            <div class="elData">
              <span v-if="!userDataLocal.contacts?.boss">Не указан</span>
              <span
                v-else
                class="boss-link"
                @click.prevent="fetchBossData(userDataLocal.contacts.boss.id)"
              >
                {{ formatBossName(userDataLocal.contacts.boss) }}
              </span>
          </div>
        </div>
      </div>
    </div>

    <div class="button-container">
      <button 
        v-if="isShowingBoss" 
        @click="returnToUser"
        class="return-button"
      >
        ← Вернуться к пользователю
      </button>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref  } from 'vue';
import { useUserStore } from '@/stores/user';
import logo from '@/assets/logo.png';

const props = defineProps({
  id: {
    type: String,
    required: true
  }
})

const userStore = useUserStore();
const isShowingBoss = ref(false);

// Вычисляемые свойства для форматирования данных
const userDataLocal = ref({
  fullName: '',
  position: '',
  department: '',
  employmentDate: '',
  contacts: {
    email: '',
    phone: '',
    telegram: '',
    boss: null
  }
});

const updateUserData = (fullName,position, department,employmentDate,email,phone,telegram,boss) => {
  userDataLocal.value = {
    fullName: fullName,//`${userStore.userData.surname} ${userStore.userData.name} ${userStore.userData.patronymic}`,
    position: position,//userStore.userData.position,
    department: department || '',//userStore.userData.department?.name || '',
    employmentDate: formatDate(employmentDate),//formatDate(userStore.userData.employment_date),
    contacts: {
      email: email,//userStore.userData.email,
      phone: phone,// userStore.userData.phone,
      telegram:telegram,// userStore.userData.tg_link,
      boss: boss//userStore.userData.boss
    }
  };
};

const returnToUser = () => {
    isShowingBoss.value = false;
    updateUserData(`${userStore.userData.surname} ${userStore.userData.name} ${userStore.userData.patronymic}`,
      userStore.userData.position,
      userStore.userData.department.name,
      userStore.userData.employment_date,
      userStore.userData.email,
      userStore.userData.phone,
      userStore.userData.telegram,
      userStore.userData.boss
    )
};

const fetchBossData = async (bossId) => {
  try {

    const response = await fetch(`http://localhost:8080/api/v1/user/${bossId}`);

    if (!response.ok) throw new Error('Ошибка загрузки данных о пользователе');

    const data = await response.json();
    if (data.user) {
  
      updateUserData(`${data.user.surname} ${data.user.name} ${data.user.patronymic}`, data.user.position,data.user.department?.name,data.user.employment_date,data.user.email,data.user.phone
        ,data.user.tg_link, data.user.boss
      )

      isShowingBoss.value = true;

    } else {
      throw new Error('Ошибка загрузки данных о пользователе');
    }

  } catch (error) {
    console.error('Ошибка загрузки данных о пользователе:', error);
    alert('Не удалось загрузить данные о пользователе');
  }
}

const fetchUserData = async () =>{
  try {

    const response = await fetch(`http://localhost:8080/api/v1/user/${userStore.userData.id}`);

    if (!response.ok) throw new Error('Ошибка загрузки данных о пользователе');

    const data = await response.json();
    if (data.user) {
      userStore.setUserData(JSON.parse(JSON.stringify(data.user)));

      updateUserData(`${data.user.surname} ${data.user.name} ${data.user.patronymic}`, data.user.position,data.user.department?.name,data.user.employment_date,data.user.email,data.user.phone
        ,data.user.tg_link, data.user.boss
      )
    } else {
      throw new Error('Ошибка загрузки данных о пользователе');
    }

  } catch (error) {
    console.error('Ошибка загрузки данных о пользователе:', error);
    alert('Не удалось загрузить данные о пользователе');
  }
}

// Функция для форматирования даты
const formatDate = (dateString) => {
  if (!dateString) return '';
  const date = new Date(dateString);
  return date.toLocaleDateString('ru-RU');
};

const formatFullName = (surname, name, patronymic) => {
  if (!surname) return '';
  
  let formatted = surname;
  if (name) {
    formatted += ` ${name.charAt(0)}.`;
    if (patronymic) {
      formatted += `${patronymic.charAt(0)}.`;
    }
  }
  return formatted;
};

const formatBossName = (boss) => {
  if (!boss || !boss.surname) return '';
  return formatFullName(boss.surname, boss.name, boss.patronymic);
};


onMounted(() => {
  fetchUserData()
})

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
            margin-left: 40px;
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

        .button-container {
          display: flex;
          justify-content: flex-end; /* Выравнивание по правому краю */
          margin-right: 20px; /* Отступ справа */
          margin-top: 20px; /* Отступ сверху */
        }

        .return-button {
          margin-top: 40px;
          padding: 8px 16px;
          background-color: #5662DE;
          color: white;
          border: none;
          border-radius: 5px;
          font-size: 16px;
          cursor: pointer;
          transition: background-color 0.3s;
        }

        .return-button:hover {
          background-color: #454fa7;
        }
</style>