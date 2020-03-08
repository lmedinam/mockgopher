<template>
  <span v-bind:class="iconBinding"></span>
</template>

<script>
export default {
  props: {
    icon: {
        type: String,
        required: false
    }
  },
  computed: {
    iconBinding: function() {
      var map = {};
      map[`mif-${this.icon}`] = true;
      return map;
    }
  }
}
</script>
