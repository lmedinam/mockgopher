<template>
  <div class="main-page">
    <TheMenuBar></TheMenuBar>
    <div class="main-page__content">
      <TheSideBar></TheSideBar>
      <main class="main-page__view-wrapper">
        <router-view></router-view>
      </main>
    </div>
    <TheStatusBar></TheStatusBar>
  </div>
</template>

<script>
import TheMenuBar from '../components/TheMenuBar';
import TheStatusBar from '../components/TheStatusBar';
import TheSideBar from '../components/TheSideBar';

export default {
  components: {
    TheMenuBar, TheStatusBar, TheSideBar
  }
}
</script>

<style lang="scss">
.main-page {
  height: 100%;
  display: flex;
  flex-direction: column;
}
.main-page__content {
  flex-grow:  1;
  display: flex;
}
.main-page__view-wrapper {
  height: 100%;
  flex-grow: 1;
}
</style>

