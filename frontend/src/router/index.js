import { createRouter, createWebHistory } from 'vue-router';
import AlbumList from '../components/AlbumList.vue'; // 相册列表页面组件
import AlbumDetail from '../components/AlbumDetail.vue'; // 相册详情页面组件
// import Home from '../components/Home.vue'; // 主页组件
// import About from '../components/About.vue'; // 关于我们页面组件

const routes = [
    // { path: '/', name: 'Home', component: Home },
    { path: '/albums', name: 'AlbumList', component: AlbumList },
    { path: '/albums/:id', name: 'AlbumDetail', component: AlbumDetail },
    // { path: '/about', name: 'About', component: About },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
