<template>
    <div class="the-status-bar">
        <div>
            <span class="the-status-bar__item">
                Nothing is running
            </span>
        </div>
        <div>
            <a href="#" class="the-status-bar__button">
                Tweet a feedback
            </a>
            <a href="#" class="the-status-bar__button">
                MockGopher Developer Preview
            </a>
        </div>
    </div>
</template>

<style lang="scss" scoped>
$color: rgba(255, 255, 255, 0.50);
$item-padding: 4px 8px;

.the-status-bar {
    display: flex;
    justify-content: space-between;
    color: $color;
    font-size: 12px;

    &__item {
        padding: $item-padding;
        display: inline-block;
    }

    .the-status-bar__button {
        display: inline-block;
        padding: $item-padding;
        text-decoration: none;
        color: $color;

        &:hover {
            background-color: lighten($color: $secondary-color, $amount: 5%);
        }
    }
}
</style>
