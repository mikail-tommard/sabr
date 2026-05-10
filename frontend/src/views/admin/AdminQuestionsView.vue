<template>
  <div class="stack" style="gap: 14px">
    <div class="admin-page-head">
      <div>
        <div class="admin-page-title">Вопросы</div>
        <div class="admin-page-sub">Модерация: pin / скрыть / удалить</div>
      </div>
      <div class="muted">Найдено: {{ filtered.length }}</div>
    </div>

    <div class="admin-table-wrap">
      <table class="admin-table">
        <thead>
          <tr>
            <th>Вопрос</th>
            <th>Автор</th>
            <th>Кампус</th>
            <th>Статус</th>
            <th>Апвоуты</th>
            <th>Ответы</th>
            <th style="text-align:right">Действия</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="q in filtered" :key="q.id">
            <td style="max-width: 520px">
              <div class="row" style="gap:8px; align-items:center">
                <span v-if="q.pinned" class="pill accent">PIN</span>
                <RouterLink :to="`/questions/${q.id}`" style="font-weight:650" class="hover-accent">{{ q.title }}</RouterLink>
              </div>
              <div class="muted" style="margin-top:4px" v-if="q.summary">{{ q.summary }}</div>
            </td>
            <td class="muted">{{ content.user(q.authorId)?.name }}</td>
            <td class="muted">{{ content.campus(q.campusId)?.name }}</td>
            <td>
              <span class="pill" :class="statusClass(q.status)">{{ q.status || 'active' }}</span>
            </td>
            <td>{{ q.upvotes }}</td>
            <td>{{ q.answersCount }}</td>
            <td style="text-align:right">
              <div class="admin-actions">
                <button class="btn sm" @click="admin.togglePinQuestion(q.id)">{{ q.pinned ? 'Unpin' : 'Pin' }}</button>
                <button class="btn sm" @click="admin.setQuestionStatus(q.id, 'hidden')">Скрыть</button>
                <button class="btn sm" @click="admin.setQuestionStatus(q.id, 'active')">Показать</button>
                <button class="btn sm danger" @click="admin.setQuestionStatus(q.id, 'deleted')">Удалить</button>
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
import type { ContentStatus } from '@/types'

const admin = useAdminStore()
const content = useContentStore()

const filtered = computed(() => {
  const q = admin.q.trim().toLowerCase()
  const all = content.questions
  if (!q) return all
  return all.filter(x => x.title.toLowerCase().includes(q) || x.summary.toLowerCase().includes(q))
})

function statusClass(status?: ContentStatus) {
  const s = status || 'active'
  if (s === 'active') return 'success'
  if (s === 'hidden') return 'accent'
  return 'danger'
}
</script>
