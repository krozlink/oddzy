<template>
  <div class="column race-item">
    <router-link v-if="!empty" :to="raceLink" :class="{imminent: secondsRemaining<300}" class="item-content has-race" alt="test">
      <span v-if="race.status === 'CLOSED'">
        {{ race.results }}
      </span>
      <span v-if="race.status === 'OPEN'">
        <!-- Race {{ race.number }} -->
        <p>
        {{ timeDisplay }}
        </p>
      </span>
    </router-link>
    <div v-if="empty" class="item-content no-race">

    </div>
  </div>
</template>

<script>
import date from '../../api/date-helper';

export default {
  name: 'ScheduleItem',
  props: ['race', 'meeting', 'empty', 'time'],
  data() {
    return {
      interval: {},
    };
  },
  computed: {
    raceLink() {
      return `/racing/${this.raceName(this.meeting.name)}/${this.race.race_id}`;
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
    raceName(name) {
      return name.toLowerCase().replace(' ', '-');
    },
  },
};
</script>

<style lang="scss" scoped>

.item-content {
  border-width: 0px 1px 1px 0px;
  border-style: none solid solid none;
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
