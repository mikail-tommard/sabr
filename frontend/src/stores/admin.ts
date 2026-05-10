import { defineStore } from 'pinia'
import type { Report, ContentStatus, UserRole } from '@/types'
import { adminService } from '@/services/adminService'
import { useContentStore } from '@/stores/content'

type AdminTab = 'dashboard'|'users'|'questions'|'posts'|'comments'|'tags'|'campuses'|'reports'|'settings'

export const useAdminStore = defineStore('admin', {
  state: () => ({
    tab: 'dashboard' as AdminTab,
    reports: [] as Report[],
    reportsLoaded: false,
    q: '',
  }),
  getters: {
    openReports: (s) => s.reports.filter(r => r.status !== 'resolved').length,
  },
  actions: {
    setTab(tab: AdminTab) { this.tab = tab },
    setQuery(q: string) { this.q = q },
    async loadReports() {
      if (this.reportsLoaded) return
      this.reports = await adminService.getReports()
      this.reportsLoaded = true
    },

    // ---- Модерация контента (локально, через content store) ----
    setQuestionStatus(id: string, status: ContentStatus) {
      const content = useContentStore()
      const q = content.qById(id)
      if (q) { (q as any).status = status; content.persist() }
    },
    setPostStatus(id: string, status: ContentStatus) {
      const content = useContentStore()
      const p = content.pById(id)
      if (p) { (p as any).status = status; content.persist() }
    },
    togglePinQuestion(id: string) {
      const content = useContentStore()
      const q = content.qById(id)
      if (q) { (q as any).pinned = !(q as any).pinned; content.persist() }
    },
    toggleFeaturedPost(id: string) {
      const content = useContentStore()
      const p = content.pById(id)
      if (p) { (p as any).featured = !(p as any).featured; content.persist() }
    },
    deleteComment(id: string) {
      const content = useContentStore()
      content.comments = content.comments.filter(c => c.id !== id)
      content.persist()
    },

    // ---- Пользователи ----
    setUserBanned(userId: string, banned: boolean) {
      const content = useContentStore()
      const u = content.users.find(x => x.id === userId)
      if (u) { u.isBanned = banned; content.persist() }
    },
    setUserRole(userId: string, role: UserRole) {
      const content = useContentStore()
      const u = content.users.find(x => x.id === userId)
      if (u) { u.role = role; content.persist() }
    },

    // ---- Теги ----
    createTag(label: string) {
      const content = useContentStore()
      const clean = label.trim().toLowerCase()
      if (!clean) return
      if (content.tags.some(t => t.label.toLowerCase() === clean)) return
      content.tags.unshift({ id: `t${content.tags.length + 1}-${Math.random().toString(16).slice(2,6)}`, label: clean, usageCount: 0 })
      content.persist()
    },
    renameTag(tagId: string, label: string) {
      const content = useContentStore()
      const t = content.tags.find(t => t.id === tagId)
      if (t) { t.label = label.trim().toLowerCase(); content.persist() }
    },
    deleteTag(tagId: string) {
      const content = useContentStore()
      content.tags = content.tags.filter(t => t.id !== tagId)
      // также убираем тег из контента
      content.questions.forEach(q => (q.tagIds = q.tagIds.filter(id => id !== tagId)))
      content.posts.forEach(p => (p.tagIds = p.tagIds.filter(id => id !== tagId)))
      content.persist()
    },

    // ---- Кампусы ----
    createCampus(payload: { name: string; city: string; description: string }) {
      const content = useContentStore()
      const id = `c${content.campuses.length + 1}-${Math.random().toString(16).slice(2,6)}`
      content.campuses.unshift({
        id,
        name: payload.name.trim(),
        city: payload.city.trim(),
        description: payload.description.trim(),
        membersCount: 0,
        postsCount: 0,
        questionsCount: 0,
      })
      content.persist()
    },
    updateCampus(id: string, patch: Partial<{ name: string; city: string; description: string }>) {
      const content = useContentStore()
      const c = content.campuses.find(c => c.id === id)
      if (!c) return
      if (patch.name != null) c.name = patch.name
      if (patch.city != null) c.city = patch.city
      if (patch.description != null) c.description = patch.description
      content.persist()
    },
    deleteCampus(id: string) {
      const content = useContentStore()
      // В MVP: запрещаем удалять, если есть пользователи
      if (content.users.some(u => u.campusId === id)) return
      content.campuses = content.campuses.filter(c => c.id !== id)
      content.persist()
    },

    // ---- Reports ----
    setReportStatus(id: string, status: Report['status']) {
      const r = this.reports.find(r => r.id === id)
      if (r) r.status = status
    },
  }
})
