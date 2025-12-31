<template>
  <div class="home-page">
    <Navbar />
    
    <div class="banner">
      <div class="container">
        <h1 class="banner-title">欢迎来到智慧社区</h1>
        <p class="banner-subtitle">让生活更便捷、更智能、更美好</p>
      </div>
    </div>
    
    <div class="container">
      <div class="quick-menu">
        <div class="quick-item" @click="$router.push('/mall')">
          <div class="quick-icon">🛒</div>
          <div class="quick-text">社区商城</div>
        </div>
        <div class="quick-item" @click="$router.push('/service/notice')">
          <div class="quick-icon">📢</div>
          <div class="quick-text">公告通知</div>
        </div>
        <div class="quick-item" @click="$router.push('/service/repair')">
          <div class="quick-icon">🔧</div>
          <div class="quick-text">报修投诉</div>
        </div>
        <div class="quick-item" @click="$router.push('/service/visitor')">
          <div class="quick-icon">👥</div>
          <div class="quick-text">访客登记</div>
        </div>
      </div>
      
      <div class="section">
        <h2 class="section-title">最新公告</h2>
        <div class="notice-list">
          <div class="notice-item card" v-for="notice in notices" :key="notice.id" @click="goToNotice(notice.id)">
            <div class="notice-title">{{ notice.title }}</div>
            <div class="notice-meta">
              <span>{{ notice.publisher }}</span>
              <span>{{ formatDate(notice.created_at) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import Navbar from '@/components/layout/Navbar.vue'
import { getNoticeList } from '@/api/service'
import dayjs from 'dayjs'

const router = useRouter()
const notices = ref([])

const formatDate = (date) => {
  return dayjs(date).format('YYYY-MM-DD')
}

const goToNotice = (id) => {
  router.push(`/service/notice?id=${id}`)
}

onMounted(async () => {
  try {
    const list = await getNoticeList()
    notices.value = list.slice(0, 5) // 只显示最新5条
  } catch (error) {
    console.error('获取公告失败:', error)
  }
})
</script>

<style scoped>
.home-page {
  min-height: 100vh;
}

.banner {
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--secondary-color) 100%);
  color: white;
  padding: 80px 0;
  margin-bottom: var(--spacing-xl);
}

.banner-title {
  font-size: 48px;
  font-weight: 700;
  margin-bottom: var(--spacing-md);
  text-align: center;
}

.banner-subtitle {
  font-size: 20px;
  text-align: center;
  opacity: 0.9;
}

.quick-menu {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: var(--spacing-lg);
  margin-bottom: var(--spacing-xl);
}

.quick-item {
  background: white;
  border-radius: var(--border-radius);
  padding: var(--spacing-xl);
  text-align: center;
  cursor: pointer;
  transition: all var(--transition-base);
  box-shadow: var(--shadow-sm);
}

.quick-item:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-md);
}

.quick-icon {
  font-size: 48px;
  margin-bottom: var(--spacing-md);
}

.quick-text {
  font-size: var(--font-size-lg);
  font-weight: 500;
  color: var(--text-primary);
}

.section {
  margin-bottom: var(--spacing-xl);
}

.section-title {
  font-size: var(--font-size-xl);
  font-weight: 600;
  margin-bottom: var(--spacing-lg);
}

.notice-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.notice-item {
  cursor: pointer;
  transition: all var(--transition-base);
}

.notice-item:hover {
  transform: translateX(4px);
}

.notice-title {
  font-size: var(--font-size-lg);
  font-weight: 500;
  margin-bottom: var(--spacing-sm);
  color: var(--text-primary);
}

.notice-meta {
  display: flex;
  justify-content: space-between;
  font-size: var(--font-size-sm);
  color: var(--text-secondary);
}

@media (max-width: 768px) {
  .quick-menu {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .banner-title {
    font-size: 32px;
  }
}
</style>
