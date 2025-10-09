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
          <li><router-link to="/support">Поддержка</router-link></li>
        </ul>
      </sidebar>

      <main class="content-area">
        <div class="page-header">
          <h1>Мой профиль</h1>
        </div>

        <div class="profile-container">
          <!-- Шапка профиля -->
          <div class="profile-header">
            <img :src="userStore.userData?.avatar" alt="Аватар" class="avatar-large">
            <div class="profile-info">
              <h2>{{ userStore.userData?.name }}</h2>
              <div class="profile-position">{{ userStore.userData?.position }}</div>
              <div class="profile-department">{{ userStore.userData?.department }}</div>
            </div>
          </div>

          <!-- Основная сетка -->
          <div class="profile-grid">
            <!-- Личная информация -->
            <div class="info-section">
              <div class="section-header">
                <h3>Личная информация</h3>
                <button class="edit-btn" @click="editPersonalInfo">Редактировать</button>
              </div>
              <div class="section-body">
                <div class="info-item">
                  <span class="info-label">Дата рождения:</span>
                  <span class="info-value">15.03.1980</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Телефон:</span>
                  <span class="info-value">+7 (912) 345-67-89</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Email:</span>
                  <span class="info-value">a.ivanov@hospital.ru</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Кабинет:</span>
                  <span class="info-value">305</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Стаж работы:</span>
                  <span class="info-value">15 лет</span>
                </div>
              </div>
            </div>

            <!-- Профессиональная информация -->
            <div class="info-section">
              <div class="section-header">
                <h3>Профессиональная информация</h3>
                <button class="edit-btn" @click="editProfessionalInfo">Редактировать</button>
              </div>
              <div class="section-body">
                <div class="info-item">
                  <span class="info-label">Отдел:</span>
                  <span class="info-value">Кардиология</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Должность:</span>
                  <span class="info-value">Ведущий специалист</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Категория:</span>
                  <span class="info-value">Высшая</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Ученая степень:</span>
                  <span class="info-value">Кандидат медицинских наук</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Дата приема:</span>
                  <span class="info-value">10.08.2008</span>
                </div>
              </div>
            </div>

            <!-- Статистика -->
            <div class="info-section">
              <div class="section-header">
                <h3>Статистика за месяц</h3>
              </div>
              <div class="section-body">
                <div class="stats-grid">
                  <div class="stat-card">
                    <div class="stat-number">42</div>
                    <div class="stat-label">Приемов</div>
                  </div>
                  <div class="stat-card">
                    <div class="stat-number">18</div>
                    <div class="stat-label">Операций</div>
                  </div>
                  <div class="stat-card">
                    <div class="stat-number">96%</div>
                    <div class="stat-label">Успешность</div>
                  </div>
                </div>
              </div>
            </div>

            <!-- ЭЦП и безопасность -->
            <div class="info-section">
              <div class="section-header">
                <h3>ЭЦП и безопасность</h3>
              </div>
              <div class="section-body">
                <div class="info-item">
                  <span class="info-label">ЭЦП:</span>
                  <span class="info-value">Активна до 15.12.2024</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Telegram:</span>
                  <span class="info-value">Привязан</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Последний вход:</span>
                  <span class="info-value">Сегодня, {{ currentTime }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Двухфакторная аутентификация:</span>
                  <span class="info-value">Включена</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Календарь -->
          <div class="calendar-section">
            <div class="section-header">
              <h3>Календарь</h3>
              <button class="edit-btn" @click="showFullCalendar">Просмотреть все</button>
            </div>
            <div class="calendar-header">
              <span>{{ currentMonth }}</span>
              <div>
                <button class="btn btn-outline" style="padding: 0.25rem 0.5rem; margin-right: 0.5rem;">←</button>
                <button class="btn btn-outline" style="padding: 0.25rem 0.5rem;">→</button>
              </div>
            </div>
            <div class="calendar-grid">
              <div 
                v-for="day in calendarDays" 
                :key="day.date"
                class="calendar-day" 
                :class="{
                  'header': day.isHeader,
                  'other-month': !day.isCurrentMonth,
                  'today': day.isToday
                }"
              >
                {{ day.display }}
                <div v-if="day.event" class="calendar-event">{{ day.event }}</div>
              </div>
            </div>
          </div>

          <!-- Аналитика и быстрые действия -->
          <div class="analytics-section">
            <div class="section-header">
              <h3>Аналитика и действия</h3>
            </div>
            <div class="analytics-grid">
              <div class="chart-placeholder">
                📊 График рабочей нагрузки<br>
                <small>Здесь будет отображаться ваша активность и статистика</small>
              </div>
              <div class="quick-actions">
                <router-link to="/tasks" class="action-btn">
                  <div class="action-icon">📋</div>
                  <div class="action-text">
                    <h4>Мои отчеты</h4>
                    <p>Просмотр и создание отчетов</p>
                  </div>
                </router-link>
                <router-link to="/tasks" class="action-btn">
                  <div class="action-icon">✅</div>
                  <div class="action-text">
                    <h4>Задачи</h4>
                    <p>Текущие и назначенные</p>
                  </div>
                </router-link>
                <a href="#" class="action-btn">
                  <div class="action-icon">📚</div>
                  <div class="action-text">
                    <h4>База знаний</h4>
                    <p>Протоколы и стандарты</p>
                  </div>
                </a>
                <router-link to="/support" class="action-btn">
                  <div class="action-icon">🛠️</div>
                  <div class="action-text">
                    <h4>Поддержка</h4>
                    <p>Техническая помощь</p>
                  </div>
                </router-link>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>

    <!-- Модальное окно редактирования личной информации -->
    <div class="modal" :class="{ show: showPersonalModal }">
      <div class="modal-content">
        <div class="modal-header">
          <h2>Редактирование личной информации</h2>
          <button class="close-btn" @click="closePersonalModal">&times;</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label for="editPhone">Телефон</label>
            <input type="tel" id="editPhone" v-model="editForm.phone">
          </div>
          <div class="form-group">
            <label for="editEmail">Email</label>
            <input type="email" id="editEmail" v-model="editForm.email">
          </div>
          <div class="form-group">
            <label for="editCabinet">Кабинет</label>
            <input type="text" id="editCabinet" v-model="editForm.cabinet">
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-outline" @click="closePersonalModal">Отмена</button>
          <button class="btn btn-primary" @click="savePersonalInfo">Сохранить</button>
        </div>
      </div>
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
const showPersonalModal = ref(false)

// Форма редактирования
const editForm = reactive({
  phone: '+7 (912) 345-67-89',
  email: 'a.ivanov@hospital.ru',
  cabinet: '305'
})

// Вычисляемые свойства
const currentTime = computed(() => {
  return new Date().toLocaleTimeString('ru-RU', { 
    hour: '2-digit', 
    minute: '2-digit' 
  })
})

const currentMonth = computed(() => {
  return new Date().toLocaleDateString('ru-RU', { 
    month: 'long', 
    year: 'numeric' 
  })
})

const calendarDays = computed(() => {
  const days = []
  const today = new Date()
  const year = today.getFullYear()
  const month = today.getMonth()
  
  // Заголовки дней недели
  const daysOfWeek = ['Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб', 'Вс']
  daysOfWeek.forEach(day => {
    days.push({ display: day, isHeader: true })
  })
  
  // Дни месяца (демо-данные)
  for (let i = 1; i <= 31; i++) {
    const isToday = i === today.getDate()
    let event = ''
    
    if (i === 6) event = 'Прием 10:00'
    else if (i === 12) event = 'Семинар 14:00'
    else if (i === 15) event = 'Консилиум 11:00'
    
    days.push({ 
      display: i, 
      isCurrentMonth: true,
      isToday,
      event
    })
  }
  
  return days
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

const editPersonalInfo = () => {
  showPersonalModal.value = true
}

const closePersonalModal = () => {
  showPersonalModal.value = false
}

const editProfessionalInfo = () => {
  alert('Редактирование профессиональной информации')
}

const showFullCalendar = () => {
  alert('Открытие полного календаря')
}

const savePersonalInfo = () => {
  alert('Личная информация сохранена')
  closePersonalModal()
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
/* Стили из вашего profile.html */
.profile-container {
  max-width: 1000px;
  margin: 0 auto;
}

.profile-header {
  display: flex;
  align-items: center;
  gap: 2rem;
  margin-bottom: 2rem;
  padding: 2rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  color: white;
}

.avatar-large {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  object-fit: cover;
  border: 4px solid rgba(255, 255, 255, 0.3);
}

.profile-info h2 {
  font-weight: 300;
  margin-bottom: 0.5rem;
  font-size: 1.8rem;
}

.profile-position {
  font-size: 1.2rem;
  margin-bottom: 0.5rem;
  opacity: 0.9;
}

.profile-department {
  background: rgba(255, 255, 255, 0.2);
  padding: 0.5rem 1rem;
  border-radius: 20px;
  font-size: 0.9rem;
  display: inline-block;
}

.profile-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
  margin-bottom: 2rem;
}

.info-section {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  border: 1px solid #e1e5e9;
  overflow: hidden;
}

.section-header {
  background: #f8f9fa;
  padding: 1.25rem;
  border-bottom: 1px solid #e1e5e9;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.section-header h3 {
  color: #2c5aa0;
  font-weight: 500;
  font-size: 1.1rem;
}

.edit-btn {
  background: #2c5aa0;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.85rem;
  transition: background 0.3s ease;
}

.edit-btn:hover {
  background: #1e3d6f;
}

.section-body {
  padding: 1.5rem;
}

.info-item {
  display: flex;
  justify-content: space-between;
  padding: 0.75rem 0;
  border-bottom: 1px solid #f0f0f0;
}

.info-item:last-child {
  border-bottom: none;
}

.info-label {
  color: #666;
  font-weight: 500;
  min-width: 140px;
}

.info-value {
  color: #333;
  text-align: right;
  flex: 1;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 1rem;
  margin-top: 1rem;
}

.stat-card {
  background: #f8f9fa;
  padding: 1.25rem;
  border-radius: 6px;
  text-align: center;
  border-left: 4px solid #2c5aa0;
}

.stat-number {
  font-size: 2rem;
  font-weight: 300;
  color: #2c5aa0;
  line-height: 1;
}

.stat-label {
  color: #666;
  font-size: 0.85rem;
  margin-top: 0.5rem;
}

.calendar-section {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  border: 1px solid #e1e5e9;
  margin-bottom: 2rem;
}

.calendar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 1.5rem;
  background: #f8f9fa;
  border-bottom: 1px solid #e1e5e9;
}

.calendar-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 1px;
  background: #e1e5e9;
  border: 1px solid #e1e5e9;
}

.calendar-day {
  background: white;
  padding: 0.75rem;
  min-height: 80px;
  font-size: 0.9rem;
}

.calendar-day.header {
  background: #f8f9fa;
  font-weight: 500;
  text-align: center;
  color: #666;
}

.calendar-day.other-month {
  color: #ccc;
  background: #fafafa;
}

.calendar-day.today {
  background: #e3f2fd;
  border: 2px solid #2c5aa0;
}

.calendar-event {
  background: #2c5aa0;
  color: white;
  padding: 0.25rem 0.5rem;
  border-radius: 4px;
  font-size: 0.75rem;
  margin-top: 0.25rem;
  cursor: pointer;
}

.analytics-section {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  border: 1px solid #e1e5e9;
  margin-bottom: 2rem;
}

.analytics-grid {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 1.5rem;
  padding: 1.5rem;
}

.chart-placeholder {
  background: #f8f9fa;
  border-radius: 6px;
  padding: 2rem;
  text-align: center;
  color: #666;
  border: 2px dashed #e1e5e9;
}

.quick-actions {
  display: grid;
  gap: 1rem;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem;
  background: #f8f9fa;
  border: 1px solid #e1e5e9;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.3s ease;
  text-decoration: none;
  color: #333;
}

.action-btn:hover {
  background: #e9ecef;
  transform: translateY(-2px);
}

.action-icon {
  width: 40px;
  height: 40px;
  background: #2c5aa0;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 1.2rem;
}

.action-text h4 {
  margin: 0;
  color: #2c5aa0;
  font-weight: 500;
}

.action-text p {
  margin: 0;
  color: #666;
  font-size: 0.85rem;
}

@media (max-width: 768px) {
  .profile-grid {
    grid-template-columns: 1fr;
  }
  .profile-header {
    flex-direction: column;
    text-align: center;
    gap: 1rem;
  }
  .stats-grid {
    grid-template-columns: 1fr;
  }
  .analytics-grid {
    grid-template-columns: 1fr;
  }
  .calendar-grid {
    grid-template-columns: repeat(7, 1fr);
    font-size: 0.8rem;
  }
  .calendar-day {
    min-height: 60px;
    padding: 0.5rem;
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