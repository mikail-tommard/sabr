<template>
  <div class="stack" style="gap: 12px">
    <div class="row" style="justify-content: space-between; gap: 10px; flex-wrap: wrap">
      <div class="title">Комментарии</div>
      <div class="muted">Удаление — локально</div>
    </div>

    <div class="card" style="padding: 10px; box-shadow:none; overflow:auto">
      <table class="table" style="width:100%; min-width: 900px">
        <thead>
          <tr>
            <th>Комментарий</th>
            <th>Автор</th>
            <th>Контекст</th>
            <th style="text-align:right">Действия</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="c in filtered" :key="c.id">
            <td style="max-width: 520px">
              <div style="white-space: pre-wrap">{{ c.content }}</div>
            </td>
            <td class="muted">{{ content.user(c.authorId)?.name }}</td>
            <td class="muted">
              <RouterLink v-if="c.parentType==='question'" :to="`/questions/${c.parentId}`" class="hover-accent">question: {{ c.parentId }}</RouterLink>
              <RouterLink v-else-if="c.parentType==='post'" :to="`/posts/${c.parentId}`" class="hover-accent">post: {{ c.parentId }}</RouterLink>
              <span v-else class="muted">answer: {{ c.parentId }}</span>
            </td>
            <td style="text-align:right">
              <button class="btn danger" @click="admin.deleteComment(c.id)">Удалить</button>
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
  if (!q) return content.comments
  return content.comments.filter(c => c.content.toLowerCase().includes(q))
})
</script>
