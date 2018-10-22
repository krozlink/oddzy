<template>
    <div id="login" class="modal" :class="{'is-active': visible}">
        <div class="modal-background"></div>
        <div class="modal-content">
            <form v-on:submit.prevent="login">
                <header class="modal-card-head">
                    <p class="modal-card-title">Login</p>
                    <button
                        v-on:click="close"
                        class="delete"
                        type="button"
                        aria-label="close"
                        :readonly="isReadonly">
                    </button>
                </header>
                <section class="modal-card-body">
                    <div class="field">
                        <p class="control has-icons-left">
                            <input
                                id="login-user"
                                class="input"
                                type="text"
                                name="username"
                                autocomplete="username"
                                placeholder="User Name"
                                :readonly="isReadonly"
                                v-model="username"
                                v-on:keyup.enter="login"
                            >
                                <span class="icon is-small is-left">
                                <i class="fas fa-user"></i>
                            </span>
                        </p>
                    </div>
                    <div class="field">
                        <p class="control has-icons-left">
                            <input
                                id="login-password"
                                class="input"
                                name="password"
                                autocomplete="current-password"
                                type="password"
                                placeholder="Password"
                                :readonly="isReadonly"
                                v-model="password"
                                v-on:keyup.enter="login"
                            >
                                <span class="icon is-small is-left">
                                <i class="fas fa-lock"></i>
                            </span>
                        </p>
                    </div>
                    <div class="field">
                        <div class="control">
                            <div id="status" class="label">{{this.$store.state.account.status_message}}</div>
                        </div>
                    </div>
                </section>
                <footer class="modal-card-foot">
                    <button type="button" class="button is-info" :readonly="isReadonly" v-on:click="login" :class="{'is-loading': isReadonly}">Login</button>
                    <button type="button" class="button" v-on:click="close" :readonly="isReadonly">Cancel</button>
                    <button type="button" class="button is-text" disabled>Forgot Password?</button>
                </footer>
            </form>
        </div>
    </div>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      password: '',
    };
  },
  computed: {
    visible() {
      return this.$store.state.account.display_login;
    },
    isReadonly() {
      return this.$store.state.account.status === 'login_submitted';
    },
  },
  methods: {
    close() {
      this.$store.dispatch('account/displayLogin', false);
    },
    login() {
      this.$store.dispatch('account/userLogin', {
        username: this.username,
        password: this.password,
      })
        .then(() => {
          if (this.$store.state.account.status === 'login_true') {
            this.username = '';
            this.password = '';
          }
        });
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

#status {
    font-size: 0.9em;
    color: red;
}
</style>
