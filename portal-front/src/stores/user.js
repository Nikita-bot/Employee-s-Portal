import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useUserStore = defineStore('user', () => {
  const user = ref(null)
  const userId = ref(null)
  
  const isAuthenticated = computed(() => !!userId.value)
  
  function setAuthData(authData) {
    userId.value = authData.userId

    localStorage.setItem('userId', userId.value)

  }
  
  async function fetchUserData() {
    if (!userId.value) {
      throw new Error('User ID not set')
    }
    
    try {
      const response = await fetch(`/api/v1/user/${userId.value}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      })
      
      if (!response.ok) {
        if (response.status === 401) {
          handleUnauthorized()
          throw new Error('Unauthorized')
        }
        throw new Error('Failed to fetch user data')
      }
      
      const userData = await response.json()
      user.value = userData
      console.log(user.value)
      localStorage.setItem('user', JSON.stringify(user.value))
      
    } catch (error) {
      console.error('Failed to fetch user data:', error)
      throw error
    }
  }
  
  async function login(loginData) {
    try {
      const response = await fetch('/api/v1/auth', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(loginData)
      })
      
      if (!response.ok) {
        throw new Error('Login failed')
      }
      
      const authData = await response.json()
      
      // Сохраняем auth данные
      setAuthData({
        userId: authData.user_id,
      })
      
      // Загружаем полные данные пользователя
      await fetchUserData()
      
    } catch (error) {
      console.error('Login error:', error)
      throw error
    }
  }
  
  function logout() {
    user.value = null
    userId.value = null
    
    localStorage.removeItem('user')
    localStorage.removeItem('userId')
  }
  
  function handleUnauthorized() {
    // Очищаем все данные пользователя из localStorage при 401 ошибке
    logout()
  }
  
  function initialize() {
    const savedUser = localStorage.getItem('user')
    const savedUserId = localStorage.getItem('userId')

    
    if (savedUserId) {
      userId.value = savedUserId
      
      if (savedUser) {
        user.value = JSON.parse(savedUser)
      } else {
        // Если есть userId но нет user данных, загружаем их
        fetchUserData().catch(console.error)
      }
    }
  }
  
  // Перехватчик для обработки 401 ошибок во всех запросах
  async function apiRequest(url, options = {}) {
    const response = await fetch(url, options)
    
    if (response.status === 401) {
      handleUnauthorized()
      throw new Error('Unauthorized')
    }
    
    return response
  }
  
  return {
    user,
    userId,
    isAuthenticated,
    setAuthData,
    fetchUserData,
    login,
    logout,
    initialize,
    apiRequest,
    handleUnauthorized
  }
})