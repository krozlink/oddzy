<template>
  <div class="column race-item">
    <router-link
      v-if="!empty"
      :to="raceLink"
      :class="{
        'closed': race.status === 'CLOSED',
        'imminent': (race.status === 'OPEN' && secondsRemaining<300),
        'interim': race.status === 'INTERIM',
        'abandoned': race.status === 'ABANDONED',
        'open': race.status === 'OPEN',
        'next': next && !(race.status === 'OPEN' && secondsRemaining<300),
        }"
      class="item-content has-race"
      alt="test">
      <span v-if="race.status === 'CLOSED'">
        {{ race.results }}
      </span>
      <span v-if="race.status === 'OPEN'">
        <!-- Race {{ race.number }} -->
        <p>
        {{ timeDisplay }}
        </p>
      </span>
      <span v-if="race.status === 'ABANDONED'">
        Abandoned
      </span>
    </router-link>
    <div v-if="empty" class="item-content no-race">

    </div>
  </div>
</template>

<script>
import date from '../../api/date-helper';
import racing from '../../api/racing';

export default {
  name: 'ScheduleItem',
  props: ['race', 'meeting', 'empty', 'time', 'next'],
  data() {
    return {
      interval: {},
    };
  },
  computed: {
    raceLink() {
      return `/racing/${racing.raceNameURL(this.meeting.name)}/${this.race.race_id}`;
    },
    secondsRemaining() {
      const remaining = parseInt(this.race.scheduled_start - this.time, 0);
      return remaining;
    },
    timeDisplay() {
      if (this.time === 0) {
        return '';
      }
      return date.formatTimeRemaining(this.secondsRemaining);
    },
  },
  methods: {

  },
};
</script>

<style lang="scss" scoped>

.next > span{
  border-bottom: $body-color;
  border-bottom-width: 2px;
  border-bottom-style: solid;
  padding-top: 2px;
}

.race-item{
  font-family: 'Open Sans';
}

.open,.interim,.imminent{
  background-color: #F7F7F7;
}

.closed,.abandoned {
  background-color: #D3D3D3;
}

.imminent {
  background-color: #3599E5;
  color: #F7F7F7;
}

.abandoned {
  color: #BB0000;
}

.item-content {
  border-width: 0px 1px 1px 0px;
  border-style: none solid solid none;
  border-color: $body-background-color;
  height: 50px;

  font-size: 0.85em;

  display:flex;
  align-items: center; /* Vertical center alignment */
  justify-content: center; /* Horizontal center alignment */
}

.has-race:hover {
  background-color: #EFEFEF;
}

a {
  color: $body-color;
}

</style>
