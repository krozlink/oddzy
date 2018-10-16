<template>
    <div id="betslip" class="tile is-vertical is-2" v-if="showBetslip"
        :class="{
          'betslip-active': showBetslip,
          'betslip-disabled': !showBetslip,
          }">

      <article
        v-if="$store.state.betslip.message.type !== ''"
        class="message is-success is-small"
        :class="{
          'is-success': $store.state.betslip.message.type === 'success',
        }"
      >
        <div class="message-body">
          <p :key="index" v-for="(line, index) in $store.state.betslip.message.lines">{{line}}</p>
        </div>
      </article>

      <betslip-item :key="bet.bet_id" :bet="bet" v-for="bet in $store.state.betslip.bets" :readonly="readonly"></betslip-item>

      <div id="bs-footer" class="card" v-if="Object.keys($store.state.betslip.bets).length > 0">
        <div id="bs-summary">
          <p>
          Number of bets: <span class="bs-bets">{{Object.keys($store.state.betslip.bets).length}}</span>
          </p>
          <p>
          Total Payout: <span class="bs-payout">{{totalPayout}}</span>
          </p>
        </div>
        <a class="button is-primary"
          :class="{'is-loading': this.$store.state.betslip.status === 'submitting'}"
          v-on:click="placeBets"
        >Place Bets</a>
      </div>
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
    totalPayout() {
      let total = 0;
      Object.values(this.$store.state.betslip.bets).forEach((b) => {
        total += b.amount * b.price;
      });
      return `$${total.toFixed(2)}`;
    },
    readonly() {
      return this.$store.state.betslip.status === 'submitting';
    },
  },
  methods: {
    placeBets() {
      this.$store.dispatch('betslip/submit');
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

  article {
    margin: 0px;

    p {
      font-size: 1.2em;
      color: green;
    }
  }
}


#bs-summary {
  padding: 10px;
}

#bs-footer {

  background-color: white;


  .button {
    width: 100%;
  }

  P {
    font-weight: 600;
    color: #333;
    padding-left: 10px;
    padding-right: 10px;
  }

  span {
    float: right;
    text-align: right;
  }
}
</style>
