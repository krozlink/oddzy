<template>
  <div>
    <div class="section">
      <div class="title is-4">
      {{ race.name}}
      </div>
      <div class="subtitle">
      Race {{ race.number}}
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'RaceCard',
  components: {
  },
  computed: {
    race() {
      return this.$store.state.racing.races[this.$route.params.id];
    },
  },
  data() {
    return {
    };
  },
};
</script>

<style lang="scss" scoped>

</style>
