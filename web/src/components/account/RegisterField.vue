<template>
    <div class="field">
        <div class="label-container">
            <label class="label">{{ field.name }}</label>
            <div class="error">{{ field.error }}</div>
        </div>
        <div class="control">
            <input class="input"
                v-on:blur="validate"
                v-on:focus="activate"
                :type = "field.type"
                v-model="field.raw_value"
                :placeholder="field.placeholder"
                v-bind:class="{'is-danger': field.showError()}"
                :readonly="readonly"
            >
        </div>
    </div>
</template>

<script>
export default {
  name: 'RegisterField',
  props: ['field', 'readonly'],
  data() {
    return {

    };
  },
  computed: {

  },
  methods: {
    validate() {
      this.field.validate();
    },
    activate() {
      if (!this.readonly) {
        this.field.activate();
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.field .label {
  display:inline;
}

.field .error {
  display: inline-block;
  margin-left: 10px;
  margin-bottom: 3px;
  vertical-align: bottom;
  font-size: 0.7em;
  color: red;
}

input.input {
  width: 290px;
}
</style>
