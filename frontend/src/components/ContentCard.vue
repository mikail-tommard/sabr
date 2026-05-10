<template>
  <div
    class="card feed-card"
    v-motion
    :initial="{ opacity: 0, y: 8 }"
    :enter="{ opacity: 1, y: 0 }"
    :hovered="{ y: -2 }"
    :transition="{ type: 'spring', stiffness: 260, damping: 22 }"
  >
    <VotePanel :value="score" :label="item.type==='question'?'апвоуты':'лайки'" @upvote="handleVote" />
    <div class="content-grow stack" style="gap:8px">
      <RouterLink :to="to" class="title" style="font-size:16px;display:block">{{ item.title }}</RouterLink>
      <div class="feed-summary">{{ item.summary }}</div>
      <div class="row" style="flex-wrap:wrap">
        <TagChip v-for="tid in item.tagIds" :key="tid" :label="tagLabel(tid)" />
      </div>
      <div class="row" style="justify-content:space-between;flex-wrap:wrap">
        <div class="row muted">
          <UserBadge :user="author" />
          <CampusBadge :name="campusName" />
          <span>{{ timeAgo(item.createdAt) }}</span>
        </div>
        <div class="row muted">
          <span v-if="item.type==='question'">💬 {{ item.answersCount }} ответов</span>
          <span v-else>💬 {{ item.commentsCount }} комм.</span>
          <span v-if="item.type==='question'">{{ item.solved ? '✅ Решено' : '⏳ Не решено' }}</span>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { computed } from 'vue'
import type { FeedItem } from '@/types'
import { useContentStore } from '@/stores/content'
import TagChip from './TagChip.vue'; import VotePanel from './VotePanel.vue'; import UserBadge from './UserBadge.vue'; import CampusBadge from './CampusBadge.vue'
import { timeAgo } from '@/utils'
const props = defineProps<{item:FeedItem}>(); const s=useContentStore()
const to = computed(()=> props.item.type==='question'?`/questions/${props.item.id}`:`/posts/${props.item.id}`)
const author = computed(()=> s.user(props.item.authorId)!)
const campusName = computed(()=> s.campus(props.item.campusId)?.name || 'Кампус')
const score = computed(()=> props.item.type==='question'?props.item.upvotes:props.item.likes)
const tagLabel = (id:string)=> s.tag(id)?.label || id
const handleVote = ()=> props.item.type==='question' ? s.voteQuestion(props.item.id) : s.likePost(props.item.id)
</script>
