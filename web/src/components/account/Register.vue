<template>
  <div id="register" class="modal" :class="{'is-active': this.visible}">
    <div class="modal-background"></div>
    <div class="modal-content">
      <header class="modal-card-head">
          <p class="modal-card-title">Create New Account</p>
          <button v-on:click="close" :disabled="isReadonly" class="delete" aria-label="close"></button>
      </header>
      <section class="modal-card-body">
        <div class="field is-horizontal">
          <div class="field-body">
            <string-field :field="fields.first_name" :readonly="isReadonly">
            </string-field>
            <string-field :field="fields.last_name" :readonly="isReadonly"></string-field>
          </div>
        </div>

        <div class="field is-horizontal">
            <div class="field-body">
              <string-field :field="fields.email_address" :readonly="isReadonly"></string-field>
            </div>
        </div>

        <div class="field is-horizontal">
            <div class="field-body">
              <string-field :field="fields.user_name" :readonly="isReadonly"></string-field>
              <string-field :field="fields.password" :readonly="isReadonly"></string-field>
            </div>
        </div>

        <string-field :field="fields.address" :readonly="isReadonly"></string-field>

        <div class="field is-horizontal">
          <div class="field-body">
            <string-field :field="fields.date_of_birth" :readonly="isReadonly"></string-field>
            <string-field :field="fields.mobile_number" :readonly="isReadonly"></string-field>
          </div>
        </div>
      </section>
      <footer class="modal-card-foot">
          <button class="button is-info"
          v-on:click="register"
          :class="{'is-loading': isReadonly}"
          :disabled="isReadonly">Register</button>
          <button class="button" v-on:click="close" :disabled="isReadonly">Cancel</button>
          <label>{{registerStatus}}</label>
      </footer>
    </div>
  </div>
</template>

<script>
import StringField from '../core/StringField.vue';
import InputValue from '../../api/input_value';

export default {
  components: {
    StringField,
  },
  data() {
    return {
      fields: {
        first_name: new InputValue('First Name', this.validateFirstName),
        last_name: new InputValue('Last Name', this.validateLastName),
        email_address: new InputValue('Email Address', this.validateEmail),
        user_name: new InputValue('User Name', this.validateUserName),
        password: new InputValue('Password', this.validatePassword, { type: 'password' }),
        address: new InputValue('Address', this.validateAddress),
        date_of_birth: new InputValue('Date of Birth', this.validateDOB, { placeholder: 'DD / MM / YYYY' }),
        mobile_number: new InputValue('Mobile', this.validateMobile, { placeholder: '04________' }),
      },
    };
  },
  computed: {
    visible() {
      return this.$store.state.account.display_register;
    },
    isReadonly() {
      return this.$store.state.account.register_status === 'registering';
    },
    registerStatus() {
      if (this.$store.state.account.register_status === 'registering') {
        return 'Registering...';
      }
      return '';
    },
  },
  methods: {
    close() {
      this.$store.dispatch('account/displayRegister', false)
        .then(() => this.reset());
    },

    reset() {
      Object.values(this.fields).forEach((f) => {
        f.reset();
      });
    },

    checkUsername() {

    },

    validateMandatory(f) {
      let isValid = false;
      const field = f;
      if (field.value.trim() === '') {
        field.error = `${field.name} is mandatory`;
      } else {
        field.error = '';
        isValid = true;
      }
      field.isValid = isValid;
      return isValid;
    },

    validateFirstName() {
      return this.validateMandatory(this.fields.first_name);
    },

    validateLastName() {
      return this.validateMandatory(this.fields.last_name);
    },

    validateEmail() {
      return this.validateMandatory(this.fields.email_address);
    },

    validateUserName() {
      return this.validateMandatory(this.fields.user_name);
    },

    validatePassword() {
      return this.validateMandatory(this.fields.password);
    },

    validateAddress() {
      return this.validateMandatory(this.fields.address);
    },

    validateDOB() {
      return this.validateMandatory(this.fields.date_of_birth);
    },

    validateMobile() {
      return this.validateMandatory(this.fields.mobile_number);
    },

    validate() {
      /* eslint-disable no-bitwise */
      let isValid = true;
      Object.values(this.fields).forEach((f) => {
        f.activate();
        isValid &= f.validate();
      });
      /* eslint-disable no-bitwise */

      return isValid;
    },
    register() {
      if (this.validate()) {
        this.$store.dispatch('account/register', this.fields);
      }
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

// div.field {
//   width: 295px;
// }


</style>
