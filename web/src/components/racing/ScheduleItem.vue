<template>
  <div class="column race-item">
    <router-link v-if="!empty" :to="raceLink" class="item-content has-race" alt="test">
      <span v-if="race.status === 'CLOSED'">
        {{ race.results }}
      </span>
      <span v-if="race.status === 'OPEN'">
        Race {{ race.number }}
      </span>
    </router-link>
    <div v-if="empty" class="item-content no-race">

    </div>
  </div>
</template>

<script>
export default {
  name: 'ScheduleItem',
  props: ['race', 'meeting', 'empty'],
  computed: {
    raceLink() {
      return `/racing/${this.raceName(this.meeting.name)}/${this.race.race_id}`;
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
