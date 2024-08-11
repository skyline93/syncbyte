import { createRouter, createWebHistory } from 'vue-router';
// import AppHome from '../views/AppHome.vue';
import LoginForm from '../components/LoginForm.vue';
import PhotoGallery from '../components/PhotoGallery.vue';
import AlbumList from '../components/AlbumList.vue';
import AlbumDetail from '../components/AlbumDetail.vue';

const routes = [
  {
    path: '/',
    name: 'AlbumList',
    component: AlbumList
  },
  {
    path: '/album/:id',
    name: 'AlbumDetail',
    component: AlbumDetail,
    props: true
  },
  { path: '/login', component: LoginForm },
  // {
  //   path: '/',
  //   component: AppHome,
  //   meta: { requiresAuth: true }
  // },
  {
    path: '/photo',
    component: PhotoGallery,
    meta: { requiresAuth: true }
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    const token = localStorage.getItem('token');
    if (!token) {
      next('/login');
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
