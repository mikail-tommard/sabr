<template>
  <section class="container admin-shell">
    <header class="card admin-topbar">
      <div class="left">
        <div class="row" style="gap:12px">
          <div class="logo-badge" aria-hidden="true">A</div>
          <div style="min-width:0">
            <div class="title" style="font-size:16px">Админ‑панель</div>
            <div class="muted" style="margin-top:2px">Модерация, пользователи, контент и настройки (локально на мок‑данных)</div>
          </div>
        </div>
      </div>

      <div class="right">
        <input v-model="admin.q" class="input" placeholder="Поиск: пользователи, вопросы, посты…" style="min-width: 260px" />
        <RouterLink to="/" class="btn sm">← В продукт</RouterLink>
      </div>
    </header>

    <div class="admin-body">
      <aside class="card admin-sidebar">
        <nav class="admin-nav">
          <RouterLink :to="{ name: 'admin-dashboard' }" :class="linkClass('dashboard')">
            <span class="label"><LayoutDashboard class="i" /><span>Дашборд</span></span>
          </RouterLink>
          <RouterLink :to="{ name: 'admin-users' }" :class="linkClass('users')">
            <span class="label"><Users class="i" /><span>Пользователи</span></span>
          </RouterLink>
          <RouterLink :to="{ name: 'admin-questions' }" :class="linkClass('questions')">
            <span class="label"><CircleHelp class="i" /><span>Вопросы</span></span>
          </RouterLink>
          <RouterLink :to="{ name: 'admin-posts' }" :class="linkClass('posts')">
            <span class="label"><FileText class="i" /><span>Публикации</span></span>
          </RouterLink>
          <RouterLink :to="{ name: 'admin-comments' }" :class="linkClass('comments')">
            <span class="label"><MessageSquare class="i" /><span>Комментарии</span></span>
          </RouterLink>
          <RouterLink :to="{ name: 'admin-tags' }" :class="linkClass('tags')">
            <span class="label"><Tags class="i" /><span>Теги</span></span>
          </RouterLink>
          <RouterLink :to="{ name: 'admin-campuses' }" :class="linkClass('campuses')">
            <span class="label"><MapPin class="i" /><span>Кампусы</span></span>
          </RouterLink>
          <RouterLink :to="{ name: 'admin-reports' }" :class="linkClass('reports')">
            <span class="label"><ShieldAlert class="i" /><span>Жалобы</span></span>
            <span v-if="admin.openReports" class="pill accent">{{ admin.openReports }}</span>
          </RouterLink>
          <RouterLink :to="{ name: 'admin-settings' }" :class="linkClass('settings')">
            <span class="label"><Settings class="i" /><span>Настройки</span></span>
          </RouterLink>
        </nav>

        <div class="line" style="margin: 12px 0"></div>
        <div class="muted">
          Доступ: <span class="kbd">user.role === Admin</span>
        </div>
      </aside>

      <main class="card admin-main">
        <RouterView />
      </main>
    </div>
  </section>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAdminStore } from '@/stores/admin'
import {
  LayoutDashboard,
  Users,
  CircleHelp,
  FileText,
  MessageSquare,
  Tags,
  MapPin,
  ShieldAlert,
  Settings,
} from 'lucide-vue-next'

const admin = useAdminStore()
const router = useRouter()

onMounted(() => {
  // В будущем: тут можно загрузить метрики/сводку с API
  admin.loadReports()
})

function linkClass(k: string) {
  const active = router.currentRoute.value.name === `admin-${k}`
  return active ? 'active' : ''
}
</script>

<style scoped>
.i { width: 16px; height: 16px; opacity: .9 }
</style>
