<template>
  <span class="button-wrapper">
    <button class="button" v-bind:class="classObject">
      <slot></slot>
    </button>
  </span>
</template>

<script>
export default {
  props: [ 'color' ],
  computed: {
    classObject: function() {
      return {
        'button--primary': this.color == 'primary',
        'button--transparent': this.color == 'transparent'
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.button {
  background-color: lighten($color: $secondary-color, $amount: 10%);
  border: none;
  border-radius: 0;
  color: white;
  padding: 15px 30px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  cursor: pointer;
  transform-origin: top;
  transition: all .5s cubic-bezier(0.075, 0.82, 0.165, 1);
  transform: perspective(4cm) rotateX(0deg);;

  &:focus {
    outline: 0;
  }

  &:active {
    transform: perspective(4cm) rotateX(-15deg) translate3d(0, 5px, -5px);
  }

  &--primary {
    background-color: $primary-color;
  }

  &--transparent {
    background-color: transparent;
  }
}
</style>
