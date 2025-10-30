<template>
  <v-app>
    <v-navigation-drawer app permanent width="240">
      <Sidebar />
    </v-navigation-drawer>

    <v-main>
      <slot />
    </v-main>

  </v-app>
</template>

<script setup>
import Sidebar from '~/components/layout/Sidebar.vue'
</script>

<style scoped>
.v-navigation-drawer {
  background-color: var(--color-surface);
}
</style>
