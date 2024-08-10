// src/components/PhotoGallery.vue
<template>
  <div class="gallery">
    <div v-if="loading" class="loading">加载中...</div>
    <div v-else>
      <div v-for="(image, index) in images" :key="index" class="gallery-item">
        <img :src="image.url" :alt="image.alt" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

const images = ref([]);
const loading = ref(true);

const fetchImages = async () => {
  try {
    const response = await axios.get('https://your-api-endpoint.com/images');
    images.value = response.data;
  } catch (error) {
    console.error('Error fetching images:', error);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchImages();
});
</script>

<style scoped>
.gallery {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.gallery-item {
  width: calc(33.33% - 10px);
}

.gallery-item img {
  width: 100%;
  height: auto;
  object-fit: cover;
  border-radius: 8px;
}

.loading {
  font-size: 1.2em;
  text-align: center;
  width: 100%;
}
</style>
