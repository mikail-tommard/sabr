<template><div style="max-width:620px;margin:40px auto"><EmptyState icon="404" title="Страница не найдена" description="Похоже, ссылка ведёт в никуда" /><div class="row" style="justify-content:center;margin-top:12px"><RouterLink to="/" class="btn primary">На главную</RouterLink></div></div></template>
<script setup lang="ts">import EmptyState from '@/components/EmptyState.vue'</script>
