<template>
  <div id="racing" class="container">
    <div class="section columns header">
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
            <div class="level-item racetype-icons">
              <div class="buttons has-addons">
              <span class="button is-primary is-light is-selected">All</span>
              <span class="button is-light">
                <span class="icon">
                  <i class="icon-thoroughbred fa-3x"></i>
                </span>
              </span>
              <span class="button is-light">
                <span class="icon">
                  <i class="icon-harness fa-3x"></i>
                </span>
              </span>
              <span class="button is-light">
                <span class="icon">
                  <i class="icon-greyhound fa-3x"></i>
                </span>
              </span>
            </div>
            </div>
          </div>
        </nav>
      </div>
    </div>

    <div class="section race-section">
      <div class="columns">
        <div class="column">
          <h6 class="title is-6">Horse Racing - Australia & New Zealand</h6>
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

    <div class="section race-section">
      <div class="columns">
        <div class="column">
          <h6 class="title is-6">Horse Racing - International</h6>
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


  </div>
</template>


<style lang="scss" scoped>

// div.container .race-location {
//   background-color: lightblue;
// }

// div.container .race-item {
//   background-color: lightcoral;
//   border-style: solid;
//   border-width: 1px;
// }

.racetype-icons .button {
  width: 70px;
}

.racetype-icons .icon {
  padding-top: 5px;
}

#racing .section  {
  padding-top: 16px;
}

#racing .header  {
  padding-bottom: 0px;
}

</style>
