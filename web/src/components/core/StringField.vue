<template>
    <div class="field">
        <div class="label-container">
            <label class="label">{{ field.name }}</label>
            <div class="error">{{ field.error }}</div>
        </div>
        <div class="control">
            <input class="input"
                v-on:blur="field.validate"
                v-on:focus="activate"
                :type = "field.type"
                v-model="field.value"
                :placeholder="field.placeholder"
                v-bind:class="{'is-danger': field.showError()}"
                :readonly="readonly"
            >
        </div>
    </div>
</template>

<script>
export default {
  name: 'StringField',
  props: ['field', 'readonly'],
  data() {
    return {

    };
  },
  computed: {

  },
  methods: {
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
  width: 150px;
  vertical-align: bottom;
  font-size: 0.7em;
  color: red;
}
</style>
