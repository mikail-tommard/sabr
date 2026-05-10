<template>
  <div v-if="p" class="page-grid" style="margin-top:16px">
    <div class="stack">
      <div class="card" style="padding:16px">
        <div style="height:180px;border:1px solid var(--border);border-radius:12px;background:linear-gradient(180deg,transparent, color-mix(in srgb, var(--accent) 10%, var(--surface)));display:grid;place-items:center" class="muted">Обложка-заглушка</div>
        <h1 style="margin:12px 0 6px">{{ p.title }}</h1>
        <div class="muted">{{ p.summary }}</div>
        <div class="row" style="justify-content:space-between;margin-top:10px;flex-wrap:wrap"><UserBadge :user="author" /><div class="row"><CampusBadge :name="s.campus(p.campusId)?.name || ''" /><span class="muted">{{ fmtDate(p.createdAt) }}</span></div></div>
        <div class="row" style="margin-top:10px;flex-wrap:wrap"><TagChip v-for="tid in p.tagIds" :key="tid" :label="s.tag(tid)?.label || tid" /></div>
      <div v-if="p.images && p.images.length" class="img-grid" style="margin-top:12px">
        <div v-for="src in p.images" :key="src" class="img-tile" style="aspect-ratio: 4/3">
          <img :src="src" alt="" />
        </div>
      </div>
        <div class="line" style="margin:12px 0"></div>
        <div v-html="html" style="line-height:1.6"></div>
        <div class="row" style="margin-top:12px;justify-content:flex-end"><button class="btn" @click="like">👍 {{ p.likes }}</button></div>
      </div>
      <CommentList :items="comments" :users="s.users" @add="txt=>s.addComment('post', p.id, meId, txt)" />
    </div>
    <aside class="stack">
      <div class="card sidebar-card"><div class="title">Похожие публикации</div><div class="list" style="margin-top:8px"><RouterLink v-for="pp in related" :key="pp.id" class="list-item" :to="`/posts/${pp.id}`">{{ pp.title }}</RouterLink></div></div>
    </aside>
  </div>
  <EmptyState v-else title="Публикация не найдена" description="Проверьте ссылку" />
</template>
<script setup lang="ts">
import { computed } from 'vue'; import { useRoute } from 'vue-router'; import { useContentStore } from '@/stores/content'; import { useAuthStore } from '@/stores/auth'; import { useUiStore } from '@/stores/ui'
import UserBadge from '@/components/UserBadge.vue'; import TagChip from '@/components/TagChip.vue'; import CampusBadge from '@/components/CampusBadge.vue'; import CommentList from '@/components/CommentList.vue'; import EmptyState from '@/components/EmptyState.vue'
import { fmtDate, mdToHtml } from '@/utils'
const route=useRoute(); const s=useContentStore(); const auth=useAuthStore(); const ui=useUiStore();
const p = computed(()=> s.pById(route.params.id as string)); const author = computed(()=> p.value ? s.user(p.value.authorId)! : s.users[0]!); const comments = computed(()=> p.value ? s.commentsFor('post', p.value.id) : []); const meId = computed(()=> auth.currentUserId || s.users[0]?.id || 'u1')
const related = computed(()=> p.value ? s.posts.filter(x=>x.id!==p.value!.id && x.tagIds.some(t=>p.value!.tagIds.includes(t))).slice(0,4) : [])
const html = computed(()=> p.value ? mdToHtml(p.value.content) : '')
const like = ()=> { if(!p.value) return; s.likePost(p.value.id); ui.toast('success','Лайк поставлен') }
</script>
