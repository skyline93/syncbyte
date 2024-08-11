<template>
    <div class="gallery">
        <div class="header">
            <h1>Welcome to the Home Page</h1>
            <p v-if="username" class="user-info">{{ username }}</p>
        </div>
        <div v-if="loading" class="loading">加载中...</div>
        <div v-else>
            <div v-for="(image, index) in images" :key="index" class="gallery-item">
                <img :src="imageUrls[index]" :alt="image.file_name" class="responsive-image"/>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axiosInstance from '../services/axiosInstance'; // 引入已配置的 axios 实例

const images = ref([]);
const imageUrls = ref([]);
const loading = ref(true);
const username = ref('');

const fetchUser = async () => {
    username.value = localStorage.getItem('username'); // 设置用户名
};

const fetchImages = async () => {
    try {
        const response = await axiosInstance.get('/api/v1/photo?album_id=1');
        images.value = response.data.data;
        imageUrls.value = await Promise.all(response.data.data.map(async (image) => {
            const url = getFullImageUrl(image.link);
            return await getImageBlob(url);
        }));
    } catch (error) {
        console.error('Error fetching images:', error);
    } finally {
        loading.value = false;
    }
};

const getFullImageUrl = (path) => {
    if (typeof path !== 'string') {
        console.error('Invalid path:', path);
        return ''; // 返回空字符串或默认图片 URL
    }

    if (path.startsWith('http')) {
        return path; // 如果路径已经是完整的 URL，直接返回
    } else {
        try {
            const trimmedPath = path.replace(/^\//, ''); // 去掉路径前的斜杠
            const url = new URL(trimmedPath, 'http://localhost:8000'); // 替换为你的基础 URL
            return url.href;
        } catch (error) {
            console.error('Invalid URL:', path, error);
            return ''; // 返回空字符串或默认图片 URL
        }
    }
};

const getImageBlob = async (url) => {
    try {
        const response = await axiosInstance.get(url, { responseType: 'blob' });
        return URL.createObjectURL(response.data);
    } catch (error) {
        console.error('Error fetching image:', error);
        return ''; // 返回空字符串或默认图片 URL
    }
};

onMounted(() => {
    fetchImages();
    fetchUser();
});
</script>

<style scoped>
.header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
}

.user-info {
    font-weight: bold;
    font-size: 16px;
    position: absolute;
    top: 10px;
    right: 10px;
}

.gallery {
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
}

.gallery-item {
    width: calc(33.33% - 10px);
}

.responsive-image {
    width: 100%;
    height: auto;
    object-fit: cover;
    /* border-radius: 8px; */
}

@media (max-width: 768px) {
  .gallery-item {
    width: calc(50% - 10px); /* 调整宽度为适应小屏幕 */
  }
}

@media (max-width: 480px) {
  .gallery-item {
    width: calc(100% - 10px); /* 单列布局 */
  }
}

.loading {
    font-size: 1.2em;
    text-align: center;
    width: 100%;
}
</style>