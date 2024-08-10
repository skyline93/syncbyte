// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import Login from '../components/Login.vue';  // 登录组件
import Photo from '../components/PhotoGallery.vue';    // 主页组件（可选）

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/photo',
    name: 'photo',
    component: Photo
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
