<template>
    <tr :class="{
            'scratched': selection.is_scratched
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
            <div :class="{
                button: true,
                'is-disable': race.status !== 'OPEN'
                }">
                4.00
            </div>
        </td>
        <td class="price fixed place">
            <div :class="{
                button: true,
                'is-disable': race.status !== 'OPEN'
                }">
                1.15
            </div>
        </td>
        <td class="price tote win">
            <div :class="{
                button: true,
                'is-disable': race.status !== 'OPEN'
                }">
                6.80
            </div>
        </td>
        <td class="price tote place">
            <div :class="{
                button: true,
                'is-disable': race.status !== 'OPEN'
                }">
                1.10
            </div>
        </td>
    </tr>
</template>

<script>
export default {
  name: 'RaceSelection',
  props: ['selection', 'race', 'meeting'],
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
  },
};
</script>

<style lang="scss" scoped>

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
    margin-top: 10px;

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
