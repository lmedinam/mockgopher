<template>
  <div class="app">
    <transition name="fade" mode="out-in">
      <router-view></router-view>
    </transition>
  </div>
</template>

<style lang="scss" scoped>
.app {
  height: 100vh;
  background-color: $secondary-color;
  font-family: "Trebuchet MS", Helvetica, sans-serif;
}
.fade-enter-active {
  animation: fade-in .5s;
}
.fade-leave-active {
  animation: fade-in .5s reverse;
}
@keyframes fade-in {
  0% {
    opacity: 0;
    transform: scale(.95);
  }
  100% {
    opacity: 1;
    transform: scale(1);
  }
}
</style>
