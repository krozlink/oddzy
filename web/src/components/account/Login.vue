<template>
    <div class="modal" :class="{'is-active': visible}">
        <div class="modal-background"></div>
        <div class="modal-content">
            <header class="modal-card-head">
                <p class="modal-card-title">Login</p>
                <button v-on:click="close" class="delete" aria-label="close"></button>
            </header>
            <section class="modal-card-body">
                <div class="field">
                    <p class="control has-icons-left">
                        <input class="input" type="text" placeholder="User Name">
                        <span class="icon is-small is-left">
                        <i class="fas fa-user"></i>
                        </span>
                    </p>
                </div>
                <div class="field">
                    <p class="control has-icons-left">
                        <input class="input" type="password" placeholder="Password">
                        <span class="icon is-small is-left">
                        <i class="fas fa-lock"></i>
                        </span>
                    </p>
                </div>
                <div class="field">
                    <div class="control">
                        <label class="checkbox">
                        <input type="checkbox">
                        Remember Me
                        </label>
                    </div>
                </div>
            </section>
            <footer class="modal-card-foot">
                <button class="button is-info">Login</button>
                <button class="button" v-on:click="close">Cancel</button>
                <button class="button is-text">Forgot Password?</button>
            </footer>
        </div>
    </div>
</template>

<script>
export default {
  data() {
    return {
    };
  },
  computed: {
    visible() {
      return this.$store.state.account.display_login;
    },
  },
  methods: {
    close() {
      this.$store.dispatch('account/displayLogin', false);
    },
  },
};
</script>

<style lang="scss" scoped>
header.modal-card-head {
    background-color: $primary;
}

.modal-card-title {
    color: white;
}

div.modal-content {
    width: 400px;
}
</style>
