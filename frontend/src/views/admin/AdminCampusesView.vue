<template>
  <div class="stack" style="gap: 12px">
    <div class="row" style="justify-content: space-between; gap: 10px; flex-wrap: wrap">
      <div class="title">Кампусы</div>
      <button class="btn primary" @click="showCreate = !showCreate">{{ showCreate ? 'Скрыть' : 'Добавить кампус' }}</button>
    </div>

    <div v-if="showCreate" class="card" style="padding: 12px; box-shadow:none">
      <div class="grid" style="grid-template-columns: 1fr 1fr; gap: 10px">
        <div>
          <div class="muted">Название</div>
          <input v-model="f.name" class="input" placeholder="Sabr Amsterdam" />
        </div>
        <div>
          <div class="muted">Город</div>
          <input v-model="f.city" class="input" placeholder="Амстердам" />
        </div>
        <div style="grid-column: 1 / -1">
          <div class="muted">Описание</div>
          <textarea v-model="f.description" class="input" style="min-height: 90px"></textarea>
        </div>
      </div>
      <div class="row" style="justify-content:flex-end; gap: 10px; margin-top: 10px">
        <button class="btn" @click="reset">Сброс</button>
        <button class="btn primary" @click="create">Создать</button>
      </div>
    </div>

    <div class="card" style="padding: 10px; box-shadow:none; overflow:auto">
      <table class="table" style="width:100%; min-width: 900px">
        <thead>
          <tr>
            <th>Название</th>
            <th>Город</th>
            <th>Участники</th>
            <th>Посты</th>
            <th>Вопросы</th>
            <th style="text-align:right">Действия</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="c in filtered" :key="c.id">
            <td><input class="input" :value="c.name" @change="admin.updateCampus(c.id, { name: ($event.target as HTMLInputElement).value })" /></td>
            <td><input class="input" :value="c.city" @change="admin.updateCampus(c.id, { city: ($event.target as HTMLInputElement).value })" /></td>
            <td class="muted">{{ c.membersCount }}</td>
            <td class="muted">{{ c.postsCount }}</td>
            <td class="muted">{{ c.questionsCount }}</td>
            <td style="text-align:right">
              <RouterLink class="btn ghost-accent" :to="`/campuses/${c.id}`">Открыть</RouterLink>
              <button class="btn danger" @click="tryDelete(c.id)">Удалить</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="muted">
      Удаление кампуса в MVP возможно только если к нему не привязаны пользователи.
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue'
import { useAdminStore } from '@/stores/admin'
import { useContentStore } from '@/stores/content'
import { useUiStore } from '@/stores/ui'

const admin = useAdminStore()
const content = useContentStore()
const ui = useUiStore()

const showCreate = ref(false)
const f = reactive({ name: '', city: '', description: '' })

const filtered = computed(() => {
  const q = admin.q.trim().toLowerCase()
  if (!q) return content.campuses
  return content.campuses.filter(c => c.name.toLowerCase().includes(q) || c.city.toLowerCase().includes(q))
})

function reset(){ f.name=''; f.city=''; f.description='' }
function create(){
  if (f.name.trim().length < 3 || f.city.trim().length < 2) {
    ui.toast('error', 'Введите название и город')
    return
  }
  admin.createCampus({ name: f.name, city: f.city, description: f.description })
  ui.toast('success', 'Кампус создан')
  reset(); showCreate.value = false
}

function tryDelete(id: string) {
  const hasUsers = content.users.some(u => u.campusId === id)
  if (hasUsers) {
    ui.toast('error', 'Нельзя удалить: есть пользователи в кампусе')
    return
  }
  admin.deleteCampus(id)
  ui.toast('success', 'Кампус удалён')
}
</script>
