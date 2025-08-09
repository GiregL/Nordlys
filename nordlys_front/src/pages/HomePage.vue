<script setup lang="ts">

import LoginForm from "../components/authentication/LoginForm.vue";
</script>

<template>
  <div class="flex flex-row gap-4 justify-between w-3/4 mx-auto">
    <div class="flex flex-col gap-4">
      <div class="card bg-base-100 shadow-sm">
        <figure>
          <img
              src="https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp"
              alt="Shoes" />
        </figure>
        <div class="card-body">
          <h2 class="card-title">Card Title</h2>
          <p>A card component has a figure, a body part, and inside body there are title and actions parts</p>
          <div class="card-actions justify-end">
            <button class="btn btn-primary">Buy Now</button>
          </div>
        </div>
      </div>

      <div class="card bg-base-100 shadow-sm">
        <figure>
          <img
              src="https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp"
              alt="Shoes" />
        </figure>
        <div class="card-body">
          <h2 class="card-title">Card Title</h2>
          <p>A card component has a figure, a body part, and inside body there are title and actions parts</p>
          <div class="card-actions justify-end">
            <button class="btn btn-primary">Buy Now</button>
          </div>
        </div>
      </div>
    </div>

    <div class="max-w-1/4 flex flex-col gap-4">
      <div class="card bg-base-100 shadow-sm">
        <figure>
          <img
              src="https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp"
              alt="Shoes" />
        </figure>
        <div class="card-body">
          <h2 class="card-title">Card Title</h2>
          <p>A card component has a figure, a body part, and inside body there are title and actions parts</p>
          <div class="card-actions justify-end">
            <button class="btn btn-primary">Buy Now</button>
          </div>
        </div>
      </div>
      <LoginForm/>
    </div>
  </div>
</template>

<style scoped>

</style>