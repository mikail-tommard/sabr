<template>
  <div class="grid-main">
    <aside class="desktop-only stack">
      <div class="card sidebar-card"><div class="title">Популярные теги</div><div class="row" style="flex-wrap:wrap;margin-top:8px"><TagChip v-for="t in s.topTags" :key="t.id" :label="t.label" /></div></div>
      <div class="card sidebar-card"><div class="title">Кампусы</div><div class="list" style="margin-top:8px"><RouterLink v-for="c in s.campuses" :key="c.id" class="list-item" :to="`/campuses/${c.id}`">{{ c.name }}</RouterLink></div></div>
    </aside>
    <main class="stack">
      <div class="card" style="padding:14px">
        <div class="tabs">
          <button v-for="t in tabs" :key="t.value" class="btn" :class="{active:s.feedTab===t.value}" @click="s.feedTab=t.value as any">{{ t.label }}</button>
        </div>
        <div class="row" style="margin-top:10px;flex-wrap:wrap">
          <select v-model="s.sort" class="select" style="max-width:180px"><option value="new">Новые</option><option value="popular">Популярные</option><option value="unanswered">Без ответа</option><option value="best-answer">С лучшим ответом</option></select>
          <select v-model="s.campusFilter" class="select" style="max-width:220px"><option value="">Все кампусы</option><option v-for="c in s.campuses" :key="c.id" :value="c.id">{{ c.name }}</option></select>
          <select v-model="s.tagFilter" class="select" style="max-width:180px"><option value="">Все теги</option><option v-for="t in s.tags" :key="t.id" :value="t.id">#{{ t.label }}</option></select>
        </div>
      </div>
      <div v-if="s.loading" class="stack"><div v-for="i in 4" :key="i" class="skeleton" style="height:112px"></div></div>
      <template v-else>
        <ContentCard v-for="item in s.feedItems" :key="`${item.type}-${item.id}`" :item="item" />
        <EmptyState v-if="!s.feedItems.length" title="Нет результатов" description="Попробуйте изменить фильтры или поиск" />
      </template>
    </main>
    <aside class="desktop-only stack">
      <div class="card sidebar-card"><div class="title">Активные пользователи</div><div class="list" style="margin-top:8px"><UserBadge v-for="u in s.topUsers" :key="u.id" :user="u" /></div></div>
      <div class="card sidebar-card"><div class="title">Топ вопросов недели</div><div class="list" style="margin-top:8px"><RouterLink v-for="q in s.topQuestions" :key="q.id" class="list-item" :to="`/questions/${q.id}`">{{ q.title }}</RouterLink></div></div>
    </aside>
  </div>
</template>
<script setup lang="ts">
import { useContentStore } from '@/stores/content'
import ContentCard from '@/components/ContentCard.vue'; import EmptyState from '@/components/EmptyState.vue'; import TagChip from '@/components/TagChip.vue'; import UserBadge from '@/components/UserBadge.vue'
const s = useContentStore();
const tabs = [ {value:'all',label:'Все'},{value:'questions',label:'Вопросы'},{value:'posts',label:'Публикации'},{value:'my-campus',label:'Мой кампус'} ]
</script>
