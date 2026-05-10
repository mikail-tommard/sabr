<template>
  <div class="card" style="padding:12px">
    <div class="row" style="justify-content:space-between"><div class="title">Комментарии</div></div>
    <div v-if="items.length" class="stack" style="margin-top:10px">
      <div v-for="c in items" :key="c.id" class="list-item surface-subtle">
        <div class="row" style="justify-content:space-between">
          <UserBadge :user="userMap[c.authorId]" />
          <span class="muted small">{{ fmtDate(c.createdAt) }}</span>
        </div>
        <div style="margin-top:6px; white-space:pre-wrap">{{ c.content }}</div>
      </div>
    </div>
    <EmptyState v-else title="Нет комментариев" description="Будьте первым" />
    <div class="row" style="margin-top:10px; align-items:flex-start">
      <textarea v-model="text" class="textarea" placeholder="Написать комментарий" style="min-height:80px"></textarea>
      <button class="btn primary" @click="submit" style="align-self:stretch">Отправить</button>
    </div>
  </div>
</template>
<script setup lang="ts">
import { computed, ref } from 'vue'
import type { Comment, User } from '@/types'
import { fmtDate } from '@/utils'
import UserBadge from './UserBadge.vue'; import EmptyState from './EmptyState.vue'
const props = defineProps<{items:Comment[]; users:User[]}>()
const emit = defineEmits<{(e:'add', text:string):void}>()
const text = ref('')
const userMap = computed(()=> Object.fromEntries(props.users.map(u=>[u.id,u])))
const submit = ()=>{ if(!text.value.trim()) return; emit('add', text.value.trim()); text.value=''}
</script>
