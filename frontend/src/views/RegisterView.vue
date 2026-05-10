<template>
  <div class="card" style="padding:18px">
    <div>
      <div class="title" style="font-size:18px">Регистрация</div>
      <div class="muted" style="margin-top:6px">UI‑заглушка без backend. Можно красиво показать будущий flow.</div>
    </div>

    <div class="stack" style="margin-top:14px">
      <div class="row" style="gap:10px">
        <div style="flex:1">
          <div class="muted">Имя</div>
          <input v-model="f.name" class="input" placeholder="Иван" />
        </div>
        <div style="flex:1">
          <div class="muted">Username</div>
          <input v-model="f.username" class="input mono" placeholder="ivan_dev" />
        </div>
      </div>

      <div>
        <div class="muted">Email</div>
        <input v-model="f.email" class="input" placeholder="you@example.com" />
      </div>

      <div class="row" style="gap:10px">
        <div style="flex:1">
          <div class="muted">Кампус</div>
          <select v-model="f.campusId" class="select">
            <option value="">Выберите кампус</option>
            <option v-for="c in s.campuses" :key="c.id" :value="c.id">{{ c.name }}</option>
          </select>
        </div>
        <div style="flex:1">
          <div class="muted">Роль</div>
          <div class="card" style="padding:10px; box-shadow:none; border-style:dashed">
            <div class="small">По умолчанию: <b>Student</b></div>
            <div class="muted" style="margin-top:6px">
              Mentor/Admin назначаются после верификации или админом — поэтому не спрашиваем роль при регистрации.
            </div>
          </div>
        </div>
      </div>

      <div>
        <div class="muted">Пароль</div>
        <input v-model="f.password" class="input mono" type="password" placeholder="••••••••" />
      </div>

      <button class="btn primary" @click="fakeRegister">Создать аккаунт (mock)</button>

      <div class="row" style="justify-content:space-between; gap:10px">
        <div class="muted">Уже есть аккаунт?</div>
        <RouterLink to="/login" class="btn ghost-accent">Войти</RouterLink>
      </div>

      <div class="card" style="padding:12px; background: color-mix(in srgb, var(--surface) 88%, var(--bg)); box-shadow:none">
        <div class="title">Как будет в backend</div>
        <div class="muted" style="margin-top:6px">
          Здесь подключим API: <span class="kbd">POST /auth/register</span> и затем <span class="kbd">POST /auth/login</span>.
          Сейчас регистрация — демонстрация UI.
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useUiStore } from '@/stores/ui'
import { useContentStore } from '@/stores/content'

const router = useRouter()
const ui = useUiStore()
const s = useContentStore()

const f = reactive({
  name: '',
  username: '',
  email: '',
  campusId: '',
  password: '',
})

function fakeRegister() {
  if (f.name.trim().length < 2 || f.username.trim().length < 3) {
    ui.toast('error', 'Заполните имя и username')
    return
  }
  ui.toast('info', 'Регистрация — UI-заглушка. Используйте mock login.')
  router.push('/login')
}
</script>
