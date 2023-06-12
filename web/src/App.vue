<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavTab from './components/NavTab.vue'
</script>

<template>
  <div class="flex">
    <div class="flex flex-col border w-1/12 gap-5">
      <NavTab to="/">Download</NavTab>
      <NavTab to="/history">History</NavTab>
      <NavTab to="/setting">Setting</NavTab>
    </div>
    <RouterView />
  </div>
</template>

