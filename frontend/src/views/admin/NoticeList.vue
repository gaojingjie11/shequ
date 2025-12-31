<template>
  <div class="admin-child-page">
    <Navbar />
    <div class="container">
      <div class="page-header">
        <h1 class="page-title">公告管理</h1>
        <button class="btn btn-primary" @click="openModal()">+ 发布公告</button>
      </div>

      <div class="table-container card">
        <table class="table">
          <thead>
            <tr>
              <th>标题</th>
              <th>内容摘要</th>
              <th>发布时间</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in notices" :key="item.id">
              <td>{{ item.title }}</td>
              <td class="content-cell">{{ item.content }}</td>
              <td>{{ formatDate(item.created_at) }}</td>
              <td>
                <button class="btn btn-sm btn-danger" @click="handleDelete(item.id)">删除</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="modal-overlay" v-if="showModal">
        <div class="modal card">
          <h3>发布公告</h3>
          <form @submit.prevent="handleSubmit">
            <div class="form-group">
              <label>标题</label>
              <input v-model="form.title" class="input" required />
            </div>
            <div class="form-group">
              <label>内容</label>
              <textarea v-model="form.content" class="input textarea" required></textarea>
            </div>
            <div class="modal-actions">
              <button type="button" class="btn btn-secondary" @click="closeModal">取消</button>
              <button type="submit" class="btn btn-primary">发布</button>
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
import { getNoticeList } from '@/api/service'
import { createNotice, deleteNotice } from '@/api/admin'
import dayjs from 'dayjs'

const notices = ref([])
const showModal = ref(false)
const form = ref({ title: '', content: '' })

const formatDate = (date) => dayjs(date).format('YYYY-MM-DD HH:mm')

const fetchNotices = async () => {
  try {
    notices.value = await getNoticeList()
  } catch (error) {
    console.error(error)
  }
}

const openModal = () => {
    form.value = { title: '', content: '' }
    showModal.value = true
}
const closeModal = () => showModal.value = false

const handleSubmit = async () => {
    try {
        await createNotice(form.value)
        alert('发布成功')
        closeModal()
        fetchNotices()
    } catch (e) {
        alert('发布失败')
    }
}

const handleDelete = async (id) => {
    if(!confirm('确定删除?')) return
    try {
        await deleteNotice(id)
        fetchNotices()
    } catch (e) {
        alert('删除失败')
    }
}

onMounted(fetchNotices)
</script>

<style scoped>
/* Reuse styles */
.admin-child-page { min-height: 100vh; padding-bottom: var(--spacing-xl); }
.page-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: var(--spacing-lg); }
.table { width: 100%; border-collapse: collapse; }
.table th, .table td { padding: 12px; border-bottom: 1px solid #eee; text-align: left; }
.content-cell { max-width: 300px; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }

.modal-overlay { position: fixed; top: 0; left: 0; right: 0; bottom: 0; background: rgba(0,0,0,0.5); display: flex; justify-content: center; align-items: center; }
.modal { padding: 24px; width: 400px; max-width: 90%; }
.form-group { margin-bottom: 16px; display: flex; flex-direction: column; }
.textarea { height: 100px; resize: vertical; }
.modal-actions { display: flex; justify-content: flex-end; gap: 12px; margin-top: 24px; }
</style>
