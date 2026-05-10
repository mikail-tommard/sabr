import { createRouter, createWebHistory } from 'vue-router'
import MainLayout from '@/layouts/MainLayout.vue'
import AuthLayout from '@/layouts/AuthLayout.vue'
import { useAuthStore } from '@/stores/auth'
import { useContentStore } from '@/stores/content'

function setMeta(title?: string, description?: string) {
  if (title) document.title = title
  if (!description) return
  let el = document.querySelector('meta[name="description"]') as HTMLMetaElement | null
  if (!el) {
    el = document.createElement('meta')
    el.name = 'description'
    document.head.appendChild(el)
  }
  el.content = description
}

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: MainLayout, children: [
      { path:'', name:'home', component: ()=>import('@/views/HomeView.vue') },
      { path:'about', name:'about', component: ()=>import('@/views/AboutView.vue'), meta: { title: 'Sabr — о проекте', description: 'Sabr — платформа для программистов из разных кампусов: вопросы, ответы, публикации, теги и кампусы.' } },
      { path:'questions/:id', name:'question', component: ()=>import('@/views/QuestionDetailsView.vue') },
      { path:'posts/:id', name:'post', component: ()=>import('@/views/PostDetailsView.vue') },
      { path:'ask', name:'ask', component: ()=>import('@/views/CreateQuestionView.vue') },
      { path:'write', name:'write', component: ()=>import('@/views/CreatePostView.vue') },
      { path:'profile/:id', name:'profile', component: ()=>import('@/views/ProfileView.vue') },
      { path:'campuses', name:'campuses', component: ()=>import('@/views/CampusesView.vue') },
      { path:'campuses/:id', name:'campus-details', component: ()=>import('@/views/CampusDetailsView.vue') },
      { path:'search', name:'search', component: ()=>import('@/views/SearchView.vue') },
    ]},
    { path: '/admin', component: ()=>import('@/views/admin/AdminLayout.vue'), meta: { requiresAdmin: true, title: 'Sabr Admin', description: 'Админ‑панель Sabr: модерация, пользователи, контент.' }, children: [
      { path: '', redirect: { name: 'admin-dashboard' } },
      { path: 'dashboard', name: 'admin-dashboard', component: ()=>import('@/views/admin/AdminDashboardView.vue'), meta: { requiresAdmin: true, title: 'Админка — Дашборд' } },
      { path: 'users', name: 'admin-users', component: ()=>import('@/views/admin/AdminUsersView.vue'), meta: { requiresAdmin: true, title: 'Админка — Пользователи' } },
      { path: 'questions', name: 'admin-questions', component: ()=>import('@/views/admin/AdminQuestionsView.vue'), meta: { requiresAdmin: true, title: 'Админка — Вопросы' } },
      { path: 'posts', name: 'admin-posts', component: ()=>import('@/views/admin/AdminPostsView.vue'), meta: { requiresAdmin: true, title: 'Админка — Публикации' } },
      { path: 'comments', name: 'admin-comments', component: ()=>import('@/views/admin/AdminCommentsView.vue'), meta: { requiresAdmin: true, title: 'Админка — Комментарии' } },
      { path: 'tags', name: 'admin-tags', component: ()=>import('@/views/admin/AdminTagsView.vue'), meta: { requiresAdmin: true, title: 'Админка — Теги' } },
      { path: 'campuses', name: 'admin-campuses', component: ()=>import('@/views/admin/AdminCampusesView.vue'), meta: { requiresAdmin: true, title: 'Админка — Кампусы' } },
      { path: 'reports', name: 'admin-reports', component: ()=>import('@/views/admin/AdminReportsView.vue'), meta: { requiresAdmin: true, title: 'Админка — Жалобы' } },
      { path: 'settings', name: 'admin-settings', component: ()=>import('@/views/admin/AdminSettingsView.vue'), meta: { requiresAdmin: true, title: 'Админка — Настройки' } },
    ]},
    { path: '/', component: AuthLayout, children: [
      { path:'login', name:'login', component: ()=>import('@/views/LoginView.vue') },
      { path:'register', name:'register', component: ()=>import('@/views/RegisterView.vue') },
    ]},
    { path:'/:pathMatch(.*)*', component: ()=>import('@/views/NotFoundView.vue') }
  ]
})

router.beforeEach(async (to) => {
  // SEO meta
  const meta = to.meta as any
  setMeta(meta?.title || 'Sabr', meta?.description)

  if (!meta?.requiresAdmin) return true

  const auth = useAuthStore()
  const content = useContentStore()
  if (!content.bootstrapped) await content.bootstrap()

  if (!auth.currentUserId) {
    return { name: 'login', query: { returnTo: to.fullPath } }
  }
  const me = content.users.find(u => u.id === auth.currentUserId)
  if (!me || me.role !== 'Admin') {
    return { name: 'home' }
  }
  return true
})

router.afterEach((to) => {
  const meta = to.meta as any
  setMeta(meta?.title || 'Sabr', meta?.description)
})
