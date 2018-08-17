<template>
  <div>

    <div v-if="loading">
      Loading
    </div>

    <div v-if="!loading" class="section">
      <div class="title is-4">
      {{ `${race.meeting.name} - Race ${race.number}`}}
      </div>
      <div class="subtitle is-6">
      {{ race.name}}
      </div>

      <div :key="selection.selection_id" v-for="selection in race.selections">
        {{`${selection.name}` }}
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'RaceCard',
  components: {
  },
  data() {
    return {
    };
  },
  computed: {
    race() {
      return this.$store.state.racing.races[this.$route.params.id];
    },
    loading() {
      return this.$store.state.racing.loadingStatus !== 'successful';
    },
  },
  created() {
    this.$store.dispatch('racing/getRaceCard', this.$route.params.id);
  },
};
</script>

<style lang="scss" scoped>

</style>
