<template>
  <v-dialog v-model="isOpen" max-width="500">
    <v-card class="editor-card">
      <v-card-title class="text-h6">
        {{ isEditMode ? 'Edit Alarm' : 'Add Alarm' }}
      </v-card-title>

      <v-card-text class="d-flex flex-column gap-4">
        <!-- Name -->
        <v-text-field
          v-model="editableAlarm.label"
          label="Name"
          variant="outlined"
          density="comfortable"
        />

        <!-- Time -->
        <v-text-field
          v-model="editableAlarm.time"
          label="Time"
          type="time"
          variant="outlined"
          density="comfortable"
        />

        <!-- Repeat Days -->
        <div>
          <div class="text-body-2 mb-1">Repeat Days</div>
          <v-btn-toggle
            v-model="editableAlarm.repeatDays"
            multiple
            variant="default"
            color="var(--color-primary)"
            class="day-toggle"
          >
            <v-btn v-for="(day, idx) in dayNames" :key="idx" :value="idx" size="small"
                          class="day-btn">
              {{ day.short }}
            </v-btn>
          </v-btn-toggle>
        </div>

        <!-- Warmup Duration -->
        <div>
          <div class="text-body-2 mb-1">Warmup Duration: {{ editableAlarm.warmupDuration }} min</div>
          <v-slider
            v-model="editableAlarm.warmupDuration"
            :min="0"
            :max="60"
            :step="5"
            color="var(--color-primary)"
          />
        </div>

        <!-- LED Color Picker -->
        <div class="mb-8">
          <div class="text-body-2 mb-1">LED Color</div>
          <v-color-picker
            v-model="editableColor"
            mode="rgb"
            hide-mode-switch
            canvas-height="120"
          />
        </div>

        <!-- Playlist Selector -->
        <v-select
          v-model="editableAlarm.musicTarget"
          label="Playlist"
          :items="playlists"
          item-title="label"
          item-value="id"
          variant="outlined"
          density="comfortable"
        />
      </v-card-text>

      <v-card-actions class="d-flex justify-end">
        <v-btn variant="text" @click="close">Cancel</v-btn>
        <v-btn color="var(--color-primary)" @click="save">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { computed, ref, watch } from 'vue'

const props = defineProps({
  modelValue: { type: Boolean, default: false },   // open/close
  alarm: { type: Object, default: null },          // existing alarm or null
  playlists: { type: Array, default: () => [] },   // [{ id, label }]
})

const emit = defineEmits(['update:modelValue', 'save'])

const isOpen = computed({
  get: () => props.modelValue,
  set: v => emit('update:modelValue', v),
})

const dayNames = [
  { short: 'M', full: 'Monday' },
  { short: 'T', full: 'Tuesday' },
  { short: 'W', full: 'Wednesday' },
  { short: 'T', full: 'Thursday' },
  { short: 'F', full: 'Friday' },
  { short: 'S', full: 'Saturday' },
  { short: 'S', full: 'Sunday' },
]

// local editable state
const editableAlarm = ref(makeEmptyAlarm())
const editableColor  = ref({ r: 255, g: 200, b: 150 })

const isEditMode = computed(() => !!editableAlarm.value.id)

// --- helpers ---
function makeEmptyAlarm () {
  return {
    id: null,
    label: '',
    time: '07:00',
    repeatDays: [],
    warmupDuration: 10,
    ledTarget: { r: 255, g: 200, b: 150 },
    musicTarget: '',
    enabled: true,
  }
}
function cloneAlarm(a) {
  return a ? JSON.parse(JSON.stringify(a)) : makeEmptyAlarm()
}
function initFromProps() {
  const next = cloneAlarm(props.alarm)
  editableAlarm.value = next
  editableColor.value = {
    r: next.ledTarget?.r ?? 255,
    g: next.ledTarget?.g ?? 200,
    b: next.ledTarget?.b ?? 150,
  }
}

// keep ledTarget in sync when user changes the picker
watch(editableColor, (rgb) => {
  editableAlarm.value.ledTarget = { r: rgb.r, g: rgb.g, b: rgb.b }
})

// re-init when: dialog opens or incoming alarm changes
watch(() => props.alarm, initFromProps, { immediate: true, deep: true })
watch(isOpen, (open) => { if (open) initFromProps() })

function close() { isOpen.value = false }

function save() {
  const payload = cloneAlarm(editableAlarm.value)
  // if editing, keep id; if new, let parent assign or do it here:
  // payload.id ||= Date.now()
  emit('save', payload)
  close()
}
</script>

<style scoped>
.editor-card {
  background-color: var(--color-bg);
  color: var(--color-surface);
  border-radius: 12px;
}

.day-toggle {
  display: flex;
  justify-content: space-between;
  width: 100%;
}

.v-color-picker {
  border-radius: 8px;
  border: 1px solid var(--color-border);
}

        .day-btn {
background-color: var(--inactive-bg);
        }
</style>
