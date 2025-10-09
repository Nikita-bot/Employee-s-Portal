import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useUserStore = defineStore('user', () => {
  const userData = ref(JSON.parse(localStorage.getItem('userData')) || null)
  
  const isAuthenticated = computed(() => !!userData.value)
  
  function setUserData(data) {
    userData.value = data
    localStorage.setItem('userData', JSON.stringify(data))
  }
  
  function clearUserData() {
    userData.value = null
    localStorage.removeItem('userData')
  }
  
  return {
    userData,
    isAuthenticated,
    setUserData,
    clearUserData
  }
})