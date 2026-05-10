<template>
  <header class="header">
    <div class="container header-row">
      <RouterLink to="/" class="logo"><span class="logo-badge">S</span><span>Sabr</span></RouterLink>

      <div class="row mobile-hide" style="gap:8px; margin-left: 8px">
        <RouterLink to="/about" class="btn ghost-accent">О проекте</RouterLink>
        <RouterLink v-if="isAdmin" to="/admin" class="btn ghost-accent">Админка</RouterLink>
      </div>

      <div class="surface-subtle row" style="flex:1; padding:6px 8px; gap:8px; min-width: 160px;">
        <Search :size="16" class="muted" />
        <input v-model="q" class="input" placeholder="Поиск по заголовкам и тегам" style="border:none;box-shadow:none;background:transparent;padding:6px 2px" @keydown.enter="goSearch" />
      </div>

      <div class="row mobile-hide">
        <RouterLink to="/ask" class="btn">Задать вопрос</RouterLink>
        <RouterLink to="/write" class="btn">Написать пост</RouterLink>
      </div>

      <div class="row">
        <button class="btn icon-btn" :title="ui.theme==='dark' ? 'Светлая тема' : 'Тёмная тема'" @click="ui.toggleTheme()">
          <Sun v-if="ui.theme==='dark'" :size="16" />
          <Moon v-else :size="16" />
        </button>

        <Menu as="div" style="position:relative">
          <MenuButton class="btn icon-btn" title="Уведомления">
            <Bell :size="16" />
            <span v-if="n.unread" class="badge-dot" style="position:absolute; top:9px; right:9px"></span>
          </MenuButton>

          <Transition
            enter-active-class="page-enter-active"
            enter-from-class="page-enter-from"
            leave-active-class="page-leave-active"
            leave-to-class="page-leave-to"
          >
            <MenuItems class="card dropdown" style="transform-origin: top right">
              <div class="row" style="justify-content:space-between;padding:6px 6px 8px">
                <b>Уведомления</b>
                <button class="btn" @click="n.markAll">Прочитать всё</button>
              </div>
              <div v-if="n.items.length">
                <MenuItem v-for="it in n.items" :key="it.id" v-slot="{ active }">
                  <div class="list-item row" :style="{justifyContent:'space-between', border:'1px solid var(--border)', marginBottom:'6px', background: !it.isRead ? 'var(--accent-soft)' : active ? 'var(--surface-2)' : 'transparent'}">
                    <RouterLink :to="it.href" style="min-width:0; flex:1">
                      <div class="small" :style="{fontWeight: it.isRead ? '500' : '650'}">{{ it.text }}</div>
                    </RouterLink>
                    <button class="btn" @click.stop="n.toggle(it.id)">{{ it.isRead ? '↺' : '✓' }}</button>
                  </div>
                </MenuItem>
              </div>
              <div v-else class="empty">Пока пусто</div>
            </MenuItems>
          </Transition>
        </Menu>

        <template v-if="me">
          <RouterLink :to="`/profile/${me.id}`" class="btn">{{ me.avatar }} {{ me.username }}</RouterLink>
          <button class="btn" @click="logout">Выйти</button>
        </template>
        <template v-else>
          <RouterLink to="/login" class="btn primary">Войти</RouterLink>
        </template>
      </div>
    </div>

    <div class="container mobile-actions" style="padding-bottom:10px">
      <RouterLink to="/ask" class="btn" style="flex:1">Задать вопрос</RouterLink>
      <RouterLink to="/write" class="btn" style="flex:1">Написать пост</RouterLink>
    </div>
  </header>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { Bell, Moon, Search, Sun } from 'lucide-vue-next'
import { Menu, MenuButton, MenuItems, MenuItem } from '@headlessui/vue'
import { useUiStore } from '@/stores/ui'
import { useNotificationsStore } from '@/stores/notifications'
import { useAuthStore } from '@/stores/auth'
import { useContentStore } from '@/stores/content'

const ui = useUiStore()
const n = useNotificationsStore()
const auth = useAuthStore()
const content = useContentStore()
const router = useRouter()

const q = ref('')

const me = computed(() => content.users.find(u => u.id === auth.currentUserId))
const isAdmin = computed(() => me.value?.role === 'Admin')
const goSearch = () => router.push({ name: 'search', query: { q: q.value } })
const logout = () => { auth.logout(); ui.toast('success', 'Вы вышли из аккаунта') }

// dropdown управляется headlessui (Menu)
</script>
