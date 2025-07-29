import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    userData: JSON.parse(localStorage.getItem('userData')) || null
  }),
  actions: {
    setUserData(data) {
      this.userData = data;
      localStorage.setItem('userData', JSON.stringify(data));
    },
    clearUserData() {
      this.userData = null;
      localStorage.removeItem('userData');
    }
  }
})