<template>
  <div class="page-grid" style="margin-top:16px; grid-template-columns: 1.25fr .75fr; align-items:start">
    <div class="card" style="padding:18px">
      <div class="row" style="justify-content:space-between; gap:12px; align-items:flex-start">
        <div>
          <h1 class="h1" style="margin:0">Создание публикации</h1>
          <div class="muted" style="margin-top:6px">Пишите аккуратно: структура, примеры, выводы. Можно добавить изображения.</div>
        </div>
        <button class="btn" @click="togglePreview">
          {{ showPreview ? 'Скрыть превью' : 'Показать превью' }}
        </button>
      </div>

      <div class="stack" style="margin-top:14px">
        <div>
          <div class="muted">Заголовок</div>
          <input v-model="f.title" class="input" placeholder="Например: Vue 3 + Pinia: удобная архитектура stores" />
        </div>

        <div>
          <div class="muted">Краткое описание</div>
          <textarea v-model="f.summary" class="textarea" style="min-height:84px" placeholder="2–3 предложения о сути публикации"></textarea>
        </div>

        <div>
          <div class="row" style="justify-content:space-between; align-items:center">
            <div class="muted">Содержимое (pseudo‑markdown)</div>
            <span class="kbd">```код```</span>
          </div>
          <textarea
            v-model="f.content"
            class="textarea mono"
            style="min-height:260px"
            placeholder="# Заголовок\n\nТекст...\n\n```ts\nconsole.log('hello')\n```"
          ></textarea>
        </div>

        <div class="card" style="padding:12px; background: color-mix(in srgb, var(--surface) 88%, var(--bg)); box-shadow:none">
          <div class="row" style="justify-content:space-between; gap:10px; align-items:center">
            <div>
              <div class="title">Изображения</div>
              <div class="muted" style="margin-top:2px">Локальная загрузка для MVP. Потом подключим upload через API.</div>
            </div>
            <label class="btn ghost-accent" style="cursor:pointer">
              <input type="file" accept="image/*" multiple style="display:none" @change="onFiles" />
              Добавить фото
            </label>
          </div>

          <div v-if="f.images.length" class="img-grid" style="margin-top:10px">
            <div v-for="(src, idx) in f.images" :key="src" class="img-tile">
              <img :src="src" alt="" />
              <button class="img-remove" title="Удалить" @click="removeImage(idx)">×</button>
            </div>
          </div>
          <div v-else class="empty" style="margin-top:10px">Пока нет изображений</div>

          <div class="muted" style="margin-top:10px">
            Подключение backend позже: <span class="kbd">POST /uploads</span> → вернуть URL и хранить в поле <span class="kbd">images[]</span>.
          </div>
        </div>

        <div class="row" style="flex-wrap:wrap; gap:10px">
          <div style="flex:1; min-width:220px">
            <div class="muted">Теги (через запятую)</div>
            <input v-model="f.tags" class="input" placeholder="vue, typescript, docker" />
          </div>
          <div style="width:280px; min-width:220px">
            <div class="muted">Кампус</div>
            <select v-model="f.campusId" class="select">
              <option v-for="c in s.campuses" :key="c.id" :value="c.id">{{ c.name }}</option>
            </select>
          </div>
        </div>

        <div class="row" style="justify-content:flex-end; gap:10px">
          <RouterLink to="/" class="btn">Отмена</RouterLink>
          <button class="btn primary" @click="submit">Опубликовать</button>
        </div>
      </div>
    </div>

    <div v-if="showPreview" class="card" style="padding:16px; position:sticky; top:92px">
      <div class="title">Превью</div>
      <h2 style="margin:10px 0 4px">{{ f.title || 'Заголовок публикации' }}</h2>
      <div class="muted">{{ f.summary || 'Краткое описание…' }}</div>

      <div v-if="f.images.length" class="img-grid" style="margin-top:12px">
        <div v-for="src in f.images.slice(0,4)" :key="src" class="img-tile" style="aspect-ratio: 4/3">
          <img :src="src" alt="" />
        </div>
      </div>

      <pre class="code" style="margin-top:12px"><code>{{ f.content }}</code></pre>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useContentStore } from '@/stores/content'
import { useAuthStore } from '@/stores/auth'
import { useUiStore } from '@/stores/ui'

const s = useContentStore()
const auth = useAuthStore()
const ui = useUiStore()
const router = useRouter()

const showPreview = ref(true)
const togglePreview = () => (showPreview.value = !showPreview.value)

const f = reactive({
  title: '',
  summary: '',
  content: '## Новый пост\n\nОпишите идею, добавьте примеры и выводы…\n\n```ts\nconsole.log("Sabr")\n```',
  tags: 'vue, typescript',
  campusId: s.campuses[0]?.id || 'c1',
  images: [] as string[],
})

function onFiles(e: Event) {
  const input = e.target as HTMLInputElement
  const files = Array.from(input.files || [])
  if (!files.length) return

  // MVP: локальные ObjectURL для превью.
  // Backend позже: загружать файлы в API и сохранять URL (см. комментарий в шаблоне).
  const urls = files.map((file) => URL.createObjectURL(file))
  f.images = [...f.images, ...urls].slice(0, 12)
  input.value = ''
}

function removeImage(idx: number) {
  const url = f.images[idx]
  try { URL.revokeObjectURL(url) } catch {}
  f.images = f.images.filter((_, i) => i !== idx)
}

const submit = () => {
  if (!auth.currentUserId) {
    ui.toast('error', 'Сначала войдите в аккаунт (mock login)')
    router.push('/login')
    return
  }
  if (f.title.trim().length < 8) {
    ui.toast('error', 'Заголовок слишком короткий')
    return
  }
  if (f.content.trim().length < 30) {
    ui.toast('error', 'Содержимое слишком короткое')
    return
  }

  const labels = f.tags.split(',').map(x => x.trim()).filter(Boolean)
  const tagIds = labels
    .map(l => s.tags.find(t => t.label.toLowerCase() === l.toLowerCase())?.id)
    .filter(Boolean) as string[]

  const p = s.createPost({
    title: f.title.trim(),
    summary: f.summary.trim() || f.content.slice(0, 140),
    content: f.content.trim(),
    images: f.images,
    tagIds: tagIds.length ? tagIds : [s.tags[0].id],
    campusId: f.campusId,
    authorId: auth.currentUserId,
  })

  ui.toast('success', 'Публикация создана')
  router.push(`/posts/${p.id}`)
}
</script>
