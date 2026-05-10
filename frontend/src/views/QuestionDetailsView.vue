<template>
  <div v-if="q" class="page-grid" style="margin-top:16px">
    <div class="stack">
      <div class="card" style="padding:16px;display:flex;gap:12px">
        <VotePanel :value="q.upvotes" label="апвоуты" @upvote="s.voteQuestion(q.id)" />
        <div class="content-grow stack" style="gap:10px">
          <div class="row" style="justify-content:space-between;align-items:flex-start"><h1 style="margin:0;font-size:22px">{{ q.title }}</h1><span class="tag">{{ q.solved ? '✅ Решено' : '⏳ Не решено' }}</span></div>
          <div>{{ q.description }}</div>
          <pre v-if="q.code" class="code"><code>{{ q.code }}</code></pre>
          <div v-if="q.tried" class="card" style="padding:10px"><b>Что пробовал:</b><div style="margin-top:6px">{{ q.tried }}</div></div>
          <div class="row" style="flex-wrap:wrap"><TagChip v-for="tid in q.tagIds" :key="tid" :label="s.tag(tid)?.label || tid" /></div>
          <div class="row" style="justify-content:space-between"><UserBadge :user="author" /><span class="muted">{{ fmtDate(q.createdAt) }}</span></div>
        </div>
      </div>

      <CommentList :items="qComments" :users="s.users" @add="txt=>s.addComment('question', q.id, meId, txt)" />

      <div class="card" style="padding:14px">
        <div class="title">Ответы ({{ ans.length }})</div>
        <div v-if="ans.length" class="stack" style="margin-top:10px">
          <div v-for="a in ans" :key="a.id" class="card" style="padding:12px;display:flex;gap:10px">
            <VotePanel :value="a.upvotes" label="апвоуты" @upvote="s.voteAnswer(a.id)" />
            <div class="content-grow stack" style="gap:8px">
              <div style="white-space:pre-wrap">{{ a.content }}</div>
              <div class="row" style="justify-content:space-between;flex-wrap:wrap"><UserBadge :user="s.user(a.authorId)!" /><div class="row"><span class="muted">{{ fmtDate(a.createdAt) }}</span><button class="btn" @click="s.markBest(q.id, a.id)">{{ a.isBest ? '⭐ Лучший ответ' : 'Сделать лучшим' }}</button></div></div>
              <CommentList :items="s.commentsFor('answer', a.id)" :users="s.users" @add="txt=>s.addComment('answer', a.id, meId, txt)" />
            </div>
          </div>
        </div>
        <EmptyState v-else title="Пока нет ответов" description="Напишите первый ответ" />
      </div>

      <div class="card" style="padding:14px">
        <div class="title">Написать ответ</div>
        <textarea v-model="answerText" class="textarea" style="margin-top:10px" placeholder="Опишите решение, объясните почему оно работает"></textarea>
        <div class="row" style="justify-content:flex-end;margin-top:10px"><button class="btn primary" @click="submit">Отправить ответ</button></div>
      </div>
    </div>
    <aside class="stack">
      <div class="card sidebar-card"><div class="title">Автор вопроса</div><div style="margin-top:10px"><UserBadge :user="author" /></div><div class="muted" style="margin-top:8px">Рейтинг: {{ author.rating }}</div><div class="muted">Кампус: {{ s.campus(author.campusId)?.name }}</div></div>
    </aside>
  </div>
  <EmptyState v-else title="Вопрос не найден" description="Возможно, ссылка устарела" />
</template>
<script setup lang="ts">
import { computed, ref } from 'vue'; import { useRoute } from 'vue-router'; import { useContentStore } from '@/stores/content'; import { useAuthStore } from '@/stores/auth'; import { useUiStore } from '@/stores/ui'
import VotePanel from '@/components/VotePanel.vue'; import TagChip from '@/components/TagChip.vue'; import UserBadge from '@/components/UserBadge.vue'; import CommentList from '@/components/CommentList.vue'; import EmptyState from '@/components/EmptyState.vue'; import { fmtDate } from '@/utils'
const route=useRoute(); const s=useContentStore(); const auth=useAuthStore(); const ui=useUiStore(); const answerText=ref('')
const q = computed(()=> s.qById(route.params.id as string)); const author = computed(()=> q.value ? s.user(q.value.authorId)! : s.users[0]!); const ans = computed(()=> q.value ? s.answersFor(q.value.id) : []); const qComments = computed(()=> q.value ? s.commentsFor('question', q.value.id) : [])
const meId = computed(()=> auth.currentUserId || s.users[0]?.id || 'u1')
const submit = ()=> { if(!q.value) return; if(!answerText.value.trim()) { ui.toast('error','Введите текст ответа'); return } s.addAnswer(q.value.id, meId.value, answerText.value.trim()); answerText.value=''; ui.toast('success','Ответ добавлен') }
</script>
