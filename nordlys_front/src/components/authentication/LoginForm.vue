<script setup lang="ts">

</script>

<template>

  <div class="card bg-base-100 card-lg shadow-sm">
    <div class="card-body">
      <h2 class="card-title">Connexion</h2>
      <form>
        <input type="text" placeholder="Adresse E-mail" class="input mb-2" />
        <input type="password" placeholder="Mot de passe" class="input" />
      </form>
      <div class="justify-between card-actions">
        <RouterLink to="/" class="btn btn-primary">Se connecter</RouterLink>
        <RouterLink to="/" class="btn btn-secondary">Annuler</RouterLink>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>