<template>
    <div class="card">
      <div class="card-header-title">
        <div class="bsi-header">
          <div class="bsi-row">
            <div class="entrant">
              <div class="runner">
                Redzel
              </div>
              <div class="meeting">
                Flemington R6
              </div>
            </div>
            <div class="odds">
              <div class="price">
                8.40
              </div>
              <div class="bet-type">
                Win Fixed
              </div>
            </div>
            <button class="delete is-small"></button>
          </div>
        </div>
      </div>
      <div class="card-content">
        something
      </div>
    </div>
</template>

<script>
export default {
  name: 'BetslipItem',
  props: 'bet',
  data() {
    return {};
  },
  computed: {

  },
  methods: {

  },
};
</script>

<style lang="scss" scoped>

.card {
  margin-left: 1px;
  margin-right: 1px;
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
