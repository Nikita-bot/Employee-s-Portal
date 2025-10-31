<template>
  <div class="support-wrapper">
    <AppHeader />
    
    <div class="main-container">
      <AppSidebar />
      
      <main class="content-area">
        <div class="page-header">
          <h1>Техническая поддержка</h1>
        </div>

        <div class="support-container">
          <div class="support-info">
            <h3>Создание заявки в техническую поддержку</h3>
            <p>Опишите проблему как можно подробнее. В зависимости от типа проблемы могут появиться дополнительные поля для уточнения.</p>
          </div>

          <form @submit.prevent="submitForm">
            <div class="form-group">
              <label for="problemType">Тип проблемы *</label>
              <select id="problemType" v-model="form.problemType" required @change="checkAdditionalFields">
                <option value="">Выберите тип проблемы</option>
                <option v-for="task in taskList" :key="task.id" :value="task.id">
                  {{ task.name }}
                </option>
              </select>
            </div>

            <div class="form-group">
              <label for="problemDescription">Подробное описание *</label>
              <textarea id="problemDescription" v-model="form.problemDescription" 
                        placeholder="Опишите проблему подробно, укажите что именно не работает, когда началась проблема..." 
                        required></textarea>
            </div>

            <!-- Дополнительные поля для принтера -->
            <div class="additional-fields" v-show="showPrinterFields">
              <h4>Информация о принтере</h4>
              <div class="form-row">
                <div class="form-group">
                  <label for="printerModel">Какой принтер</label>
                  <select id="printerModel" v-model="form.printerModel">
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
                  <select id="printerLocation" v-model="form.printerLocation">
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
            <div class="additional-fields" v-show="showComputerFields">
              <h4>Информация о компьютере</h4>
              <div class="form-group">
                <label for="computerLocation">Где стоит компьютер</label>
                <select id="computerLocation" v-model="form.computerLocation">
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
            <div class="additional-fields" v-show="showSoftwareFields">
              <h4>Информация о программном обеспечении</h4>
              <div class="form-group">
                <label for="computerName">Имя компьютера *</label>
                <input type="text" id="computerName" v-model="form.computerName" 
                       placeholder="Например: PC-ADMIN-01 или USER-305">
              </div>
            </div>

            <div class="urgent-checkbox">
              <input type="checkbox" id="urgentRequest" v-model="form.urgent">
              <label for="urgentRequest">Срочная заявка (критическая проблема)</label>
            </div>

            <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
              {{ isSubmitting ? 'Отправка...' : 'Создать заявку' }}
            </button>

            <div class="success-message" v-show="showSuccess">
              Заявка успешно создана! Номер вашей заявки: <strong>{{ ticketNumber }}</strong>
            </div>

            <div class="error-message" v-show="showError">
              Ошибка при создании заявки: {{ errorMessage }}
            </div>
          </form>
        </div>
      </main>
    </div>

    <AppFooter />
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useUserStore } from '@/stores/user'
import AppHeader from '@/components/AppHeader.vue'
import AppSidebar from '@/components/AppSidebar.vue'
import AppFooter from '@/components/AppFooter.vue'

const showPrinterFields = ref(false)
const showComputerFields = ref(false)
const showSoftwareFields = ref(false)
const showSuccess = ref(false)
const showError = ref(false)
const isSubmitting = ref(false)
const ticketNumber = ref('')
const taskList = ref([])
const errorMessage = ref('')

const userStore = useUserStore()
const userData = reactive({
  id: userStore.user.user.id, 
  branchId: userStore.user.user.employee.branch_id 
})

const form = reactive({
  problemType: '',
  problemDescription: '',
  printerModel: '',
  printerLocation: '',
  computerLocation: '',
  computerName: '',
  urgent: false
})

const checkAdditionalFields = () => {
  const selectedTask = taskList.value.find(task => task.id === parseInt(form.problemType))
  if (!selectedTask) return
  console.log(selectedTask)
  const taskName = selectedTask.name.toLowerCase()
  showPrinterFields.value = taskName.includes('принтер') || taskName.includes('картридж')
  showComputerFields.value = taskName.includes('компьютер')
  showSoftwareFields.value = taskName.includes('программ') || taskName.includes('ариадн')
}

const buildDescription = () => {
  let description = form.problemDescription

  if (showPrinterFields.value) {
    description += `\n\nДополнительная информация о принтере:`
    if (form.printerModel) description += `\nМодель принтера: ${form.printerModel}`
    if (form.printerLocation) description += `\nМестоположение: ${form.printerLocation}`
  }

  if (showComputerFields.value && form.computerLocation) {
    description += `\n\nМестоположение компьютера: ${form.computerLocation}`
  }

  if (showSoftwareFields.value && form.computerName) {
    description += `\n\nИмя компьютера: ${form.computerName}`
  }

  if (form.urgent) {
    description += `\n\n!!! СРОЧНАЯ ЗАЯВКА !!!`
  }

  return description
}

const submitForm = () => {
  if (!form.problemType || !form.problemDescription) {
    showError.value = true
    errorMessage.value = 'Пожалуйста, заполните все обязательные поля'
    return
  }

  if (showSoftwareFields.value && !form.computerName) {
    showError.value = true
    errorMessage.value = 'Для программных проблем обязательно укажите имя компьютера'
    return
  }

  createSupportTicket()
}

const createSupportTicket = async () => {
  isSubmitting.value = true
  showError.value = false

  const taskData = {
    task_id: parseInt(form.problemType),
    executor: 0, // Автоматически выберется
    initiator: userData.id,
    description: buildDescription(),
    status: 0, // Создана
    priority: form.urgent ? 3 : 2, // Высокий при срочной, иначе средний
    branch_id: userData.branchId,
    create_date: new Date().toISOString().split('T')[0]
  }

  try {
    const response = await fetch('/api/v1/tasks', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(taskData)
    })

    if (response.ok) {
      ticketNumber.value = 'T' + Date.now().toString().slice(-6)
      showSuccess.value = true
      console.log('Заявка создана:', taskData)
      
      setTimeout(() => {
        resetForm()
      }, 5000)
    } else {
      throw new Error('Ошибка сервера: ' + response.status)
    }
  } catch (error) {
    showError.value = true
    errorMessage.value = error.message
    console.error('Ошибка создания заявки:', error)
  } finally {
    isSubmitting.value = false
  }
}

const loadTaskList = async () => {
  try {
    const response = await fetch('/api/v1/tasks/support')
    if (response.ok) {
      const data = await response.json()
      taskList.value = data.task_list || []
    } else {
      throw new Error('Ошибка загрузки списка задач')
    }
  } catch (error) {
    console.error('Ошибка загрузки задач:', error)
    showError.value = true
    errorMessage.value = 'Не удалось загрузить список задач'
  }
}

const resetForm = () => {
  Object.assign(form, {
    problemType: '',
    problemDescription: '',
    printerModel: '',
    printerLocation: '',
    computerLocation: '',
    computerName: '',
    urgent: false
  })
  
  showPrinterFields.value = false
  showComputerFields.value = false
  showSoftwareFields.value = false
  showSuccess.value = false
  showError.value = false
}

onMounted(() => {
  loadTaskList()
})
</script>

<style scoped>
.support-wrapper {
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

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  transition: all 0.3s ease;
}

.btn-primary {
  background: #2c5aa0;
  color: white;
}

.btn-primary:hover {
  background: #1e3d6f;
}

.btn-primary:disabled {
  background: #ccc;
  cursor: not-allowed;
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

.error-message {
  background: #fde8e8;
  color: #e74c3c;
  padding: 1rem;
  border-radius: 6px;
  margin-top: 1rem;
}

@media (max-width: 768px) {
  .main-container {
    grid-template-columns: 1fr;
  }
  
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .support-container {
    max-width: 100%;
  }
}
</style>