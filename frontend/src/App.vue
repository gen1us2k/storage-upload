<template>
  <div @drop="handleDrop" @dragover.prevent="allowDrop" @click="openFileInput" class="drop-area">
    Drag files here or click to select
    <input type="file" ref="fileInput" style="display: none" @change="handleFileChange" multiple />
  </div>

  <ul>
    <li v-for="file in files" :key="file.id">
      <span>
        {{ file.name }}
        <span v-if="file.progress !== undefined">- {{ file.progress }}%</span>
      </span>
      <button @click="removeFile(file.id)" :disabled="file.uploading">Remove</button>
    </li>
  </ul>

  <div v-if="error" class="error-message">{{ error }}</div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

type FileWithProgress = {
  id: number
  name: string
  progress?: number
  uploading?: boolean
}

const files = ref<FileWithProgress[]>([])
const fileInput = ref<HTMLInputElement>()
const error = ref<string>('')

const handleDrop = (event: DragEvent) => {
  event.preventDefault()
  const droppedFiles = event.dataTransfer?.files
  handleFiles(droppedFiles as FileList)
}

const allowDrop = (event: DragEvent) => {
  event.preventDefault()
}

const openFileInput = () => {
  const input = fileInput.value as HTMLInputElement
  input.click()
}

const handleFileChange = () => {
  const input = fileInput.value as HTMLInputElement
  const selectedFiles = input.files
  handleFiles(selectedFiles)
}

const handleFiles = async (fileList: FileList | null) => {
  if (!fileList) return

  const uploadPromises = []
  error.value = ''

  for (let i = 0; i < fileList.length; i++) {
    const file = fileList[i]
    const fileWithProgress: FileWithProgress = {
      id: files.value.length + 1,
      name: file.name,
      progress: 0,
      uploading: true
    }
    files.value.push(fileWithProgress)

    const uploadPromise = uploadFile(file, fileWithProgress)
    uploadPromises.push(uploadPromise)

    try {
      await Promise.all(uploadPromises)
    } catch (uploadError: any) {
      fileWithProgress.uploading = false
      error.value = uploadError.message
    }
  }
}

const uploadFile = async (file: File, fileWithProgress: FileWithProgress) => {
  const formData = new FormData()
  formData.append('file', file)

  try {
    const response = await fetch('/api/upload', {
      method: 'POST',
      body: formData
    })

    fileWithProgress.progress = 100
    fileWithProgress.uploading = false

    // Handle successful upload response if needed
    await response.json()
  } catch (uploadError: any) {
    throw new Error(`Error uploading file ${file.name}: ${uploadError.message}`)
  }
}

const removeFile = (id: number) => {
  files.value = files.value.filter((file) => file.id !== id)
}
</script>
