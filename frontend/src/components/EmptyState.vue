<template><div class="empty"><div style="font-size:28px">{{ icon }}</div><div class="title" style="margin-top:6px">{{ title }}</div><div class="muted" style="margin-top:4px">{{ description }}</div></div></template>
<script setup lang="ts">withDefaults(defineProps<{title:string;description:string;icon?:string}>(),{icon:'📭'})</script>
