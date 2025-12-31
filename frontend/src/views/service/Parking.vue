<template>
  <div class="parking-page">
    <Navbar />
    
    <div class="container">
      <h1 class="page-title">车位管理</h1>
      
      <div class="parking-info card" v-if="parking">
        <h3>我的车位信息</h3>
        <div class="info-grid">
          <div class="info-item">
            <span class="info-label">车位号：</span>
            <span class="info-value">{{ parking.parking_no }}</span>
          </div>
          <div class="info-item">
            <span class="info-label">状态：</span>
            <span class="tag" :class="parking.status === 1 ? 'tag-success' : 'tag-warning'">
              {{ parking.status === 1 ? '已占用' : '空闲' }}
            </span>
          </div>
          <div class="info-item">
            <span class="info-label">车牌号：</span>
            <span class="info-value">{{ parking.car_plate || '未绑定' }}</span>
          </div>
        </div>
        
        <div class="bind-form">
          <h4>{{ parking.car_plate ? '更新车牌' : '绑定车牌' }}</h4>
          <form @submit.prevent="handleBindCar">
            <div class="form-group">
              <input 
                v-model="carPlate" 
                class="input" 
                placeholder="请输入车牌号，如：辽A88888"
                required
              />
            </div>
            <button type="submit" class="btn btn-primary" :disabled="loading">
              {{ loading ? '提交中...' : parking.car_plate ? '更新' : '绑定' }}
            </button>
          </form>
        </div>
      </div>
      
      <div class="empty-state card" v-else>
        <div class="empty-state-icon">🅿️</div>
        <p>您还没有分配车位</p>
        <p class="text-secondary">请联系物业管理员分配车位</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Navbar from '@/components/layout/Navbar.vue'
import { getMyParking, bindCar } from '@/api/service'

const parking = ref(null)
const carPlate = ref('')
const loading = ref(false)

const fetchParking = async () => {
  try {
    const data = await getMyParking()
    parking.value = data
    if (data.car_plate) {
      carPlate.value = data.car_plate
    }
  } catch (error) {
    console.log('未分配车位或获取失败')
  }
}

const handleBindCar = async () => {
  if (!carPlate.value.trim()) {
    alert('请输入车牌号')
    return
  }
  
  loading.value = true
  try {
    await bindCar({ car_plate: carPlate.value })
    alert(parking.value.car_plate ? '更新成功！' : '绑定成功！')
    await fetchParking()
  } catch (error) {
    alert('操作失败：' + (error.response?.data?.msg || error.message))
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchParking()
})
</script>

<style scoped>
.parking-page {
  min-height: 100vh;
  padding-bottom: var(--spacing-xl);
}

.parking-info {
  max-width: 600px;
  margin: 0 auto;
}

.parking-info h3 {
  font-size: var(--font-size-xl);
  margin-bottom: var(--spacing-lg);
}

.info-grid {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-xl);
  padding-bottom: var(--spacing-lg);
  border-bottom: 1px solid var(--border-color);
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.info-label {
  font-weight: 500;
  color: var(--text-secondary);
}

.info-value {
  font-size: var(--font-size-lg);
  font-weight: 600;
  color: var(--text-primary);
}

.bind-form h4 {
  font-size: var(--font-size-lg);
  margin-bottom: var(--spacing-md);
}

.bind-form form {
  display: flex;
  gap: var(--spacing-md);
  align-items: flex-end;
}

.bind-form .form-group {
  flex: 1;
}

.empty-state {
  max-width: 400px;
  margin: var(--spacing-xl) auto;
  text-align: center;
  padding: var(--spacing-xl);
}
</style>
