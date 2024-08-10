// src/services/apiService.js
import axios from 'axios';

const API_URL = 'http://192.168.209.130:8000';

const getAuthToken = () => {
  return localStorage.getItem('authToken');
};

const axiosInstance = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json'
  }
});

axiosInstance.interceptors.request.use(
  config => {
    const token = getAuthToken();
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

axiosInstance.interceptors.response.use(
  response => response,
  error => {
    if (error.response && error.response.status === 401) {
      // 处理401 Unauthorized错误
      localStorage.removeItem('authToken');
      // 跳转到登录页面或进行其他处理
    }
    return Promise.reject(error);
  }
);

export const fetchProtectedData = async () => {
  try {
    const response = await axiosInstance.get('/protected');
    return response.data;
  } catch (error) {
    console.error('Failed to fetch protected data:', error);
    throw error;
  }
};
