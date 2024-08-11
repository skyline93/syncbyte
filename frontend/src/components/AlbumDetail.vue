<template>
    <div>
        <button @click="goBack" class="back-button">Back to Albums</button>
        <h1>Album Photos</h1>
        <div v-if="loading" class="loading">加载中...</div>
        <div v-else class="gallery-container">
            <div ref="lightGalleryRef" class="gallery">
                <a v-for="(image, index) in images" :key="index" :href="imageUrls[index]" class="gallery-item">
                    <img :src="imageUrls[index]" :alt="image.file_name" />
                </a>
            </div>
        </div>
    </div>
</template>

<script setup>
import lightGallery from 'lightgallery';
import { ref, onMounted, nextTick } from 'vue';
import axiosInstance from '../services/axiosInstance';
import { BASE_URL } from '../config';
import { useRoute, useRouter } from 'vue-router';

const images = ref([]);
const imageUrls = ref([]);
const loading = ref(true);
const lightGalleryRef = ref(null);
const route = useRoute();
const router = useRouter();
const albumId = route.params.id;

const fetchImages = async () => {
    try {
        const response = await axiosInstance.get(`/api/v1/photo?album_id=${albumId}`);
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
        return '';
    }

    if (path.startsWith('http')) {
        return path;
    } else {
        try {
            const trimmedPath = path.replace(/^\//, '');
            const url = new URL(trimmedPath, BASE_URL);
            return url.href;
        } catch (error) {
            console.error('Invalid URL:', path, error);
            return '';
        }
    }
};

const getImageBlob = async (url) => {
    try {
        const response = await axiosInstance.get(url, { responseType: 'blob' });
        return URL.createObjectURL(response.data);
    } catch (error) {
        console.error('Error fetching image:', error);
        return '';
    }
};

const goBack = () => {
    router.push({ name: 'AlbumList' });
};

onMounted(async () => {
    const lgThumbnail = await import('lightgallery/plugins/thumbnail/lg-thumbnail.es5.js');
    const lgFullscreen = await import('lightgallery/plugins/fullscreen/lg-fullscreen.es5.js');

    nextTick(() => {
        lightGallery(lightGalleryRef.value, {
            plugins: [lgThumbnail.default, lgFullscreen.default],
            thumbnail: true,
            fullscreen: true,
        });
    });

    fetchImages();
});
</script>

<style scoped>
.back-button {
    margin: 10px 0;
    padding: 10px;
    background-color: #42b983;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

.back-button:hover {
    background-color: #369b72;
}

.gallery-container {
    padding: 10px;
}

.gallery {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    gap: 5px;
}

.gallery-item {
    width: 100%;
}

.gallery-item img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.loading {
    font-size: 1.2em;
    text-align: center;
    width: 100%;
}
</style>
