<template>
  <div class="news-wrapper">
    <AppHeader />
    
    <div class="main-container">
      <AppSidebar />
      
      <main class="content-area">
        <div class="page-header">
          <h1>Новости и объявления</h1>
          <div class="news-stats" v-if="!isLoading">
            Показано {{ startItem }}-{{ endItem }} из {{ filteredNews.length }} новостей
          </div>
        </div>

        <div class="news-controls">
          <button class="btn btn-primary" @click="showCreateNewsModal">Создать новость</button>
          <input type="text" class="search-box" v-model="searchQuery" placeholder="Поиск по новостям...">
          <select class="filter-select" v-model="dateFilter" @change="applyFilters">
            <option value="all">За все время</option>
            <option value="today">Сегодня</option>
            <option value="week">За неделю</option>
            <option value="month">За месяц</option>
          </select>
          <select class="filter-select" v-model="sortBy" @change="applyFilters">
            <option value="newest">Сначала новые</option>
            <option value="oldest">Сначала старые</option>
          </select>
        </div>

        <div class="news-grid">
          <div v-for="news in paginatedNews" :key="news.id" 
               :class="['news-card', { featured: news.important }]">
            <div class="news-card-header">
              <div class="news-meta">
                <span class="news-date">{{ formatDate(news.date) }}</span>
              </div>
              <div class="news-title">{{ news.title }}</div>
              <div class="news-preview">{{ getPreview(news.content) }}</div>
            </div>
            <div class="news-card-body">
              <div class="news-content">{{ getShortContent(news.content) }}</div>
              <div v-if="news.attachments && news.attachments.length > 0" class="news-attachments">
                <a v-for="att in news.attachments" :key="att.name" href="#" class="attachment">
                  📎 {{ att.name }}
                </a>
              </div>
            </div>
            <div class="news-card-footer">
              <span class="news-author">Автор: {{ news.author?.name || 'Неизвестный автор' }}</span>
              <span class="read-more" @click="openNewsModal(news.id)">Читать полностью</span>
            </div>
          </div>
          <div v-if="paginatedNews.length === 0 && !isLoading" class="no-news">
            Новостей пока нет
          </div>
          <div v-if="isLoading" class="loading-news">
            Загрузка новостей...
          </div>
        </div>

        <!-- Пагинация -->
        <div class="pagination" v-if="totalPages > 1 && !isLoading">
          <button class="page-btn" @click="prevPage" :disabled="currentPage === 1">
            ← Назад
          </button>
          
          <button v-for="page in visiblePages" 
                  :key="page"
                  :class="['page-btn', { active: page === currentPage }]"
                  @click="goToPage(page)">
            {{ page }}
          </button>
          
          <button class="page-btn" @click="nextPage" :disabled="currentPage === totalPages">
            Вперед →
          </button>
        </div>

        <div class="pagination-info" v-if="totalPages > 1 && !isLoading">
          Страница {{ currentPage }} из {{ totalPages }}
        </div>
      </main>
    </div>

    <!-- Модальное окно создания новости -->
    <div class="modal" :class="{ show: showCreateModal }" @click="closeCreateNewsModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h2>Создание новости</h2>
          <button class="close-btn" @click="closeCreateNewsModal">&times;</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label for="newsTitle">Заголовок новости *</label>
            <input type="text" id="newsTitle" v-model="newNews.title" placeholder="Введите заголовок" required>
          </div>
          <div class="form-group">
            <label for="newsContent">Содержание новости *</label>
            <textarea id="newsContent" v-model="newNews.content" 
                      placeholder="Подробное содержание новости..." required></textarea>
          </div>
          <div v-if="createError" class="error-message">
            {{ createError }}
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-outline" @click="closeCreateNewsModal">Отмена</button>
          <button class="btn btn-primary" @click="createNewNews" :disabled="isCreating">
            {{ isCreating ? 'Создание...' : 'Опубликовать' }}
          </button>
        </div>
      </div>
    </div>

    <!-- Модальное окно просмотра новости -->
    <div class="modal" :class="{ show: showNewsModal }" @click="closeNewsModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h2>{{ selectedNews.title }}</h2>
          <button class="close-btn" @click="closeNewsModal">&times;</button>
        </div>
        <div class="modal-body">
          <div class="full-news-meta">
            <span>{{ formatDate(selectedNews.date) }}</span>
            <span>Автор: {{ selectedNews.author?.name || 'Неизвестный автор' }}</span>
          </div>
          <div class="full-news-content" v-html="selectedNews.content"></div>
          <div v-if="selectedNews.attachments && selectedNews.attachments.length > 0" class="news-attachments">
            <h3>Прикрепленные файлы:</h3>
            <a v-for="att in selectedNews.attachments" :key="att.name" href="#" class="attachment">
              📎 {{ att.name }}
            </a>
          </div>
        </div>
      </div>
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

const userStore = useUserStore()
const showCreateModal = ref(false)
const showNewsModal = ref(false)
const searchQuery = ref('')
const dateFilter = ref('all')
const sortBy = ref('newest')
const isLoading = ref(false)
const isCreating = ref(false)
const createError = ref('')

// Пагинация
const currentPage = ref(1)
const pageSize = 10 // Новостей на странице

const selectedNews = reactive({
  id: 0,
  title: '',
  content: '',
  date: '',
  author: null,
  attachments: []
})

const newNews = reactive({
  title: '',
  content: '',
})

const newsData = ref([])

// Загрузка всех новостей с API
const fetchNews = async () => {
  isLoading.value = true
  try {
    const response = await fetch('/api/v1/news')
    if (!response.ok) {
      throw new Error('Ошибка загрузки новостей')
    }
    
    const data = await response.json()
    newsData.value = data.news || []
    currentPage.value = 1 // Сбрасываем на первую страницу
    
  } catch (error) {
    console.error('Error fetching news:', error)
    newsData.value = []
  } finally {
    isLoading.value = false
  }
}

// Вычисляемые свойства для фильтрации и пагинации
const filteredNews = computed(() => {
  let filtered = [...newsData.value]

  // Поиск
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(news => 
      news.title.toLowerCase().includes(query) ||
      news.content.toLowerCase().includes(query)
    )
  }

  // Фильтрация по дате
  if (dateFilter.value !== 'all') {
    const now = new Date()
    filtered = filtered.filter(news => {
      if (!news.date) return false
      
      const newsDate = new Date(news.date)
      
      switch (dateFilter.value) {
        case 'today':
          return newsDate.toDateString() === now.toDateString()
        case 'week':
          const weekAgo = new Date(now.getTime() - 7 * 24 * 60 * 60 * 1000)
          return newsDate >= weekAgo
        case 'month':
          const monthAgo = new Date(now.getTime() - 30 * 24 * 60 * 60 * 1000)
          return newsDate >= monthAgo
        default:
          return true
      }
    })
  }

  // Сортировка
  filtered.sort((a, b) => {
    const dateA = new Date(a.date || 0)
    const dateB = new Date(b.date || 0)
    
    if (sortBy.value === 'newest') {
      return dateB - dateA
    } else {
      return dateA - dateB
    }
  })

  return filtered
})

// Пагинация
const totalPages = computed(() => Math.ceil(filteredNews.value.length / pageSize))

const startItem = computed(() => {
  return (currentPage.value - 1) * pageSize + 1
})

const endItem = computed(() => {
  const end = currentPage.value * pageSize
  return end > filteredNews.value.length ? filteredNews.value.length : end
})

const paginatedNews = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  const end = start + pageSize
  return filteredNews.value.slice(start, end)
})

// Видимые страницы в пагинации
const visiblePages = computed(() => {
  const pages = []
  const total = totalPages.value
  const current = currentPage.value
  
  if (total <= 5) {
    for (let i = 1; i <= total; i++) {
      pages.push(i)
    }
  } else {
    let start = Math.max(1, current - 2)
    let end = Math.min(total, current + 2)
    
    if (current <= 3) {
      end = 5
    } else if (current >= total - 2) {
      start = total - 4
    }
    
    for (let i = start; i <= end; i++) {
      pages.push(i)
    }
  }
  
  return pages
})

// Методы пагинации
const goToPage = (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const prevPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const applyFilters = () => {
  currentPage.value = 1 // Сбрасываем на первую страницу при изменении фильтров
}

const getShortContent = (content) => {
  if (!content) return 'Нет содержания'
  const text = content.replace(/<[^>]*>/g, '')
  return text.substring(0, 150) + '...'
}

const getPreview = (content) => {
  if (!content) return 'Нет содержания'
  const text = content.replace(/<[^>]*>/g, '')
  return text.length > 100 ? text.substring(0, 100) + '...' : text
}

const formatDate = (dateString) => {
  if (!dateString) return 'Дата не указана'
  try {
    const date = new Date(dateString)
    return date.toLocaleDateString('ru-RU', {
      day: 'numeric',
      month: 'long',
      year: 'numeric'
    })
  } catch {
    return 'Неверная дата'
  }
}

// Модальные окна
const showCreateNewsModal = () => {
  showCreateModal.value = true
  createError.value = ''
}

const closeCreateNewsModal = () => {
  showCreateModal.value = false
  newNews.title = ''
  newNews.content = ''
  createError.value = ''
}

const createNewNews = async () => {
  if (!newNews.title || !newNews.content) {
    createError.value = 'Пожалуйста, заполните все обязательные поля'
    return
  }

  isCreating.value = true
  createError.value = ''

  try {
    const response = await fetch('/api/v1/news', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        title: newNews.title,
        content: newNews.content
      })
    })

    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.message || 'Ошибка создания новости')
    }

    // Обновляем список новостей
    await fetchNews()
    closeCreateNewsModal()
    
  } catch (error) {
    console.error('Error creating news:', error)
    createError.value = error.message || 'Ошибка при создании новости'
  } finally {
    isCreating.value = false
  }
}

const openNewsModal = (newsId) => {
  const news = newsData.value.find(n => n.id === newsId)
  if (news) {
    Object.assign(selectedNews, news)
    showNewsModal.value = true
  }
}

const closeNewsModal = () => {
  showNewsModal.value = false
}

// Загружаем новости при монтировании компонента
onMounted(() => {
  fetchNews()
})
</script>

<style scoped>
/* Стили для пагинации */
.news-stats {
  color: #666;
  font-size: 0.9rem;
  margin-top: 0.5rem;
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 0.5rem;
  margin-top: 2rem;
  padding-top: 2rem;
  border-top: 1px solid #e1e5e9;
  flex-wrap: wrap;
}

.page-btn {
  padding: 0.5rem 1rem;
  border: 1px solid #ddd;
  background: white;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 0.9rem;
  min-width: 44px;
  text-align: center;
}

.page-btn:hover:not(:disabled) {
  background: #f8f9fa;
  border-color: #2c5aa0;
}

.page-btn.active {
  background: #2c5aa0;
  color: white;
  border-color: #2c5aa0;
}

.page-btn:disabled {
  background: #f8f9fa;
  color: #999;
  cursor: not-allowed;
  border-color: #e1e5e9;
}

.pagination-info {
  text-align: center;
  color: #666;
  font-size: 0.9rem;
  margin-top: 1rem;
}

/* Остальные стили */
.error-message {
  background: #ffe6e6;
  color: #d63031;
  padding: 0.75rem;
  border-radius: 6px;
  margin-bottom: 1rem;
  display: block;
}

.no-news, .loading-news {
  text-align: center;
  color: #666;
  padding: 2rem;
  font-style: italic;
  grid-column: 1 / -1;
}

.news-wrapper {
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

.news-controls {
  display: flex;
  gap: 1rem;
  margin-bottom: 2rem;
  align-items: center;
}

.btn {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
  transition: all 0.3s ease;
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

.filter-select {
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  background: white;
  min-width: 150px;
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
  max-width: 800px;
  max-height: 90vh;
  overflow-y: auto;
}

.modal-header {
  padding: 1.5rem;
  border-bottom: 1px solid #e1e5e9;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-header h2 {
  color: #2c5aa0;
  font-weight: 500;
}

.close-btn {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: #666;
}

.modal-body {
  padding: 1.5rem;
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
}

.form-group textarea {
  resize: vertical;
  min-height: 120px;
}

.modal-footer {
  padding: 1rem 1.5rem;
  border-top: 1px solid #e1e5e9;
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
}

.btn-outline {
  background: transparent;
  border: 1px solid #2c5aa0;
  color: #2c5aa0;
}

@media (max-width: 768px) {
  .main-container {
    grid-template-columns: 1fr;
  }
  
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
  
  .pagination {
    gap: 0.25rem;
  }
  
  .page-btn {
    padding: 0.4rem 0.8rem;
    font-size: 0.8rem;
    min-width: 40px;
  }
}
</style>