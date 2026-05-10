import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import { router } from './router'
import './styles.css'
import { MotionPlugin } from '@vueuse/motion'
import { useUiStore } from './stores/ui'
import { useContentStore } from './stores/content'
import { useAuthStore } from './stores/auth'

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.use(MotionPlugin)

const ui = useUiStore()
ui.initTheme()

const content = useContentStore()
void content.bootstrap()

const auth = useAuthStore()
auth.init()

// Легкий prefetch популярных страниц после первого рендера (ускоряет последующую навигацию)
const prefetchRoutes = () => {
  const jobs = [
    () => import('@/views/HomeView.vue'),
    () => import('@/views/SearchView.vue'),
    () => import('@/views/LoginView.vue'),
    () => import('@/views/CreateQuestionView.vue'),
    () => import('@/views/CreatePostView.vue')
  ]
  jobs.forEach(job => { try { void job() } catch {} })
}

const ric = (window as Window & { requestIdleCallback?: (cb: () => void) => number }).requestIdleCallback
if (typeof ric === 'function') ric(prefetchRoutes)
else setTimeout(prefetchRoutes, 600)

app.mount('#app')
