<template>
  <div id="racing" class="container">
    <div class="section columns header">
      <div class="column is-one-third">
        <h4 class="title is-4">Today's Racing</h4>
      </div>
      <div class="column">
        <nav class="level">
          <div class="level-left">
            <div class="level-item">
              <div class="buttons has-addons">
                <router-link active-class="is-primary" exact class="button" to="/racing">
                  <span>Today</span>
                </router-link>
                <router-link active-class="is-primary" class="button" :to="{name: 'racing-date', params: {date: getTomorrow()}}">
                  <span>Tomorrow</span>
                </router-link>
                <router-link active-class="is-primary" class="button" :to="{name: 'racing-date', params: {date: getOvermorrow()}}">
                  <span v-text="getOvermorrowDay()"></span>
                </router-link>
                <span class="button is-light is-disable">
                  <span class="icon is-large">
                    <i class="fas fa-lg fa-calendar-alt"></i>
                  </span>
                  <span>Select Date</span>
                </span>
              </div>
            </div>
          </div>
          <div class="level-right">
            <div class="level-item racetype-icons">
              <div class="buttons has-addons">
                <span v-on:click="filterAll" v-bind:class="{'is-primary': filterType === 'all'}" class="button">All</span>
                <span v-on:click="filterHorseRacing" v-bind:class="{'is-primary': filterType === 'horse-racing'}" class="button">
                  <span class="icon">
                    <i class="icon-thoroughbred fa-3x"></i>
                  </span>
                </span>
                <span v-on:click="filterHarness" v-bind:class="{'is-primary': filterType === 'harness'}" class="button">
                  <span class="icon">
                    <i class="icon-harness fa-3x"></i>
                  </span>
                </span>
                <span v-on:click="filterGreyhounds" v-bind:class="{'is-primary': filterType === 'greyhounds'}" class="button">
                  <span class="icon">
                    <i class="icon-greyhound fa-3x"></i>
                  </span>
                </span>
              </div>
            </div>
          </div>
        </nav>
      </div>
    </div>

    <div v-if="loading">
      Loading
    </div>

    <div v-if="!loading">
      <race-section
        v-if="this.filterType === 'all' || this.filterType === 'horse-racing'"
        v-bind:racedate="this.filterDate"
        racetype='horse-racing'
        v-bind:racelocal="true">
      </race-section>

      <race-section
        v-if="this.filterType === 'all' || this.filterType === 'horse-racing'"
        v-bind:racedate="this.filterDate"
        racetype='horse-racing'
        v-bind:racelocal="false">
      </race-section>

      <race-section
        v-if="this.filterType === 'all' || this.filterType === 'harness'"
        v-bind:racedate="this.filterDate"
        racetype='harness'
        v-bind:racelocal="true">
      </race-section>

      <race-section
        v-if="this.filterType === 'all' || this.filterType === 'harness'"
        v-bind:racedate="this.filterDate"
        racetype='harness'
        v-bind:racelocal="false">
      </race-section>

      <race-section
        v-if="this.filterType === 'all' || this.filterType === 'greyhounds'"
        v-bind:racedate="this.filterDate"
        racetype='greyhounds'
        v-bind:racelocal="true">
      </race-section>

      <race-section
        v-if="this.filterType === 'all' || this.filterType === 'greyhounds'"
        v-bind:racedate="this.filterDate"
        racetype='greyhounds'
        v-bind:racelocal="false">
      </race-section>
    </div>
  </div>
</template>

<script>
import RaceSection from '../components/RaceSection.vue';
import DateHelper from '../api/date-helper';

export default {
  name: 'Racing',
  components: {
    RaceSection,
  },
  data() {
    return {
      filterType: 'all',
      filterDate: this.$route.params.date || this.getToday(),
      lastUpdate: 0,
    };
  },

  computed: {
    loading() {
      return this.$store.state.racing.loadingStatus !== 'successful';
    },
  },

  created() {
    this.$store.dispatch('racing/getRacingSchedule', this.$route.params.date || this.getToday());
  },
  watch: {
    $route() {
      this.filterDate = this.$route.params.date || this.getToday();
      this.$store.dispatch('racing/getRacingSchedule', this.filterDate);
    },
  },

  methods: {

    getToday() {
      return DateHelper.formatDate(DateHelper.todayDate());
    },
    getTomorrow() {
      return DateHelper.formatDate(DateHelper.tomorrowDate());
    },
    getOvermorrow() {
      return DateHelper.formatDate(DateHelper.overmorrowDate());
    },
    getOvermorrowDay() {
      return DateHelper.getDayString(DateHelper.overmorrowDate().getDay());
    },

    filterAll() {
      this.filterType = 'all';
    },
    filterHorseRacing() {
      this.filterType = 'horse-racing';
    },
    filterHarness() {
      this.filterType = 'harness';
    },
    filterGreyhounds() {
      this.filterType = 'greyhounds';
    },
  },
};
</script>

<style lang="scss" scoped>
.racetype-icons .button {
  width: 70px;
}

.racetype-icons .icon {
  padding-top: 5px;
}

#racing .section {
  padding-top: 16px;
}

#racing .header {
  padding-bottom: 0px;
}

</style>
