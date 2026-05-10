<template><RouterLink :to="`/profile/${user.id}`" class="row"><span>{{ user.avatar }}</span><div><div class="small">{{ user.name }}</div><div class="muted small">@{{ user.username }} · {{ user.role }}</div></div></RouterLink></template>
<script setup lang="ts">import type { User } from '@/types'; defineProps<{user:User}>()</script>
