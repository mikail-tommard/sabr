<template>
  <section class="container" style="padding: 18px 0">
    <div class="grid" style="grid-template-columns: 1.2fr .8fr; gap: 14px">
      <div class="card" style="padding: 18px">
        <div class="title" style="font-size: 22px">Sabr — dev‑community платформа для кампусов</div>
        <div class="muted" style="margin-top: 8px; max-width: 70ch">
          Sabr — это MVP прототип комьюнити‑платформы для студентов и разработчиков из разных кампусов.
          Формат — смесь Habr + Stack Overflow + campus community: вопросы, ответы, публикации, профили, теги и кампусы.
        </div>

        <div class="line" style="margin: 14px 0"></div>

        <div class="title">Что умеет MVP</div>
        <div class="stack" style="margin-top: 10px">
          <div class="card" style="padding: 12px; box-shadow:none">
            <div class="row" style="justify-content: space-between; gap: 10px">
              <b>Вопросы и ответы</b>
              <span class="pill">Q&A</span>
            </div>
            <div class="muted" style="margin-top: 6px">
              Создание вопросов, ответы, комментарии, апвоуты, метка «лучший ответ», статус «решено/не решено».
            </div>
          </div>

          <div class="card" style="padding: 12px; box-shadow:none">
            <div class="row" style="justify-content: space-between; gap: 10px">
              <b>Публикации</b>
              <span class="pill">Posts</span>
            </div>
            <div class="muted" style="margin-top: 6px">
              Посты с тегами и кампусом, лайки, комментарии, похожие публикации, загрузка изображений (локально).
            </div>
          </div>

          <div class="card" style="padding: 12px; box-shadow:none">
            <div class="row" style="justify-content: space-between; gap: 10px">
              <b>Поиск и фильтры</b>
              <span class="pill">Search</span>
            </div>
            <div class="muted" style="margin-top: 6px">
              Глобальный поиск по заголовкам и тегам, фильтры по типу/кампусу/тегам, сортировки (новые, популярные и т.д.).
            </div>
          </div>

          <div class="card" style="padding: 12px; box-shadow:none">
            <div class="row" style="justify-content: space-between; gap: 10px">
              <b>Админка</b>
              <span class="pill">Admin</span>
            </div>
            <div class="muted" style="margin-top: 6px">
              Модерация контента, управление пользователями, тегами, кампусами, просмотр жалоб/репортов.
              Всё работает на мок‑данных и локальном состоянии.
            </div>
          </div>
        </div>

        <div class="line" style="margin: 14px 0"></div>

        <div class="title">Технологии</div>
        <div class="row" style="margin-top: 10px; flex-wrap: wrap; gap: 10px">
          <span class="pill">Vue 3</span>
          <span class="pill">TypeScript</span>
          <span class="pill">Vite</span>
          <span class="pill">Vue Router</span>
          <span class="pill">Pinia</span>
          <span class="pill">Headless UI</span>
          <span class="pill">VueUse Motion</span>
        </div>

        <div class="muted" style="margin-top: 10px">
          SEO‑страница нужна для продвижения: здесь собраны ключевые слова и описание продукта.
          В реальном продакшене лучше добавить SSR/пререндер, но для MVP достаточно мета‑тегов и лендинга.
        </div>
      </div>

      <aside class="stack" style="gap: 12px">
        <div class="card" style="padding: 14px">
          <div class="title">Ключевые слова (SEO)</div>
          <div class="muted" style="margin-top: 8px">
            комьюнити для программистов, вопросы и ответы, dev community, кампусы, студенты, публикации, теги, обмен опытом
          </div>
        </div>

        <div class="card" style="padding: 14px">
          <div class="title">Как подключим backend</div>
          <div class="muted" style="margin-top: 8px">
            Слой сервисов уже выделен. Дальше заменяем fakeDelay на реальные запросы:
            <span class="kbd">/services/contentService</span>, <span class="kbd">/services/adminService</span>, <span class="kbd">/services/auth</span>.
          </div>
        </div>

        <RouterLink to="/" class="btn primary" style="justify-content:center">Перейти в ленту</RouterLink>
      </aside>
    </div>
  </section>
</template>
