import { createRouter, createWebHistory } from 'vue-router';
import AppHome from '../views/AppHome.vue';
import LoginForm from '../components/LoginForm.vue';

const routes = [
  { path: '/login', component: LoginForm },
  {
    path: '/',
    component: AppHome,
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
