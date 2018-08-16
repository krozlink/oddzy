<template>
    <div class="section race-section" v-if="display">
      <div class="columns">
        <div class="column">
          <h6 class="title is-6">{{ raceTypeDisplay }} - {{ raceLocalDisplay }}</h6>
        </div>
      </div>

      <div :key="index"  v-for="(m, index) in meetings()" class="columns race-catagory ">
        <div class="column race-location is-one-fifth">
          {{ m.name }}
        </div>
        <div class="column race-list">
          <div class="columns">
            <div :key="index" v-for="(r, index) in getRaces(m.meeting_id)" class="column race-item">
              <schedule-item :race="r"> </schedule-item>
            </div>
            <div :key="'empty-'+ i" v-for="i in maxRaces() - getRaces(m.meeting_id).length" class="column race-item">
              <schedule-item :race="{}" :empty="true"> </schedule-item>
            </div>
          </div>
        </div>
      </div>
    </div>

</template>

<script>
import { mapGetters } from 'vuex';
import ScheduleItem from './ScheduleItem.vue';

export default {
  name: 'ScheduleSection',
  components: {
    ScheduleItem,
  },
  props: ['racedate', 'racetype', 'racelocal'],
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
    maxRaces() {
      // console.debug('calculating max races');
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


<style lang="sass" scoped>

</style>
