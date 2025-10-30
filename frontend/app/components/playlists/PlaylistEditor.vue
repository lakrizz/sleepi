<template>
  <v-dialog v-model="isOpen" max-width="900">
    <v-card class="editor-card">
      <v-card-title class="text-h6 d-flex justify-space-between align-center">
        <span>{{ isEditMode ? 'Edit Playlist' : 'New Playlist' }}</span>
        <v-btn icon="mdi-close" variant="text" @click="close" />
      </v-card-title>

      <v-card-text>
        <v-text-field
          v-model="editablePlaylist.name"
          label="Playlist Name"
          variant="outlined"
          density="comfortable"
          class="mb-6"
        />

        <v-row>
          <!-- Current Playlist -->
          <v-col cols="12" md="6">
            <h3 class="text-h6 mb-2">Playlist Files</h3>
            <v-sheet class="playlist-box">
              <draggable
                v-model="editablePlaylist.files"
                item-key="id"
                handle=".drag-handle"
                ghost-class="drag-ghost"
              >
                <template #item="{ element }">
                  <v-list-item class="file-item">
                    <template #prepend>
                      <v-icon class="drag-handle" size="18" icon="mdi-drag" />
                    </template>
                    <v-list-item-title>{{ element.title }}</v-list-item-title>
                    <v-list-item-subtitle>
                      {{ element.artist || 'Unknown Artist' }} — {{ element.length }}
                    </v-list-item-subtitle>
                    <template #append>
                      <v-btn
                        icon="mdi-minus"
                        variant="text"
                        color="var(--color-primary-dark)"
                        @click="removeFile(element)"
                      />
                    </template>
                  </v-list-item>
                </template>
              </draggable>
              <div v-if="!editablePlaylist.files.length" class="empty-text">
                No files yet. Add from library →
              </div>
            </v-sheet>
          </v-col>

          <!-- Library -->
          <v-col cols="12" md="6">
            <h3 class="text-h6 mb-2">Available Files</h3>
            <v-sheet class="playlist-box">
              <v-list-item
                v-for="file in availableFiles"
                :key="file.id"
                class="file-item"
              >
                <v-list-item-title>{{ file.title }}</v-list-item-title>
                <v-list-item-subtitle>
                  {{ file.artist || 'Unknown Artist' }} — {{ file.length }}
                </v-list-item-subtitle>
                <template #append>
                  <v-btn
                    icon="mdi-plus"
                    variant="text"
                    color="var(--color-primary)"
                    @click="addFile(file)"
                  />
                </template>
              </v-list-item>
              <div v-if="!availableFiles.length" class="empty-text">
                All files are in the playlist.
              </div>
            </v-sheet>
          </v-col>
        </v-row>
      </v-card-text>

      <v-card-actions class="d-flex justify-end mt-2">
        <v-btn variant="text" @click="close">Cancel</v-btn>
        <v-btn color="var(--color-primary)" @click="save">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { computed, ref, watch } from 'vue'
import draggable from 'vuedraggable'

const props = defineProps({
  modelValue: Boolean,
  playlist: Object,
})

const emit = defineEmits(['update:modelValue', 'save'])

const isOpen = computed({
  get: () => props.modelValue,
  set: (v) => emit('update:modelValue', v),
})

const editablePlaylist = ref({ id: null, name: '', files: [] })

const isEditMode = computed(() => !!editablePlaylist.value.id)

// Mock global library
const library = ref([
  { id: 101, title: 'Soft Dawn', artist: 'Aurora Fields', length: '03:42' },
  { id: 102, title: 'Sunrise Intro', artist: 'Low Tide', length: '02:10' },
  { id: 103, title: 'Waking Light', artist: 'Low Tide', length: '04:07' },
  { id: 104, title: 'Ocean Breath', artist: 'Somnus', length: '06:03' },
  { id: 105, title: 'Drifting', artist: 'Quiet Forms', length: '05:27' },
  { id: 106, title: 'Through Mist', artist: 'Lunaris', length: '07:45' },
  { id: 107, title: 'Binary Bloom', artist: 'Eleven', length: '04:33' },
])

const availableFiles = computed(() => {
  const usedIds = editablePlaylist.value.files.map(f => f.id)
  return library.value.filter(f => !usedIds.includes(f.id))
})

// Initialize dialog state
watch(() => props.playlist, (pl) => {
  if (pl) editablePlaylist.value = JSON.parse(JSON.stringify(pl))
  else editablePlaylist.value = { id: null, name: '', files: [] }
}, { immediate: true })

function addFile(file) {
  editablePlaylist.value.files.push(file)
}

function removeFile(file) {
  editablePlaylist.value.files = editablePlaylist.value.files.filter(f => f.id !== file.id)
}

function close() {
  emit('update:modelValue', false)
}

function save() {
  emit('save', { ...editablePlaylist.value })
  close()
}
</script>

<style scoped>
.editor-card {
  background-color: var(--color-bg);
  color: var(--color-surface);
  border-radius: 12px;
}

.playlist-box {
  border: 1px solid var(--color-border);
  border-radius: 8px;
  min-height: 250px;
  background-color: var(--color-text);
  overflow-y: auto;
  padding: 0.5rem 0;
}

.file-item {
  border-bottom: 1px solid var(--color-border);
  color: var(--color-dark);
}

.drag-handle {
  cursor: grab;
  color: var(--color-surface-alt);
  margin-right: 0.3rem;
}

.drag-ghost {
  opacity: 0.6;
}

.empty-text {
  text-align: center;
  padding: 1rem;
  color: var(--color-surface-alt);
}

</style>

