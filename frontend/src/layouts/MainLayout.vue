<template>
  <div>
    <AppHeader />
    <div class="container" style="padding-bottom:24px"><RouterView /></div>
    <div class="toast-host">
      <div v-for="t in ui.toasts" :key="t.id" class="toast" :class="t.type">{{ t.text }}</div>
    </div>
  </div>
</template>
<script setup lang="ts">
import AppHeader from '@/components/AppHeader.vue'
import { useUiStore } from '@/stores/ui'
const ui = useUiStore()
</script>
