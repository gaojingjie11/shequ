<template>
  <div class="login-page">
    <div class="login-container">
      <div class="login-card card">
        <h2 class="login-title">欢迎回来</h2>
        <p class="login-subtitle">登录智慧社区</p>
        
        <form @submit.prevent="handleLogin" class="login-form">
          <div class="form-group">
            <label>手机号</label>
            <input 
              v-model="form.mobile" 
              type="tel" 
              class="input" 
              placeholder="请输入手机号"
              required
            />
          </div>
          
          <div class="form-group">
            <label>密码</label>
            <input 
              v-model="form.password" 
              type="password" 
              class="input" 
              placeholder="请输入密码"
              required
            />
          </div>
          
          <div class="form-footer">
            <router-link to="/register" class="link">还没有账号？注册</router-link>
          </div>
          
          <button type="submit" class="btn btn-primary btn-lg" :disabled="loading">
            <span v-if="!loading">登录</span>
            <span v-else class="loading"></span>
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()

const form = ref({
  mobile: '',
  password: ''
})

const loading = ref(false)

const handleLogin = async () => {
  loading.value = true
  try {
    await userStore.login(form.value)
    alert('登录成功！')
    router.push('/home')
  } catch (error) {
    alert(error.response?.data?.message || '登录失败，请检查用户名和密码')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, var(--primary-light) 0%, var(--secondary-light) 100%);
}

.login-container {
  width: 100%;
  max-width: 400px;
  padding: var(--spacing-lg);
}

.login-card {
  animation: fadeIn 0.5s ease;
}

.login-title {
  font-size: var(--font-size-2xl);
  font-weight: 600;
  text-align: center;
  margin-bottom: var(--spacing-sm);
  color: var(--text-primary);
}

.login-subtitle {
  text-align: center;
  color: var(--text-secondary);
  margin-bottom: var(--spacing-xl);
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.form-group label {
  font-weight: 500;
  color: var(--text-primary);
}

.form-footer {
  display: flex;
  justify-content: flex-end;
}

.link {
  color: var(--primary-color);
  text-decoration: none;
  font-size: var(--font-size-sm);
}

.link:hover {
  text-decoration: underline;
}

.btn-lg {
  width: 100%;
}
</style>
