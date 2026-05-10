<template>
  <div v-if="u" class="stack" style="margin-top:16px">
    <div class="card" style="padding:16px">
      <div class="row" style="justify-content:space-between;align-items:flex-start;flex-wrap:wrap">
        <div class="row"><div style="font-size:42px">{{ u.avatar }}</div><div><div class="title" style="font-size:22px">{{ u.name }}</div><div class="muted">@{{ u.username }} · {{ u.role }} · {{ s.campus(u.campusId)?.name }}</div><div class="row" style="margin-top:6px"><span class="tag" v-for="b in u.badges" :key="b">🏅 {{ b }}</span></div></div></div>
        <div class="card" style="padding:10px 12px">Рейтинг: <b>{{ u.rating }}</b></div>
      </div>
      <div class="kv" style="margin-top:12px">
        <div class="item"><div class="muted">Вопросы</div><div class="title">{{ stats.q }}</div></div>
        <div class="item"><div class="muted">Ответы</div><div class="title">{{ stats.a }}</div></div>
        <div class="item"><div class="muted">Публикации</div><div class="title">{{ stats.p }}</div></div>
        <div class="item"><div class="muted">Комментарии</div><div class="title">{{ stats.c }}</div></div>
      </div>
    </div>
    <div class="card" style="padding:14px">
      <div class="tabs"><button class="btn" :class="{active:tab==='questions'}" @click="tab='questions'">Вопросы</button><button class="btn" :class="{active:tab==='answers'}" @click="tab='answers'">Ответы</button><button class="btn" :class="{active:tab==='posts'}" @click="tab='posts'">Публикации</button><button class="btn" :class="{active:tab==='comments'}" @click="tab='comments'">Комментарии</button></div>
      <div class="stack" style="margin-top:12px">
        <template v-if="tab==='questions'"><ContentCard v-for="q in myQuestions" :key="q.id" :item="q" /><EmptyState v-if="!myQuestions.length" title="Нет вопросов" description="Пользователь еще не публиковал вопросы" /></template>
        <template v-else-if="tab==='posts'"><ContentCard v-for="p in myPosts" :key="p.id" :item="p" /><EmptyState v-if="!myPosts.length" title="Нет публикаций" description="Пусто" /></template>
        <template v-else-if="tab==='answers'"><div v-for="a in myAnswers" :key="a.id" class="card" style="padding:12px"><RouterLink :to="`/questions/${a.questionId}`" class="title">Ответ в вопросе #{{ a.questionId }}</RouterLink><div style="margin-top:6px">{{ a.content }}</div></div><EmptyState v-if="!myAnswers.length" title="Нет ответов" description="Пусто" /></template>
        <template v-else><div v-for="c in myComments" :key="c.id" class="card" style="padding:12px">{{ c.content }}</div><EmptyState v-if="!myComments.length" title="Нет комментариев" description="Пусто" /></template>
      </div>
    </div>
  </div>
  <EmptyState v-else title="Пользователь не найден" description="Проверьте ссылку" />
</template>
<script setup lang="ts">
import { computed, ref } from 'vue'; import { useRoute } from 'vue-router'; import { useContentStore } from '@/stores/content'; import ContentCard from '@/components/ContentCard.vue'; import EmptyState from '@/components/EmptyState.vue'
const route=useRoute(); const s=useContentStore(); const tab=ref<'questions'|'answers'|'posts'|'comments'>('questions')
const u = computed(()=> s.user(route.params.id as string));
const myQuestions = computed(()=> u.value ? s.questions.filter(x=>x.authorId===u.value!.id) : []); const myPosts = computed(()=> u.value ? s.posts.filter(x=>x.authorId===u.value!.id) : []); const myAnswers = computed(()=> u.value ? s.answers.filter(x=>x.authorId===u.value!.id) : []); const myComments = computed(()=> u.value ? s.comments.filter(x=>x.authorId===u.value!.id) : [])
const stats = computed(()=> ({ q: myQuestions.value.length, p: myPosts.value.length, a: myAnswers.value.length, c: myComments.value.length }))
</script>
