<template>
    <tr :class="{
            'scratched': selection.is_scratched,
            'selection-row': true,
        }">
        <th class="sel-number">
            <div>
                {{selection.number}}
            </div>
            <div class="result">
                <span v-if="placed || win" :class="{
                    'tag': true,
                    'is-warning': win,
                    'is-success': placed,
                }">{{result}}<sup>{{resultOrdinal}}</sup></span>
            </div>
        </th>
        <td class="sel-img">
            <img :src="selection.image_url" class="image" alt="">
        </td>
        <td class = "sel-info">
            <div class="name">{{selection.name}} <span v-if="selection.barrier">({{selection.barrier}})</span></div>
            <div v-if="selection.jockey" class="jockey"><span class="j-icon">J</span>{{selection.jockey}} {{selection.jockey_weight}}</div>
        </td>
        <td class="price fixed win">
            <price-button :selection="selection" :race="race" :meeting="meeting" :betType="'fixed'" :winType="'win'" :price="priceWin"></price-button>
        </td>
        <td class="price fixed place">
            <price-button :selection="selection" :race="race" :meeting="meeting" :betType="'fixed'" :winType="'place'" :price="pricePlace"></price-button>
        </td>
        <td class="price tote win">
            <price-button :selection="selection" :race="race" :meeting="meeting" :betType="'tote'" :winType="'win'" :price="'SP'"></price-button>
        </td>
        <td class="price tote place">
            <price-button :selection="selection" :race="race" :meeting="meeting" :betType="'tote'" :winType="'place'" :price="'SP'"></price-button>
        </td>
    </tr>
</template>

<script>
import test from '../../api/test-prices';
import PriceButton from '../prices/PriceButton.vue';

export default {
  name: 'RaceSelection',
  props: ['selection', 'race', 'meeting'],
  components: {
    PriceButton,
  },
  computed: {
    open() {
      return this.race.status === 'OPEN';
    },
    win() {
      const results = this.race.results.split(',');
      return results.length > 0 && this.selection.number.toString() === results[0];
    },
    placed() {
      const results = this.race.results.split(',');
      return results.length > 0 && results.slice(1).includes(this.selection.number.toString());
    },
    result() {
      const results = this.race.results.split(',');
      return results.indexOf(this.selection.number.toString()) + 1;
    },
    resultOrdinal() {
      return this.getOrdinal(this.result);
    },
    price() {
      if (this.race.status !== 'OPEN') {
        return 'SUSP';
      }
      return '4.00';
    },
    priceWin() {
      if (this.race.status !== 'OPEN') {
        return 'SUSP';
      }

      const p = this.$store.state.racing.prices[this.selection.number];
      if (p && p.win) {
        return p.win.price;
      }
      return this.getTestPrice(this.selection.number, true);
    },
    pricePlace() {
      if (this.race.status !== 'OPEN') {
        return 'SUSP';
      }

      const p = this.$store.state.racing.prices[this.selection.number];
      if (p && p.place) {
        return p.place.price;
      }
      return this.getTestPrice(this.selection.number, false);
    },
  },
  methods: {
    getOrdinal(result) {
      if (result === 1) {
        return 'st';
      } else if (result === 2) {
        return 'nd';
      } else if (result === 3) {
        return 'rd';
      }
      return 'th';
    },
    getTestPrice(barrier, isWin) {
      if (barrier > 12) {
        return isWin ? 101 : 15;
      }

      return isWin ? test.WinPrices[barrier - 1] : test.PlacePrices[barrier - 1];
    },
  },
};
</script>

<style lang="scss" scoped>

@keyframes oddsincrease {
    0% {background-color: white;}
    50% {background-color: green;}
    100% {background-color: white;}
}

@keyframes oddsdecrease {
    0% {background-color: white;}
    50% {background-color: red;}
    100% {background-color: white;}
}


.odds-increased{
  animation-name: oddsincrease;
  animation-duration: 3s;
}

.odds-decreased{
  animation-name: oddsdecrease;
  animation-duration: 3s;
}

.price.fixed.place {
    border-right: 1px;
    border-right-style: solid;
    border-right-color: $grey-lighter
}

.price {
    min-width: 70px;
}

.price div {
    border-style: solid;
    border-color: #AAA;
    border-width: 2px;
    width: 55px;
    height: 30px;
    // margin-top: 10px;

    align-self: center;
    vertical-align: middle;

    display:flex;
    align-items: center; /* Vertical center alignment */
    justify-content: center; /* Horizontal center alignment */

    cursor: pointer;
}

.price div:hover{
    background-color: $primary;
    color: white;
    border-color: $primary;
}


tr{
    height: 70px;
}

.selection-row th {
    vertical-align: middle;
}

.selection-row td {
    vertical-align: middle;
}

.sel-number div {
    text-align: center;
    width: 30px;
}


.sel-info {
    width: 100%;
}

.sel-img {
    min-width: 65px;
}

.sel-img img {
  width: 40px;
}

.jockey {
    font-size: 0.7em;
}

.j-icon {
    font-family: 'Quicksand';
    font-size: 0.8em;
    display: inline-block;
    background-color: $primary;
    color: white;
    width: 12px;
    text-align: center;
    border-radius: 100%;
    margin-right: 5px;
}

</style>
