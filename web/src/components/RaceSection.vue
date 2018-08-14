<template>
    <div class="section race-section" v-if="display">
      <div class="columns">
        <div class="column">
          <h6 class="title is-6">{{ raceTypeDisplay }} - {{ raceLocalDisplay }}</h6>
        </div>
      </div>

      <div class="columns race-catagory ">
        <div class="column race-location is-one-fifth">
          Flemington
        </div>
        <div class="column race-list">
          <div class="columns">
            <div class="column race-item">
              Race 1
            </div>
            <div class="column race-item">
              Race 2
            </div>
            <div class="column race-item">
              Race 3
            </div>
            <div class="column race-item">
              Race 4
            </div>
            <div class="column race-item">
              Race 5
            </div>
            <div class="column race-item">
              Race 6
            </div>
            <div class="column race-item">
              Race 7
            </div>
            <div class="column race-item">
              Race 8
            </div>
            <div class="column race-item">
              Race 9
            </div>
            <div class="column race-item">
              Race 10
            </div>
            <div class="column race-item">
              Race 11
            </div>
            <div class="column race-item">
              Race 12
            </div>
          </div>
        </div>
      </div>
    </div>

</template>

<script>
export default {
  props: ['racedate', 'racetype', 'racelocal'],

  computed: {
    raceTypeDisplay() {
      if (this.racetype === 'horse-racing') {
        return 'Horse Racing';
      } else if (this.racetype === 'harness') {
        return 'Harness';
      } else if (this.racetype === 'greyhounds') {
        return 'Greyhounds';
      }

      return 'Undefined';
    },
    raceLocalDisplay() {
      return this.racelocal ? 'Australia & New Zealand' : 'International';
    },
    display() {
      const meetings = this.$store.getters['racing/filterMeetings'](this.racetype, this.racedate, this.racelocal);
      return meetings.length > 0;
    },
  },
};
</script>


<style lang="sass" scoped>

</style>
