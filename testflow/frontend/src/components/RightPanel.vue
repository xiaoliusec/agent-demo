<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { ElTabs, ElTabPane, ElTable, ElTableColumn, ElProgress, ElTag, ElButton } from 'element-plus'

interface TestCase {
  id: string
  module: string
  name: string
  status: string
  priority: string
  method: string
}

const testCases = ref<TestCase[]>([])
const moduleStats = ref<{name: string, total: number, passed: number, failed: number}[]>([])

const activeTab = ref('list')

const statusMap: Record<string, string> = {
  passed: 'success',
  failed: 'danger',
  blocked: 'warning'
}

const priorityMap: Record<string, string> = {
  high: 'danger',
  medium: 'warning',
  low: 'info'
}

const groupByModule = computed(() => {
  const groups: Record<string, TestCase[]> = {}
  testCases.value.forEach(tc => {
    const module = tc.module || '未分类'
    if (!groups[module]) {
      groups[module] = []
    }
    groups[module].push(tc)
  })
  return groups
})

const totalProgress = computed(() => {
  if (testCases.value.length === 0) return 0
  const passed = testCases.value.filter(tc => tc.status === 'passed').length
  return Math.round((passed / testCases.value.length) * 100)
})

const exportReport = () => {
  // TODO: 调用后端导出 Excel
  console.log('Export report')
}

onMounted(() => {
  // @ts-ignore
  if (window.runtime) {
    // @ts-ignore
    window.runtime.Events.On('test_case_generated', (data: any) => {
      testCases.value.push({
        id: data.id || `TC-${testCases.value.length + 1}`,
        module: data.module || '登录模块',
        name: data.name || '测试用例',
        status: data.status || 'passed',
        priority: data.priority || 'medium',
        method: data.method || '场景法'
      })
    })
    
    // @ts-ignore
    window.runtime.Events.On('test_complete', () => {
      updateModuleStats()
    })
  }
})

const updateModuleStats = () => {
  const stats: Record<string, {total: number, passed: number, failed: number}> = {}
  testCases.value.forEach(tc => {
    if (!stats[tc.module]) {
      stats[tc.module] = { total: 0, passed: 0, failed: 0 }
    }
    stats[tc.module].total++
    if (tc.status === 'passed') {
      stats[tc.module].passed++
    } else if (tc.status === 'failed') {
      stats[tc.module].failed++
    }
  })
  moduleStats.value = Object.entries(stats).map(([name, s]) => ({
    name,
    ...s
  }))
}
</script>

<template>
  <div class="right-panel-container">
    <div class="header">
      <h3>测试进度</h3>
      <el-button type="primary" size="small" @click="exportReport">导出报告</el-button>
    </div>
    
    <div class="progress-section">
      <el-progress 
        :percentage="totalProgress" 
        :stroke-width="12"
        :color="totalProgress >= 80 ? '#67c23a' : totalProgress >= 50 ? '#409eff' : '#e6a23c'"
      />
      <div class="progress-stats">
        <span>总用例: {{ testCases.length }}</span>
        <span class="passed">通过: {{ testCases.filter(tc => tc.status === 'passed').length }}</span>
        <span class="failed">失败: {{ testCases.filter(tc => tc.status === 'failed').length }}</span>
      </div>
    </div>
    
    <el-tabs v-model="activeTab" class="case-tabs">
      <el-tab-pane label="用例列表" name="list">
        <el-table :data="testCases" size="small" max-height="400">
          <el-table-column prop="id" label="编号" width="80" />
          <el-table-column prop="name" label="用例名称" min-width="120" show-overflow-tooltip />
          <el-table-column prop="module" label="模块" width="80" show-overflow-tooltip />
          <el-table-column prop="status" label="状态" width="70">
            <template #default="{ row }">
              <el-tag :type="statusMap[row.status]" size="small">
                {{ row.status }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="priority" label="优先级" width="70">
            <template #default="{ row }">
              <el-tag :type="priorityMap[row.priority]" size="small">
                {{ row.priority }}
              </el-tag>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
      
      <el-tab-pane label="模块统计" name="modules">
        <div class="module-list">
          <div v-for="stat in moduleStats" :key="stat.name" class="module-item">
            <div class="module-name">{{ stat.name }}</div>
            <div class="module-stats">
              <span class="total">总计: {{ stat.total }}</span>
              <span class="passed">通过: {{ stat.passed }}</span>
              <span class="failed">失败: {{ stat.failed }}</span>
            </div>
            <el-progress 
              :percentage="stat.total > 0 ? Math.round((stat.passed / stat.total) * 100) : 0"
              :stroke-width="6"
              :show-text="false"
            />
          </div>
          <div v-if="moduleStats.length === 0" class="empty-hint">
            暂无统计数据
          </div>
        </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<style scoped>
.right-panel-container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  border-bottom: 1px solid #e0e0e0;
}

.header h3 {
  margin: 0;
  font-size: 16px;
  color: #303133;
}

.progress-section {
  padding: 16px;
  border-bottom: 1px solid #e0e0e0;
}

.progress-stats {
  display: flex;
  justify-content: space-between;
  margin-top: 8px;
  font-size: 12px;
  color: #606266;
}

.progress-stats .passed {
  color: #67c23a;
}

.progress-stats .failed {
  color: #f56c6c;
}

.case-tabs {
  flex: 1;
  padding: 0 16px;
}

.case-tabs :deep(.el-tabs__header) {
  margin-bottom: 12px;
}

.module-list {
  padding: 8px 0;
}

.module-item {
  padding: 12px 0;
  border-bottom: 1px solid #f5f5f5;
}

.module-item:last-child {
  border-bottom: none;
}

.module-name {
  font-weight: 500;
  color: #303133;
  margin-bottom: 8px;
}

.module-stats {
  display: flex;
  gap: 16px;
  font-size: 12px;
  color: #909399;
  margin-bottom: 8px;
}

.module-stats .passed {
  color: #67c23a;
}

.module-stats .failed {
  color: #f56c6c;
}

.empty-hint {
  text-align: center;
  color: #c0c4cc;
  padding: 40px 0;
}
</style>
