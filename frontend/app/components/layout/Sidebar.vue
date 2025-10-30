<template>
  <div>
      <v-img
        class="my-4"
        height="150"
        src="~/assets/logo.png"
      />
  </div>
    <nav class="nav">
      <NuxtLink to="/" class="nav-item">
        <v-icon icon="mdi-view-dashboard" class="nav-icon" />
        <span>Dashboard</span>
      </NuxtLink>

      <NuxtLink to="/alarms" class="nav-item">
        <v-icon icon="mdi-alarm" class="nav-icon" />
        <span>Alarms</span>
      </NuxtLink>

      <NuxtLink to="/sleepscapes" class="nav-item">
        <v-icon icon="mdi-weather-night" class="nav-icon" />
        <span>Sleepscapes</span>
      </NuxtLink>

      <NuxtLink to="/playlists" class="nav-item">
        <v-icon icon="mdi-playlist-music" class="nav-icon" />
        <span>Playlists</span>
      </NuxtLink>

      <NuxtLink to="/files" class="nav-item">
        <v-icon icon="mdi-music" class="nav-icon" />
        <span>Local Library</span>
      </NuxtLink>

      <NuxtLink to="/system" class="nav-item">
        <v-icon icon="mdi-chip" class="nav-icon" />
        <span>System Info</span>
      </NuxtLink>

      <NuxtLink to="/settings" class="nav-item">
        <v-icon icon="mdi-cog" class="nav-icon" />
        <span>Settings</span>
      </NuxtLink>
    </nav>
</template>

<style scoped>
.logo {
  display: flex;
  align-items: center;
  gap: 0.6rem;
  padding: 0 1.2rem 1.2rem 1.2rem;
}

.logo img {
  width: 32px;
  height: 32px;
}

.title {
  font-weight: 600;
  font-size: 1.3rem;
}

.nav {
  display: flex;
  flex-direction: column;
  gap: 0.3rem;
  margin: 1rem;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 0.8rem;
  color: var(--color-text);
  text-decoration: none;
  padding: 0.6rem 1.2rem;
  border-radius: 8px;
  transition: background-color var(--transition-fast);
}

.nav-item:hover {
  background-color: var(--hover-bg);
}

.nav-item.router-link-active {
  background-color: var(--active-bg);
}

.nav-icon {
  font-size: 22px;
}
</style>
