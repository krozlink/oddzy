<template>
  <div id="account" class="container">
    <div class="section columns header">
      <div class="column">
        <h4 class="title is-4">My Account</h4>
      </div>
    </div>
    <div>
      <div v-if="this.$store.state.account.authenticated" class="section">
        <div class="field">
          <label class="label">Name</label>
          <div class="control">
            <input class="input" type="text" disabled :value="this.name">
          </div>
        </div>

        <div class="field">
          <label class="label">Username</label>
          <div class="control">
            <input class="input" type="text" disabled :value="this.$store.state.account.user_details.username">
          </div>
        </div>

        <div class="field">
          <label class="label">Email Address</label>
          <div class="control">
            <input class="input" type="text" disabled :value="this.$store.state.account.user_details.email_address">
          </div>
        </div>

        <div class="field">
          <label class="label">Mobile Number</label>
          <div class="control">
            <input class="input" type="text" disabled :value="this.mobileNumber">
          </div>
        </div>

        <div class="field">
          <label class="label">Date Of Birth</label>
          <div class="control">
            <input class="input" type="text" disabled :value="this.dateOfBirth">
          </div>
        </div>

        <div class="field">
          <label class="label">Address</label>
          <div class="control">
            <input class="input" type="text" disabled :value="this.$store.state.account.user_details.address">
          </div>
        </div>
      </div>
      <div v-else  class="section">
        Login to view your account details
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Account',
  data() {
    return {

    };
  },

  computed: {
    name() {
      if (!this.$store.state.account.authenticated) return '';
      return `${this.$store.state.account.user_details.first_name} ${this.$store.state.account.user_details.last_name}`;
    },

    dateOfBirth() {
      if (!this.$store.state.account.authenticated) return '';
      const [year, month, day] = this.$store.state.account.user_details.date_of_birth.split('-');
      return `${day}/${month}/${year}`;
    },

    mobileNumber() {
      if (!this.$store.state.account.authenticated) return '';
      return `0${this.$store.state.account.user_details.phone_number.substring(3)}`;
    },
  },

  methods: {


  },
};
</script>

<style lang="scss" scoped>


#account.container{
  @media screen and (min-width: 1472px) {
    width:1100px;
  }
}

#account .header {
  padding-bottom: 0px;
}


#account .section {
  padding-top: 16px;
}

</style>
