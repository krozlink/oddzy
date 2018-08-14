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
                <!-- <span v-on:click="dateToday" v-bind:class="{'is-primary': filterDate ==='today'}" class="button is-light">Today</span> -->
                <router-link active-class="is-primary" class="button" :to="{name: 'racing-date', params: {date: getTomorrow()}}">
                  <span>Tomorrow</span>
                </router-link>
                <!-- <span v-on:click="dateTomorrow" v-bind:class="{'is-primary': filterDate ==='tomorrow'}" class="button is-light">Tomorrow</span> -->
                <router-link active-class="is-primary" class="button" :to="{name: 'racing-date', params: {date: getOvermorrow()}}">
                  <span v-text="getOvermorrowDay()"></span>
                </router-link>

                <!-- <span v-on:click="dateOvermorrow" v-bind:class="{'is-primary': filterDate ==='overmorrow'}" class="button is-light">Saturday</span> -->
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

    <div class="section race-section">
      <div class="columns">
        <div class="column">
          <h6 class="title is-6">Horse Racing - Australia & New Zealand</h6>
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

    <div class="section race-section">
      <div class="columns">
        <div class="column">
          <h6 class="title is-6">Horse Racing - International</h6>
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

  </div>
</template>

<script>
import date from '../api/date-helper';

export default {
  created() {
    // if last update > 1 minute ago then update racing schedule
    console.log(this.$route);
    this.$store.dispatch('racing/getRacingSchedule', this.$route.params.date || this.getToday());
  },
  watch: {
    $route() {
      this.$store.dispatch('racing/getRacingSchedule', this.$route.params.date || this.getToday());
    },
  },

  methods: {
    getToday() {
      return date.formatDate(date.todayDate());
    },
    getTomorrow() {
      return date.formatDate(date.tomorrowDate());
    },
    getOvermorrow() {
      return date.formatDate(date.overmorrowDate());
    },
    getOvermorrowDay() {
      return date.getDayString(date.overmorrowDate().getDay());
    },

    dateToday() {
      this.filterDate = 'today';
    },
    dateTomorrow() {
      this.filterDate = 'tomorrow';
    },
    dateOvermorrow() {
      this.filterDate = 'overmorrow';
    },
    dateCustom() {
      // todo
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

  data() {
    return {
      filterType: 'all',
      filterDate: 'today',
      lastUpdate: 0,
    };
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
