<template>
  <v-container fluid class="sleepscapes-page">
    <v-row class="align-center mb-4">
      <v-col cols="12" class="d-flex align-center justify-space-between">
        <h2 class="text-h5" style="color: var(--color-surface);">Sleepscapes</h2>

        <v-btn
          color="var(--color-primary)"
          prepend-icon="mdi-plus"
          class="add-btn"
          @click="onAdd"
        >
          Add Sleepscape
        </v-btn>
      </v-col>
    </v-row>

    <!-- List of sleepscapes -->
    <v-table class="sleepscapes-table">
      <thead>
        <tr>
          <th>Name</th>
          <th>Audio Source</th>
          <th>Type</th>
          <th>Expression</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="sc in sleepscapes" :key="sc.id">
          <td>{{ sc.name }}</td>
          <td>{{ sc.sourceName || '—' }}</td>
          <td>{{ sc.sourceType }}</td>
          <td>
            <code class="expr-snippet">
              {{ truncate(sc.expression, 50) }}
            </code>
          </td>
          <td class="actions">
            <v-btn
              icon="mdi-pencil"
              variant="text"
              color="var(--color-primary)"
              @click="onEdit(sc)"
            />
            <v-btn
              icon="mdi-delete"
              variant="text"
              color="var(--color-primary-dark)"
              @click="onDelete(sc.id)"
            />
          </td>
        </tr>
        <tr v-if="!sleepscapes.length">
          <td colspan="5" class="empty-text">No sleepscapes yet.</td>
        </tr>
      </tbody>
    </v-table>

    <!-- Add/Edit dialog -->
    <v-dialog v-model="showEditor" max-width="900">
      <v-card class="editor-card">
        <v-card-title class="d-flex justify-space-between align-center">
          <span>{{ editing ? 'Edit Sleepscape' : 'New Sleepscape' }}</span>
          <v-btn icon="mdi-close" variant="text" @click="closeEditor" />
        </v-card-title>

        <v-card-text>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="editable.name"
                label="Name"
                variant="outlined"
                density="comfortable"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-select
                v-model="editable.sourceType"
                :items="['playlist', 'file']"
                label="Source Type"
                variant="outlined"
                density="comfortable"
              />
            </v-col>

            <v-col cols="12" v-if="editable.sourceType">
              <v-text-field
                v-model="editable.sourceName"
                :label="editable.sourceType === 'playlist' ? 'Playlist Name' : 'File Name'"
                variant="outlined"
                density="comfortable"
              />
            </v-col>

            <!-- Expression Editor -->
            <v-col cols="12">
              <v-textarea
                v-model="editable.expression"
                label="LED Expression"
                variant="outlined"
                rows="8"
                hint="Use Sleepi Light Expressions (Expr syntax). Example: for(i,led_count(),set_led(i,sin(t+i/10)*127+128,0,255))"
              />
            </v-col>

            <!-- Preview placeholder -->
            <v-col cols="12" class="preview-area">
              <v-card variant="tonal" class="preview-card">
                <v-card-text>
                  <div class="text-body-1 mb-2">Preview (future LED simulation)</div>
                  <div class="led-preview">
                    <div
                      v-for="i in previewLeds"
                      :key="i"
                      class="led-dot"
                      :style="{ backgroundColor: ledColor(i) }"
                    />
                  </div>
                </v-card-text>
              </v-card>
            </v-col>
          </v-row>
        </v-card-text>

        <v-card-actions class="d-flex justify-end">
          <v-btn variant="text" @click="closeEditor">Cancel</v-btn>
          <v-btn color="var(--color-primary)" @click="saveSleepscape">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup>
import { ref } from 'vue'

const sleepscapes = ref([
  {
    id: 1,
    name: 'Nebula Drift',
    sourceType: 'playlist',
    sourceName: 'Deep Sleep Waves',
    expression: 'for(i,led_count(),set_led(i,sin(t+i/10)*127+128,20,255))',
  },
  {
    id: 2,
    name: 'Pulsar Blue',
    sourceType: 'file',
    sourceName: 'blue_drift.mp3',
    expression: 'r=127+127*sin(t);for(i,led_count(),set_led(i,r,0,255-r))',
  },
])

const showEditor = ref(false)
const editing = ref(false)
const editable = ref({
  id: null,
  name: '',
  sourceType: '',
  sourceName: '',
  expression: '',
})

const onAdd = () => {
  editing.value = false
  editable.value = { id: null, name: '', sourceType: '', sourceName: '', expression: '' }
  showEditor.value = true
}

const onEdit = (sc) => {
  editable.value = JSON.parse(JSON.stringify(sc))
  editing.value = true
  showEditor.value = true
}

const onDelete = (id) => {
  sleepscapes.value = sleepscapes.value.filter(s => s.id !== id)
}

const closeEditor = () => (showEditor.value = false)

const saveSleepscape = () => {
  if (editing.value) {
    const idx = sleepscapes.value.findIndex(s => s.id === editable.value.id)
    if (idx !== -1) sleepscapes.value[idx] = { ...editable.value }
  } else {
    editable.value.id = Date.now()
    sleepscapes.value.push({ ...editable.value })
  }
  showEditor.value = false
}

function truncate(str, n) {
  return str.length > n ? str.substring(0, n) + '…' : str
}

/* --- Mock LED preview --- */
const previewLeds = 24
function ledColor(i) {
  const hue = (i * 15 + Date.now() / 100) % 360
  return `hsl(${hue}, 80%, 60%)`
}
</script>

<style scoped>
.sleepscapes-page {
  background-color: var(--color-bg);
  min-height: 100vh;
  color: var(--color-surface);
}

.sleepscapes-table {
  width: 100%;
  background-color: #fff;
  color: var(--color-surface);
  border-radius: 8px;
  overflow: hidden;
}

.sleepscapes-table th {
  text-align: left;
  padding: 0.8rem;
  background-color: var(--color-surface);
  color: var(--color-text);
}

.sleepscapes-table td {
  padding: 0.8rem;
  border-top: 1px solid var(--color-border);
  vertical-align: middle;
}

.expr-snippet {
  font-size: 0.8rem;
  background-color: rgba(31,44,58,0.06);
  padding: 2px 6px;
  border-radius: 4px;
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

.editor-card {
  background-color: var(--color-bg);
  color: var(--color-surface);
  border-radius: 12px;
}

.preview-area {
  margin-top: 1rem;
}

.preview-card {
  border-radius: 10px;
  text-align: center;
}

.led-preview {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 6px;
}

.led-dot {
  width: 16px;
  height: 16px;
  border-radius: 50%;
  box-shadow: 0 0 4px rgba(0,0,0,0.3);
  transition: background-color 0.2s ease;
}

.add-btn {
  background-color: var(--color-primary);
  color: var(--color-text);
  text-transform: none;
  font-weight: 500;
}
</style>

