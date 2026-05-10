<template>
  <div class="vote">
    <button class="btn" @click="$emit('upvote')" title="Апвоут">▲</button>
    <div class="title">{{ value }}</div>
    <div class="muted small">{{ label }}</div>
  </div>
</template>
<script setup lang="ts">
defineProps<{ value:number; label?:string }>()
defineEmits<{ (e:'upvote'):void }>()
</script>
