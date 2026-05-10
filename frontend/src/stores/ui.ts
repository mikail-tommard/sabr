import { defineStore } from 'pinia'
export const useUiStore = defineStore('ui', {
  state: ()=>({ theme: (localStorage.getItem('sabr-theme') as 'light'|'dark') || 'light', toasts: [] as {id:string;type:'success'|'error'|'info';text:string}[] }),
  actions: {
    initTheme(){ document.documentElement.classList.toggle('dark', this.theme==='dark') },
    toggleTheme(){ this.theme=this.theme==='dark'?'light':'dark'; localStorage.setItem('sabr-theme',this.theme); this.initTheme() },
    toast(type:'success'|'error'|'info', text:string){ const id=Math.random().toString(36).slice(2); this.toasts.push({id,type,text}); setTimeout(()=> this.toasts=this.toasts.filter(t=>t.id!==id), 2200) }
  }
})
