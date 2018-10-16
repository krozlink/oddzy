<template>
    <div :class="{
        button: true,
        'is-disable': race.status !== 'OPEN',
        'odds-increased': race.status === 'OPEN' && increased,
        'odds-decreased': race.status === 'OPEN' && decreased,
    }"
    v-on:click="addToBetslip"
    >
    {{ price }}
    </div>
</template>

<script>
import Bet from '../../api/bet';

export default {
  name: 'PriceButton',
  props: ['selection', 'race', 'meeting', 'price', 'betType', 'winType'],
  computed: {
    increased() {
      if (this.betType === 'tote') return false;

      const p = this.$store.state.racing.prices[this.selection.number];
      return (p && p[this.winType] && p[this.winType].change === 'increase');
    },
    decreased() {
      if (this.betType === 'tote') return false;

      const p = this.$store.state.racing.prices[this.selection.number];
      return (p && p[this.winType] && p[this.winType].change === 'decrease');
    },
  },
  methods: {
    addToBetslip() {
      this.$store.dispatch('betslip/addToBetslip', new Bet({
        runner_number: this.selection.number,
        runner_name: this.selection.name,
        meeting_name: this.meeting.name,
        meeting_number: this.race.number,
        price: this.price,
        bet_type: this.betType,
        win_type: this.winType,
        amount: 0,
        selection_id: this.selection.selection_id,
      }));
    },
  },
};
</script>

<style lang="scss" scoped>

</style>
