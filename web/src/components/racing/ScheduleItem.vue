<template>
  <div class="race-item">
    <span v-if="race.status === 'CLOSED'">
      {{ race.results }}
    </span>
    <span v-if="race.status === 'OPEN'">
      Race {{ race.number }}
    </span>
  </div>
</template>

<script>
export default {
  name: 'ScheduleItem',
  props: ['race'],
};
</script>

<style lang="scss" scoped>


</style>
