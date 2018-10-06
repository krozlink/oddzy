<template>
    <div class="modal" :class="{'is-active': this.visible}">
        <div class="modal-background"></div>
        <div class="modal-content">
            <header class="modal-card-head">
                <p class="modal-card-title">Create New Account</p>
                <button v-on:click="close" class="delete" aria-label="close"></button>
            </header>
            <section class="modal-card-body">
                <div class="field is-horizontal">
                    <div class="field-body">
                        <div class="field">
                            <label class="label">First Name</label>
                            <div class="control">
                                <input class="input" type="text" placeholder="First Name">
                            </div>
                        </div>
                        <div class="field">
                            <label class="label">Last Name</label>
                            <div class="control">
                                <input class="input" type="text" placeholder="Last Name">
                            </div>
                        </div>
                    </div>
                </div>

                <div class="field is-horizontal">
                    <div class="field-body">
                        <div class="field">
                        <label class="label">Email Address</label>
                            <div class="control">
                                <input class="input" type="email" placeholder="Email Address">
                            </div>
                        </div>
                    </div>
                </div>

                <div class="field is-horizontal">
                    <div class="field-body">
                        <div class="field">
                            <label class="label">User Name</label>
                            <div class="control">
                                <input class="input" type="text" placeholder="User Name">
                            </div>
                        </div>
                        <div class="field">
                            <label class="label">Password</label>
                            <div class="control">
                                <input class="input" type="password" placeholder="Password">
                            </div>
                        </div>
                    </div>
                </div>

                <div class="field">
                    <label class="label">Address</label>
                    <div class="control">
                        <input class="input" type="text" placeholder="Address">
                    </div>
                </div>
                <div class="field is-horizontal">
                    <div class="field-body">
                        <div class="field">
                            <label class="label">Date Of Birth</label>
                            <div class="control">
                                <input class="input" type="text" placeholder="DD / MM / YYYY">
                            </div>
                        </div>
                        <div class="field">
                            <label class="label">Mobile Number</label>
                            <div class="control">
                                <input class="input" type="text" placeholder="04________">
                            </div>
                        </div>
                    </div>
                </div>
            </section>
            <footer class="modal-card-foot">
                <button class="button is-info">Register</button>
                <button class="button" v-on:click="close">Cancel</button>
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
      return this.$store.state.account.display_register;
    },
  },
  methods: {
    close() {
      this.$store.dispatch('account/displayRegister', false);
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
</style>
