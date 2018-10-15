<template>
  <div id="home" class="container">
    <div class="section">
      <div class="tile is-ancestor">
        <div class="tile is-parent is-vertical">
          <article class="tile  notification is-primary is-child">
            <div class="title is-6">
              Featured Events
            </div>
            <ul>
              <li>List Events Here</li>
              <li>Today's Racing</li>
              <li>AFL Matches</li>
              <li>NRL Matches</li>
              <li>Premier League</li>
              <li>MLB Matches</li>
            </ul>
          </article>
          <article class="tile is-child  notification is-primary">
            <div class="subtitle">
              Next To Jump
            </div>
            <div>
              List Races
            </div>
          </article>
        </div>
        <div class="tile is-7 is-parent is-vertical">
          <article class="tile is-child  notification is-primary">
            <div class="title">
              Promotion Banner
            </div>
            <div class="subtitle">
              Subtitle
            </div>
            <div class="subtitle is-6">
              More details
            </div>
          </article>

          <article class="tile is-child  notification is-primary">
            <div class="subtitle">Live Events</div>
            <ul>
              <li>Sample Live Events</li>
              <li>Chelsea vs Manchester United</li>
              <li>Leicester v Wolverhampton</li>
              <li>3rd Test England v India</li>
            </ul>
          </article>
        </div>
        <div class="tile is-parent is-vertical">
            <article class="tile is-child  notification is-primary">
              <div class="subtitle is-6">
                Rewards Program
              </div>
            </article>
            <article class="tile is-child  notification is-primary">
              <div class="subtitle is-6">
                Download App
              </div>
            </article>
            <article class="tile is-child  notification is-primary">
              <div class="subtitle is-6">
                Invite Friends
              </div>
            </article>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'home',
  components: {},
};
</script>

<style lang="scss" scoped>
  #home.container{
    @media screen and (min-width: 1472px) {
      width:1100px;
    }
  }

  #home .section {
    padding-top: 16px;
  }
</style>

