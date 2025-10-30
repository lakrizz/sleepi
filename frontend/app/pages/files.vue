<template>
  <v-container fluid class="library-page">
    <v-row class="align-center mb-4">
      <v-col cols="12" class="d-flex align-center justify-space-between">
        <h2 class="text-h5" style="color: var(--color-surface);">Library</h2>
        <v-btn
          color="var(--color-primary)"
          prepend-icon="mdi-upload"
          class="upload-btn"
          @click="openFilePicker"
        >
          Upload File
        </v-btn>
      </v-col>
    </v-row>

    <!-- Dropzone -->
    <v-sheet
      class="dropzone"
      :class="{ 'dropzone--active': isDragActive }"
      @dragover.prevent="onDragOver"
      @dragleave.prevent="onDragLeave"
      @drop.prevent="onDrop"
    >
      <v-icon size="36" color="var(--color-primary)">mdi-cloud-upload</v-icon>
      <p>Drag & drop audio files here, or click “Upload File” above</p>
      <input ref="fileInput" type="file" multiple accept="audio/*,video/*" @change="onFileSelect" hidden />
    </v-sheet>

    <v-divider class="my-6" />

    <!-- Table -->
    <v-table class="library-table">
      <thead>
        <tr>
          <th>Type</th>
          <th>Title</th>
          <th>Artist</th>
          <th>Size</th>
          <th>Length</th>
          <th>Uploaded</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="file in files" :key="file.id">
          <!-- File type column -->
          <td class="type-icon">
            <v-icon :icon="fileIcon(file.type)" size="20" color="var(--color-surface-alt)" />
          </td>

          <!-- Editable title -->
          <td>
            <v-text-field
              v-if="editingId === file.id"
              v-model="file.title"
              dense
              variant="outlined"
              hide-details
              class="rename-input"
            />
            <span v-else>{{ file.title }}</span>
          </td>

          <!-- Editable artist -->
          <td>
            <v-text-field
              v-if="editingId === file.id"
              v-model="file.artist"
              dense
              variant="outlined"
              hide-details
              class="rename-input"
            />
            <span v-else>{{ file.artist || '—' }}</span>
          </td>

          <td>{{ formatSize(file.size) }}</td>
          <td>{{ file.length }}</td>
          <td>{{ new Date(file.uploadedAt).toLocaleString() }}</td>

          <!-- Actions -->
          <td class="actions">
            <v-btn
              v-if="editingId !== file.id"
              icon="mdi-pencil"
              variant="text"
              color="var(--color-primary)"
              @click="startEditing(file.id)"
            />
            <v-btn
              v-else
              icon="mdi-check"
              variant="text"
              color="var(--color-accent)"
              @click="stopEditing"
            />
            <v-btn
              icon="mdi-delete"
              variant="text"
              color="var(--color-primary-dark)"
              @click="deleteFile(file.id)"
            />
          </td>
        </tr>

        <tr v-if="!files.length">
          <td colspan="7" class="empty-text">No files in your library yet.</td>
        </tr>
      </tbody>
    </v-table>
  </v-container>
</template>

<script setup>
import { ref } from 'vue'

const files = ref([
  {
    id: 1,
    title: 'Soft Dawn',
    artist: 'Aurora Fields',
    size: 3950000,
    length: '03:42',
    type: 'audio/mp3',
    uploadedAt: Date.now() - 86400000,
  },
  {
    id: 2,
    title: 'Ocean Breath',
    artist: 'Somnus',
    size: 6300000,
    length: '06:03',
    type: 'audio/wav',
    uploadedAt: Date.now() - 43200000,
  },
  {
    id: 3,
    title: 'Night Motion',
    artist: 'Lunaris',
    size: 7450000,
    length: '07:45',
    type: 'video/mp4',
    uploadedAt: Date.now(),
  },
])

const fileInput = ref(null)
const isDragActive = ref(false)
const editingId = ref(null)

const onDragOver = () => (isDragActive.value = true)
const onDragLeave = () => (isDragActive.value = false)

function openFilePicker() {
  fileInput.value.click()
}

function onDrop(event) {
  isDragActive.value = false
  handleFiles(event.dataTransfer.files)
}

function onFileSelect(event) {
  handleFiles(event.target.files)
}

function handleFiles(fileList) {
  Array.from(fileList).forEach((file) => {
    const newFile = {
      id: Date.now() + Math.random(),
      title: file.name.replace(/\.[^/.]+$/, ''),
      artist: 'Unknown',
      size: file.size,
      length: mockLength(),
      type: file.type || 'audio/mpeg',
      uploadedAt: Date.now(),
    }
    files.value.push(newFile)
  })
}

function mockLength() {
  const m = Math.floor(Math.random() * 5) + 2
  const s = String(Math.floor(Math.random() * 60)).padStart(2, '0')
  return `${m}:${s}`
}

function fileIcon(type) {
  if (type.includes('video')) return 'mdi-file-video'
  if (type.includes('wav')) return 'mdi-music'
  if (type.includes('mp3') || type.includes('audio')) return 'mdi-music-note'
  return 'mdi-file'
}

function startEditing(id) {
  editingId.value = id
}

function stopEditing() {
  editingId.value = null
}

function deleteFile(id) {
  files.value = files.value.filter(f => f.id !== id)
}

function formatSize(bytes) {
  if (bytes < 1024) return bytes + ' B'
  if (bytes < 1048576) return (bytes / 1024).toFixed(1) + ' KB'
  return (bytes / 1048576).toFixed(2) + ' MB'
}
</script>

<style scoped>
.library-page {
  background-color: var(--color-bg);
  min-height: 100vh;
  color: var(--color-surface);
}

.upload-btn {
  background-color: var(--color-primary);
  color: var(--color-text);
  text-transform: none;
  font-weight: 500;
}

/* Dropzone */
.dropzone {
  border: 2px dashed var(--color-border);
  border-radius: 12px;
  background-color: #fff;
  text-align: center;
  padding: 2rem;
  transition: background-color var(--transition-fast), border-color var(--transition-fast);
  cursor: pointer;
}

.dropzone--active {
  border-color: var(--color-primary);
  background-color: rgba(198, 61, 61, 0.08);
}

.dropzone p {
  color: var(--color-surface-alt);
  margin-top: 0.5rem;
}

/* Table */
.library-table {
  width: 100%;
  border-radius: 8px;
  overflow: hidden;
  background-color: #fff;
  color: var(--color-surface);
}

.library-table th {
  text-align: left;
  padding: 0.8rem;
  background-color: var(--color-surface);
  color: var(--color-text);
}

.library-table td {
  padding: 0.8rem;
  border-top: 1px solid var(--color-border);
  vertical-align: middle;
  color: var(--color-surface);
}

.type-icon {
  text-align: center;
}

.rename-input {
  max-width: 200px;
  color: var(--color-surface);
}

.actions {
  display: flex;
  gap: 0.4rem;
}

.empty-text {
  text-align: center;
  padding: 1.5rem;
  color: var(--color-surface-alt);
}
</style>
