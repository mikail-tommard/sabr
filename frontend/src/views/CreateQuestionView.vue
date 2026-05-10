<template>
  <div class="card" style="padding:16px;margin-top:16px;max-width:860px">
    <h1 style="margin-top:0">Создание вопроса</h1>
    <div class="stack">
      <div><div class="muted">Заголовок</div><input v-model="f.title" class="input" /></div>
      <div><div class="muted">Описание проблемы</div><textarea v-model="f.description" class="textarea"></textarea></div>
      <div><div class="muted">Код</div><textarea v-model="f.code" class="textarea" style="min-height:100px"></textarea></div>
      <div><div class="muted">Что пробовал</div><textarea v-model="f.tried" class="textarea" style="min-height:90px"></textarea></div>
      <div class="row" style="flex-wrap:wrap">
        <div style="flex:1;min-width:220px"><div class="muted">Теги (через запятую)</div><input v-model="f.tags" class="input" placeholder="vue, typescript" /></div>
        <div style="width:260px"><div class="muted">Кампус</div><select v-model="f.campusId" class="select"><option v-for="c in s.campuses" :key="c.id" :value="c.id">{{ c.name }}</option></select></div>
      </div>
      <div class="row" style="justify-content:flex-end"><button class="btn primary" @click="submit">Опубликовать вопрос</button></div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { reactive } from 'vue'; import { useRouter } from 'vue-router'; import { useContentStore } from '@/stores/content'; import { useAuthStore } from '@/stores/auth'; import { useUiStore } from '@/stores/ui'
const s=useContentStore(); const auth=useAuthStore(); const ui=useUiStore(); const router=useRouter();
const f = reactive({ title:'', description:'', code:'', tried:'', tags:'vue, typescript', campusId: s.campuses[0]?.id || 'c1' })
const submit = ()=>{
  if(f.title.trim().length<10){ ui.toast('error','Заголовок должен быть не короче 10 символов'); return }
  if(f.description.trim().length<20){ ui.toast('error','Описание слишком короткое'); return }
  const labels=f.tags.split(',').map(x=>x.trim()).filter(Boolean)
  const tagIds = labels.map(l=> s.tags.find(t=>t.label===l)?.id).filter(Boolean) as string[]
  const q=s.createQuestion({ title:f.title.trim(), description:f.description.trim(), code:f.code.trim(), tried:f.tried.trim(), tagIds: tagIds.length?tagIds:[s.tags[0].id], campusId:f.campusId, authorId: auth.currentUserId || s.users[0].id })
  ui.toast('success','Вопрос создан'); router.push(`/questions/${q.id}`)
}
</script>
