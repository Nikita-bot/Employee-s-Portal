<template>
  <div class="mainBlock">
    <!-- Весь блок будет динамическим -->
    <div class="infoBlock">
      <img class="photoProfileMain" :src="logo" alt="Фото профиля">
      <div class="info">
        <!-- Объединенные блоки данных -->
        <div class="infoItem">
          <span class="label">ФИО:</span>
          <span class="value">{{ userDataLocal.fullName || 'Не указан' }}</span>
        </div>
        <div class="infoItem">
          <span class="label">Должность:</span>
          <span class="value">
            {{ userDataLocal.position || 'Не указан' }}
            <span v-if="userDataLocal.other_position?.length" 
              class="position-info-icon" 
              @mouseover="showTooltip = true"
              @mouseleave="showTooltip = false">
              <i class="bi bi-info-circle"></i>
              <div v-if="showTooltip" class="position-tooltip">
                <div v-for="(pos, index) in userDataLocal.other_position" :key="index" class="position-item">
                  <strong>{{ pos.department }}:</strong> {{ pos.name }}
                </div>
              </div>
            </span>
          </span>
        </div>
        <div class="infoItem">
          <span class="label">Отделение:</span>
          <span class="value">{{ userDataLocal.department || 'Не указан' }}</span>
        </div>
        <div class="infoItem">
          <span class="label">Работает с:</span>
          <span class="value">{{ userDataLocal.employmentDate || 'Не указан' }}</span>
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
              <span v-if="!userDataLocal.contacts?.boss">Каменева Е.А.</span>
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
import "bootstrap-icons/font/bootstrap-icons.css";
import { onMounted, ref  } from 'vue';
import { useUserStore } from '@/stores/user';
import { useRoute } from 'vue-router'
import logo from '@/assets/logo.png';

const props = defineProps({
  id: {
    type: String,
    required: true
  }
})

const userStore = useUserStore();
if (!userStore.userData.id) {
  router.push('/login')
}
const route = useRoute()
const isShowingBoss = ref(false);
const showTooltip = ref(false);

// Вычисляемые свойства для форматирования данных
const userDataLocal = ref({
  fullName: '',
  position: '',
  department: '',
  employmentDate: '',
  other_position: [],
  contacts: {
    email: '',
    phone: '',
    telegram: '',
    boss: null
  }
});

const updateUserData = (fullName,position, department,employmentDate,email,phone,telegram,boss, other_position) => {
  userDataLocal.value = {
    fullName: fullName,
    position: position,
    department: department || '',
    employmentDate: formatDate(employmentDate),
    other_position: other_position,
    contacts: {
      email: email,
      phone: phone,
      telegram:telegram,
      boss: boss
    }
  };
};

const returnToUser = () => {
    isShowingBoss.value = false;
    updateUserData(`${userStore.userData.surname} ${userStore.userData.name} ${userStore.userData.patronymic}`,
      userStore.userData.employee.position,
      userStore.userData.employee.department,
      userStore.userData.start_date,
      userStore.userData.email,
      userStore.userData.phone,
      userStore.userData.tg_link,
      userStore.userData.employee.boss,
      userStore.userData.other_position
    )
};

const fetchBossData = async (bossId) => {

  if (!bossId){
    return 
  }
  try {

    const response = await fetch(`/api/v1/user/${bossId}`);

    if (!response.ok) throw new Error('Ошибка загрузки данных о пользователе');

    const data = await response.json();
    if (data.user) {
  
      updateUserData(`${data.user.surname} ${data.user.name} ${data.user.patronymic}`, data.user.employee.position, data.user.employee.department, data.user.employee.start_date,data.user.email,data.user.phone
        ,data.user.tg_link, data.user.employee.boss.name, data.user.employee.other_position
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

    const response = await fetch(`/api/v1/user/${route.params.id}`);

    if (!response.ok) throw new Error('Ошибка загрузки данных о пользователе');

    const data = await response.json();
    if (data.user) {
      userStore.setUserData(JSON.parse(JSON.stringify(data.user)));
      console.log(data.user)
      updateUserData(`${data.user.surname} ${data.user.name} ${data.user.patronymic}`, data.user.employee.position, data.user.employee.department, data.user.employee.start_date,data.user.email,data.user.phone
        ,data.user.tg_link, data.user.employee.boss.name, data.user.employee.other_position
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
  return dateString;
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


onMounted(async () => {
  if (!userStore.userData || Object.keys(userStore.userData).length === 0) {
    try {
      fetchUserData();
      console.log('Данные пользователя успешно загружены');
    } catch (error) {
      console.error('Ошибка загрузки данных:', error);
    }
  } else {
    console.log('Данные уже загружены');
    updateUserData(`${userStore.userData.surname} ${userStore.userData.name} ${userStore.userData.patronymic}`, userStore.userData.employee.position, userStore.userData.employee.department, userStore.userData.employee.start_date,userStore.userData.email,userStore.userData.phone
        ,userStore.userData.tg_link, userStore.userData.employee.boss.name, userStore.userData.employee.other_position
      )
  }
});

</script>

<style scoped>
        .mainBlock{
            background-color: #FFFFFF;
            width: 60%;
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

        .info {
          margin-left: 40px;
        }

        .infoItem {
          display: flex;
          margin-top: 15px;
        }

        .infoItem:not(:first-child) {
          margin-top: 30px;
        }

        .label {
          color: #5662DE;
          width: 150px; /* Фиксированная ширина для выравнивания */
          text-align: right;
          margin-right: 50px;
          flex-shrink: 0; /* Запрещаем сжатие метки */
        }

        .value {
          flex-grow: 1; /* Занимает все доступное пространство */
          word-break: break-word; /* Перенос длинных значений */
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

        .position-info-icon {
          position: relative;
          display: inline-block;
          margin-left: 5px;
          cursor: help;
          color: #3498db;
        }

        .position-tooltip {
          position: absolute;
          bottom: 100%;
          left: 50%;
          transform: translateX(-50%);
          background: white;
          border: 1px solid #ddd;
          border-radius: 4px;
          padding: 10px;
          box-shadow: 0 2px 10px rgba(0,0,0,0.1);
          z-index: 100;
          min-width: 250px;
          font-size: 14px;
        }

        .position-item {
          padding: 5px 0;
          border-bottom: 1px solid #eee;
        }

        .position-item:last-child {
          border-bottom: none;
        }

        .position-info-icon .bi {
          transition: all 0.2s;
        }

        .position-info-icon:hover .bi {
          color: #2980b9;
        }
</style>