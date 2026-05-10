<template>
  <div class="stack" style="gap: 12px">
    <div class="row" style="gap: 12px; flex-wrap: wrap">
      <div class="card" style="padding: 12px; flex: 1; min-width: 220px; box-shadow:none">
        <div class="muted">Пользователи</div>
        <div class="title" style="font-size: 22px">{{ content.users.length }}</div>
      </div>
      <div class="card" style="padding: 12px; flex: 1; min-width: 220px; box-shadow:none">
        <div class="muted">Вопросы</div>
        <div class="title" style="font-size: 22px">{{ content.questions.length }}</div>
      </div>
      <div class="card" style="padding: 12px; flex: 1; min-width: 220px; box-shadow:none">
        <div class="muted">Публикации</div>
        <div class="title" style="font-size: 22px">{{ content.posts.length }}</div>
      </div>
      <div class="card" style="padding: 12px; flex: 1; min-width: 220px; box-shadow:none">
        <div class="muted">Открытые жалобы</div>
        <div class="title" style="font-size: 22px">{{ admin.openReports }}</div>
      </div>
    </div>

    <div class="card" style="padding: 12px; box-shadow:none">
      <div class="title">Быстрые действия</div>
      <div class="row" style="gap: 10px; margin-top: 10px; flex-wrap: wrap">
        <RouterLink class="btn primary" to="/admin/reports">Разобрать жалобы</RouterLink>
        <RouterLink class="btn" to="/admin/questions">Модерация вопросов</RouterLink>
        <RouterLink class="btn" to="/admin/posts">Модерация публикаций</RouterLink>
        <RouterLink class="btn" to="/admin/users">Управление пользователями</RouterLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAdminStore } from '@/stores/admin'
import { useContentStore } from '@/stores/content'

const admin = useAdminStore()
const content = useContentStore()
</script>
