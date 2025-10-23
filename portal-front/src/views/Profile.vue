<template>
  <div class="profile-wrapper">
    <AppHeader />
    
    <div class="main-container">
      <AppSidebar />
      
      <main class="content-area">
        <div class="page-header">
          <h1>Мой профиль</h1>
        </div>

        <div class="profile-container">
          <!-- Шапка профиля -->
          <div class="profile-header">
            <img :src="user.avatar" alt="Аватар" class="avatar-large">
            <div class="profile-info">
              <h2>{{ user.name }}</h2>
              <div class="profile-position">{{ user.position }}</div>
              <div class="profile-department">{{ user.department }}</div>
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
                  <span class="info-value">{{ user.birthDate }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Телефон:</span>
                  <span class="info-value">{{ user.phone }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Email:</span>
                  <span class="info-value">{{ user.email }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Кабинет:</span>
                  <span class="info-value">{{ user.office }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Стаж работы:</span>
                  <span class="info-value">{{ user.experience }}</span>
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
                  <span class="info-value">{{ user.department }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Должность:</span>
                  <span class="info-value">{{ user.position }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Категория:</span>
                  <span class="info-value">{{ user.category }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Ученая степень:</span>
                  <span class="info-value">{{ user.degree }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Дата приема:</span>
                  <span class="info-value">{{ user.hireDate }}</span>
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
                    <div class="stat-number">{{ stats.appointments }}</div>
                    <div class="stat-label">Приемов</div>
                  </div>
                  <div class="stat-card">
                    <div class="stat-number">{{ stats.operations }}</div>
                    <div class="stat-label">Операций</div>
                  </div>
                  <div class="stat-card">
                    <div class="stat-number">{{ stats.successRate }}</div>
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
                  <span class="info-value">{{ security.ecp }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Telegram:</span>
                  <span class="info-value">{{ security.telegram }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Последний вход:</span>
                  <span class="info-value">{{ security.lastLogin }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Двухфакторная аутентификация:</span>
                  <span class="info-value">{{ security.twoFactor }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Календарь -->
          <div class="calendar-section">
            <div class="section-header">
              <h3>Календарь</h3>
              <button class="edit-btn" @click="showCalendar">Просмотреть все</button>
            </div>
            <div class="calendar-header">
              <span>{{ currentMonth }}</span>
              <div>
                <button class="btn btn-outline" style="padding: 0.25rem 0.5rem; margin-right: 0.5rem;">←</button>
                <button class="btn btn-outline" style="padding: 0.25rem 0.5rem;">→</button>
              </div>
            </div>
            <div class="calendar-grid">
              <div v-for="day in calendar.daysOfWeek" :key="day" class="calendar-day header">
                {{ day }}
              </div>
              <div v-for="day in calendar.days" :key="day.number" 
                   :class="['calendar-day', day.class]">
                {{ day.number }}
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
                <a href="#" class="action-btn">
                  <div class="action-icon">📋</div>
                  <div class="action-text">
                    <h4>Мои отчеты</h4>
                    <p>Просмотр и создание отчетов</p>
                  </div>
                </a>
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
    <div class="modal" :class="{ show: showPersonalModal }" @click="closeEditPersonalModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h2>Редактирование личной информации</h2>
          <button class="close-btn" @click="closeEditPersonalModal">&times;</button>
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
            <input type="text" id="editCabinet" v-model="editForm.office">
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-outline" @click="closeEditPersonalModal">Отмена</button>
          <button class="btn btn-primary" @click="savePersonalInfo">Сохранить</button>
        </div>
      </div>
    </div>

    <AppFooter />
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { useRouter } from 'vue-router'
import AppHeader from '@/components/AppHeader.vue'
import AppSidebar from '@/components/AppSidebar.vue'
import AppFooter from '@/components/AppFooter.vue'

const router = useRouter()
const showPersonalModal = ref(false)

const user = reactive({
  name: 'Иванов Александр Сергеевич',
  avatar: '/default-avatar.png',
  position: 'Врач-кардиолог высшей категории',
  department: 'Кардиологическое отделение',
  birthDate: '15.03.1980',
  phone: '+7 (912) 345-67-89',
  email: 'a.ivanov@hospital.ru',
  office: '305',
  experience: '15 лет',
  category: 'Высшая',
  degree: 'Кандидат медицинских наук',
  hireDate: '10.08.2008'
})

const stats = reactive({
  appointments: 42,
  operations: 18,
  successRate: '96%'
})

const security = reactive({
  ecp: 'Активна до 15.12.2024',
  telegram: 'Привязан',
  lastLogin: 'Сегодня, 10:30',
  twoFactor: 'Включена'
})

const editForm = reactive({
  phone: user.phone,
  email: user.email,
  office: user.office
})

const calendar = reactive({
  daysOfWeek: ['Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб', 'Вс'],
  days: [
    { number: 1, class: 'other-month' }, { number: 2, class: 'other-month' }, 
    { number: 3, class: 'other-month' }, { number: 4, class: 'other-month' }, 
    { number: 5, class: 'other-month' }, { number: 6, class: 'today', event: 'Прием 10:00' },
    { number: 7, class: '' }, { number: 8, class: '' }, { number: 9, class: '' }, 
    { number: 10, class: '' }, { number: 11, class: '' }, { number: 12, class: '', event: 'Семинар 14:00' },
    { number: 13, class: '' }, { number: 14, class: '' }, { number: 15, class: '', event: 'Консилиум 11:00' },
    // ... остальные дни
  ]
})

const currentMonth = computed(() => {
  return new Date().toLocaleDateString('ru-RU', { month: 'long', year: 'numeric' })
})

const editPersonalInfo = () => {
  editForm.phone = user.phone
  editForm.email = user.email
  editForm.office = user.office
  showPersonalModal.value = true
}

const closeEditPersonalModal = () => {
  showPersonalModal.value = false
}

const editProfessionalInfo = () => {
  alert('Редактирование профессиональной информации')
}

const showCalendar = () => {
  alert('Открытие полного календаря')
}

const savePersonalInfo = () => {
  user.phone = editForm.phone
  user.email = editForm.email
  user.office = editForm.office
  alert('Личная информация сохранена')
  closeEditPersonalModal()
}
</script>

<style scoped>
.profile-wrapper {
  display: grid;
  grid-template-rows: auto 1fr auto;
  min-height: 100vh;
  background-color: #f8f9fa;
  color: #333;
  line-height: 1.6;
}

.main-container {
  display: grid;
  grid-template-columns: 250px 1fr;
  gap: 0;
}

.content-area {
  padding: 2rem;
  background-color: #fff;
  min-height: calc(100vh - 140px);
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 2rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid #eee;
}

.page-header h1 {
  font-weight: 300;
  color: #2c5aa0;
}

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

/* Modal Styles */
.modal {
  display: none;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  z-index: 1000;
  align-items: center;
  justify-content: center;
}

.modal.show {
  display: flex;
}

.modal-content {
  background: white;
  border-radius: 8px;
  width: 90%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
}

.modal-header {
  padding: 1.5rem;
  border-bottom: 1px solid #e1e5e9;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #f8f9fa;
  border-radius: 8px 8px 0 0;
}

.modal-header h2 {
  color: #2c5aa0;
  font-weight: 500;
  font-size: 1.25rem;
}

.close-btn {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: #666;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 4px;
}

.close-btn:hover {
  background: #e9ecef;
}

.modal-body {
  padding: 1.5rem;
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
  transition: all 0.3s ease;
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #2c5aa0;
  box-shadow: 0 0 0 3px rgba(44, 90, 160, 0.1);
}

.modal-footer {
  padding: 1rem 1.5rem;
  border-top: 1px solid #e1e5e9;
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  transition: all 0.3s ease;
  min-width: 100px;
}

.btn-primary {
  background: #2c5aa0;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background: #1e3d6f;
}

.btn-outline {
  background: transparent;
  border: 1px solid #2c5aa0;
  color: #2c5aa0;
}

.btn-outline:hover {
  background: #2c5aa0;
  color: white;
}

@media (max-width: 768px) {
  .main-container {
    grid-template-columns: 1fr;
  }
  
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
</style>