<template>
  <div class="stack" style="gap: 14px">
    <div class="admin-page-head">
      <div>
        <div class="admin-page-title">Пользователи</div>
        <div class="admin-page-sub">Роли, бан, быстрый переход в профиль</div>
      </div>
      <div class="muted">Найдено: {{ filtered.length }}</div>
    </div>

    <div class="admin-table-wrap">
      <table class="admin-table" style="min-width: 860px">
        <thead>
          <tr>
            <th>Пользователь</th>
            <th>Кампус</th>
            <th>Роль</th>
            <th>Рейтинг</th>
            <th>Статус</th>
            <th style="text-align:right">Действия</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="u in filtered" :key="u.id">
            <td>
              <div class="row" style="gap:10px">
                <span style="font-size:18px">{{ u.avatar }}</span>
                <div>
                  <div style="font-weight:650">{{ u.name }}</div>
                  <div class="muted">@{{ u.username }}</div>
                </div>
              </div>
            </td>
            <td class="muted">{{ content.campus(u.campusId)?.name }}</td>
            <td>
              <select class="select" :value="u.role" @change="admin.setUserRole(u.id, ($event.target as HTMLSelectElement).value as any)">
                <option value="Student">Student</option>
                <option value="Mentor">Mentor</option>
                <option value="Alumni">Alumni</option>
                <option value="Admin">Admin</option>
              </select>
            </td>
            <td>{{ u.rating }}</td>
            <td>
              <span class="pill" :class="u.isBanned ? 'danger' : 'success'">{{ u.isBanned ? 'Banned' : 'Active' }}</span>
            </td>
            <td style="text-align:right">
              <div class="admin-actions">
                <button class="btn sm" @click="admin.setUserBanned(u.id, !u.isBanned)">{{ u.isBanned ? 'Разбанить' : 'Забанить' }}</button>
                <RouterLink class="btn sm ghost-accent" :to="`/profile/${u.id}`">Профиль</RouterLink>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAdminStore } from '@/stores/admin'
import { useContentStore } from '@/stores/content'

const admin = useAdminStore()
const content = useContentStore()

const filtered = computed(() => {
  const q = admin.q.trim().toLowerCase()
  if (!q) return content.users
  return content.users.filter(u =>
    u.name.toLowerCase().includes(q) ||
    u.username.toLowerCase().includes(q)
  )
})
</script>
