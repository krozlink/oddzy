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
              Race goes here
            </div>
            <div>
              And another race here
            </div>
          </article>
        </div>
        <div class="tile is-7 is-parent is-vertical">
          <article class="tile is-child  notification is-primary">
            <div class="title">
              Some Big Promotion
            </div>
            <div class="subtitle">
              It's really awesome
            </div>
            <div class="subtitle">
              Click here
            </div>
          </article>

          <article class="tile is-child  notification is-primary">
            <div class="subtitle">Live Events</div>
            <ul>
              <li>Chelsea vs Manchester United</li>
              <li>Leicester v Wolverhampton</li>
              <li>3rd Test England v India</li>
            </ul>
          </article>
        </div>
        <div class="tile is-parent is-vertical">
            <article class="tile is-child  notification is-primary">
              <div class="subtitle">
                Rewards Program
              </div>
            </article>
            <article class="tile is-child  notification is-primary">
              <div class="subtitle">
                Download App
              </div>
            </article>
            <article class="tile is-child  notification is-primary">
              <div class="subtitle">
                Gamble Responsibly
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
    padding-top: 20px;
  }
</style>

