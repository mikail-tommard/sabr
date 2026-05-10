<template>
  <div class="stack" style="margin-top:16px">
    <div class="card" style="padding:14px"><div class="row"><input v-model="q" class="input" placeholder="Поиск по кампусам" /><RouterLink to="/" class="btn">На главную</RouterLink></div></div>
    <div class="stack">
      <RouterLink v-for="c in filtered" :key="c.id" :to="`/campuses/${c.id}`" class="card" style="padding:14px;display:block">
        <div class="row" style="justify-content:space-between;flex-wrap:wrap"><div><div class="title">{{ c.name }}</div><div class="muted">{{ c.city }} · {{ c.description }}</div></div><span class="tag">{{ c.membersCount }} участников</span></div>
        <div class="row muted" style="margin-top:10px"><span>Публикации: {{ c.postsCount }}</span><span>Вопросы: {{ c.questionsCount }}</span></div>
      </RouterLink>
    </div>
    <EmptyState v-if="!filtered.length" title="Нет результатов" description="Попробуйте другой запрос" />
  </div>
</template>
<script setup lang="ts">import { computed, ref } from 'vue'; import { useContentStore } from '@/stores/content'; import EmptyState from '@/components/EmptyState.vue'; const s=useContentStore(); const q=ref(''); const filtered=computed(()=> s.campuses.filter(c => (c.name+c.city+c.description).toLowerCase().includes(q.value.toLowerCase())))</script>
