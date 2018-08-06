<template>
  <div id="app">

    <section class="hero is-dark">
      <div class="column hero-body is-offset-1">
        <h1 class="title">
          Oddzy
        </h1>
      </div>
    </section>


    <div class="columns">
      <div class="column is-1">
      </div>
      <div class="column is-2">
        <nav class="menu">
          <ul class="menu-list">
            <li>
              <a>
                <span class="icon">
                  <i class="fas fa-home"></i>
                </span>
              Home
              </a>
            </li>
            <li>
              <a>
                <span class="icon">
                  <i class="fab fa-sticker-mule"></i>
                </span>
              Racing
              </a>
            </li>
            <li>
              <a class="is-disable">
                <span class="icon">
                  <i class="fas fa-football-ball"></i>
                </span>
                Sports
              </a>
            </li>
            <li>
              <a class="is-disable">
                <span class="icon">
                  <i class="fas fa-dollar-sign"></i>
                </span>
                Promotions
              </a>
            </li>
            <li>
              <a class="is-disable">
                <span class="icon">
                  <i class="fas fa-user"></i>
                </span>
                My Account
              </a>
            </li>
            <li>
              <a class="is-disable">
                <span class="icon">
                  <i class="fas fa-phone"></i>
                </span>
                Contact Us
              </a>
            </li>
          </ul>
        </nav>
      </div>
      <div class="column" id="content">
        
      </div>
      <div class="column is-1"></div>
    </div>


  </div>
</template>

<style lang="scss">

nav li {
  height: 50px;
}

.is-disable{
  pointer-events: none;
  opacity: .65;
}

// #content {
//   background-color: cornsilk;
// }



// nav {
//   border-width: 1px;
//   border-style: solid;
// }

// div {
//   border-width: 1px;
//   border-style: solid;
// }
</style>
