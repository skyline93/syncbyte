// src/components/PhotoGallery.vue
<template>
  <div class="gallery">
    <div v-if="loading" class="loading">加载中...</div>
    <div v-else>
      <div v-for="(image, index) in images" :key="index" class="gallery-item">
        <img :src="image.link" :alt="image.file_name" />
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
  const token = localStorage.getItem('authToken');

  try {
    const response = await axios.get('http://192.168.209.130:8000/api/v1/photo?album_id=2', {
      headers: {
        Authorization: `Bearer ${token}` // 将 token 添加到请求头
      }
    });
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
