<script setup lang="ts">

import { useRouter } from 'vue-router'

const router = useRouter()
const goToHome = () => {
  router.push('/')
}

</script>

<template>
  <div class="navbar bg-base-100 shadow-sm">
    <div class="flex-1">
      <a class="btn btn-ghost text-xl" @click="goToHome">Nordlys</a>
    </div>
    <div class="flex-none">
      <ul class="menu menu-horizontal px-1">
        <li><RouterLink to="/">Accueil</RouterLink></li>
        <li><RouterLink to="/about">A propos</RouterLink></li>
        <li><RouterLink to="/login">Connexion</RouterLink></li>
        <li><RouterLink to="/register">S'enregistrer</RouterLink></li>
      </ul>
    </div>
  </div>
</template>

<style scoped>

</style>