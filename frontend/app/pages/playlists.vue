<template>
  <v-container fluid class="playlist-page">
    <v-row class="align-center mb-4">
      <v-col cols="12" class="d-flex align-center justify-space-between">
        <h2 class="text-h5" style="color: var(--color-surface);">Playlists</h2>

        <v-btn
          color="var(--color-primary)"
          prepend-icon="mdi-plus"
          class="add-btn"
          @click="onAddPlaylist"
        >
          Add Playlist
        </v-btn>
      </v-col>
    </v-row>

    <v-row>
      <v-col
        v-for="pl in playlists"
        :key="pl.id"
        cols="12"
        md="6"
        lg="4"
      >
        <v-card class="playlist-card">
          <v-card-title class="d-flex align-center justify-space-between">
            <div>
              <div class="playlist-name">{{ pl.name }}</div>
              <div class="playlist-meta">
                {{ pl.files.length }} files • {{ formatTotalLength(pl.files) }}
              </div>
            </div>
            <div class="d-flex align-center gap-1">
              <v-btn
                icon="mdi-pencil"
                variant="text"
                color="var(--color-primary)"
                @click="editPlaylist(pl)"
              />
              <v-btn
                icon="mdi-delete"
                variant="text"
                color="var(--color-primary-dark)"
                @click="deletePlaylist(pl.id)"
              />
            </div>
          </v-card-title>
        </v-card>
      </v-col>
    </v-row>

    <PlaylistEditor
      v-model="showEditor"
      :playlist="selectedPlaylist"
      @save="savePlaylist"
    />
  </v-container>
</template>

<script setup>
import { ref } from 'vue'
import PlaylistEditor from '~/components/playlists/PlaylistEditor.vue'

const playlists = ref([
  {
    id: 1,
    name: 'Morning Vibes',
    files: [
      { id: 1, title: 'Soft Dawn', artist: 'Aurora Fields', length: '03:42' },
      { id: 2, title: 'Sunrise Intro', artist: 'Low Tide', length: '02:10' },
      { id: 3, title: 'Waking Light', artist: 'Low Tide', length: '04:07' },
    ],
  },
  {
    id: 2,
    name: 'Deep Sleep',
    files: [
      { id: 1, title: 'Ocean Breath', artist: 'Somnus', length: '06:03' },
      { id: 2, title: 'Drifting', artist: 'Quiet Forms', length: '05:27' },
      { id: 3, title: 'Through Mist', artist: 'Lunaris', length: '07:45' },
    ],
  },
])

const showEditor = ref(false)
const selectedPlaylist = ref(null)

const onAddPlaylist = () => {
  selectedPlaylist.value = { id: null, name: '', files: [] }
  showEditor.value = true
}

const editPlaylist = (pl) => {
  selectedPlaylist.value = JSON.parse(JSON.stringify(pl))
  showEditor.value = true
}

const savePlaylist = (payload) => {
  if (payload.id) {
    const idx = playlists.value.findIndex(p => p.id === payload.id)
    if (idx !== -1) playlists.value[idx] = payload
  } else {
    payload.id = Date.now()
    playlists.value.push(payload)
  }
}

const deletePlaylist = (id) => {
  playlists.value = playlists.value.filter(p => p.id !== id)
}

const formatTotalLength = (files) => {
  const totalSec = files.reduce((acc, f) => {
    const [m, s] = f.length.split(':').map(Number)
    return acc + m * 60 + s
  }, 0)
  const mm = Math.floor(totalSec / 60)
  const ss = (totalSec % 60).toString().padStart(2, '0')
  return `${mm}:${ss}`
}
</script>

<style scoped>
.playlist-page {
  background-color: var(--color-bg);
  min-height: 100vh;
}

.playlist-card {
  background-color: var(--color-text);
  color: var(--color-surface);
  border-radius: 12px;
  transition: box-shadow var(--transition-fast);
  padding-bottom: 0.5rem;
}

.playlist-card:hover {
  box-shadow: 0 4px 10px rgba(31, 44, 58, 0.15);
}

.playlist-name {
  font-weight: 600;
  font-size: 1.2rem;
}

.playlist-meta {
  font-size: 0.85rem;
  color: var(--color-surface-alt);
}

.add-btn {
  background-color: var(--color-primary);
  color: var(--color-bg);
  text-transform: none;
  font-weight: 500;
}
</style>
