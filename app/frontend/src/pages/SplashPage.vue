<template>
  <div class="splash-page">
    <div class="splash-page__content">
      <div class="splash-page__logo"></div>
      <h1 class="splash-page__title">MockGopher</h1>
    </div>
    <span class="splash-page__credits">Made by Luis Medina</span>
  </div>
</template>

<style lang="scss">
.splash-page {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  &__content {
    text-align: center;
  }
  &__logo {
    margin: auto;
    width: 200px;
    height: 200px;
    background-color: #1692B3;
    border-radius: 999px;
  }
  &__title {
    font-size: 48px;
    font-weight: bold;
    margin-top: 24px;
    color: #FFF;
  }
  &__credits {
    position: absolute;
    right: 0;
    left: 0;
    text-align: center;
    bottom: 15%;
    color: #FFF;
  }
}
</style>
