<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column is-one-third">
          <h4 class="title is-4">Today's Racing</h4>
        </div>
        <div class="column">
          <nav class="level">
            <div class="level-left">
              <div class="level-item">
                <div class="buttons has-addons">
                  <span class="button is-primary is-light is-selected">Today</span>
                  <span class="button is-light">Tomorrow</span>
                  <span class="button is-light">Saturday</span>
                  <span class="button is-light">
                    <span class="icon is-large">
                      <i class="fas fa-lg fa-calendar-alt"></i>
                    </span>
                    <span>Select Date</span>
                  </span>
                </div>
              </div>
            </div>
            <div class="level-right">
              <div class="level-item">
                <div class="buttons has-addons">
                <span class="button is-primary is-light is-selected">All</span>
                <span class="button is-light">Horses</span>
                <span class="button is-light">Harness</span>
                <span class="button is-light">Greyhounds</span>
              </div>
              </div>
            </div>
          </nav>
        </div>
      </div>

      <div class="columns race-catagory ">
        <div class="column race-location is-one-fifth">
          Flemington
        </div>
        <div class="column race-list">
          <div class="columns">
            <div class="column race-item">
              Race 1
            </div>
            <div class="column race-item">
              Race 2
            </div>
            <div class="column race-item">
              Race 3
            </div>
            <div class="column race-item">
              Race 4
            </div>
            <div class="column race-item">
              Race 5
            </div>
            <div class="column race-item">
              Race 6
            </div>
            <div class="column race-item">
              Race 7
            </div>
            <div class="column race-item">
              Race 8
            </div>
            <div class="column race-item">
              Race 9
            </div>
            <div class="column race-item">
              Race 10
            </div>
            <div class="column race-item">
              Race 11
            </div>
            <div class="column race-item">
              Race 12
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>


<style lang="scss" scoped>

div.container .race-location {
  background-color: lightblue;
}

div.container .race-item {
  background-color: lightcoral;
  border-style: solid;
  border-width: 1px;
}

</style>