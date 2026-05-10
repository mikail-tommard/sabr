<template>
  <div class="stack" style="gap: 12px">
    <div class="row" style="justify-content: space-between; gap: 10px; flex-wrap: wrap">
      <div class="title">Модерация публикаций</div>
      <div class="muted">Feature / скрыть / удалить</div>
    </div>

    <div class="card" style="padding: 10px; box-shadow:none; overflow:auto">
      <table class="table" style="width:100%; min-width: 900px">
        <thead>
          <tr>
            <th>Публикация</th>
            <th>Автор</th>
            <th>Кампус</th>
            <th>Статус</th>
            <th>Лайки</th>
            <th>Комменты</th>
            <th style="text-align:right">Действия</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="p in filtered" :key="p.id">
            <td style="max-width: 420px">
              <div class="row" style="gap:8px; align-items:center">
                <span v-if="p.featured" class="pill">FEATURED</span>
                <RouterLink :to="`/posts/${p.id}`" style="font-weight:650" class="hover-accent">{{ p.title }}</RouterLink>
              </div>
              <div class="muted" style="margin-top:4px">{{ p.summary }}</div>
            </td>
            <td class="muted">{{ content.user(p.authorId)?.name }}</td>
            <td class="muted">{{ content.campus(p.campusId)?.name }}</td>
            <td>
              <span class="pill" :style="pillStyle(p.status)">{{ p.status || 'active' }}</span>
            </td>
            <td>{{ p.likes }}</td>
            <td>{{ p.commentsCount }}</td>
            <td style="text-align:right">
              <button class="btn" @click="admin.toggleFeaturedPost(p.id)">{{ p.featured ? 'Unfeature' : 'Feature' }}</button>
              <button class="btn" @click="admin.setPostStatus(p.id, 'hidden')">Скрыть</button>
              <button class="btn" @click="admin.setPostStatus(p.id, 'active')">Показать</button>
              <button class="btn danger" @click="admin.setPostStatus(p.id, 'deleted')">Удалить</button>
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
import type { ContentStatus } from '@/types'

const admin = useAdminStore()
const content = useContentStore()

const filtered = computed(() => {
  const q = admin.q.trim().toLowerCase()
  const all = content.posts
  if (!q) return all
  return all.filter(x => x.title.toLowerCase().includes(q) || x.summary.toLowerCase().includes(q))
})

function pillStyle(status?: ContentStatus) {
  const s = status || 'active'
  if (s === 'active') return ''
  if (s === 'hidden') return 'background: color-mix(in srgb, var(--accent) 12%, transparent); color: var(--accent)'
  return 'background: color-mix(in srgb, var(--danger) 20%, transparent); color: var(--danger)'
}
</script>
