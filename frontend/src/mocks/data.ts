import type { Answer, Campus, Comment, NotificationItem, Post, Question, Tag, User, Report } from '@/types'
const now = Date.now(); const d=(n:number)=> new Date(now-n*86400000).toISOString(); const h=(n:number)=> new Date(now-n*3600000).toISOString();

export const campuses: Campus[] = [
  {id:'c1',name:'Sabr Moscow',city:'Москва',membersCount:1420,postsCount:210,questionsCount:580,description:'Frontend + backend + алгоритмы'},
  {id:'c2',name:'Sabr Kazan',city:'Казань',membersCount:980,postsCount:160,questionsCount:430,description:'Mobile и backend community'},
  {id:'c3',name:'Sabr SPB',city:'Санкт-Петербург',membersCount:1260,postsCount:200,questionsCount:510,description:'DevOps, data и web'},
  {id:'c4',name:'Sabr Innopolis',city:'Иннополис',membersCount:1100,postsCount:250,questionsCount:470,description:'Продуктовая и системная разработка'},
  {id:'c5',name:'Sabr Novosibirsk',city:'Новосибирск',membersCount:730,postsCount:120,questionsCount:300,description:'Алгоритмы и C++'},
]
export const users: User[] = [
  // Админ-аккаунт для MVP админки (mock)
  {id:'u0',name:'Админ Sabr',username:'admin',avatar:'🛡️',campusId:'c1',role:'Admin',rating:9999,badges:['Admin']},
  {id:'u1',name:'Микаил С.',username:'mika_front',avatar:'🧑‍💻',campusId:'c1',role:'Student',rating:1420,badges:['TS','Vue']},
  {id:'u2',name:'Аян Т.',username:'ayan_backend',avatar:'👨‍🔧',campusId:'c4',role:'Mentor',rating:3810,badges:['Backend','Top Answer']},
  {id:'u3',name:'Лена К.',username:'lena_ui',avatar:'👩‍🎨',campusId:'c3',role:'Student',rating:1985,badges:['UI']},
  {id:'u4',name:'Амир Н.',username:'amir_cpp',avatar:'👨‍💻',campusId:'c5',role:'Alumni',rating:2750,badges:['Algorithms']},
  {id:'u5',name:'Саша Р.',username:'sasha_devops',avatar:'🧑‍🔧',campusId:'c3',role:'Mentor',rating:3560,badges:['DevOps']},
  {id:'u6',name:'Дина В.',username:'dina_mobile',avatar:'👩‍💻',campusId:'c2',role:'Student',rating:1210,badges:['Mobile']},
  {id:'u7',name:'Вера Л.',username:'vera_sql',avatar:'🗃️',campusId:'c1',role:'Mentor',rating:3120,badges:['SQL']},
  {id:'u8',name:'Рома П.',username:'roma_js',avatar:'⚙️',campusId:'c2',role:'Student',rating:870,badges:['JS']}
]
export const tags: Tag[] = ['vue','typescript','javascript','react','nodejs','postgresql','docker','algorithms','css','tailwind','git','api'].map((t,i)=>({id:'t'+(i+1),label:t,usageCount:50+i*7}))
export const questions: Question[] = [
  {id:'q1',type:'question',title:'Почему computed в Vue 3 не обновляется после изменения массива?',summary:'Меняю элемент по индексу и computed ведёт себя странно.',description:'Есть reactive список. После обновления элемента по индексу UI местами не перерисовывается. В каких случаях лучше ref([]) вместо reactive([])?',code:`const state = reactive({ items: [{ done: false }] })\nstate.items[0].done = true`,tried:'Пробовал spread и nextTick.',authorId:'u1',campusId:'c1',tagIds:['t1','t2','t3'],createdAt:d(1),upvotes:12,answersCount:2,commentsCount:2,solved:true,bestAnswerId:'a1', pinned:true},
  {id:'q2',type:'question',title:'Как правильно проектировать REST endpoint для batch update?',summary:'Нужно обновлять 50+ сущностей и обрабатывать частичные ошибки.',description:'Как лучше вернуть результат: per-item result или отдельные endpoint? Что делать с частичной ошибкой?',authorId:'u2',campusId:'c4',tagIds:['t12','t5'],createdAt:d(2),upvotes:18,answersCount:2,commentsCount:1,solved:false},
  {id:'q3',type:'question',title:'PostgreSQL: индекс не используется в запросе с ILIKE',summary:'Planner идёт в seq scan. Как ускорить поиск по title?',description:'Нужен поиск по title через ILIKE %term%. Какие варианты: trigram index, full-text?',authorId:'u7',campusId:'c1',tagIds:['t6','t12'],createdAt:d(3),upvotes:21,answersCount:2,commentsCount:2,solved:true,bestAnswerId:'a4'},
  {id:'q4',type:'question',title:'Tailwind: как организовать токены для light/dark темы?',summary:'Не хочу дублировать dark:* классы в каждом компоненте.',description:'Как удобнее хранить semantic tokens?',authorId:'u3',campusId:'c3',tagIds:['t10','t9','t1'],createdAt:d(4),upvotes:16,answersCount:1,commentsCount:0,solved:false},
  {id:'q5',type:'question',title:'Docker compose для monorepo: как подружить hot reload?',summary:'HMR нестабилен, особенно на Windows.',description:'Frontend + backend + db в docker compose. Иногда Vite HMR не видит изменения.',authorId:'u5',campusId:'c3',tagIds:['t7','t5'],createdAt:d(5),upvotes:14,answersCount:1,commentsCount:1,solved:true,bestAnswerId:'a6'},
  {id:'q6',type:'question',title:'Как объяснить debounce vs throttle на собеседовании?',summary:'Нужна короткая и понятная формулировка с примерами.',description:'Ищу чёткое объяснение и реальные кейсы для UI.',authorId:'u8',campusId:'c2',tagIds:['t2','t3'],createdAt:d(1),upvotes:9,answersCount:2,commentsCount:0,solved:true,bestAnswerId:'a7'},
  {id:'q7',type:'question',title:'Git rebase vs merge в командной работе',summary:'Какой workflow выбрать для студенческой команды?',description:'Нужны практические правила и компромисс.',authorId:'u6',campusId:'c2',tagIds:['t11'],createdAt:d(6),upvotes:11,answersCount:1,commentsCount:0,solved:false},
  {id:'q8',type:'question',title:'Почему useEffect в React срабатывает дважды в dev?',summary:'StrictMode пугает новичков.',description:'Как корректно объяснить двойной вызов эффектов в dev?',authorId:'u3',campusId:'c3',tagIds:['t4','t3'],createdAt:d(2),upvotes:17,answersCount:1,commentsCount:1,solved:true,bestAnswerId:'a9'},
  {id:'q9',type:'question',title:'DFS или Union-Find для поиска цикла в графе?',summary:'Хочу не путаться на собеседованиях.',description:'Когда что выбирать в зависимости от типа графа?',authorId:'u4',campusId:'c5',tagIds:['t8'],createdAt:d(7),upvotes:24,answersCount:1,commentsCount:1,solved:false},
  {id:'q10',type:'question',title:'Vitest: как тестировать composables с localStorage?',summary:'Нужны изолированные тесты useTheme.',description:'Как мокать localStorage и чистить состояние между тестами?',authorId:'u1',campusId:'c1',tagIds:['t1','t2'],createdAt:d(3),upvotes:8,answersCount:1,commentsCount:0,solved:false},
]
export const answers: Answer[] = [
  {id:'a1',questionId:'q1',authorId:'u2',content:'Чаще всего проблема не в computed, а в том, как читается реактивная зависимость. Для списка ref([]) обычно проще.',createdAt:d(1),upvotes:14,isBest:true},
  {id:'a2',questionId:'q1',authorId:'u3',content:'Проверь деструктуризацию и промежуточные переменные. Иногда полезно делать map и возвращать новый массив.',createdAt:d(1),upvotes:7,isBest:false},
  {id:'a3',questionId:'q2',authorId:'u2',content:'Обычно 200 + массив результатов по id. Это проще для клиента.',createdAt:d(2),upvotes:11,isBest:false},
  {id:'a4',questionId:'q3',authorId:'u7',content:'Смотри pg_trgm + GIN/GiST индекс. Для ILIKE %term% это частый путь.',createdAt:d(2),upvotes:19,isBest:true},
  {id:'a5',questionId:'q4',authorId:'u3',content:'CSS variables + semantic tokens, а Tailwind только как utilities.',createdAt:d(4),upvotes:12,isBest:false},
  {id:'a6',questionId:'q5',authorId:'u5',content:'Проверь CHOKIDAR_USEPOLLING=true и volume mounts.',createdAt:d(5),upvotes:16,isBest:true},
  {id:'a7',questionId:'q6',authorId:'u2',content:'Debounce ждёт паузу (поиск), throttle ограничивает частоту (scroll).',createdAt:d(1),upvotes:15,isBest:true},
  {id:'a8',questionId:'q7',authorId:'u5',content:'Feature branches + PR + squash merge — хороший базовый вариант.',createdAt:d(6),upvotes:12,isBest:false},
  {id:'a9',questionId:'q8',authorId:'u4',content:'StrictMode в dev повторно запускает для выявления side effects.',createdAt:d(2),upvotes:18,isBest:true},
  {id:'a10',questionId:'q9',authorId:'u4',content:'Для неориентированного графа при добавлении рёбер удобен Union-Find; иначе DFS.',createdAt:d(7),upvotes:20,isBest:false},
  {id:'a11',questionId:'q10',authorId:'u1',content:'Мокай localStorage и сбрасывай state перед каждым тестом.',createdAt:d(3),upvotes:6,isBest:false}
]
export const posts: Post[] = [
  {id:'p1',type:'post',title:'Как мы организовали code review в студенческой команде',summary:'Простые правила, которые реально работают и не тормозят релизы.',content:'## Code review\n\n1. Маленькие PR\n2. Ясные описания\n3. Checklists\n\n```ts\nconsole.log("review")\n```',authorId:'u5',campusId:'c3',tagIds:['t11','t5'],createdAt:d(1),likes:34,commentsCount:2},
  {id:'p2',type:'post',title:'Vue 3 + Pinia: структура stores для MVP',summary:'Как не перегрузить архитектуру и оставить путь для роста.',content:'## Pinia в MVP\n\nРазделите auth/ui/content stores.\n\n```ts\nexport const useContentStore = defineStore(...)\n```',authorId:'u1',campusId:'c1',tagIds:['t1','t2'],createdAt:d(2),likes:28,commentsCount:2},
  {id:'p3',type:'post',title:'Минималистичный UI для dev-community',summary:'Почему спокойная палитра и типографика выигрывают у яркости.',content:'## UI\n\nНейтральные цвета + 1 акцент.\nМного воздуха, меньше шума.',authorId:'u3',campusId:'c3',tagIds:['t9','t10'],createdAt:d(3),likes:41,commentsCount:2, featured:true},
  {id:'p4',type:'post',title:'PostgreSQL индексы: быстрый практический гид',summary:'B-Tree, GIN, partial indexes и типичные ошибки.',content:'## Индексы\n\nСмотрите EXPLAIN ANALYZE перед оптимизацией.',authorId:'u7',campusId:'c1',tagIds:['t6'],createdAt:d(4),likes:39,commentsCount:1},
  {id:'p5',type:'post',title:'Docker для локальной разработки',summary:'Стабильная конфигурация frontend/backend/db в compose.',content:'## Docker\n\nСледите за volumes и watchers.',authorId:'u5',campusId:'c3',tagIds:['t7','t5'],createdAt:d(5),likes:22,commentsCount:1},
  {id:'p6',type:'post',title:'Как писать хорошие вопросы в тех-комьюнити',summary:'Шаблон вопроса, который ускоряет получение полезных ответов.',content:'## Хороший вопрос\n\nКонтекст, ожидаемый результат, что пробовал.',authorId:'u2',campusId:'c4',tagIds:['t12'],createdAt:d(6),likes:45,commentsCount:3},
  {id:'p7',type:'post',title:'Tailwind + CSS variables: light/dark без шума',summary:'Практика для чистых тем и понятных токенов.',content:'## Темы\n\nИспользуйте semantic variables.',authorId:'u3',campusId:'c3',tagIds:['t10','t9'],createdAt:d(7),likes:26,commentsCount:1},
  {id:'p8',type:'post',title:'От pet-project к портфолио',summary:'Что показать наставнику: README, демо, архитектуру.',content:'## Портфолио\n\nПокажите структуру, demo и reasoning.',authorId:'u2',campusId:'c4',tagIds:['t11','t12'],createdAt:d(8),likes:52,commentsCount:2}
]

// Моковые жалобы/репорты — для админки
export const reports: Report[] = [
  { id:'r1', createdAt:h(18), reporterId:'u6', reason:'spam', targetType:'comment', targetId:'cm4', note:'Похоже на накрутку/спам', status:'open' },
  { id:'r2', createdAt:h(30), reporterId:'u3', reason:'abuse', targetType:'user', targetId:'u8', note:'Грубость в комментариях', status:'in_review' },
  { id:'r3', createdAt:h(44), reporterId:'u1', reason:'other', targetType:'post', targetId:'p6', note:'Проверить корректность ссылок/формулировок', status:'resolved' },
]
export const comments: Comment[] = [
  {id:'cm1',parentType:'question',parentId:'q1',authorId:'u3',content:'Покажи, как объявлен computed?',createdAt:h(12)},
  {id:'cm2',parentType:'answer',parentId:'a1',authorId:'u1',content:'С ref([]) стало понятнее, спасибо!',createdAt:h(11)},
  {id:'cm3',parentType:'post',parentId:'p2',authorId:'u5',content:'Нравится разделение stores на content/ui/auth.',createdAt:h(20)},
  {id:'cm4',parentType:'post',parentId:'p3',authorId:'u1',content:'Отличный разбор палитры и отступов.',createdAt:h(30)}
]
export const notifications: NotificationItem[] = [
  {id:'n1',userId:'u1',text:'На ваш вопрос ответили',isRead:false,href:'/questions/q1'},
  {id:'n2',userId:'u1',text:'Ваш ответ получил апвоут',isRead:false,href:'/questions/q6'},
  {id:'n3',userId:'u1',text:'Комментарий к публикации',isRead:true,href:'/posts/p2'}
]
