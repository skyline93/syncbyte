// src/services/authService.js
import axios from 'axios';

const API_URL = 'http://192.168.209.130:8000';

export const login = async (username, password) => {
  try {
    const response = await axios.post(`${API_URL}/login`, {
      username,
      password
    }, {
      headers: { 'Content-Type': 'application/json' }
    });

    const { token } = response.data.token;
    localStorage.setItem('authToken', token); // 保存令牌到本地存储
    return token;
  } catch (error) {
    console.error('Login failed:', error);
    throw error;
  }
};
