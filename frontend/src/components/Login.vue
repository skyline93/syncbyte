<template>
  <div>
    <h1>Login</h1>
    <form @submit.prevent="handleLogin">
      <input v-model="username" type="text" placeholder="Username" required />
      <input v-model="password" type="password" placeholder="Password" required />
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { login } from '../services/authService'; // 登录服务

const username = ref('admin');
const password = ref('admin123');

const handleLogin = async () => {
  try {
    await login(username.value, password.value);
    alert('Login successful!');
    // 登录成功后重定向到主页或其他页面
    window.location.href = '/';
  } catch (error) {
    alert('Login failed!');
  }
};
</script>
