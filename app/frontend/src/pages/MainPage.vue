<template>
    <div class="main-page">
        <TheMenuBar></TheMenuBar>
        <div class="main-page__content">

        </div>
        <TheStatusBar></TheStatusBar>
    </div>
</template>

<script>
import TheMenuBar from '../components/TheMenuBar';
import TheStatusBar from '../components/TheStatusBar';

export default {
    components: {
        TheMenuBar, TheStatusBar
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
}
</style>

