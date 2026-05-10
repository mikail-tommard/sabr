<template>
  <div class="stack" style="gap: 12px">
    <div class="row" style="justify-content: space-between; gap: 10px; flex-wrap: wrap">
      <div class="title">Теги</div>
      <div class="row" style="gap:10px">
        <input v-model="newTag" class="input mono" placeholder="например: graphql" style="min-width:220px" />
        <button class="btn primary" @click="create">Добавить</button>
      </div>
    </div>

    <div class="card" style="padding: 10px; box-shadow:none; overflow:auto">
      <table class="table" style="width:100%; min-width: 760px">
        <thead>
          <tr>
            <th>Тег</th>
            <th>Использований</th>
            <th style="text-align:right">Действия</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="t in filtered" :key="t.id">
            <td>
              <input class="input mono" :value="t.label" @change="admin.renameTag(t.id, ($event.target as HTMLInputElement).value)" />
            </td>
            <td class="muted">{{ t.usageCount }}</td>
            <td style="text-align:right">
              <button class="btn danger" @click="admin.deleteTag(t.id)">Удалить</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useAdminStore } from '@/stores/admin'
import { useContentStore } from '@/stores/content'

const admin = useAdminStore()
const content = useContentStore()
const newTag = ref('')

const filtered = computed(() => {
  const q = admin.q.trim().toLowerCase()
  if (!q) return content.tags
  return content.tags.filter(t => t.label.toLowerCase().includes(q))
})

function create() {
  admin.createTag(newTag.value)
  newTag.value = ''
}
</script>
