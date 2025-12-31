<template>
  <div class="data-screen">
    <div class="screen-header">
      <div class="header-left" @click="$router.push('/admin')">
        <el-icon><ArrowLeft /></el-icon> 返回管理后台
      </div>
      <div class="header-center">智慧社区数据可视化大屏</div>
      <div class="header-right">{{ currentTime }}</div>
    </div>

    <div class="screen-body">
      <!-- Row 1: Key Metrics -->
      <div class="metrics-row">
        <el-row :gutter="20">
          <el-col :span="6">
            <div class="metric-card">
              <div class="metric-icon user-bg"><el-icon><User /></el-icon></div>
              <div class="metric-info">
                <div class="label">总用户数</div>
                <div class="value">{{ stats.totalUsers || 0 }}</div>
              </div>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="metric-card">
              <div class="metric-icon order-bg"><el-icon><ShoppingCart /></el-icon></div>
              <div class="metric-info">
                <div class="label">今日订单</div>
                <div class="value">{{ stats.todayOrders || 0 }}</div>
              </div>
            </div>
          </el-col>
           <el-col :span="6">
            <div class="metric-card">
              <div class="metric-icon parking-bg"><el-icon><Van /></el-icon></div>
              <div class="metric-info">
                <div class="label">车位占用率</div>
                <div class="value">{{ stats.parkingRate || '0%' }}</div>
              </div>
            </div>
          </el-col>
           <el-col :span="6">
            <div class="metric-card">
              <div class="metric-icon money-bg"><el-icon><Money /></el-icon></div>
              <div class="metric-info">
                <div class="label">本月营收</div>
                <div class="value">¥{{ stats.monthIncome || 0 }}</div>
              </div>
            </div>
          </el-col>
        </el-row>
      </div>

      <!-- Row 2: Charts -->
      <el-row :gutter="20" class="charts-row">
        <el-col :span="16">
          <div class="chart-card">
            <div class="chart-title">近7日交易走势</div>
            <div ref="lineChartRef" class="chart-container"></div>
          </div>
        </el-col>
        <el-col :span="8">
           <div class="chart-card">
            <div class="chart-title">费用构成分析</div>
            <div ref="pieChartRef" class="chart-container"></div>
          </div>
        </el-col>
      </el-row>

       <!-- Row 3: More Charts -->
       <el-row :gutter="20" class="charts-row">
        <el-col :span="12">
           <div class="chart-card">
            <div class="chart-title">报修类型分布</div>
             <div ref="barChartRef" class="chart-container"></div>
          </div>
        </el-col>
         <el-col :span="12">
           <div class="chart-card">
            <div class="chart-title">实时访客记录</div>
            <div class="visitor-list">
              <el-table :data="visitorList" style="width: 100%" :row-class-name="tableRowClassName">
                <el-table-column prop="visitor_name" label="访客姓名" width="100" />
                <el-table-column prop="visit_time" label="来访时间" />
                 <el-table-column prop="reason" label="事由" />
                <el-table-column label="状态">
                   <template #default="scope">
                      <el-tag size="small" :type="scope.row.status === 1 ? 'success' : 'warning'">
                        {{ scope.row.status === 1 ? '已通过' : '待审核' }}
                      </el-tag>
                   </template>
                </el-table-column>
              </el-table>
            </div>
          </div>
        </el-col>
       </el-row>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import * as echarts from 'echarts'
import dayjs from 'dayjs'
import { getAdminParkingList } from '@/api/admin' 
// Note: We might need more dedicated stats APIs, but for now mocking/reusing what we have.

const currentTime = ref(dayjs().format('YYYY-MM-DD HH:mm:ss'))
let timer = null
const stats = ref({
  totalUsers: 142,
  todayOrders: 23,
  parkingRate: '65%',
  monthIncome: '12,450'
})

const visitorList = ref([
    { visitor_name: '张三', visit_time: '2024-12-31 10:30', reason: '快递', status: 1 },
    { visitor_name: '李四', visit_time: '2024-12-31 11:15', reason: '外卖', status: 1 },
    { visitor_name: '王五', visit_time: '2024-12-31 11:20', reason: '探亲', status: 0 },
    { visitor_name: '赵六', visit_time: '2024-12-31 11:45', reason: '维修', status: 0 },
])

const lineChartRef = ref(null)
const pieChartRef = ref(null)
const barChartRef = ref(null)
let lineChart = null
let pieChart = null
let barChart = null

const initCharts = () => {
    // Line Chart
    lineChart = echarts.init(lineChartRef.value)
    lineChart.setOption({
        tooltip: { trigger: 'axis' },
        grid: { left: '3%', right: '4%', bottom: '3%', containLabel: true },
        xAxis: { type: 'category', data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'] },
        yAxis: { type: 'value' },
        series: [{
            data: [150, 230, 224, 218, 135, 147, 260],
            type: 'line',
            smooth: true,
            areaStyle: {},
            itemStyle: { color: '#409EFF' }
        }]
    })

    // Pie Chart
    pieChart = echarts.init(pieChartRef.value)
    pieChart.setOption({
        tooltip: { trigger: 'item' },
        legend: { top: '5%', left: 'center' },
        series: [{
            name: 'Access From',
            type: 'pie',
            radius: ['40%', '70%'],
            avoidLabelOverlap: false,
            itemStyle: { borderRadius: 10, borderColor: '#fff', borderWidth: 2 },
            label: { show: false, position: 'center' },
            emphasis: { label: { show: true, fontSize: 20, fontWeight: 'bold' } },
            labelLine: { show: false },
            data: [
                { value: 1048, name: '物业费' },
                { value: 735, name: '停车费' },
                { value: 580, name: '商城自营' },
                { value: 484, name: '其他收入' },
            ]
        }]
    })

    // Bar Chart
     barChart = echarts.init(barChartRef.value)
     barChart.setOption({
        tooltip: { trigger: 'axis', axisPointer: { type: 'shadow' } },
        grid: { left: '3%', right: '4%', bottom: '3%', containLabel: true },
        xAxis: [{ type: 'category', data: ['水电', '门窗', '暖气', '公共区域', '电梯', '其他'] }],
        yAxis: [{ type: 'value' }],
        series: [{
            name: '报修数量',
            type: 'bar',
            barWidth: '60%',
            data: [10, 52, 200, 334, 390, 330],
            itemStyle: { color: '#E6A23C' }
        }]
     })
}

// Need resize listener
const handleResize = () => {
    lineChart && lineChart.resize()
    pieChart && pieChart.resize()
    barChart && barChart.resize()
}

onMounted(() => {
    timer = setInterval(() => {
        currentTime.value = dayjs().format('YYYY-MM-DD HH:mm:ss')
    }, 1000)
    initCharts()
    window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
    clearInterval(timer)
    window.removeEventListener('resize', handleResize)
    lineChart && lineChart.dispose()
    pieChart && pieChart.dispose()
    barChart && barChart.dispose()
})

const tableRowClassName = ({ rowIndex }) => {
  if (rowIndex === 1) {
    return 'warning-row'
  } else if (rowIndex === 3) {
    return 'success-row'
  }
  return ''
}
</script>

<style scoped>
.data-screen {
    width: 100%;
    min-height: 100vh;
    background-color: #0b1120; /* Dark Theme */
    color: #fff;
    padding: 20px;
}

.screen-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
    border-bottom: 1px solid #1f2937;
    padding-bottom: 15px;
}

.header-left {
    cursor: pointer;
    display: flex;
    align-items: center;
    gap: 5px;
    font-size: 14px;
    color: #9ca3af;
}

.header-center {
    font-size: 24px;
    font-weight: bold;
    background: linear-gradient(to right, #409EFF, #00f2fe);
    -webkit-background-clip: text;
    color: transparent;
}

.header-right {
    font-family: monospace;
    font-size: 16px;
    color: #409EFF;
}

.metrics-row {
  margin-bottom: 20px;
}

.metric-card {
    background-color: #1f2937;
    padding: 20px;
    border-radius: 8px;
    display: flex;
    align-items: center;
    gap: 15px;
    transition: transform 0.3s;
}

.metric-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 4px 12px rgba(0,0,0,0.5);
}

.metric-icon {
    width: 50px;
    height: 50px;
    border-radius: 12px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 24px;
    color: white;
}
.user-bg { background: linear-gradient(135deg, #667eea, #764ba2); }
.order-bg { background: linear-gradient(135deg, #ff9a9e, #fecfef); }
.parking-bg { background: linear-gradient(135deg, #a18cd1, #fbc2eb); }
.money-bg { background: linear-gradient(135deg, #84fab0, #8fd3f4); }

.metric-info .label {
    color: #9ca3af;
    font-size: 12px;
    margin-bottom: 4px;
}
.metric-info .value {
    font-size: 24px;
    font-weight: 700;
}

.charts-row {
    margin-bottom: 20px;
}

.chart-card {
    background-color: #1f2937;
    padding: 20px;
    border-radius: 8px;
    height: 350px;
}

.chart-title {
    font-size: 16px;
    font-weight: 600;
    margin-bottom: 15px;
    border-left: 4px solid #409EFF;
    padding-left: 10px;
}

.chart-container {
    width: 100%;
    height: 300px;
}

/* Element UI Overrides for Dark Mode if needed, but simplistic approach here */
:deep(.el-table) {
    --el-table-bg-color: #1f2937;
    --el-table-tr-bg-color: #1f2937;
    --el-table-header-bg-color: #111827;
    --el-table-text-color: #e5e7eb;
    --el-table-header-text-color: #9ca3af;
    --el-table-row-hover-bg-color: #374151;
    --el-table-border-color: #374151;
    background-color: transparent !important;
}

:deep(.el-table tr) {
    background-color: #1f2937 !important;
}
</style>
