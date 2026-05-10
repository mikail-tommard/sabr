<template>
  <div class="stack" style="gap: 12px">
    <div class="title">Настройки</div>
    <div class="muted">В MVP — это UI-заглушка. В будущем здесь будут настройки модерации, лимиты, интеграции, аудит‑лог.</div>

    <div class="grid" style="grid-template-columns: 1fr 1fr; gap: 12px">
      <div class="card" style="padding: 12px; box-shadow:none">
        <div class="title">Политики</div>
        <div class="stack" style="margin-top: 10px">
          <label class="row" style="gap: 10px"><input type="checkbox" checked /> Авто‑фильтр спама</label>
          <label class="row" style="gap: 10px"><input type="checkbox" checked /> Лимит частоты публикаций</label>
          <label class="row" style="gap: 10px"><input type="checkbox" /> Предмодерация для новых аккаунтов</label>
        </div>
      </div>

      <div class="card" style="padding: 12px; box-shadow:none">
        <div class="title">Интеграции</div>
        <div class="muted" style="margin-top: 10px">Sentry / PostHog / Email‑уведомления — подключим на backend.</div>
        <div class="row" style="gap: 10px; margin-top: 10px; flex-wrap: wrap">
          <span class="pill">Sentry</span>
          <span class="pill">PostHog</span>
          <span class="pill">SMTP</span>
        </div>
      </div>
    </div>

    <div class="card" style="padding: 12px; box-shadow:none">
      <div class="title">Где подключать API</div>
      <div class="muted" style="margin-top: 8px">
        1) <span class="kbd">src/services/adminService.ts</span> — reports/users/content moderation
        <br />
        2) <span class="kbd">src/services/contentService.ts</span> — feed/questions/posts
      </div>
    </div>
  </div>
</template>
