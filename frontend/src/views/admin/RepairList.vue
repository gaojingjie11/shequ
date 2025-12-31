<template>
  <div class="admin-child-page">
    <Navbar />
    <div class="container">
      <div class="page-header">
        <h1 class="page-title">报修处理</h1>
      </div>

      <div class="table-container card">
        <table class="table">
          <thead>
            <tr>
              <th>提交人</th>
              <th>电话</th>
              <th>内容</th>
              <th>状态</th>
              <th>提交时间</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in list" :key="item.id">
              <td>{{ item.user_id }} (需关联名)</td>
              <td>{{ item.phone }}</td>
              <td class="content-cell">{{ item.content }}</td>
              <td>
                <span class="tag" :class="getStatusClass(item.status)">
                  {{ getStatusText(item.status) }}
                </span>
              </td>
              <td>{{ formatDate(item.created_at) }}</td>
              <td>
                <button 
                  v-if="item.status === 0" 
                  class="btn btn-sm btn-primary" 
                  @click="openProcess(item)"
                >
                  处理
                </button>
                <div v-else>
                    <small>{{ item.result }}</small>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="modal-overlay" v-if="showModal">
        <div class="modal card">
          <h3>处理报修</h3>
          <p class="mb-4">报修内容: {{ currentItem?.content }}</p>
          <form @submit.prevent="handleSubmit">
            <div class="form-group">
              <label>处理结果/反馈</label>
              <textarea v-model="processForm.result" class="input textarea" required placeholder="请输入处理结果..."></textarea>
            </div>
             <div class="modal-actions">
              <button type="button" class="btn btn-secondary" @click="closeModal">取消</button>
              <button type="submit" class="btn btn-primary">提交</button>
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
import { getAdminRepairList, processRepair } from '@/api/admin'
import dayjs from 'dayjs'

const list = ref([])
const showModal = ref(false)
const currentItem = ref(null)
const processForm = ref({ result: '' })

const formatDate = (date) => dayjs(date).format('YYYY-MM-DD HH:mm')
const getStatusText = (s) => (s===0 ? '待处理' : '已处理')
const getStatusClass = (s) => (s===0 ? 'tag-warning' : 'tag-success')

const fetchList = async () => {
    try {
        const res = await getAdminRepairList()
        // API 返回可能带了用户信息，这里简化
        list.value = res
    } catch (e) {
        console.error(e)
    }
}

const openProcess = (item) => {
    currentItem.value = item
    processForm.value.result = ''
    showModal.value = true
}
const closeModal = () => showModal.value = false

const handleSubmit = async () => {
    try {
        await processRepair({
            id: currentItem.value.id,
            feedback: processForm.value.result, // 改为 feedback
            status: 1 // 标记为已处理
        })
        alert('处理成功')
        closeModal()
        fetchList()
    } catch(e) {
        alert('提交失败')
    }
}

onMounted(fetchList)
</script>

<style scoped>
/* Reuse styles */
.admin-child-page { min-height: 100vh; padding-bottom: var(--spacing-xl); }
.page-header { display: flex; justify-content: space-before; margin-bottom: var(--spacing-lg); }
.table { width: 100%; border-collapse: collapse; }
.table th, .table td { padding: 12px; border-bottom: 1px solid #eee; text-align: left; }
.content-cell { max-width: 250px; }
.mb-4 { margin-bottom: 16px; }

.modal-overlay { position: fixed; top: 0; left: 0; right: 0; bottom: 0; background: rgba(0,0,0,0.5); display: flex; justify-content: center; align-items: center; }
.modal { padding: 24px; width: 400px; max-width: 90%; }
.form-group { margin-bottom: 16px; display: flex; flex-direction: column; }
.textarea { height: 100px; resize: vertical; }
.modal-actions { display: flex; justify-content: flex-end; gap: 12px; margin-top: 24px; }
</style>
