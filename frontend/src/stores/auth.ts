import { defineStore } from 'pinia'
const KEY='sabr-user-id'
export const useAuthStore = defineStore('auth', {
  state: ()=>({ currentUserId: localStorage.getItem(KEY) || '' }),
  getters: { isLoggedIn: s=> !!s.currentUserId },
  actions: {
    init(){ this.currentUserId = localStorage.getItem(KEY) || '' },
    // совместимость с ранними версиями UI
    loginAs(id:string){ this.currentUserId=id; localStorage.setItem(KEY,id) },
    login(id:string){ this.currentUserId=id; localStorage.setItem(KEY,id) },
    logout(){ this.currentUserId=''; localStorage.removeItem(KEY) }
  }
})
