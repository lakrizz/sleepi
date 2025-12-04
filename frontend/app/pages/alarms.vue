<script setup>
import { ref } from "vue";
import AlarmEditor from "~/components/alarms/AlarmEditor.vue";
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { AlarmService } from "@/api/gen/v1/alarms_pb";

// state
const alarms = ref([]);
const selectedAlarm = ref(null);
const showEditor = ref(false);

// client
const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});
const client = createClient(AlarmService, transport);

// ---- Load alarms on mount ----
const loadAlarms = async () => {
  const resp = await client.listAlarms({});
  console.log(resp);
  alarms.value = resp.alarms ?? [];
};
loadAlarms();

// ---- Open / Edit / Add ----
const onAddAlarm = () => {
  selectedAlarm.value = null; // new entry
  showEditor.value = true;
};

const editAlarm = async (alarm) => {
  // load from backend to avoid stale data
  const resp = await client.getAlarm({ id: alarm.id });
  selectedAlarm.value = structuredClone(resp.alarm);
  showEditor.value = true;
};

// ---- Save (Create or Update) ----
const saveAlarm = async (alarm) => {
  if (!alarm.id) {
    // CREATE
    console.log("wanting to save", alarm)
    client.createAlarm({alarm}, (err, res) => {
      console.log("err", err, "res", res)
      if (!err) {
        alarms.value.push(resp.alarm);
      }
    })
    
  } else {
    // UPDATE
    const resp = await client.updateAlarm({ alarm });
    const idx = alarms.value.findIndex((a) => a.id === alarm.id);
    if (idx !== -1) alarms.value[idx] = resp.alarm;
  }

  showEditor.value = false;
  selectedAlarm.value = null;
};

// ---- Delete ----
const deleteAlarm = async (id) => {
  await client.deleteAlarm({ id });
  alarms.value = alarms.value.filter((a) => a.id !== id);
};

// ---- Formatting helpers ----
const formatTime = (t) => t;

const formatRepeat = (days) => {
  const d = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
  return days.length ? days.map((x) => d[x]).join(", ") : "One-time";
};

// playlists (static example)
const playlistOptions = [
  { id: "1c0f8591-28c3-4a51-8088-99ce60ef5665", label: "Morning Vibes" },
  { id: "669aa8ca-c3a6-4cd9-abe9-d449f8cff60d", label: "Energetic Boost" },
  { id: "a4d2693a-61a3-4c95-aafc-bcf2fee79735", label: "Relax & Dream" },
];
</script>



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
          :style="{
            borderLeft: alarm.enabled
              ? '4px solid var(--color-primary)'
              : '4px solid var(--color-border)'
          }"
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
              Warmup:
              {{ alarm.warmupDuration?.seconds ?? 0 }} sec
            </div>

            <div class="detail led-strip">
              <v-icon icon="mdi-lightbulb-on" size="18" />
              <span>LED:</span>

              <v-tooltip
                :text="`rgb(${alarm.ledTarget?.r}, ${alarm.ledTarget?.g}, ${alarm.ledTarget?.b})`"
                location="top"
              >
                <template #activator="{ props }">
                  <div
                    v-bind="props"
                    class="led-color"
                    :style="{
                      backgroundColor: `rgb(${alarm.ledTarget?.r}, ${alarm.ledTarget?.g}, ${alarm.ledTarget?.b})`
                    }"
                  />
                </template>
              </v-tooltip>
            </div>

            <div class="detail">
              <v-icon icon="mdi-music" size="18" />
              Playlist: {{ alarm.playableId }}
            </div>
          </v-card-text>

          <v-card-actions>
            <v-btn
              icon="mdi-pencil"
              variant="text"
              @click="editAlarm(alarm)"
              color="var(--color-primary)"
            />
            <v-btn
              icon="mdi-delete"
              variant="text"
              @click="deleteAlarm(alarm.id)"
              color="var(--color-primary-dark)"
            />
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

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
