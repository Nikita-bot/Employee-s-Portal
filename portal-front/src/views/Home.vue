<template>
  <div class="home-wrapper">
    <AppHeader />
    
    <div class="main-container">
      <AppSidebar />
      
      <main class="content-area">
        <div class="content">
          <section class="welcome-banner">
            <h1>Добро пожаловать в корпоративный портал</h1>
            <p>Сегодня: {{ currentDate }}</p>
            <div class="user-role">{{ userStore.user.user.name + ' ' + userStore.user.user.patronymic + ' | ' + (userStore.user.user.employee?.department || '') }}</div>
          </section>
          
          <!-- Быстрый доступ -->
          <section class="quick-access-grid">
            <div class="quick-card" @click="$router.push('/tasks')">
              <h3>Мои задачи</h3>
              <p>Текущие и назначенные задания</p>
            </div>
            <div class="quick-card">
              <h3>Электронный документооборот</h3>
              <p>Работа с документами и подписями</p>
            </div>
            <div class="quick-card" @click="$router.push('/profile')">
              <h3>Мой профиль</h3>
              <p>ЭЦП, аналитика, календарь</p>
            </div>
            <div class="quick-card" @click="openDocsPortal">
              <h3>База знаний</h3>
              <p>Протоколы и стандарты лечения</p>
            </div>
            <div class="quick-card" @click="openExternalPortal">
              <h3>Внешние порталы</h3>
              <p>Доступ к смежным системам</p>
            </div>
            <div class="quick-card" @click="$router.push('/support')">
              <h3>Техподдержка</h3>
              <p>Помощь с техническими вопросами</p>
            </div>
          </section>
          
          <!-- Новостной блок на всю ширину -->
          <section class="news-section">
            <div class="section-header">
              <h2>Последние новости</h2>
              <router-link to="/news" class="view-all">Все новости</router-link>
            </div>
            <div class="news-grid">
              <div v-for="newsItem in news" :key="newsItem.id" class="news-item">
                <div class="news-meta">
                  <span class="news-date">{{ formatDate(newsItem.date) }}</span>
                  <span class="news-author">{{ newsItem.author.name }}</span>
                </div>
                <div class="news-title">{{ newsItem.title }}</div>
                <div class="news-preview">{{ getPreview(newsItem.content) }}</div>
              </div>
              <div v-if="news.length === 0 && !isLoading" class="no-news">
                Новостей пока нет
              </div>
              <div v-if="isLoading" class="loading-news">
                Загрузка новостей...
              </div>
            </div>
          </section>
        </div>
      </main>
    </div>
    
    <AppFooter />
  </div>
</template>

<script setup>
import { computed, ref, onMounted } from 'vue'
import { useUserStore } from '@/stores/user'
import AppHeader from '@/components/AppHeader.vue'
import AppSidebar from '@/components/AppSidebar.vue'
import AppFooter from '@/components/AppFooter.vue'

const userStore = useUserStore()
const news = ref([])
const isLoading = ref(false)

const currentDate = computed(() => {
  return new Date().toLocaleDateString('ru-RU', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
})

const openExternalPortal = () => {
  window.open('https://portals.kkbsmp.ru/', '_blank')
}

const openDocsPortal = () => {
  window.open('http://documents.kkbsmp.ru/view/list', '_blank')
}

const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('ru-RU', {
    day: 'numeric',
    month: 'long',
    year: 'numeric'
  })
}

const getPreview = (content) => {
  const text = content.replace(/<[^>]*>/g, '')
  return text.length > 150 ? text.substring(0, 150) + '...' : text
}

const fetchNews = async () => {
  isLoading.value = true
  try {
    const response = await fetch('/api/v1/news')
    if (!response.ok) {
      throw new Error('Failed to fetch news')
    }
    
    const data = await response.json()
    news.value = data.news.slice(0, 3)
    
  } catch (error) {
    console.error('Error fetching news:', error)
    news.value = []
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  fetchNews()
})
</script>

<style scoped>
body{
    margin: 0;
    padding: 0;
}

.home-wrapper {
  display: grid;
  grid-template-rows: auto 1fr auto;
  min-height: 100vh;
  background-color: #f8f9fa;
  color: #333;
  line-height: 1.6;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
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

.content {
  max-width: 1200px;
  margin: 0 auto;
}

.welcome-banner {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 1.5rem;
  border-radius: 8px;
  margin-bottom: 2rem;
}

.welcome-banner h1 {
  font-weight: 300;
  margin-bottom: 0.5rem;
  font-size: 1.5rem;
}

.user-role {
  background: rgba(255, 255, 255, 0.2);
  padding: 0.25rem 0.75rem;
  border-radius: 20px;
  font-size: 0.9rem;
  display: inline-block;
  margin-top: 0.5rem;
}

/* Сетка быстрого доступа */
.quick-access-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
}

.quick-card {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  padding: 1.25rem;
  text-align: center;
  transition: all 0.3s ease;
  cursor: pointer;
  border: 1px solid #e1e5e9;
  min-height: 120px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.quick-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.quick-card h3 {
  color: #2c5aa0;
  margin-bottom: 0.5rem;
  font-weight: 500;
  font-size: 1rem;
}

.quick-card p {
  color: #666;
  font-size: 0.85rem;
  line-height: 1.4;
}

/* Новостной блок */
.news-section {
  background: #f8f9fa;
  border-radius: 8px;
  padding: 1.5rem;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
  padding-bottom: 0.75rem;
  border-bottom: 2px solid #e1e5e9;
}

.section-header h2 {
  color: #2c5aa0;
  font-weight: 500;
  font-size: 1.25rem;
}

.news-grid {
  display: grid;
  gap: 1rem;
}

.news-item {
  background: white;
  padding: 1.25rem;
  border-radius: 6px;
  border-left: 4px solid #2c5aa0;
  transition: all 0.3s ease;
}

.news-item:hover {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.news-meta {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.5rem;
  font-size: 0.9rem;
  color: #666;
}

.news-date {
  font-weight: 500;
}

.news-author {
  color: #2c5aa0;
}

.news-title {
  font-weight: 500;
  margin-bottom: 0.5rem;
  color: #333;
  font-size: 1.1rem;
}

.news-preview {
  color: #666;
  font-size: 0.9rem;
  line-height: 1.5;
}

.view-all {
  color: #2c5aa0;
  text-decoration: none;
  font-weight: 500;
}

.view-all:hover {
  text-decoration: underline;
}

.no-news, .loading-news {
  text-align: center;
  color: #666;
  padding: 2rem;
  font-style: italic;
}

/* Адаптивность */
@media (max-width: 768px) {
  .main-container {
    grid-template-columns: 1fr;
  }
  
  .content-area {
    padding: 1rem;
  }
  
  .quick-access-grid {
    grid-template-columns: 1fr 1fr;
  }
  
  .welcome-banner h1 {
    font-size: 1.25rem;
  }
  
  .news-meta {
    flex-direction: column;
    gap: 0.25rem;
  }
}

@media (max-width: 480px) {
  .quick-access-grid {
    grid-template-columns: 1fr;
  }
  
  .content-area {
    padding: 0.75rem;
  }
  
  .welcome-banner {
    padding: 1rem;
  }
  
  .welcome-banner h1 {
    font-size: 1.1rem;
  }
}
</style>