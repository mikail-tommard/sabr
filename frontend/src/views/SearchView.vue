<template>
  <div class="stack" style="margin-top:16px">
    <div class="card" style="padding:14px">
      <div class="row" style="flex-wrap:wrap"><input v-model="s.search" class="input" placeholder="Поиск..." /><select v-model="s.typeFilter" class="select" style="max-width:200px"><option value="">Все типы</option><option value="question">Вопросы</option><option value="post">Публикации</option></select><select v-model="s.campusFilter" class="select" style="max-width:240px"><option value="">Все кампусы</option><option v-for="c in s.campuses" :key="c.id" :value="c.id">{{ c.name }}</option></select></div>
    </div>
    <ContentCard v-for="item in s.feedItems" :key="`${item.type}-${item.id}`" :item="item" />
    <EmptyState v-if="!s.feedItems.length" title="Нет результатов поиска" description="Измените запрос или фильтры" />
  </div>
</template>
<script setup lang="ts">import { onMounted } from 'vue'; import { useRoute } from 'vue-router'; import { useContentStore } from '@/stores/content'; import ContentCard from '@/components/ContentCard.vue'; import EmptyState from '@/components/EmptyState.vue'; const s=useContentStore(); const route=useRoute(); onMounted(()=> { const q=route.query.q; if(typeof q==='string') s.search=q })</script>
