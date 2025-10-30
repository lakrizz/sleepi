<template>
  <v-container fluid class="alarm-list">
    <AlarmEditor
      v-model="showEditor"
      :alarm="selectedAlarm"
      :key="selectedAlarm ? selectedAlarm.id : 'new'"
      :playlists="playlistOptions"
      @save="saveAlarm"
    />

    <v-row>
      <v-col cols="12" class="d-flex align-center justify-space-between mb-4">
        <h2 class="text-h3" style="color: var(--color-surface);">Alarms</h2>

        <v-btn
          color="var(--color-primary)"
          prepend-icon="mdi-plus"
          class="add-btn"
          @click="onAddAlarm"
        >
          Add Alarm
        </v-btn>
      </v-col>
    </v-row>

    <v-row>
      <v-col
        v-for="alarm in alarms"
        :key="alarm.id"
        cols="12"
        md="6"
        lg="4"
      >
        <v-card
          class="alarm-card"
          :style="{ borderLeft: alarm.enabled ? '4px solid var(--color-primary)' : '4px solid var(--color-border)' }"
        >
          <v-card-title class="d-flex align-center justify-space-between">
            <div>
              <div class="alarm-time">{{ formatTime(alarm.time) }}</div>
              <div class="alarm-label">{{ alarm.label }}</div>
            </div>
            <v-switch v-model="alarm.enabled" color="var(--color-primary)" hide-details />
          </v-card-title>

          <v-card-subtitle>
            {{ formatRepeat(alarm.repeatDays) }}
          </v-card-subtitle>

          <v-card-text class="alarm-details">
            <div class="detail">
              <v-icon icon="mdi-timer-sand" size="18" />
              Warmup: {{ alarm.warmupDuration }} min
            </div>

            <!-- LED Color Swatch -->
            <div class="detail led-strip">
              <v-icon icon="mdi-lightbulb-on" size="18" />
              <span>LED:</span>

              <v-tooltip text="rgb({{ alarm.ledTarget.r }}, {{ alarm.ledTarget.g }}, {{ alarm.ledTarget.b }})" location="top">
                <template #activator="{ props }">
                  <div
                    v-bind="props"
                    class="led-color"
                    :style="{
                      backgroundColor: `rgb(${alarm.ledTarget.r}, ${alarm.ledTarget.g}, ${alarm.ledTarget.b})`
                    }"
                  />
                </template>
              </v-tooltip>
            </div>

            <div class="detail">
              <v-icon icon="mdi-music" size="18" />
              Playlist: {{ alarm.musicTarget }}
            </div>
          </v-card-text>

          <v-card-actions>
            <v-btn icon="mdi-pencil" variant="text" @click="editAlarm(alarm)" color="var(--color-primary)" />
            <v-btn icon="mdi-delete" variant="text" @click="deleteAlarm(alarm.id)" color="var(--color-primary-dark)" />
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref } from 'vue'
import AlarmEditor from '~/components/alarms/AlarmEditor.vue'

const showEditor = ref(false)
const selectedAlarm = ref(null)

const playlistOptions = [
  { id: 'playlist_morning_vibes', label: 'Morning Vibes' },
  { id: 'playlist_energetic', label: 'Energetic Boost' },
  { id: 'playlist_relax', label: 'Relax & Dream' },
]

const onAddAlarm = () => {
  selectedAlarm.value = null
  showEditor.value = true
}

const editAlarm = (alarm) => {
  selectedAlarm.value = { ...alarm }
  showEditor.value = true
}

const saveAlarm = (newAlarm) => {
console.log("newAlarm", newAlarm)
  if (newAlarm.id) {
  console.log("updating", "newAlarm")
    // update existing
    const idx = alarms.value.findIndex(a => a.id === newAlarm.id)
    if (idx !== -1) alarms.value[idx] = newAlarm
  } else {
  console.log("creating", newAlarm)
    // create new
    newAlarm.id = Date.now()
    alarms.value.push(newAlarm)
  }
}

const alarms = ref([
  {
    id: 1,
    label: 'Morning Routine',
    time: '07:30',
    repeatDays: [1, 2, 3, 4, 5],
    enabled: true,
    warmupDuration: 15,
    ledTarget: { r: 255, g: 180, b: 120 },
    musicTarget: 'playlist_morning_vibes',
  },
  {
    id: 2,
    label: 'Workout',
    time: '18:00',
    repeatDays: [1, 3, 5],
    enabled: false,
    warmupDuration: 5,
    ledTarget: { r: 100, g: 255, b: 180 },
    musicTarget: 'playlist_energetic',
  },
  {
    id: 3,
    label: 'Weekend Chill',
    time: '09:00',
    repeatDays: [6, 0],
    enabled: true,
    warmupDuration: 20,
    ledTarget: { r: 255, g: 120, b: 120 },
    musicTarget: 'playlist_relax',
  },
])


const deleteAlarm = (id) => {
  console.log('Delete alarm id:', id)
}

const formatTime = (time) => {
  const [h, m] = time.split(':')
  return `${h}:${m}`
}

const formatRepeat = (days) => {
  const dayNames = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']
  return days && days.length ? days.map(d => dayNames[d]).join(', ') : 'One-time'
}
</script>

<style scoped>
.alarm-list {
  background-color: var(--color-bg);
  min-height: 100vh;
}

.alarm-card {
  background-color: var(--color-text);
  color: var(--color-surface);
  border-radius: 12px;
  transition: box-shadow var(--transition-fast);
}

.alarm-card:hover {
  box-shadow: 0 4px 10px rgba(31, 44, 58, 0.15);
}

.alarm-time {
  font-size: 1.6rem;
  font-weight: 600;
}

.alarm-label {
  font-size: 0.95rem;
  color: var(--color-surface-alt);
}

.add-btn {
  background-color: var(--color-primary);
  color: var(--color-bg);
  text-transform: none;
  font-weight: 500;
}

.alarm-details {
  display: flex;
  flex-direction: column;
  gap: 0.3rem;
  margin-top: 0.5rem;
}

.detail {
  display: flex;
  align-items: center;
  gap: 0.4rem;
  font-size: 0.9rem;
  color: var(--color-surface-alt);
}

/* LED color swatch strip */
.led-strip {
  display: flex;
  align-items: center;
  gap: 0.6rem;
}

.led-color {
  width: 48px;
  height: 12px;
  border-radius: 4px;
  border: 1px solid var(--color-border);
  cursor: pointer;
}
</style>
