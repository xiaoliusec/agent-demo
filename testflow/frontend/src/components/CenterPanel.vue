<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElCard, ElScrollbar } from 'element-plus'

const screenshotUrl = ref('')
const operationLogs = ref<{time: string, action: string, result: string}[]>([])

onMounted(() => {
  // @ts-ignore
  if (window.runtime) {
    // @ts-ignore
    window.runtime.Events.On('screenshot_update', (data: number[]) => {
      const blob = new Blob([new Uint8Array(data)], { type: 'image/png' })
      screenshotUrl.value = URL.createObjectURL(blob)
    })
    
    // @ts-ignore
    window.runtime.Events.On('step_complete', (data: any) => {
      operationLogs.value.push({
        time: new Date().toLocaleTimeString(),
        action: data.result?.Output || '完成',
        result: data.result?.OK ? '成功' : '失败'
      })
    })
  }
})
</script>

<template>
  <div class="center-panel-container">
    <div class="browser-section">
      <div class="browser-header">
        <span>浏览器预览</span>
        <div class="browser-controls">
          <span class="control-dot"></span>
          <span class="control-dot"></span>
          <span class="control-dot"></span>
        </div>
      </div>
      <div class="browser-content">
        <img v-if="screenshotUrl" :src="screenshotUrl" alt="Browser Screenshot" class="screenshot" />
        <div v-else class="placeholder">
          <p>暂无截图</p>
          <p class="hint">开始测试后将显示浏览器画面</p>
        </div>
      </div>
    </div>
    
    <div class="operation-section">
      <div class="section-header">
        <span>操作日志</span>
        <el-button text size="small" @click="operationLogs = []">清空</el-button>
      </div>
      <el-scrollbar height="200px">
        <div v-if="operationLogs.length === 0" class="empty-hint">
          暂无操作记录
        </div>
        <div v-for="(log, idx) in operationLogs" :key="idx" class="operation-item">
          <span class="time">{{ log.time }}</span>
          <span class="action">{{ log.action }}</span>
          <span class="result" :class="log.result">{{ log.result }}</span>
        </div>
      </el-scrollbar>
    </div>
  </div>
</template>

<style scoped>
.center-panel-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 16px;
  gap: 16px;
}

.browser-section {
  flex: 1;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.browser-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: #f5f5f5;
  border-bottom: 1px solid #e0e0e0;
  font-size: 13px;
  font-weight: 500;
}

.browser-controls {
  display: flex;
  gap: 6px;
}

.control-dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background: #ddd;
}

.control-dot:nth-child(1) { background: #ff5f57; }
.control-dot:nth-child(2) { background: #ffbd2e; }
.control-dot:nth-child(3) { background: #28c940; }

.browser-content {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #fafafa;
  overflow: hidden;
}

.screenshot {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
}

.placeholder {
  text-align: center;
  color: #909399;
}

.placeholder p {
  margin: 8px 0;
}

.placeholder .hint {
  font-size: 12px;
  color: #c0c4cc;
}

.operation-section {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  padding: 12px 16px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  font-size: 13px;
  font-weight: 500;
}

.empty-hint {
  text-align: center;
  color: #c0c4cc;
  font-size: 12px;
  padding: 20px;
}

.operation-item {
  display: flex;
  align-items: center;
  padding: 6px 0;
  font-size: 12px;
  border-bottom: 1px solid #f5f5f5;
}

.operation-item:last-child {
  border-bottom: none;
}

.time {
  color: #909399;
  width: 80px;
  flex-shrink: 0;
}

.action {
  flex: 1;
  color: #303133;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.result {
  width: 40px;
  text-align: right;
  flex-shrink: 0;
}

.result.成功 {
  color: #67c23a;
}

.result.失败 {
  color: #f56c6c;
}
</style>
