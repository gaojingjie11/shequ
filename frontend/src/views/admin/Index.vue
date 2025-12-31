<template>
  <div class="admin-page">
    <Navbar />
    
    <div class="container">
      <h1 class="page-title">管理后台</h1>
      
      <div class="admin-menu">
        <div class="admin-card card" v-if="hasPermission(['admin'])">
          <div class="admin-icon">👥</div>
          <h3>用户管理</h3>
          <p>管理社区用户</p>
          <button class="btn btn-sm btn-primary" @click="$router.push('/admin/users')">进入</button>
        </div>
        
        <div class="admin-card card" v-if="hasPermission(['admin', 'store'])">
          <div class="admin-icon">🛒</div>
          <h3>商品管理</h3>
          <p>管理商城商品</p>
          <button class="btn btn-sm btn-primary" @click="$router.push('/admin/products')">进入</button>
        </div>
        
        <div class="admin-card card" v-if="hasPermission(['admin', 'store'])">
          <div class="admin-icon">📦</div>
          <h3>订单管理</h3>
          <p>查看所有订单</p>
          <button class="btn btn-sm btn-primary" @click="$router.push('/admin/orders')">进入</button>
        </div>
        
        <div class="admin-card card" v-if="hasPermission(['admin', 'store'])">
          <div class="admin-icon">🏪</div>
          <h3>门店管理</h3>
          <p>管理社区门店</p>
          <button class="btn btn-sm btn-primary" @click="$router.push('/admin/stores')">进入</button>
        </div>
        
        <div class="admin-card card" v-if="hasPermission(['admin', 'property'])">
          <div class="admin-icon">📢</div>
          <h3>公告管理</h3>
          <p>发布和管理公告</p>
          <button class="btn btn-sm btn-primary" @click="$router.push('/admin/notices')">进入</button>
        </div>
        
        <div class="admin-card card" v-if="hasPermission(['admin', 'property'])">
          <div class="admin-icon">🔧</div>
          <h3>报修处理</h3>
          <p>处理报修投诉</p>
          <button class="btn btn-sm btn-primary" @click="$router.push('/admin/repairs')">进入</button>
        </div>
        
        <div class="admin-card card" v-if="hasPermission(['admin', 'property'])">
          <div class="admin-icon">👥</div>
          <h3>访客审核</h3>
          <p>审核访客申请</p>
          <button class="btn btn-sm btn-primary" @click="$router.push('/admin/visitors')">进入</button>
        </div>

        <div class="admin-card card" v-if="hasPermission(['admin', 'property'])">
          <div class="admin-icon">🅿️</div>
          <h3>车位管理</h3>
          <p>管理车位及分配</p>
          <button class="btn btn-sm btn-primary" @click="$router.push('/admin/parking')">进入</button>
        </div>

        <div class="admin-card card" v-if="hasPermission(['admin', 'property'])">
          <div class="admin-icon">💰</div>
          <h3>物业费</h3>
          <p>账单管理与发布</p>
          <button class="btn btn-sm btn-primary" @click="$router.push('/admin/property-fee')">进入</button>
        </div>
        
        <div class="admin-card card" v-if="hasPermission(['admin', 'store', 'property'])">
          <div class="admin-icon">📈</div>
          <h3>数据大屏</h3>
          <p>查看社区数据可视化</p>
          <button class="btn btn-sm btn-primary" @click="$router.push('/admin/data')">进入</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import Navbar from '@/components/layout/Navbar.vue'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const role = userStore.userInfo.role || 'user'

const hasPermission = (allowedRoles) => {
  return allowedRoles.includes(role)
}
</script>

<style scoped>
.admin-page {
  min-height: 100vh;
  padding-bottom: var(--spacing-xl);
}

.admin-menu {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: var(--spacing-lg);
}

.admin-card {
  text-align: center;
  padding: var(--spacing-xl);
  transition: all var(--transition-base);
}

.admin-card:hover {
  transform: translateY(-4px);
}

.admin-icon {
  font-size: 48px;
  margin-bottom: var(--spacing-md);
}

.admin-card h3 {
  font-size: var(--font-size-lg);
  font-weight: 600;
  margin-bottom: var(--spacing-sm);
}

.admin-card p {
  color: var(--text-secondary);
  margin-bottom: var(--spacing-lg);
}
</style>
