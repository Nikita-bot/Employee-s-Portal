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
          <li><router-link to="/news" class="active">Новости</router-link></li>
          <li><router-link to="/support">Поддержка</router-link></li>
        </ul>
      </sidebar>

      <main class="content-area">
        <div class="page-header">
          <h1>Новости и объявления</h1>
        </div>

        <div class="news-controls">
          <button class="btn btn-primary" @click="showCreateNewsModal">Создать новость</button>
          <input 
            type="text" 
            class="search-box" 
            placeholder="Поиск по новостям..." 
            v-model="searchQuery"
            @input="searchNews"
          >
          <select class="filter-select" v-model="categoryFilter" @change="filterNews">
            <option value="all">Все новости</option>
            <option value="announcement">Объявления</option>
            <option value="event">События</option>
            <option value="medical">Медицинские</option>
            <option value="technical">Технические</option>
          </select>
          <select class="filter-select" v-model="dateFilter" @change="filterNews">
            <option value="all">За все время</option>
            <option value="today">Сегодня</option>
            <option value="week">За неделю</option>
            <option value="month">За месяц</option>
          </select>
        </div>

        <div class="news-grid">
          <div 
            v-for="news in filteredNews" 
            :key="news.id"
            class="news-card" 
            :class="{ featured: news.important }"
          >
            <div class="news-card-header">
              <div class="news-meta">
                <span class="news-date">{{ news.date }}</span>
                <span class="news-category" :class="`category-${news.category}`">
                  {{ categoryText[news.category] }}
                </span>
              </div>
              <div class="news-title">{{ news.title }}</div>
              <div class="news-preview">{{ news.preview }}</div>
            </div>
            <div class="news-card-body">
              <div class="news-content">{{ truncateContent(news.content) }}</div>
              <div v-if="news.attachments.length > 0" class="news-attachments">
                <a 
                  v-for="attachment in news.attachments" 
                  :key="attachment.name"
                  href="#" 
                  class="attachment"
                >
                  📎 {{ attachment.name }}
                </a>
              </div>
            </div>
            <div class="news-card-footer">
              <span class="news-author">Автор: {{ news.author }}</span>
              <span class="read-more" @click="openNewsModal(news.id)">Читать полностью</span>
            </div>
          </div>
        </div>

        <div class="pagination">
          <button 
            v-for="page in totalPages" 
            :key="page"
            class="page-btn" 
            :class="{ active: currentPage === page }"
            @click="changePage(page)"
          >
            {{ page }}
          </button>
          <button class="page-btn" @click="nextPage">→</button>
        </div>
      </main>
    </div>

    <!-- Модальное окно создания новости -->
    <div class="modal" :class="{ show: showCreateModal }">
      <div class="modal-content">
        <div class="modal-header">
          <h2>Создание новости</h2>
          <button class="close-btn" @click="closeCreateNewsModal">&times;</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label for="newsTitle">Заголовок новости</label>
            <input 
              type="text" 
              id="newsTitle" 
              v-model="newNews.title"
              placeholder="Введите заголовок"
            >
          </div>
          <div class="form-group">
            <label for="newsCategory">Категория</label>
            <select id="newsCategory" v-model="newNews.category">
              <option value="announcement">Объявление</option>
              <option value="event">Событие</option>
              <option value="medical">Медицинская</option>
              <option value="technical">Техническая</option>
            </select>
          </div>
          <div class="form-group">
            <label>
              <input type="checkbox" v-model="newNews.important"> Важная новость
            </label>
          </div>
          <div class="form-group">
            <label for="newsContent">Содержание новости</label>
            <textarea 
              id="newsContent" 
              v-model="newNews.content"
              placeholder="Подробное содержание новости..."
            ></textarea>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-outline" @click="closeCreateNewsModal">Отмена</button>
          <button class="btn btn-primary" @click="createNewNews">Опубликовать</button>
        </div>
      </div>
    </div>

    <!-- Модальное окно просмотра новости -->
    <div class="modal" :class="{ show: showViewModal }">
      <div class="modal-content">
        <div class="modal-header">
          <h2>{{ selectedNews?.title }}</h2>
          <button class="close-btn" @click="closeNewsModal">&times;</button>
        </div>
        <div class="modal-body">
          <div class="full-news-meta">
            <span>{{ selectedNews?.date }}</span>
            <span>{{ categoryText[selectedNews?.category] }}</span>
            <span>Автор: {{ selectedNews?.author }}</span>
          </div>
          <div class="full-news-content" v-html="selectedNews?.content"></div>
          <div v-if="selectedNews?.attachments.length > 0" class="news-attachments">
            <h3>Прикрепленные файлы:</h3>
            <a 
              v-for="attachment in selectedNews?.attachments" 
              :key="attachment.name"
              href="#" 
              class="attachment"
            >
              📎 {{ attachment.name }}
            </a>
          </div>
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
const showCreateModal = ref(false)
const showViewModal = ref(false)
const searchQuery = ref('')
const categoryFilter = ref('all')
const dateFilter = ref('all')
const currentPage = ref(1)
const itemsPerPage = 5

// Данные
const newsData = ref([
  {
    id: 1,
    title: "Обновление регламента электронного документооборота",
    preview: "С 15 октября вводится новый порядок подписания медицинских документов...",
    content: `<p>С 15 октября 2023 года в нашей больнице вводится новый порядок подписания медицинских документов с использованием усиленной электронной подписи.</p>
              <h3>Основные изменения:</h3>
              <p>Все медицинские документы должны быть подписаны в течение 24 часов с момента создания. Внедряется двухэтапная система проверки для особо важных случаев.</p>
              <h3>Обучение:</h3>
              <p>Обязательные обучающие семинары для всех сотрудников пройдут с 10 по 12 октября в конференц-зале главного корпуса.</p>`,
    date: "05.10.2023",
    category: "announcement",
    author: "Администрация",
    important: true,
    attachments: [
      { name: "Новый регламент.pdf", type: "pdf" },
      { name: "График обучения.xlsx", type: "excel" }
    ]
  },
  {
    id: 2,
    title: "Плановые технические работы 15 октября",
    preview: "15 октября с 22:00 до 02:00 будут проводиться технические работы...",
    content: `<p>Уважаемые сотрудники! Сообщаем вам о проведении плановых технических работ.</p>
              <p><strong>Дата и время:</strong> 15 октября 2023 года с 22:00 до 02:00</p>
              <p><strong>Что будет обновлено:</strong></p>
              <ul>
                <li>Система электронного документооборота</li>
                <li>База данных пациентов</li>
                <li>Резервное копирование</li>
              </ul>
              <p>В указанное время корпоративный портал и смежные системы будут недоступны. Пожалуйста, спланируйте свою работу accordingly.</p>`,
    date: "03.10.2023",
    category: "technical",
    author: "IT-отдел",
    important: true,
    attachments: []
  },
  {
    id: 3,
    title: "Семинар по новым медицинским стандартам в кардиологии",
    preview: "Приглашаем всех врачей-кардиологов на семинар 12 октября в 14:00...",
    content: `<p>Приглашаем всех врачей-кардиологов и заинтересованных специалистов на научно-практический семинар.</p>
              <p><strong>Тема:</strong> "Современные подходы к диагностике и лечению сердечно-сосудистых заболеваний"</p>
              <p><strong>Дата и время:</strong> 12 октября 2023 года, 14:00</p>
              <p><strong>Место:</strong> Конференц-зал главного корпуса</p>
              <p><strong>Спикеры:</strong></p>
              <ul>
                <li>Проф. Иванов А.С. - "Инновации в кардиохирургии"</li>
                <li>Доц. Петрова М.И. - "Новые протоколы медикаментозной терапии"</li>
                <li>К.м.н. Сидоров В.П. - "Реабилитация кардиологических пациентов"</li>
              </ul>`,
    date: "01.10.2023",
    category: "medical",
    author: "Научный отдел",
    important: false,
    attachments: [
      { name: "Программа семинара.pdf", type: "pdf" }
    ]
  }
])

const newNews = reactive({
  title: '',
  category: 'announcement',
  content: '',
  important: false
})

const selectedNewsId = ref(null)

// Тексты для отображения
const categoryText = {
  'announcement': 'Объявление',
  'event': 'Событие', 
  'medical': 'Медицинское',
  'technical': 'Техническое'
}

// Вычисляемые свойства
const filteredNews = computed(() => {
  let filtered = newsData.value
  
  // Фильтрация по категории
  if (categoryFilter.value !== 'all') {
    filtered = filtered.filter(news => news.category === categoryFilter.value)
  }
  
  // Поиск
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(news => 
      news.title.toLowerCase().includes(query) ||
      news.content.toLowerCase().includes(query) ||
      news.preview.toLowerCase().includes(query)
    )
  }
  
  // Пагинация
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return filtered.slice(start, end)
})

const totalPages = computed(() => {
  return Math.ceil(newsData.value.length / itemsPerPage)
})

const selectedNews = computed(() => {
  return newsData.value.find(news => news.id === selectedNewsId.value)
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

const showCreateNewsModal = () => {
  showCreateModal.value = true
}

const closeCreateNewsModal = () => {
  showCreateModal.value = false
  Object.assign(newNews, {
    title: '',
    category: 'announcement',
    content: '',
    important: false
  })
}

const createNewNews = () => {
  if (newNews.title && newNews.content) {
    const news = {
      id: newsData.value.length + 1,
      title: newNews.title,
      preview: newNews.content.substring(0, 100) + '...',
      content: newNews.content,
      date: new Date().toLocaleDateString('ru-RU'),
      category: newNews.category,
      author: userStore.userData?.name,
      important: newNews.important,
      attachments: []
    }
    
    newsData.value.unshift(news)
    closeCreateNewsModal()
  }
}

const openNewsModal = (newsId) => {
  selectedNewsId.value = newsId
  showViewModal.value = true
}

const closeNewsModal = () => {
  showViewModal.value = false
  selectedNewsId.value = null
}

const searchNews = () => {
  currentPage.value = 1 // Сброс пагинации при поиске
}

const filterNews = () => {
  currentPage.value = 1 // Сброс пагинации при фильтрации
}

const changePage = (page) => {
  currentPage.value = page
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

const truncateContent = (content) => {
  // Удаляем HTML теги для превью
  const text = content.replace(/<[^>]*>/g, '')
  return text.length > 150 ? text.substring(0, 150) + '...' : text
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
/* Стили из вашего news.html */
.news-controls {
  display: flex;
  gap: 1rem;
  margin-bottom: 2rem;
  align-items: center;
}

.search-box {
  flex: 1;
  max-width: 300px;
  padding: 0.5rem 1rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
}

.news-grid {
  display: grid;
  gap: 1.5rem;
}

.news-card {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  border: 1px solid #e1e5e9;
  transition: all 0.3s ease;
  overflow: hidden;
}

.news-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transform: translateY(-2px);
}

.news-card.featured {
  border-left: 4px solid #e74c3c;
}

.news-card-header {
  padding: 1.5rem;
  border-bottom: 1px solid #f0f0f0;
}

.news-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.5rem;
}

.news-date {
  color: #666;
  font-size: 0.9rem;
  font-weight: 500;
}

.news-category {
  padding: 0.25rem 0.75rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 500;
}

.category-announcement {
  background: #e3f2fd;
  color: #2c5aa0;
}

.category-event {
  background: #fff3e0;
  color: #f39c12;
}

.category-medical {
  background: #e8f5e8;
  color: #27ae60;
}

.category-technical {
  background: #f3e5f5;
  color: #9c27b0;
}

.news-title {
  font-weight: 500;
  color: #2c5aa0;
  font-size: 1.2rem;
  margin-bottom: 0.5rem;
  line-height: 1.4;
}

.news-preview {
  color: #666;
  line-height: 1.5;
}

.news-card-body {
  padding: 1.5rem;
}

.news-content {
  color: #333;
  line-height: 1.6;
  margin-bottom: 1rem;
}

.news-attachments {
  margin-top: 1rem;
  padding-top: 1rem;
  border-top: 1px solid #f0f0f0;
}

.attachment {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background: #f8f9fa;
  border-radius: 6px;
  margin-right: 0.5rem;
  margin-bottom: 0.5rem;
  text-decoration: none;
  color: #333;
  font-size: 0.9rem;
}

.attachment:hover {
  background: #e9ecef;
}

.news-card-footer {
  padding: 1rem 1.5rem;
  background: #f8f9fa;
  border-top: 1px solid #f0f0f0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.news-author {
  color: #666;
  font-size: 0.9rem;
}

.read-more {
  color: #2c5aa0;
  text-decoration: none;
  font-weight: 500;
  cursor: pointer;
}

.read-more:hover {
  text-decoration: underline;
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  margin-top: 2rem;
  padding-top: 2rem;
  border-top: 1px solid #e1e5e9;
}

.page-btn {
  padding: 0.5rem 1rem;
  border: 1px solid #ddd;
  background: white;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.page-btn.active {
  background: #2c5aa0;
  color: white;
  border-color: #2c5aa0;
}

.page-btn:hover:not(.active) {
  background: #f8f9fa;
}

.full-news-meta {
  display: flex;
  gap: 2rem;
  margin-bottom: 1.5rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid #e1e5e9;
  color: #666;
  font-size: 0.9rem;
}

.full-news-content {
  line-height: 1.7;
  color: #333;
}

.full-news-content h3 {
  color: #2c5aa0;
  margin: 1.5rem 0 0.5rem 0;
}

.full-news-content p {
  margin-bottom: 1rem;
}

@media (max-width: 768px) {
  .news-controls {
    flex-direction: column;
    align-items: stretch;
  }
  .search-box {
    max-width: none;
  }
  .news-meta {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.5rem;
  }
  .full-news-meta {
    flex-direction: column;
    gap: 0.5rem;
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