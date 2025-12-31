<template>
  <div class="notice-page">
    <Navbar />
    
    <div class="container">
      <h1 class="page-title">公告通知</h1>
      
      <div class="notice-list">
        <div class="notice-card card" v-for="notice in notices" :key="notice.id">
          <div class="notice-header">
            <h3 class="notice-title">{{ notice.title }}</h3>
            <div class="notice-meta">
              <span class="publisher">{{ notice.publisher }}</span>
              <span class="date">{{ formatDate(notice.created_at) }}</span>
            </div>
          </div>
          <p class="notice-content">{{ notice.content }}</p>
          <div class="notice-footer">
            <span class="views">👁 {{ notice.view_count }} 次浏览</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Navbar from '@/components/layout/Navbar.vue'
import { getNoticeList } from '@/api/service'
import dayjs from 'dayjs'

const notices = ref([])

const formatDate = (date) => {
  return dayjs(date).format('YYYY-MM-DD HH:mm')
}

onMounted(async () => {
  try {
    notices.value = await getNoticeList()
  } catch (error) {
    console.error('获取公告失败:', error)
  }
})
</script>

<style scoped>
.notice-page {
  min-height: 100vh;
  padding-bottom: var(--spacing-xl);
}

.notice-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.notice-card {
  padding: var(--spacing-xl);
}

.notice-header {
  margin-bottom: var(--spacing-md);
}

.notice-title {
  font-size: var(--font-size-xl);
  font-weight: 600;
  margin-bottom: var(--spacing-sm);
}

.notice-meta {
  display: flex;
  gap: var(--spacing-lg);
  font-size: var(--font-size-sm);
  color: var(--text-secondary);
}

.publisher {
  color: var(--primary-color);
  font-weight: 500;
}

.notice-content {
  line-height: 1.8;
  color: var(--text-primary);
  margin-bottom: var(--spacing-md);
}

.notice-footer {
  padding-top: var(--spacing-md);
  border-top: 1px solid var(--border-color);
  font-size: var(--font-size-sm);
  color: var(--text-light);
}
</style>
