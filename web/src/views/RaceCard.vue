<template>
  <div id="race-card" class="container">

    <div v-if="loading"  class="section">
      Loading
    </div>
    <div v-else class="section">
      <div class="race-title">
        <div class="title is-5 meeting-name">
          {{ `${race.meeting.name} - Race ${race.number}`}}
        </div>
        <div class="subtitle is-6 race-name">
          {{ race.name}}
        </div>
      </div>

      <div class="race-status">
        <div class="title is-5">
          {{ formattedTime }}
        </div>
        <div class="subtitle is-6">
          <span v-if="race.status !== 'OPEN'" class="highlight">
            {{ formattedStatus }}
          </span>
        </div>
      </div>

      <table class="table is-striped is-hoverable is-fullwidth">
        <thead>
          <tr class="top">
            <th rowspan="2"></th>
            <th v-if="isHorseRacing" rowspan="2"></th>
            <th rowspan="2">Runner</th>
            <th colspan="2" class="bet-type">Fixed</th>
            <th colspan="2" class="bet-type">Tote</th>
          </tr>
          <tr class="bottom">
            <th>Win</th>
            <th>Place</th>
            <th>Win</th>
            <th>Place</th>
          </tr>
        </thead>
        <tbody>
            <race-selection :key="selection.selection_id" v-for="selection in race.selections" :race="race" :meeting="race.meeting" :selection="selection"></race-selection>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import RaceSelection from '../components/racing/RaceSelection.vue';
import date from '../api/date-helper';

export default {
  name: 'RaceCard',
  components: {
    RaceSelection,
  },
  data() {
    return {
    };
  },
  sockets: {
    connect() {
    },
    updatePrice(val) {
      this.$store.dispatch('racing/updatePrice', val);
    },
  },
  computed: {
    race() {
      return this.$store.state.racing.races[this.$route.params.id];
    },
    loading() {
      return this.$store.state.racing.loadingStatus !== 'successful';
    },
    formattedTime() {
      return date.formatTime(this.$store.state.racing.races[this.$route.params.id].scheduled_start * 1000);
    },
    formattedStatus() {
      const { status } = this.$store.state.racing.races[this.$route.params.id];
      if (status === 'CLOSED') {
        return 'RACE CLOSED';
      } else if (status === 'ABANDONED') {
        return 'RACE ABANDONED';
      }
      return status;
    },
  },
  methods: {
    isHorseRacing() {
      return this.race().meeting.race_type === 'horse-racing';
    },
    formatTime() {
      return date.formatTime(this.race().scheduled_start);
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

  .race-title {
    width: 550px;
    height: 50px;
  }

  .race-status,.race-title {
    display: inline-block;
  }

  .highlight {
    color: $danger;
  }

  .race-status {
    margin-right: 5px;
    float: right;
    text-align: right;
  }

  .race-name {
    text-transform: capitalize;
    // font-variant-alternates: annotation()
  }

  #race-card > .section {
    padding-top: 20px;
  }

  table {
    width: 100%;
  }

  thead {
    font-size: 1em;
  }

  .top th {
    vertical-align: middle;
  }

  th.bet-type {
    text-align:center;
  }

  .bet-type {
    font-size: 0.9em;
  }

  thead th {
    background-color: #D3D3D3;
  }

  thead tr.top th {
    padding: 2px;
  }

  thead tr.bottom th {
    padding: 0px;
    height: 20px;
    font-size: 0.8em;
    text-align: center;
  }

  table thead tr.top th:first-child{
    border-style: solid;
    border-radius:10px 0px 0px 0px;
  }


  table thead tr.top th:last-child{
    border-style: solid;
    border-radius:0 10px 0px 0;
  }

</style>
