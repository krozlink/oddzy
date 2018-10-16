<template>
    <div class="card" :class="{'readonly': readonly}">
      <div class="card-header-title">
        <div class="bsi-header">
          <div class="bsi-row">
            <div class="entrant">
              <div class="runner">
                {{this.bet.runner_name}}
              </div>
              <div class="meeting">
                {{this.bet.meeting_name}} R{{this.bet.meeting_number}}
              </div>
            </div>
            <div class="odds">
              <div class="price">
                {{formatPrice}}
              </div>
              <div class="bet-type">
                {{ this.bet.win_type}} {{this.bet.bet_type}}
              </div>
            </div>
            <button class="delete is-small" v-on:click="removeFromBetslip"></button>
          </div>
        </div>
      </div>
      <div class="card-content">
        <div class="field has-addons">
          <p class="control">
            <a class="button is-static">$</a>
          </p>
          <p class="control">
            <input
              type="number"
              v-model="amount"
              class="input" min="0"
              :class="{'is-danger': error}"
              v-on:blur="validateError"
              step="1"
              placeholder="Amount"
              :readonly="readonly"
            >
          </p>
        </div>
        <p class="help is-danger" v-if="error">{{this.bet.message}}</p>
        <p class="help is-success" v-if="!error">Estimated Payout: ${{payout}}</p>
      </div>
    </div>
</template>

<script>
export default {
  name: 'BetslipItem',
  props: ['bet', 'readonly'],
  data() {
    return {
      valid: true,
      message: '',
    };
  },
  computed: {
    error() {
      return this.bet.message !== '';
    },
    amount: {
      get() {
        return this.$store.state.betslip.bets[this.bet.bet_id].amount;
      },
      set(value) {
        this.$store.dispatch('betslip/updateBetAmount', {
          betId: this.bet.bet_id,
          amount: value,
        });
      },
    },
    formatPrice() {
      if (this.bet.price < 10) {
        return this.bet.price.toFixed(2);
      }

      return this.bet.price;
    },
    payout() {
      return (this.amount * this.bet.price).toFixed(2);
    },
  },
  methods: {
    validateError() {
      if (this.bet.message !== '') {
        this.bet.validate();
      }
    },
    removeFromBetslip() {
      this.$store.dispatch('betslip/removeFromBetslip', this.bet.bet_id);
    },
  },
};
</script>

<style lang="scss" scoped>

.card {
  margin-left: 1px;
  margin-right: 1px;
}

.card.readonly {
  background-color: #EEE;

  input {
    background-color: #E7E7E7;
  }
}

.card-content {
  padding: 10px;
}

.card-content p.help {
  font-size: 0.9em;
  font-weight: 600;
}

.bsi-header {
  display:table;
  width: 100%;
}

.bsi-row {
  display:table-row;
}

.bsi-row .delete {
  float:right;
}

.bsi-delete {
  width: 16px;
  height: 16px;
}

.entrant, .odds {
  display:table-cell;
}

.meeting,.bet-type {
  font-weight: 400;
  font-size: 0.8em;
  text-transform: uppercase;
}

.odds{
  text-align: right;
}

.runner,.price {
  display:inline-block;
  position: relative;
}

.card-header-title {
  display:inline-block;
  width: 100%;
}

</style>
