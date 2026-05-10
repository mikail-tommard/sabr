<template>
  <div class="stack" style="gap: 12px">
    <div class="row" style="justify-content: space-between; gap: 10px; flex-wrap: wrap">
      <div class="title">Жалобы (Reports)</div>
      <div class="muted">Моковые данные, статусы меняются локально</div>
    </div>

    <div class="card" style="padding: 10px; box-shadow:none; overflow:auto">
      <table class="table" style="width:100%; min-width: 960px">
        <thead>
          <tr>
            <th>ID</th>
            <th>Причина</th>
            <th>Цель</th>
            <th>От</th>
            <th>Заметка</th>
            <th>Статус</th>
            <th style="text-align:right">Действия</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="r in filtered" :key="r.id">
            <td class="mono">{{ r.id }}</td>
            <td><span class="pill">{{ r.reason }}</span></td>
            <td>
              <span class="pill">{{ r.targetType }}</span>
              <span class="mono" style="margin-left: 8px">{{ r.targetId }}</span>
            </td>
            <td class="muted">{{ content.user(r.reporterId)?.name }}</td>
            <td class="muted" style="max-width: 320px">{{ r.note }}</td>
            <td>
              <select class="select" :value="r.status" @change="admin.setReportStatus(r.id, ($event.target as HTMLSelectElement).value as any)">
                <option value="open">open</option>
                <option value="in_review">in_review</option>
                <option value="resolved">resolved</option>
              </select>
            </td>
            <td style="text-align:right">
              <RouterLink v-if="r.targetType==='question'" class="btn ghost-accent" :to="`/questions/${r.targetId}`">Открыть</RouterLink>
              <RouterLink v-else-if="r.targetType==='post'" class="btn ghost-accent" :to="`/posts/${r.targetId}`">Открыть</RouterLink>
              <RouterLink v-else-if="r.targetType==='user'" class="btn ghost-accent" :to="`/profile/${r.targetId}`">Профиль</RouterLink>
              <button v-else class="btn" disabled>—</button>
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
  if (!q) return admin.reports
  return admin.reports.filter(r =>
    r.id.toLowerCase().includes(q) ||
    r.targetId.toLowerCase().includes(q) ||
    (r.note || '').toLowerCase().includes(q)
  )
})
</script>
