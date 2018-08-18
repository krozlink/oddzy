<template>
  <div id="race-card" class="container">

    <div v-if="loading">
      Loading
    </div>

    <div v-if="!loading" class="section">
      <div class="title is-5">
      {{ `${race.meeting.name} - Race ${race.number}`}}
      </div>
      <div class="subtitle is-6">
      {{ race.name}}
      </div>

      <table class="table is-striped is-hoverable is-fullwidth">
        <thead>
          <tr>
            <th> </th>
            <th> </th>
            <th>Runner</th>
            <th>Win</th>
            <th>Place</th>
            <th>Win</th>
            <th>Place</th>
          </tr>
        </thead>
        <tbody>
            <race-selection :key="selection.selection_id" v-for="selection in race.selections" :selection="selection"></race-selection>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import RaceSelection from '../components/racing/RaceSelection.vue';

export default {
  name: 'RaceCard',
  components: {
    RaceSelection,
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
  #race-card.container{
    @media screen and (min-width: 1472px) {
      width:1100px;
    }
  }

  table {
    width: 100%;
  }
</style>
