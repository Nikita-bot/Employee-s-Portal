import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import './assets/main.css'

const app = createApp(App)
const pinia = createPinia()
app.use(pinia)
app.use(router)

// Инициализируем хранилище
import { useUserStore } from './stores/user'
const userStore = useUserStore()
userStore.initialize()

app.mount('#app')