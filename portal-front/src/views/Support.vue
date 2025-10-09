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

    <div class="main-container">
      <sidebar>
        <ul class="nav-links">
          <li><router-link to="/">Главная</router-link></li>
          <li><router-link to="/tasks">Задачи</router-link></li>
          <li><a href="#">ЭДО</a></li>
          <li><a href="#">База знаний</a></li>
          <li><a href="#">Порталы</a></li>
          <li><router-link to="/news">Новости</router-link></li>
          <li><router-link to="/support" class="active">Поддержка</router-link></li>
        </ul>
      </sidebar>

      <main class="content-area">
        <div class="page-header">
          <h1>Техническая поддержка</h1>
        </div>

        <div class="support-container">
          <div class="support-info">
            <h3>Создание заявки в техническую поддержку</h3>
            <p>Опишите проблему как можно подробнее. В зависимости от типа проблемы могут появиться дополнительные поля для уточнения.</p>
          </div>

          <form @submit.prevent="submitSupportRequest">
            <div class="form-group">
              <label for="problemType">Тип проблемы *</label>
              <select 
                id="problemType" 
                v-model="supportRequest.problemType" 
                required 
                @change="checkAdditionalFields"
              >
                <option value="">Выберите тип проблемы</option>
                <option value="printer">Проблема с принтером/картриджем</option>
                <option value="computer">Проблема с компьютером</option>
                <option value="software">Проблема с программой/Ариадной</option>
                <option value="network">Проблема с сетью/интернетом</option>
                <option value="other">Другая проблема</option>
              </select>
            </div>

            <div class="form-group">
              <label for="problemDescription">Подробное описание *</label>
              <textarea 
                id="problemDescription" 
                v-model="supportRequest.problemDescription"
                placeholder="Опишите проблему подробно, укажите что именно не работает, когда началась проблема..." 
                required
              ></textarea>
            </div>

            <!-- Дополнительные поля для принтера -->
            <div 
              class="additional-fields" 
              v-show="showPrinterFields"
            >
              <h4>Информация о принтере</h4>
              <div class="form-row">
                <div class="form-group">
                  <label for="printerModel">Какой принтер</label>
                  <select id="printerModel" v-model="supportRequest.printerModel">
                    <option value="">Выберите модель</option>
                    <option value="hp-laserjet">HP LaserJet Pro</option>
                    <option value="canon-mf">Canon i-SENSYS MF</option>
                    <option value="epson-workforce">Epson Workforce</option>
                    <option value="xerox-versalink">Xerox VersaLink</option>
                    <option value="other">Другая модель</option>
                  </select>
                </div>
                <div class="form-group">
                  <label for="printerLocation">Где стоит принтер</label>
                  <select id="printerLocation" v-model="supportRequest.printerLocation">
                    <option value="">Выберите местоположение</option>
                    <option value="reception-1">Регистратура 1 этаж</option>
                    <option value="reception-2">Регистратура 2 этаж</option>
                    <option value="cardiology">Кардиологическое отделение</option>
                    <option value="surgery">Хирургическое отделение</option>
                    <option value="therapy">Терапевтическое отделение</option>
                    <option value="administration">Администрация</option>
                    <option value="other">Другое место</option>
                  </select>
                </div>
              </div>
            </div>

            <!-- Дополнительные поля для компьютера -->
            <div 
              class="additional-fields" 
              v-show="showComputerFields"
            >
              <h4>Информация о компьютере</h4>
              <div class="form-group">
                <label for="computerLocation">Где стоит компьютер</label>
                <select id="computerLocation" v-model="supportRequest.computerLocation">
                  <option value="">Выберите местоположение</option>
                  <option value="reception-1">Регистратура 1 этаж</option>
                  <option value="reception-2">Регистратура 2 этаж</option>
                  <option value="doctor-office-305">Кабинет 305 (Кардиология)</option>
                  <option value="doctor-office-412">Кабинет 412 (Хирургия)</option>
                  <option value="nurses-station">Пост медсестер</option>
                  <option value="administration">Администрация</option>
                  <option value="other">Другое место</option>
                </select>
              </div>
            </div>

            <!-- Дополнительные поля для программ -->
            <div 
              class="additional-fields" 
              v-show="showSoftwareFields"
            >
              <h4>Информация о программном обеспечении</h4>
              <div class="form-group">
                <label for="computerName">Имя компьютера *</label>
                <input 
                  type="text" 
                  id="computerName" 
                  v-model="supportRequest.computerName"
                  placeholder="Например: PC-ADMIN-01 или USER-305"
                >
              </div>
            </div>

            <div class="urgent-checkbox">
              <input type="checkbox" id="urgentRequest" v-model="supportRequest.urgent">
              <label for="urgentRequest">Срочная заявка (критическая проблема)</label>
            </div>

            <button 
              type="submit" 
              class="btn btn-primary" 
              :disabled="submitting"
            >
              {{ submitting ? 'Отправка...' : 'Создать заявку' }}
            </button>

            <div 
              class="success-message" 
              v-show="showSuccessMessage"
            >
              Заявка успешно создана! Номер вашей заявки: <strong>{{ ticketNumber }}</strong>
            </div>
          </form>
        </div>
      </main>
    </div>

    <footer>
      <p>© 2023 Городская Больница. Корпоративная информационная система. Версия 2.1</p>
    </footer>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()

// Состояние
const dropdownOpen = ref(false)
const submitting = ref(false)
const showSuccessMessage = ref(false)
const ticketNumber = ref('')

// Форма поддержки
const supportRequest = reactive({
  problemType: '',
  problemDescription: '',
  printerModel: '',
  printerLocation: '',
  computerLocation: '',
  computerName: '',
  urgent: false
})

// Вычисляемые свойства для отображения дополнительных полей
const showPrinterFields = computed(() => {
  return supportRequest.problemType === 'printer'
})

const showComputerFields = computed(() => {
  return supportRequest.problemType === 'computer'
})

const showSoftwareFields = computed(() => {
  return supportRequest.problemType === 'software'
})

// Методы
const toggleDropdown = () => {
  dropdownOpen.value = !dropdownOpen.value
}

const closeDropdown = () => {
  dropdownOpen.value = false
}

const logout = () => {
  userStore.clearUserData()
  router.push('/login')
  closeDropdown()
}

const checkAdditionalFields = () => {
  // Логика автоматически обрабатывается через computed свойства
}

const submitSupportRequest = async () => {
  // Валидация
  if (!supportRequest.problemType || !supportRequest.problemDescription) {
    alert('Пожалуйста, заполните все обязательные поля')
    return
  }

  // Дополнительная валидация для программных проблем
  if (supportRequest.problemType === 'software' && !supportRequest.computerName) {
    alert('Для программных проблем обязательно укажите имя компьютера')
    return
  }

  submitting.value = true

  try {
    // Имитация отправки на сервер
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    // Генерация номера заявки
    ticketNumber.value = 'T' + Date.now().toString().slice(-6)
    showSuccessMessage.value = true

    // Логирование данных (в реальном приложении - отправка на сервер)
    console.log('Создана заявка поддержки:', {
      number: ticketNumber.value,
      ...supportRequest,
      user: userStore.userData?.name
    })

    // Сброс формы через 5 секунд
    setTimeout(() => {
      resetForm()
    }, 5000)

  } catch (error) {
    alert('Ошибка при создании заявки')
  } finally {
    submitting.value = false
  }
}

const resetForm = () => {
  Object.assign(supportRequest, {
    problemType: '',
    problemDescription: '',
    printerModel: '',
    printerLocation: '',
    computerLocation: '',
    computerName: '',
    urgent: false
  })
  showSuccessMessage.value = false
  ticketNumber.value = ''
}

// Обработчики событий
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
/* Стили из вашего support.html */
.support-container {
  max-width: 600px;
  margin: 0 auto;
}

.support-info {
  background: #f8f9fa;
  padding: 1.5rem;
  border-radius: 8px;
  margin-bottom: 2rem;
  border-left: 4px solid #2c5aa0;
}

.support-info h3 {
  color: #2c5aa0;
  margin-bottom: 0.5rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: #555;
  font-weight: 500;
}

.form-group input,
.form-group select,
.form-group textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.3s ease;
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #2c5aa0;
}

.form-group textarea {
  resize: vertical;
  min-height: 120px;
}

.additional-fields {
  background: #f8f9fa;
  padding: 1.5rem;
  border-radius: 6px;
  margin-top: 1rem;
  border-left: 3px solid #f39c12;
}

.additional-fields h4 {
  color: #f39c12;
  margin-bottom: 1rem;
  font-weight: 500;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
}

.urgent-checkbox {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.urgent-checkbox input {
  width: auto;
}

.success-message {
  background: #e8f5e8;
  color: #27ae60;
  padding: 1rem;
  border-radius: 6px;
  margin-top: 1rem;
}

@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }
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

sidebar {
  background-color: #2c5aa0;
  color: white;
  padding: 2rem 0;
}

footer {
  background-color: #1a1a1a;
  color: #999;
  text-align: center;
  padding: 1.5rem;
  font-size: 0.9rem;
}
</style>