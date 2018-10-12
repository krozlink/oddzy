<template>
  <div id="register" class="modal" :class="{'is-active': this.visible}">
    <div class="modal-background"></div>
    <div class="modal-content">
      <form>
        <header class="modal-card-head">
            <p class="modal-card-title">Create New Account</p>
            <button v-on:click="close" :disabled="isReadonly" class="delete" aria-label="close"></button>
        </header>
        <section class="modal-card-body">
          <div class="field is-horizontal">
            <div class="field-body">
              <register-field :field="fields.first_name" :readonly="isReadonly"></register-field>
              <register-field :field="fields.last_name" :readonly="isReadonly"></register-field>
            </div>
          </div>

          <div class="field is-horizontal">
              <div class="field-body">
                <register-field :field="fields.email_address" :readonly="isReadonly"></register-field>
                <register-field :field="fields.user_name" :readonly="isReadonly"></register-field>
              </div>
          </div>

          <div class="field is-horizontal">
              <div class="field-body">
                <register-field :field="fields.password" :readonly="isReadonly"></register-field>
                <register-field :field="fields.confirm_password" :readonly="isReadonly"></register-field>
              </div>
          </div>

          <div class="field is-horizontal">
              <div class="field-body">
                <register-field :field="fields.address" :readonly="isReadonly"></register-field>
                <register-field :field="fields.mobile_number" :readonly="isReadonly"></register-field>
              </div>
          </div>

          <div class="field is-horizontal">
            <div class="field-body">
              <register-field :field="fields.date_of_birth" :readonly="isReadonly"></register-field>
            </div>
          </div>

          <div class="field">
              <div class="control">
                  <label class="checkbox">
                  <input type="checkbox"
                    v-model="fields.agree.raw_value"
                    v-on:blur="validateAgree"
                    :disabled="isReadonly"
                    >
                  I agree to the terms and conditions
                  </label>
                  <div class="error">{{ fields.agree.error }}</div>
              </div>
          </div>

        </section>
        <footer class="modal-card-foot">
            <button class="button is-info"
            v-on:click="register"
            type="button"
            :class="{'is-loading': isReadonly}"
            :disabled="isReadonly">Register</button>
            <button class="button" v-on:click="close" :disabled="isReadonly">Cancel</button>
            <label>{{this.$store.state.account.status_message}}</label>
        </footer>
      </form>
    </div>
  </div>
</template>

<script>
import RegisterField from './RegisterField.vue';
import InputValue from '../../api/input_value';
import Validation from '../../api/input_validation';
import DateHelper from '../../api/date-helper';


export default {
  components: {
    RegisterField,
  },
  data() {
    return {
      fields: {
        first_name: new InputValue('First Name', 'given_name', [Validation.Mandatory]),
        last_name: new InputValue('Last Name', 'family_name', [Validation.Mandatory]),
        email_address: new InputValue('Email Address', 'email', [Validation.Mandatory, Validation.EmailAddress]),
        user_name: new InputValue('User Name', null, [Validation.Mandatory]),
        password: new InputValue('Password', null, [Validation.Mandatory, Validation.Password], { type: 'password' }),
        confirm_password: new InputValue('Confirm Password', null, [this.validatePasswordsMatch], { type: 'password' }),
        address: new InputValue('Address', 'address', [Validation.Mandatory]),
        date_of_birth: new InputValue('Date of Birth', 'birthdate', [this.validateDOB], this.dobOptions()),
        mobile_number: new InputValue('Mobile Number', 'phone_number', [Validation.Mandatory, Validation.MobileNumber], this.mobileOptions()),
        agree: new InputValue('Agree', null, [this.validateIAgreeSelected], { type: 'checkbox' }),
      },
    };
  },
  computed: {
    visible() {
      return this.$store.state.account.display_register;
    },
    isReadonly() {
      return this.$store.state.account.status === 'registration_submitted';
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
      // TODO - ensure username is unique
    },

    mobileOptions() {
      return {
        placeholder: '04________',
        parser: val => `+61${val.substring(1)}`,
      };
    },

    dobOptions() {
      return {
        placeholder: 'DD / MM / YYYY',
        parser: (val) => {
          const date = DateHelper.parseDateString(val);
          const year = date.getFullYear().toString();
          const month = (date.getMonth() + 1).toString().padStart(2, '0');
          const day = date.getDate().toString().padStart(2, '0');
          return `${year}-${month}-${day}`;
        },
      };
    },

    validatePasswordsMatch() {
      return Validation.Match(this.fields.password, this.fields.confirm_password);
    },

    validateDOB() {
      return Validation.Mandatory(this.fields.date_of_birth)
        && Validation.IsDate(this.fields.date_of_birth)
        && Validation.MinimumAge(this.fields.date_of_birth, 18);
    },

    validateAgree() {
      this.fields.agree.validate();
    },

    validateIAgreeSelected() {
      return Validation.IsTrue(this.fields.agree, 'You must agree to the terms and conditions');
    },

    validateAll() {
      let isValid = true;
      Object.values(this.fields).forEach((f) => {
        f.activate();
        isValid = f.validate() && isValid;
      });

      return isValid;
    },
    register() {
      if (this.validateAll()) {
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

.field .error {
  display: inline-block;
  margin-left: 10px;
  margin-bottom: 3px;
  vertical-align: bottom;
  font-size: 0.7em;
  color: red;
}

.password-instruction {
  font-size: 0.7em;
  // line-height: 16px;
  // flex-grow: 1;
  // flex-shrink: 1;
  // margin-bottom: 0px;
  // position: relative;
  // max-width: 100%;
  width: 290px;
  align-items: flex-start;
}


</style>
