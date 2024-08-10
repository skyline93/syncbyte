// src/App.vue
<template>
  <div id="app">
    <h1>相册</h1>
    <PhotoGallery />
  </div>
</template>

<script>
import PhotoGallery from './components/PhotoGallery.vue';

export default {
  name: 'App',
  components: {
    PhotoGallery
  }
};
</script>

<style>
/* 全局样式 */
</style>
