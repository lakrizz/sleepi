<template>
  <v-container fluid class="system-page">
    <h2 class="text-h5 mb-6" style="color: var(--color-surface);">System Information</h2>

    <!-- General Info -->
    <v-card class="info-card mb-6">
      <v-card-title class="text-h6">General</v-card-title>
      <v-card-text>
        <v-row>
          <v-col cols="12" md="6">
            <div class="info-line"><strong>App Version:</strong> {{ system.version }}</div>
            <div class="info-line"><strong>Build Date:</strong> {{ system.buildDate }}</div>
            <div class="info-line"><strong>Uptime:</strong> {{ formatUptime(system.uptimeSeconds) }}</div>
            <div class="info-line"><strong>Operating System:</strong> {{ system.os }}</div>
            <div class="info-line"><strong>Architecture:</strong> {{ system.arch }}</div>
          </v-col>

          <v-col cols="12" md="6">
            <div class="info-line">
              <strong>Update Status:</strong>
              <v-chip
                v-if="system.updateAvailable"
                color="var(--color-primary)"
                text-color="var(--color-text)"
                label
                size="small"
              >
                Update Available
              </v-chip>
              <v-chip v-else color="var(--color-accent)" text-color="var(--color-text)" label size="small">
                Up to Date
              </v-chip>
            </div>
            <div v-if="system.updateAvailable" class="mt-2">
              <div><strong>New Version:</strong> {{ system.newVersion }}</div>
              <div><strong>Release Notes:</strong></div>
              <ul class="release-notes">
                <li v-for="(note, i) in system.releaseNotes" :key="i">{{ note }}</li>
              </ul>
              <v-btn color="var(--color-primary)" class="mt-2">Install Update</v-btn>
            </div>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>

    <!-- Disk & Memory -->
    <v-card class="info-card mb-6">
      <v-card-title class="text-h6">System Resources</v-card-title>
      <v-card-text>
        <div class="resource">
          <div class="label">Disk Usage</div>
          <v-progress-linear
            :model-value="system.diskUsedPct"
            height="12"
            color="var(--color-primary)"
            rounded
          />
          <div class="caption">{{ system.diskUsed }} / {{ system.diskTotal }}</div>
        </div>

        <div class="resource mt-4">
          <div class="label">Memory Usage</div>
          <v-progress-linear
            :model-value="system.memoryUsedPct"
            height="12"
            color="var(--color-accent)"
            rounded
          />
          <div class="caption">{{ system.memoryUsed }} / {{ system.memoryTotal }}</div>
        </div>

        <div class="resource mt-4">
          <div class="label">CPU Usage</div>
          <v-progress-linear
            :model-value="system.cpuLoadPct"
            height="12"
            color="var(--color-primary-dark)"
            rounded
          />
          <div class="caption">{{ system.cpuLoadPct.toFixed(1) }}%</div>
        </div>
      </v-card-text>
    </v-card>

    <!-- App Statistics -->
    <v-card class="info-card mb-6">
      <v-card-title class="text-h6">App Statistics</v-card-title>
      <v-card-text>
        <v-row>
          <v-col cols="12" md="3" v-for="stat in stats" :key="stat.label">
            <v-card class="stat-card" variant="tonal">
              <v-card-text class="d-flex flex-column align-center justify-center py-4">
                <v-icon :icon="stat.icon" size="32" class="mb-2" />
                <div class="text-h5">{{ stat.value }}</div>
                <div class="text-body-2">{{ stat.label }}</div>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>

    <!-- Environment -->
    <v-card class="info-card">
      <v-card-title class="text-h6">Environment</v-card-title>
      <v-card-text>
        <div class="info-line"><strong>Network:</strong> {{ system.network }}</div>
        <div class="info-line"><strong>IP Address:</strong> {{ system.ip }}</div>
        <div class="info-line"><strong>Temperature Sensor:</strong> {{ system.temperature }} °C</div>
        <div class="info-line"><strong>Last Boot:</strong> {{ system.lastBoot }}</div>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script setup>
import { ref } from 'vue'

const system = ref({
  version: '1.4.2',
  newVersion: '1.5.0',
  buildDate: '2025-10-18',
  updateAvailable: true,
  releaseNotes: [
    'Improved power management for Raspberry Pi 4.',
    'Enhanced NFC tag stability.',
    'Added dark theme option.',
    'Minor bug fixes and optimizations.'
  ],
  os: 'Raspberry Pi OS (Debian 12 Bookworm)',
  arch: 'arm64',
  uptimeSeconds: 86400 * 3 + 4520, // 3 days 1h15min
  diskUsed: '6.8 GB',
  diskTotal: '16 GB',
  diskUsedPct: 42,
  memoryUsed: '1.9 GB',
  memoryTotal: '4 GB',
  memoryUsedPct: 47,
  cpuLoadPct: 23.6,
  network: 'eth0 (LAN)',
  ip: '192.168.1.45',
  temperature: 47.8,
  lastBoot: '2025-10-26 09:22',
})

const stats = ref([
  { label: 'Alarms Configured', value: 8, icon: 'mdi-alarm' },
  { label: 'Playlists', value: 4, icon: 'mdi-music' },
  { label: 'Files in Library', value: 32, icon: 'mdi-file-music' },
  { label: 'Sleepscapes', value: 3, icon: 'mdi-weather-night' },
])

function formatUptime(seconds) {
  const days = Math.floor(seconds / 86400)
  const hours = Math.floor((seconds % 86400) / 3600)
  const mins = Math.floor((seconds % 3600) / 60)
  return `${days}d ${hours}h ${mins}m`
}
</script>

<style scoped>
.system-page {
  background-color: var(--color-bg);
  min-height: 100vh;
  color: var(--color-surface);
}

.info-card {
  background-color: #fff;
  border-radius: 12px;
  color: var(--color-surface);
}

.stat-card {
  background-color: rgba(31, 44, 58, 0.05);
  border-radius: 10px;
  text-align: center;
  color: var(--color-surface);
}

.info-line {
  margin-bottom: 0.4rem;
}

.release-notes {
  margin-top: 0.3rem;
  margin-left: 1rem;
  color: var(--color-surface-alt);
}

.resource {
  margin-bottom: 0.5rem;
}

.label {
  font-weight: 600;
  margin-bottom: 0.3rem;
}

.caption {
  font-size: 0.85rem;
  color: var(--color-surface-alt);
  margin-top: 0.2rem;
}
</style>
}
