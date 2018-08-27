<template>
    <div class="section race-section" v-if="display">
      <div class="columns race-section-title">
        <div class="column">
          <h6 class="title is-6">{{ raceTypeDisplay }} - {{ raceLocalDisplay }}</h6>
        </div>
      </div>

      <div :key="index"  v-for="(m, index) in meetings()" class="columns race-catagory ">
        <router-link :to="nextRaceLink(m)" class="column race-location is-one-fifth">
          {{ m.name }}
        </router-link>
        <div class="column race-list">
          <div class="columns">
              <schedule-item :time="time" :key="ri" v-for="(r, ri) in getRaces(m.meeting_id)" :next="isNextRace(ri, m.meeting_id)" :race="r" :meeting="m" :empty="false"> </schedule-item>
              <schedule-item :key="'empty-'+ i" v-for="i in maxRaces() - getRaces(m.meeting_id).length" :race="{}" :meeting="m" :empty="true"> </schedule-item>
          </div>
        </div>
      </div>
    </div>

</template>

<script>
import { mapGetters } from 'vuex';
import racing from '../../api/racing';
import ScheduleItem from './ScheduleItem.vue';

export default {
  name: 'ScheduleSection',
  components: {
    ScheduleItem,
  },
  props: ['racedate', 'racetype', 'racelocal', 'time'],
  data() {
    return {
    };
  },
  computed: {
    ...mapGetters({
      getMeetings: 'racing/filterMeetings',
      getRaces: 'racing/getRacesForMeeting',
    }),
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
      return this.getMeetings(this.racetype, this.racedate, this.racelocal).length > 0;
    },
  },
  methods: {
    nextRaceLink(meeting) {
      return `/racing/${racing.raceNameURL(meeting.name)}/${this.getNextRace(meeting.meeting_id)}`;
    },
    getNextRace(meetingId) {
      const races = this.getRaces(meetingId);
      for (let i = 0; i < races.length; i += 1) {
        if (races[i].status !== 'CLOSED' && races[i].status !== 'ABANDONED') {
          return races[i].race_id;
        }
      }

      return races[races.length - 1].race_id;
    },
    isNextRace(index, meetingId) {
      const races = this.getRaces(meetingId);

      const { status } = races[index];
      if (status === 'CLOSED' || status === 'ABANDONED') {
        return false;
      }

      if (index === 0) return true;

      return races[index - 1].status === 'CLOSED' || races[index - 1].status === 'ABANDONED';
    },
    maxRaces() {
      let max = 0;
      const meetings = this.meetings();
      meetings.forEach((m) => {
        if (m.race_ids.length > max) {
          max = m.race_ids.length;
        }
      });
      return max;
    },
    meetings() {
      return this.getMeetings(this.racetype, this.racedate, this.racelocal);
    },
    races(meetingId) {
      return this.getRaces(meetingId);
    },
  },
};
</script>


<style lang="scss" scoped>

  #location-header>.loc-title {
    padding-left: 12px;
  }

  #location-header {
    display: none;
    color: $body-color;
    font-size: 0.9em;
    height: 20px;
  }

  #location-header .race-header{
    background-color: rgb(211, 211, 211);

    border-bottom-color: $grey-lighter;
    border-bottom-width: 1px;
    border-bottom-style: solid;
  }

  @media screen and (min-width: 768px) {
    #location-header {
        display: flex;
    }
  }

  .location-numbers {
    margin-left: 12px;
    margin-right: 12px;
  }

  .race-header {
    text-align: center;
    margin-top: 12px;
  }

  .race-catagory .column {
    padding: 0px;
    font-size: 0.95em;
  }

  .section.race-section {
    padding-left: 12px;
    padding-right: 12px;
    margin-left: 24px;
    margin-right: 24px;
  }

  .race-catagory .race-location {
    padding-left: 12px;
    padding-top: 12px;
    height: 50px;
  }

  .race-list .columns {
    margin: 0px;
  }

  .race-section-title {
    background-color: $primary;
  }

  .race-section-title .title {
    color: #F7F7F7;
    font-weight: 500;
  }

  .race-section .race-location {
    font-size: 0.95em;
  }

  .race-section .race-location:hover {
    background-color: #EFEFEF;
  }

  a {
    color: $body-color;
  }


  .race-section {
    background-color: #F7F7F7;
  }
</style>
