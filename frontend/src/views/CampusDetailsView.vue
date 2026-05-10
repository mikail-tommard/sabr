<template>
  <div v-if="campus" class="stack" style="margin-top:16px">
    <div class="card" style="padding:16px"><h1 style="margin:0">{{ campus.name }}</h1><div class="muted">{{ campus.city }} · {{ campus.description }}</div><div class="row muted" style="margin-top:10px"><span>Участников: {{ campus.membersCount }}</span><span>Публикаций: {{ campus.postsCount }}</span><span>Вопросов: {{ campus.questionsCount }}</span></div></div>
    <div class="page-grid">
      <div class="stack"><div class="card" style="padding:14px"><div class="title">Вопросы кампуса</div><div class="stack" style="margin-top:10px"><ContentCard v-for="q in qItems" :key="q.id" :item="q" /><EmptyState v-if="!qItems.length" title="Нет вопросов" description="Пусто" /></div></div></div>
      <div class="card sidebar-card"><div class="title">Публикации</div><div class="list" style="margin-top:8px"><RouterLink v-for="p in pItems" :key="p.id" class="list-item" :to="`/posts/${p.id}`">{{ p.title }}</RouterLink></div></div>
    </div>
  </div>
  <EmptyState v-else title="Кампус не найден" description="Проверьте ссылку" />
</template>
<script setup lang="ts">import { computed } from 'vue'; import { useRoute } from 'vue-router'; import { useContentStore } from '@/stores/content'; import ContentCard from '@/components/ContentCard.vue'; import EmptyState from '@/components/EmptyState.vue'; const route=useRoute(); const s=useContentStore(); const campus=computed(()=> s.campus(route.params.id as string)); const qItems=computed(()=> s.questions.filter(q=>q.campusId===route.params.id)); const pItems=computed(()=> s.posts.filter(p=>p.campusId===route.params.id).slice(0,6))</script>
