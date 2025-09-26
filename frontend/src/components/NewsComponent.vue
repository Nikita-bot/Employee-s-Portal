<template>
  <div class="news-page">
    <div class="news-header">
      <button v-if="hasAdminRole" class="add-news-btn" @click="createNews">
        + Добавить новость
      </button>
      <!-- <button class="add-news-btn" @click="createNews">
        + Добавить новость
      </button> -->
    </div>
    <div class="news-list-container">
      <div class="news-list">
        <div v-for="news in newsList" :key="news.id" class="news-item">
          <div class="news-date">{{news.date }}</div>
          <h3 class="news-title">{{ news.title }}</h3>
          <div class="news-content">{{ news.content }}</div>
          <div class="news-author">{{ formatAuthor(news.author) }}</div>
        </div>
      </div>
    </div>
    <CreateNewsModal 
      :isOpen="isCreateModalOpen" 
      @close="isCreateModalOpen = false"
      @created="fetchNews"
    />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useUserStore } from '@/stores/user';
import CreateNewsModal from '@/components/CreateNewsModal.vue';

const isCreateModalOpen = ref(false);

const userStore = useUserStore();
if (!userStore.userData.id) {
  router.push('/login')
}
const newsList = ref([]);


const hasAdminRole = computed(() => {
  return userStore.userData?.roles?.some(role => role.name === 'admin');
});


const fetchNews = async () => {
  try {
    const response = await fetch('/api/v1/news');
    if (!response.ok) throw new Error('Ошибка загрузки новостей');
    const data = await response.json();
    newsList.value = data.news;
  } catch (error) {
    console.error('Ошибка при загрузке новостей:', error);
    alert('Не удалось загрузить новости');
  }
};


const createNews = async () => {
  isCreateModalOpen.value = true;
};



const formatAuthor = (author) => {
  if (!author) return 'Неизвестный автор';
  return `${author.surname} ${author.name?.charAt(0) || ''}.${author.patronymic?.charAt(0) || ''}.`;
};


onMounted(fetchNews);
</script>

<style scoped>
.news-page {
  max-width: 800px;
  margin: 0px auto;
  padding: 0 20px;
}

.news-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
}

.news-header h1 {
  color: #5662DE;
  font-size: 28px;
}

.add-news-btn {
  background-color: #5662DE;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  margin-top: 20px;
}

.news-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
  padding-right: 10px;
}

.news-list-container {
  max-height: 80vh; /* или любая другая подходящая высота */
  overflow-y: auto;
  
  /* Стилизация скроллбара (невидимый) */
  scrollbar-width: none; /* Для Firefox */
  -ms-overflow-style: none; /* Для IE и Edge */
}



.news-item {
  background-color: white;
  border-radius: 10px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  position: relative;
  min-width: 720px;
}

.news-date {
  position: absolute;
  top: 15px;
  right: 15px;
  color: #888;
  font-size: 13px;
}

.news-title {
  color: #333;
  font-size: 18px;
  margin-bottom: 10px;
  padding-right: 60px;
}

.news-content {
  color: #555;
  line-height: 1.5;
  margin-bottom: 15px;
}

.news-author {
  color: #5662DE;
  font-size: 14px;
  font-style: italic;
}
</style>