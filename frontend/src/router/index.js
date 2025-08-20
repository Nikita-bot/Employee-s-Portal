import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '@/views/LoginPage.vue'
import TaskPage from '@/views/TaskPage.vue'
import EdoPage from '@/views/EdoPage.vue'
import EcpPage from '@/views/EcpPage.vue'
import CalendarPage from '@/views/CalendarPage.vue'
import AnalyticsPage from '@/views/AnalyticsPage.vue'
import UserPage from '@/views/UserPage.vue'
import NewsPage from '@/views/NewsPage.vue'
import ChatPage from '@/views/ChatPage.vue'
import SupportPage from '@/views/SupportPage.vue'

// Определяем маршруты
const routes = [
  {
    path: '/',
    redirect: '/login' 
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginPage
  },
  {
    path: '/news',
    name: 'News',
    component: NewsPage
  },
  {
    path: '/tasks',
    name: 'Tasks',
    component: TaskPage
  },
  {
    path: '/ecp',
    name: 'ECP',
    component: EcpPage
  },
  {
    path: '/analytics',
    name: 'Analytics',
    component: AnalyticsPage
  },
  {
    path: '/calendar',
    name: 'Calendar',
    component: CalendarPage
  },
  {
    path: '/user/:id', 
    component: UserPage,
    beforeEnter: (to, from, next) => {
      const userId = to.params.id
      next()
    }
  },
  {
    path: '/support',
    component: SupportPage,
  },
  {
    path: '/chat', 
    component: ChatPage
  },
  
]

// Создаем экземпляр роутера
const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router