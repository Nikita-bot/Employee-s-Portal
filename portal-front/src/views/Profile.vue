<template>
  <div class="profile-wrapper">
    <AppHeader />
    
    <div class="main-container">
      <AppSidebar />
      
      <main class="content-area">
        <div class="page-header">
          <h1>Мой профиль</h1>
          <button class="btn btn-outline" @click="changePassword" v-if="!showPasswordForm">
            Сменить пароль
          </button>
        </div>

        <!-- Форма смены пароля -->
        <div class="password-form" v-if="showPasswordForm">
          <div class="form-section">
            <h3>Смена пароля</h3>
            <div class="form-group">
              <label for="newPassword">Новый пароль</label>
              <input 
                type="password" 
                id="newPassword" 
                v-model="passwordForm.newPassword"
                placeholder="Введите новый пароль"
              >
            </div>
            <div class="form-group">
              <label for="confirmPassword">Подтвердите пароль</label>
              <input 
                type="password" 
                id="confirmPassword" 
                v-model="passwordForm.confirmPassword"
                placeholder="Подтвердите новый пароль"
              >
            </div>
            <div class="form-actions">
              <button class="btn btn-outline" @click="cancelPasswordChange">Отмена</button>
              <button class="btn btn-primary" @click="savePassword" :disabled="!isPasswordValid">
                Сохранить пароль
              </button>
            </div>
            <div class="error-message" v-if="passwordError">
              {{ passwordError }}
            </div>
            <div class="success-message" v-if="passwordSuccess">
              Пароль успешно изменен!
            </div>
          </div>
        </div>

        <div class="profile-container" v-if="!showPasswordForm">
          <!-- Шапка профиля -->
          <div class="profile-header">
            <img :src="user.avatar" alt="Аватар" class="avatar-large">
            <div class="profile-info">
              <h2>{{ fullName }}</h2>
              <div class="profile-position">{{ user.position }}</div>
              <div class="profile-department">{{ user.department }}</div>
              <div class="profile-branch" v-if="user.branch">Филиал: {{ user.branch }}</div>
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
                  <span class="info-label">Телефон:</span>
                  <span class="info-value">{{ user.phone || 'Не указан' }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Email:</span>
                  <span class="info-value">{{ user.email || 'Не указан' }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Telegram:</span>
                  <span class="info-value">
                    <a v-if="user.tgId" :href="user.tgLink" target="_blank" class="tg-link">
                      {{ user.tgId }}
                    </a>
                    <span v-else>Не указан</span>
                  </span>
                </div>
                <div class="info-item" v-if="user.tabNumber">
                  <span class="info-label">Табельный номер:</span>
                  <span class="info-value">{{ user.tabNumber }}</span>
                </div>
                <div class="info-item" v-if="user.address">
                  <span class="info-label">Адрес:</span>
                  <span class="info-value">{{ user.address }}</span>
                </div>
                <div class="info-item" v-if="user.passport">
                  <span class="info-label">Паспорт:</span>
                  <span class="info-value">{{ user.passport }}</span>
                </div>
              </div>
            </div>

            <!-- Профессиональная информация -->
            <div class="info-section">
              <div class="section-header">
                <h3>Профессиональная информация</h3>
              </div>
              <div class="section-body">
                <div class="info-item">
                  <span class="info-label">Отдел:</span>
                  <span class="info-value">{{ user.department || 'Не указан' }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">Должность:</span>
                  <span class="info-value">{{ user.position || 'Не указана' }}</span>
                </div>
                <div class="info-item" v-if="user.employment">
                  <span class="info-label">Занятость:</span>
                  <span class="info-value">{{ user.employment }}</span>
                </div>
                <div class="info-item" v-if="user.startDate">
                  <span class="info-label">Дата приема:</span>
                  <span class="info-value">{{ formatDate(user.startDate) }}</span>
                </div>
                <div class="info-item" v-if="otherPositions.length > 0">
                  <span class="info-label">Другие должности:</span>
                  <span class="info-value">
                    <div v-for="pos in otherPositions" :key="pos.name" class="other-position">
                      {{ pos.name }} ({{ pos.department }})
                    </div>
                  </span>
                </div>
              </div>
            </div>
          </div>

          <!-- Быстрые действия -->
          <div class="analytics-section">
            <div class="section-header">
              <h3>Действия</h3>
            </div>
            <div class="analytics-grid">
              <div class="quick-actions">
                <router-link to="/tasks" class="action-btn">
                  <div class="action-icon">✅</div>
                  <div class="action-text">
                    <h4>Задачи</h4>
                    <p>Текущие и назначенные</p>
                  </div>
                </router-link>
                <router-link to="/support" class="action-btn">
                  <div class="action-icon">🛠️</div>
                  <div class="action-text">
                    <h4>Поддержка</h4>
                    <p>Техническая помощь</p>
                  </div>
                </router-link>
                <a href="#" class="action-btn" @click.prevent="changePassword">
                  <div class="action-icon">🔒</div>
                  <div class="action-text">
                    <h4>Безопасность</h4>
                    <p>Сменить пароль</p>
                  </div>
                </a>
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
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user.js'
import AppHeader from '@/components/AppHeader.vue'
import AppSidebar from '@/components/AppSidebar.vue'
import AppFooter from '@/components/AppFooter.vue'

const router = useRouter()
const userStore = useUserStore()

const showPersonalModal = ref(false)
const showPasswordForm = ref(false)
const passwordError = ref('')
const passwordSuccess = ref('')

const user = reactive({
  id: null,
  name: '',
  surname: '',
  patronymic: '',
  login: '',
  phone: '',
  email: '',
  tgLink: '',
  tgId: '',
  address: '',
  passport: '',
  position: '',
  department: '',
  branch: '',
  employment: '',
  startDate: '',
  tabNumber: '',
  roles: [],
  otherPositions: [],
  avatar: '/default-avatar.png'
})

const editForm = reactive({
  phone: '',
  email: ''
})

const passwordForm = reactive({
  newPassword: '',
  confirmPassword: ''
})

const fullName = computed(() => {
  return `${user.surname} ${user.name} ${user.patronymic}`.trim()
})

const otherPositions = computed(() => {
  return user.otherPositions || []
})

const isPasswordValid = computed(() => {
  return passwordForm.newPassword && 
         passwordForm.confirmPassword && 
         passwordForm.newPassword === passwordForm.confirmPassword
})

const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleDateString('ru-RU')
}

const formatPassport = (passData) => {
  if (!passData.ser || !passData.num) return ''
  return `${passData.ser} ${passData.num}`
}

const loadUserData = async () => {
  if (!userStore.userId) {
    console.error('User ID not available')
    return
  }

  try {
    const response = await fetch(`/api/v1/user/${userStore.userId}`)
    if (!response.ok) {
      throw new Error('Failed to fetch user data')
    }

    const data = await response.json()
    const userData = data.user

    // Заполняем данные пользователя
    Object.assign(user, {
      id: userData.id,
      name: userData.name,
      surname: userData.surname,
      patronymic: userData.patronymic,
      login: userData.login,
      phone: userData.phone,
      email: userData.email,
      tgLink: userData.tg_link,
      tgId: userData.tg_id,
      address: userData.adress,
      passport: formatPassport({
        ser: userData.pasport_ser,
        num: userData.pasport_num
      }),
      roles: userData.roles || []
    })

    // Заполняем данные сотрудника
    if (userData.employee) {
      const emp = userData.employee
      Object.assign(user, {
        position: emp.position,
        department: emp.department,
        branch: emp.branch,
        employment: emp.zanyatost,
        startDate: emp.start_date,
        tabNumber: emp.tab_num,
        otherPositions: emp.other_position || []
      })
    }

    // Расчет опыта
    if (user.startDate) {
      const start = new Date(user.startDate)
      const now = new Date()
      const experience = now.getFullYear() - start.getFullYear()
      user.experience = experience > 0 ? experience : 0
    }

    user.rolesCount = user.roles.length

  } catch (error) {
    console.error('Error loading user data:', error)
  }
}

const changePassword = () => {
  showPasswordForm.value = true
  passwordError.value = ''
  passwordSuccess.value = ''
  passwordForm.newPassword = ''
  passwordForm.confirmPassword = ''
}

const cancelPasswordChange = () => {
  showPasswordForm.value = false
  passwordError.value = ''
  passwordSuccess.value = ''
}

const savePassword = async () => {
  if (!isPasswordValid.value) {
    passwordError.value = 'Пароли не совпадают или слишком короткие (мин. 6 символов)'
    return
  }
  console.log(userStore.userId)
  const formData = new FormData();
  formData.append('id', userStore.userId);
  formData.append('pass', passwordForm.newPassword);
  try {
    const response = await fetch('/api/v1/user/pass', {
      method: 'PATCH',
      body: formData
    })

    console.log(response)
    if (response.ok) {
      passwordSuccess.value = true
      passwordError.value = ''
      
      setTimeout(() => {
        showPasswordForm.value = false
        passwordSuccess.value = false
      }, 2000)
    } else {
      const errorText = await response.text()
      throw new Error(errorText || 'Ошибка при смене пароля')
    }
  } catch (error) {
    passwordError.value = error.message
    passwordSuccess.value = false
  }
}

const editPersonalInfo = () => {
  editForm.phone = user.phone
  editForm.email = user.email
  showPersonalModal.value = true
}

const closeEditPersonalModal = () => {
  showPersonalModal.value = false
}

const savePersonalInfo = () => {
  user.phone = editForm.phone
  user.email = editForm.email
  // Здесь можно добавить вызов API для сохранения изменений
  closeEditPersonalModal()
}

onMounted(() => {
  if (userStore.isAuthenticated) {
    loadUserData()
  }
})
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

.password-form {
  max-width: 500px;
  margin: 0 auto 2rem;
  background: #fff;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.form-section h3 {
  color: #2c5aa0;
  margin-bottom: 1.5rem;
}

.form-actions {
  display: flex;
  gap: 1rem;
  margin-top: 1.5rem;
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
  margin-right: 0.5rem;
}

.profile-branch {
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

.other-position {
  margin-bottom: 0.25rem;
  font-size: 0.9rem;
}

.analytics-section {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  border: 1px solid #e1e5e9;
  margin-bottom: 2rem;
}

.analytics-grid {
  padding: 1.5rem;
}

.quick-actions {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
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

.error-message {
  background: #fde8e8;
  color: #e74c3c;
  padding: 1rem;
  border-radius: 6px;
  margin-top: 1rem;
}

.success-message {
  background: #e8f5e8;
  color: #27ae60;
  padding: 1rem;
  border-radius: 6px;
  margin-top: 1rem;
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

.form-group input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 1rem;
  transition: all 0.3s ease;
}

.form-group input:focus {
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

.btn-primary:disabled {
  background: #ccc;
  cursor: not-allowed;
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
  
  .quick-actions {
    grid-template-columns: 1fr;
  }
  
  .info-item {
    flex-direction: column;
    gap: 0.25rem;
  }
  
  .info-label, .info-value {
    text-align: left;
  }

.tg-link {
  color: #2c5aa0;
  text-decoration: none;
  border-bottom: 1px dashed #2c5aa0;
}

.tg-link:hover {
  color: #1e3d6f;
  border-bottom: 1px solid #1e3d6f;
}
}
</style>