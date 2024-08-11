<template>
  <div class="container">
    <!-- 左侧菜单栏 -->
    <aside class="sidebar">
      <h2>Syncbyte</h2>
      <ul>
        <li @click="navigateTo('AlbumList')">相册</li>
        <!-- 根据需要添加更多菜单项 -->
      </ul>
    </aside>

    <!-- 动态内容显示区域 -->
    <div class="main-content">
      <router-view></router-view>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

// 导航到不同的页面
const navigateTo = (routeName) => {
  router.push({ name: routeName });
};
</script>

<style scoped>
.container {
  display: flex;
}

.sidebar {
  width: 200px;
  background-color: #f5f5f5;
  padding: 20px;
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.1);
}

.sidebar h2 {
  margin-top: 0;
}

.sidebar ul {
  list-style: none;
  padding: 0;
}

.sidebar ul li {
  cursor: pointer;
  padding: 10px 0;
  border-bottom: 1px solid #ccc;
}

.sidebar ul li:hover {
  background-color: #ddd;
}

.main-content {
  flex: 1;
  padding: 20px;
}
</style>
