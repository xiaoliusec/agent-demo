<script setup lang="ts">
import { ref } from 'vue'
import LeftPanel from './components/LeftPanel.vue'
import CenterPanel from './components/CenterPanel.vue'
import RightPanel from './components/RightPanel.vue'

const leftWidth = ref(350)
const rightWidth = ref(350)
const isDraggingLeft = ref(false)
const isDraggingRight = ref(false)

const startDragLeft = () => {
  isDraggingLeft.value = true
}

const startDragRight = () => {
  isDraggingRight.value = true
}

const onMouseMove = (e: MouseEvent) => {
  if (isDraggingLeft.value) {
    const newWidth = Math.max(250, Math.min(600, e.clientX))
    leftWidth.value = newWidth
  }
  if (isDraggingRight.value) {
    const newWidth = Math.max(250, Math.min(600, window.innerWidth - e.clientX))
    rightWidth.value = newWidth
  }
}

const onMouseUp = () => {
  isDraggingLeft.value = false
  isDraggingRight.value = false
}
</script>

<template>
  <div class="app-container" @mousemove="onMouseMove" @mouseup="onMouseUp">
    <div class="left-panel" :style="{ width: leftWidth + 'px' }">
      <LeftPanel />
    </div>
    
    <div 
      class="resize-handle left" 
      @mousedown="startDragLeft"
      :class="{ active: isDraggingLeft }"
    ></div>
    
    <div class="center-panel">
      <CenterPanel />
    </div>
    
    <div 
      class="resize-handle right" 
      @mousedown="startDragRight"
      :class="{ active: isDraggingRight }"
    ></div>
    
    <div class="right-panel" :style="{ width: rightWidth + 'px' }">
      <RightPanel />
    </div>
  </div>
</template>

<style scoped>
.app-container {
  display: flex;
  height: 100vh;
  width: 100vw;
  overflow: hidden;
  background: #f5f5f5;
}

.left-panel {
  background: white;
  border-right: 1px solid #e0e0e0;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.center-panel {
  flex: 1;
  background: #fafafa;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.right-panel {
  background: white;
  border-left: 1px solid #e0e0e0;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.resize-handle {
  width: 4px;
  background: #e0e0e0;
  cursor: col-resize;
  transition: background 0.2s;
}

.resize-handle:hover,
.resize-handle.active {
  background: #409eff;
}
</style>
