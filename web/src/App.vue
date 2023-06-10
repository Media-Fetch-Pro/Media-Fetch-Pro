<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<template>
  <div class="flex">
    <div class="flex flex-col border w-3/12">
      <router-link to="/">下载</router-link>
      <router-link to="/history">下载历史</router-link>
      <router-link to="/setting">设置</router-link>
    </div>
    <RouterView />
  </div>
</template>

