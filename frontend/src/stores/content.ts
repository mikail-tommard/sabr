import { defineStore } from 'pinia'
import { contentService } from '@/services/contentService'
import type { Answer, Campus, Comment, FeedItem, Post, Question, Tag, User } from '@/types'
import { uid } from '@/utils'

type Tab='all'|'questions'|'posts'|'my-campus'
type Sort='new'|'popular'|'unanswered'|'best-answer'

// ВАЖНО: ключ версионируется. Если вы обновили мок-данные/структуру,
// просто увеличьте версию — это автоматически сбросит старый кеш.
const STORAGE_KEY = 'sabr-content-cache-v2'

type ContentState = {
  users: User[]
  campuses: Campus[]
  tags: Tag[]
  questions: Question[]
  answers: Answer[]
  posts: Post[]
  comments: Comment[]
  loading: boolean
  bootstrapped: boolean
  feedTab: Tab
  search: string
  campusFilter: string
  tagFilter: string
  typeFilter: ''|'question'|'post'
  sort: Sort
}

export const useContentStore = defineStore('content', {
  state: (): ContentState => ({
    users: [], campuses: [], tags: [], questions: [], answers: [], posts: [], comments: [],
    loading: false, bootstrapped: false,
    feedTab: 'all', search: '', campusFilter: '', tagFilter: '', typeFilter: '', sort: 'new'
  }),
  getters: {
    meCampus(state){ return state.campusFilter || state.campuses[0]?.id || '' },
    usersById: s => Object.fromEntries(s.users.map(u => [u.id, u])) as Record<string, User>,
    campusesById: s => Object.fromEntries(s.campuses.map(c => [c.id, c])) as Record<string, Campus>,
    tagsById: s => Object.fromEntries(s.tags.map(t => [t.id, t])) as Record<string, Tag>,
    feedItems(state): FeedItem[] {
      const tagLookup = Object.fromEntries(state.tags.map(t => [t.id, t.label.toLowerCase()])) as Record<string,string>
      // Скрытый/удалённый контент не показываем в публичной ленте.
      // Админка сможет показывать всё через отдельные таблицы.
      let items: FeedItem[] = [...state.questions, ...state.posts]
        .filter((i:any) => !i.status || i.status === 'active')

      if (state.feedTab === 'questions') items = items.filter(i => i.type === 'question')
      if (state.feedTab === 'posts') items = items.filter(i => i.type === 'post')
      if (state.feedTab === 'my-campus' && state.campusFilter) items = items.filter(i => i.campusId === state.campusFilter)
      if (state.typeFilter) items = items.filter(i => i.type === state.typeFilter)
      if (state.campusFilter) items = items.filter(i => i.campusId === state.campusFilter)
      if (state.tagFilter) items = items.filter(i => i.tagIds.includes(state.tagFilter))

      const query = state.search.trim().toLowerCase()
      if (query) {
        items = items.filter(i =>
          i.title.toLowerCase().includes(query) ||
          i.summary.toLowerCase().includes(query) ||
          i.tagIds.some(tid => (tagLookup[tid] || '').includes(query))
        )
      }

      const score = (i: FeedItem) => i.type === 'question' ? i.upvotes : i.likes
      items.sort((a, b) => {
        // pinned/featured выше всего
        const pinA = (a.type === 'question' && (a as any).pinned) ? 1 : 0
        const pinB = (b.type === 'question' && (b as any).pinned) ? 1 : 0
        if (pinA !== pinB) return pinB - pinA
        const featA = (a.type === 'post' && (a as any).featured) ? 1 : 0
        const featB = (b.type === 'post' && (b as any).featured) ? 1 : 0
        if (featA !== featB) return featB - featA

        if (state.sort === 'new') return +new Date(b.createdAt) - +new Date(a.createdAt)
        if (state.sort === 'popular') return score(b) - score(a)
        if (state.sort === 'unanswered') return ((b.type === 'question' && b.answersCount === 0) ? 1 : 0) - ((a.type === 'question' && a.answersCount === 0) ? 1 : 0)
        return ((b.type === 'question' && !!b.bestAnswerId) ? 1 : 0) - ((a.type === 'question' && !!a.bestAnswerId) ? 1 : 0)
      })
      return items
    },
    topTags: s => [...s.tags].sort((a,b)=>b.usageCount-a.usageCount).slice(0,8),
    topUsers: s => [...s.users].sort((a,b)=>b.rating-a.rating).slice(0,6),
    topQuestions: s => [...s.questions].sort((a,b)=>b.upvotes-a.upvotes).slice(0,5)
  },
  actions: {
    persist() {
      const payload = {
        users: this.users, campuses: this.campuses, tags: this.tags, questions: this.questions,
        answers: this.answers, posts: this.posts, comments: this.comments
      }
      localStorage.setItem(STORAGE_KEY, JSON.stringify(payload))
    },
    hydrateFromCache() {
      try {
        const raw = localStorage.getItem(STORAGE_KEY)
        if (!raw) return false
        const parsed = JSON.parse(raw)
        Object.assign(this, parsed)

        // Санити-чек: если в кеше нет админ-аккаунта, вероятно это старый кеш.
        // В таком случае — игнорируем кеш и загрузим актуальные моки.
        const hasAdmin = Array.isArray(this.users) && this.users.some((u:any) => u?.role === 'Admin')
        if (!hasAdmin) return false

        this.bootstrapped = true
        if (!this.campusFilter) this.campusFilter = this.campuses[0]?.id || ''
        return true
      } catch {
        return false
      }
    },
    async bootstrap() {
      if (this.bootstrapped) return
      if (this.hydrateFromCache()) return
      this.loading = true
      const d = await contentService.getBootstrapData()
      Object.assign(this, d)
      this.campusFilter = this.campuses[0]?.id || ''
      this.bootstrapped = true
      this.loading = false
      this.persist()
    },
    user(id:string){ return this.usersById[id] },
    campus(id:string){ return this.campusesById[id] },
    tag(id:string){ return this.tagsById[id] },
    qById(id:string){ return this.questions.find(x=>x.id===id) },
    pById(id:string){ return this.posts.find(x=>x.id===id) },
    answersFor(qid:string){ return this.answers.filter(a=>a.questionId===qid) },
    commentsFor(type:'question'|'answer'|'post', pid:string){ return this.comments.filter(c=>c.parentType===type && c.parentId===pid) },
    voteQuestion(id:string){ const q=this.qById(id); if(q){ q.upvotes++; this.persist() } },
    voteAnswer(id:string){ const a=this.answers.find(a=>a.id===id); if(a){ a.upvotes++; this.persist() } },
    likePost(id:string){ const p=this.pById(id); if(p){ p.likes++; this.persist() } },
    createQuestion(payload:{title:string;description:string;code?:string;tried?:string;tagIds:string[];campusId:string;authorId:string}){
      const q:Question={id:uid('q'),type:'question',title:payload.title,summary:payload.description.slice(0,140),description:payload.description,code:payload.code,tried:payload.tried,authorId:payload.authorId,campusId:payload.campusId,tagIds:payload.tagIds,createdAt:new Date().toISOString(),upvotes:0,answersCount:0,commentsCount:0,solved:false, status:'active'}
      this.questions.unshift(q)
      this.persist()
      return q
    },
    createPost(payload:{title:string;summary:string;content:string;images?:string[];tagIds:string[];campusId:string;authorId:string}){
      const p:Post={id:uid('p'),type:'post',title:payload.title,summary:payload.summary,content:payload.content,images:payload.images || [],authorId:payload.authorId,campusId:payload.campusId,tagIds:payload.tagIds,createdAt:new Date().toISOString(),likes:0,commentsCount:0, status:'active'}
      this.posts.unshift(p)
      this.persist()
      return p
    },
    addAnswer(questionId:string, authorId:string, content:string){
      const a:Answer={id:uid('a'),questionId,authorId,content,createdAt:new Date().toISOString(),upvotes:0,isBest:false}
      this.answers.unshift(a)
      const q=this.qById(questionId); if(q) q.answersCount++
      this.persist()
      return a
    },
    addComment(parentType:'question'|'answer'|'post', parentId:string, authorId:string, content:string){
      this.comments.unshift({id:uid('cm'),parentType,parentId,authorId,content,createdAt:new Date().toISOString()})
      const q=this.qById(parentId); const p=this.pById(parentId)
      if(q) q.commentsCount++
      if(p) p.commentsCount++
      this.persist()
    },
    markBest(questionId:string, answerId:string){
      this.answers.filter(a=>a.questionId===questionId).forEach(a=>a.isBest=(a.id===answerId))
      const q=this.qById(questionId)
      if(q){ q.bestAnswerId=answerId; q.solved=true }
      this.persist()
    }
  }
})
