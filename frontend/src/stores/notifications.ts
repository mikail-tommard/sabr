import { defineStore } from 'pinia'
import { notifications } from '@/mocks/data'
export const useNotificationsStore = defineStore('notifs', {
  state: ()=>({ items: structuredClone(notifications) }),
  getters: { unread: s => s.items.filter(i=>!i.isRead).length },
  actions: { toggle(id:string){ const n=this.items.find(x=>x.id===id); if(n) n.isRead=!n.isRead }, markAll(){ this.items.forEach(x=>x.isRead=true) } }
})
