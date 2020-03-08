<template>
  <div class="the-menubar">
    <div class="the-menubar__item">
      <AppButton color="primary">
        <AppIcon icon="plus" class="the-menubar__icon"></AppIcon>New route
      </AppButton>
      <AppButton>Run</AppButton>
    </div>
    <h1 class="the-menubar__title">My Workspace</h1>
    <div>
      <AppButton color="transparent">
        <AppIcon icon="cog"></AppIcon>
      </AppButton>
    </div>
  </div>
</template>

<style lang="scss">
.the-menubar {
  display: flex;
  justify-content: space-between;
  padding: 10px;
  align-items: center;

  &__item {
    * + * {
      margin-left: 10px;
    }
  }

  &__icon {
    margin-right: 10px;
  }

  &__title {
    font-size: 24px;
    font-weight: bold;
    color: #FFF;
  }
}
</style>
