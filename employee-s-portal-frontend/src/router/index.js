import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '@/views/LoginPage.vue'
import TaskPage from '@/views/TaskPage.vue'
import EdoPage from '@/views/EdoPage.vue'
import EcpPage from '@/views/EcpPage.vue'
import CalendarPage from '@/views/CalendarPage.vue'
import KnowledgePage from '@/views/KnowledgePage.vue'
import AnalyticsPage from '@/views/AnalyticsPage.vue'
import UserPage from '@/views/UserPage.vue'

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
    path: '/tasks',
    name: 'Tasks',
    component: TaskPage
  },
  {
    path: '/edo',
    name: 'EDO',
    component: EdoPage
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
    path: '/knowledge',
    name: 'Knowledge',
    component: KnowledgePage
  },
  {
    path: '/user', // Ловит любой несоответствующий маршрут
    component: UserPage 
  },
  {
    path: '/:pathMatch(.*)*', // Ловит любой несоответствующий маршрут
    redirect: '/tasks' 
  },
  
  
]

// Создаем экземпляр роутера
const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router