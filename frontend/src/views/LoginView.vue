<template>
  <div class="card" style="padding:18px">
    <div>
      <div class="title" style="font-size:18px">Вход</div>
      <div class="muted" style="margin-top:6px">Mock‑авторизация для MVP: выберите тестового пользователя.</div>
    </div>

    <div class="stack" style="margin-top:14px">
      <div>
        <div class="muted">Пользователь</div>
        <select v-model="selected" class="select">
          <option value="">Выберите пользователя</option>
          <option v-for="u in users" :key="u.id" :value="u.id">{{ u.avatar }} {{ u.name }} (@{{ u.username }})</option>
        </select>
      </div>

      <div class="row" style="gap:10px">
        <button class="btn primary" style="flex:1" @click="submit">Войти</button>
        <RouterLink to="/" class="btn" style="flex:1; text-align:center">На главную</RouterLink>
      </div>

      <div class="line"></div>

      <div class="row" style="justify-content:space-between; gap:10px">
        <div class="muted">Нет аккаунта?</div>
        <RouterLink to="/register" class="btn ghost-accent">Регистрация</RouterLink>
      </div>

      <div class="card" style="padding:12px; background: color-mix(in srgb, var(--surface) 88%, var(--bg)); box-shadow:none">
        <div class="title">Заметка</div>
        <div class="muted" style="margin-top:6px">
          В реальном проекте здесь будет форма (email/пароль) и запрос к API.
          Сейчас login просто сохраняет userId в <span class="kbd">localStorage</span>.
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useContentStore } from '@/stores/content'
import { useUiStore } from '@/stores/ui'

const router = useRouter()
const route = useRoute()
const auth = useAuthStore()
const content = useContentStore()
const ui = useUiStore()

const users = computed(() => content.users)
const selected = ref('')

function submit() {
  if (!selected.value) {
    ui.toast('error', 'Выберите пользователя')
    return
  }
  auth.loginAs(selected.value)
  ui.toast('success', 'Вы вошли (mock)')
  const returnTo = typeof route.query.returnTo === 'string' ? route.query.returnTo : '/'
  router.push(returnTo)
}
</script>
