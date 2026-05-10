<template>
  <div class="auth-shell">
    <div class="auth-left">
      <div class="auth-brand">
        <div class="logo-s21">S</div>
        <div>
          <div class="brand-name">Sabr</div>
          <div class="brand-sub">Campus dev community • Q&A • Posts</div>
        </div>
      </div>

      <div class="auth-hints">
        <div class="hint">
          <div class="hint-title">Стиль “School 21”</div>
          <div class="muted">Темный фон, строгая сетка, умеренный неон‑акцент, моно‑детали.</div>
        </div>
        <div class="hint">
          <div class="hint-title">MVP без backend</div>
          <div class="muted">Mock login сохраняется в localStorage. Данные — заглушки внутри проекта.</div>
        </div>
      </div>

      <div class="auth-footer muted">
        Совет: для демо выберите тестового пользователя — это “вход” в MVP.
      </div>
    </div>

    <div class="auth-right">
      <div class="auth-card">
        <RouterView />
      </div>
    </div>
  </div>
</template>
