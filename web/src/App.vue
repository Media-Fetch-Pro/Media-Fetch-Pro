<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<template>
  <div class="flex">
    <div class="flex flex-col border w-1/12">
      <router-link to="/">Download</router-link>
      <router-link to="/history">History</router-link>
      <router-link to="/setting">Setting</router-link>
    </div>
    <RouterView />
  </div>
</template>

