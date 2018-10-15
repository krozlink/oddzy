<template>
    <div id="betslip" class="tile is-vertical is-2" v-if="showBetslip"
        :class="{
          'betslip-active': showBetslip,
          'betslip-disabled': !showBetslip,
          }">
      <betslip-item></betslip-item>
    </div>
</template>

<script>
import BetslipItem from './BetslipItem.vue';

export default {
  components: {
    BetslipItem,
  },
  computed: {
    showBetslip() {
      return this.$store.state.betslip.show;
    },
  },
};
</script>

<style lang="scss" scoped>
#betslip {
  background-color: #EEE;
  width: $bs-width;
  min-height: calc(100vh - 85px);
  float:right;
  position: absolute;
  top: $nav-height;
  right: 0px;

  z-index: 1000;
}
</style>
