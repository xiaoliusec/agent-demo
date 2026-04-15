<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElForm, ElFormItem, ElInput, ElSelect, ElOption, ElButton, ElScrollbar } from 'element-plus'

interface Config {
  target_url: string
  username: string
  password: string
  ai_provider: string
  api_key: string
  api_endpoint: string
  model_name: string
  max_cases: number
}

interface TestStatus {
  state: string
  current_step: number
  total_cases: number
  passed_cases: number
  failed_cases: number
  current_url: string
  last_error: string
}

const config = ref<Config>({
  target_url: '',
  username: '',
  password: '',
  ai_provider: 'openai',
  api_key: '',
  api_endpoint: '',
  model_name: 'gpt-4',
  max_cases: 100
})

const status = ref<TestStatus>({
  state: 'idle',
  current_step: 0,
  total_cases: 0,
  passed_cases: 0,
  failed_cases: 0,
  current_url: '',
  last_error: ''
})

const logs = ref<string[]>([])
const aiResponses = ref<any[]>([])
const goal = ref('')
const isRunning = ref(false)

const providers = [
  { value: 'openai', label: 'OpenAI' },
  { value: 'claude', label: 'Claude' },
  { value: 'deepseek', label: 'DeepSeek' },
  { value: 'doubao', label: '豆包' },
  { value: 'ollama', label: 'Ollama (本地)' },
  { value: 'vllm', label: 'vLLM (本地)' }
]

const startTest = async () => {
  if (!goal.value) {
    window.message.warning('请输入测试目标')
    return
  }
  
  logs.value.push(`[${new Date().toLocaleTimeString()}] 开始测试: ${goal.value}`)
  isRunning.value = true
  
  try {
    // 模拟开始测试
    // @ts-ignore
    if (window.go) {
      // @ts-ignore
      await window.go.testflow.App.StartTest(goal.value)
    }
  } catch (e) {
    logs.value.push(`[ERROR] ${e}`)
  }
}

const stopTest = async () => {
  try {
    // @ts-ignore
    if (window.go) {
      // @ts-ignore
      await window.go.testflow.App.StopTest()
    }
  } catch (e) {
    logs.value.push(`[ERROR] ${e}`)
  }
  isRunning.value = false
}

const loadConfig = async () => {
  try {
    // @ts-ignore
    if (window.go) {
      // @ts-ignore
      const savedConfig = await window.go.testflow.App.GetConfig()
      if (savedConfig) {
        config.value = savedConfig
      }
    }
  } catch (e) {
    console.error('Failed to load config:', e)
  }
}

onMounted(() => {
  loadConfig()
  
  // @ts-ignore
  if (window.runtime) {
    // @ts-ignore
    window.runtime.Events.On('ai_decision', (data: any) => {
      aiResponses.value.push(data)
      logs.value.push(`[AI] ${data.thought_summary}`)
    })
    
    // @ts-ignore
    window.runtime.Events.On('step_complete', (data: any) => {
      logs.value.push(`[STEP ${data.step}] 完成`)
    })
    
    // @ts-ignore
    window.runtime.Events.On('test_complete', (data: string) => {
      logs.value.push(`[完成] ${data}`)
      isRunning.value = false
    })
    
    // @ts-ignore
    window.runtime.Events.On('test_case_generated', (data: any) => {
      logs.value.push(`[用例] ${data.name || '新用例'}`)
    })
  }
})
</script>

<template>
  <div class="left-panel-container">
    <div class="header">
      <h2>TestFlow</h2>
      <span class="subtitle">AI 自动化测试工具</span>
    </div>
    
    <el-scrollbar class="scroll-area">
      <div class="config-section">
        <h3>系统配置</h3>
        <el-form label-position="top" size="small">
          <el-form-item label="目标 URL">
            <el-input v-model="config.target_url" placeholder="https://example.com" />
          </el-form-item>
          
          <el-form-item label="用户名">
            <el-input v-model="config.username" placeholder="可选" />
          </el-form-item>
          
          <el-form-item label="密码">
            <el-input v-model="config.password" type="password" placeholder="可选" show-password />
          </el-form-item>
          
          <el-form-item label="AI 提供商">
            <el-select v-model="config.ai_provider" style="width: 100%">
              <el-option
                v-for="p in providers"
                :key="p.value"
                :label="p.label"
                :value="p.value"
              />
            </el-select>
          </el-form-item>
          
          <el-form-item label="API Key">
            <el-input v-model="config.api_key" type="password" placeholder="可选（本地部署无需填写）" show-password />
          </el-form-item>
          
          <el-form-item label="API Endpoint">
            <el-input v-model="config.api_endpoint" placeholder="可选，自定义 API 地址" />
          </el-form-item>
          
          <el-form-item label="模型名称">
            <el-input v-model="config.model_name" placeholder="gpt-4" />
          </el-form-item>
          
          <el-form-item label="最大用例数">
            <el-input-number v-model="config.max_cases" :min="1" :max="1000" />
          </el-form-item>
        </el-form>
      </div>
      
      <div class="test-section">
        <h3>测试控制</h3>
        <el-input
          v-model="goal"
          type="textarea"
          :rows="3"
          placeholder="输入测试目标，如：测试用户登录功能"
          :disabled="isRunning"
        />
        <div class="button-group">
          <el-button 
            type="primary" 
            @click="startTest" 
            :disabled="isRunning || !goal"
            style="flex: 1"
          >
            开始测试
          </el-button>
          <el-button 
            type="danger" 
            @click="stopTest" 
            :disabled="!isRunning"
            style="flex: 1"
          >
            停止
          </el-button>
        </div>
      </div>
      
      <div class="status-section">
        <h3>状态监控</h3>
        <div class="status-grid">
          <div class="status-item">
            <span class="label">状态</span>
            <span class="value" :class="status.state">{{ status.state }}</span>
          </div>
          <div class="status-item">
            <span class="label">步数</span>
            <span class="value">{{ status.current_step }}</span>
          </div>
          <div class="status-item">
            <span class="label">用例</span>
            <span class="value">{{ status.total_cases }}</span>
          </div>
          <div class="status-item">
            <span class="label">通过</span>
            <span class="value passed">{{ status.passed_cases }}</span>
          </div>
          <div class="status-item">
            <span class="label">失败</span>
            <span class="value failed">{{ status.failed_cases }}</span>
          </div>
        </div>
      </div>
      
      <div class="ai-response-section">
        <h3>AI 响应</h3>
        <el-scrollbar height="200px">
          <div v-for="(resp, idx) in aiResponses" :key="idx" class="ai-response-item">
            <div class="thought">{{ resp.thought_summary }}</div>
            <div v-if="resp.next_action" class="action">
              {{ resp.next_action.tool_name }}
            </div>
          </div>
        </el-scrollbar>
      </div>
      
      <div class="log-section">
        <h3>日志</h3>
        <el-scrollbar height="150px">
          <div v-for="(log, idx) in logs" :key="idx" class="log-item">
            {{ log }}
          </div>
        </el-scrollbar>
      </div>
    </el-scrollbar>
  </div>
</template>

<style scoped>
.left-panel-container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.header {
  padding: 16px;
  background: linear-gradient(135deg, #409eff, #66b1ff);
  color: white;
}

.header h2 {
  margin: 0;
  font-size: 20px;
}

.header .subtitle {
  font-size: 12px;
  opacity: 0.9;
}

.scroll-area {
  flex: 1;
  padding: 16px;
}

h3 {
  margin: 0 0 12px 0;
  font-size: 14px;
  color: #409eff;
  border-bottom: 1px solid #e0e0e0;
  padding-bottom: 8px;
}

.config-section,
.test-section,
.status-section,
.ai-response-section,
.log-section {
  margin-bottom: 20px;
}

.button-group {
  display: flex;
  gap: 8px;
  margin-top: 12px;
}

.status-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 8px;
}

.status-item {
  background: #f5f7fa;
  padding: 8px 12px;
  border-radius: 4px;
}

.status-item .label {
  font-size: 11px;
  color: #909399;
  display: block;
}

.status-item .value {
  font-size: 16px;
  font-weight: bold;
  color: #303133;
}

.status-item .value.running {
  color: #409eff;
}

.status-item .value.completed {
  color: #67c23a;
}

.status-item .value.error {
  color: #f56c6c;
}

.status-item .value.passed {
  color: #67c23a;
}

.status-item .value.failed {
  color: #f56c6c;
}

.ai-response-item {
  background: #f5f7fa;
  padding: 8px 12px;
  margin-bottom: 8px;
  border-radius: 4px;
  font-size: 12px;
}

.ai-response-item .thought {
  color: #303133;
}

.ai-response-item .action {
  color: #409eff;
  margin-top: 4px;
  font-family: monospace;
}

.log-item {
  font-size: 11px;
  color: #606266;
  padding: 2px 0;
  font-family: monospace;
}
</style>
