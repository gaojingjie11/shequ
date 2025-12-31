<template>
  <div class="admin-child-page">
    <Navbar />
    <div class="container">
      <div class="page-header">
        <h1 class="page-title">门店管理</h1>
        <button class="btn btn-primary" @click="openModal()">+ 添加门店</button>
      </div>

      <div class="table-container card">
        <table class="table">
          <thead>
            <tr>
              <th>ID</th>
              <th>名称</th>
              <th>地址</th>
              <th>电话</th>
              <th>营业时间</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="store in stores" :key="store.id">
              <td>{{ store.id }}</td>
              <td>{{ store.name }}</td>
              <td>{{ store.address }}</td>
              <td>{{ store.phone }}</td>
              <td>{{ store.business_hours }}</td>
              <td>
                <button class="btn btn-sm btn-info" @click="openModal(store)">编辑</button>
                <button class="btn btn-sm btn-danger ml-2" @click="handleDelete(store.id)">删除</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- 弹窗 -->
      <div class="modal-overlay" v-if="showModal">
        <div class="modal card">
          <h3>{{ isEdit ? '编辑门店' : '添加门店' }}</h3>
          <form @submit.prevent="handleSubmit">
            <div class="form-group">
              <label>门店名称</label>
              <input v-model="form.name" class="input" required />
            </div>
            <div class="form-group">
              <label>地址</label>
              <input v-model="form.address" class="input" required />
            </div>
            <div class="form-group">
              <label>电话</label>
              <input v-model="form.phone" class="input" required />
            </div>
             <div class="form-group">
              <label>区域</label>
              <input v-model="form.region" class="input" placeholder="如: A区" />
            </div>
            <div class="form-group">
              <label>营业时间</label>
              <input v-model="form.business_hours" class="input" placeholder="如: 09:00 - 22:00" />
            </div>
            
            <div class="modal-actions">
              <button type="button" class="btn btn-secondary" @click="closeModal">取消</button>
              <button type="submit" class="btn btn-primary">保存</button>
            </div>
          </form>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Navbar from '@/components/layout/Navbar.vue'
import request from '@/utils/request' // 没有封装store list API，这里直接调或者补充API
import { createStore, updateStore, deleteStore } from '@/api/admin' 

// 临时补充 getStores，因为 api/service 里好像没暴露给admin用
const getStores = () => {
    return request({ url: '/stores', method: 'get' })
}

const stores = ref([])
const showModal = ref(false)
const isEdit = ref(false)
const form = ref({
    id: 0,
    name: '',
    address: '',
    phone: '',
    region: '',
    business_hours: ''
})

const fetchStores = async () => {
    try {
        const res = await getStores()
        stores.value = res
    } catch (e) {
        console.error(e)
    }
}

const openModal = (store = null) => {
    isEdit.value = !!store
    if(store) {
        form.value = { ...store }
    } else {
        form.value = { id: 0, name: '', address: '', phone:'', region:'', business_hours:'' }
    }
    showModal.value = true
}

const closeModal = () => showModal.value = false

const handleSubmit = async () => {
    try {
        if(isEdit.value) {
            await updateStore(form.value)
        } else {
            await createStore(form.value)
        }
        alert('保存成功')
        closeModal()
        fetchStores()
    } catch (e) {
        alert('操作失败')
    }
}

const handleDelete = async (id) => {
    if(!confirm('确定删除?')) return
    try {
        await deleteStore(id)
        fetchStores()
    } catch (e) {
        alert('删除失败')
    }
}

onMounted(fetchStores)
</script>

<style scoped>
.admin-child-page {
  min-height: 100vh;
  padding-bottom: var(--spacing-xl);
}
.page-header {
  display: flex; justify-content: space-between; align-items: center; margin-bottom: var(--spacing-lg);
}
.table { width: 100%; border-collapse: collapse; }
.table th, .table td { padding: 12px; border-bottom: 1px solid #eee; text-align: left; }
.ml-2 { margin-left: 8px; }

/* Modal Styles - Reuse */
.modal-overlay {
  position: fixed; top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(0,0,0,0.5);
  display: flex; justify-content: center; align-items: center;
}
.modal { padding: 24px; width: 400px; max-width: 90%; }
.form-group { margin-bottom: 16px; display: flex; flex-direction: column; }
.modal-actions { display: flex; justify-content: flex-end; gap: 12px; margin-top: 24px; }
</style>
