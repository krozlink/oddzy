<template>
    <div class="field">
        <div class="label-container">
            <label class="label">{{ field.description }}</label>
            <div class="error">{{ field.error }}</div>
        </div>
        <div class="control">
            <input class="input"
                @blur="validate"
                @focus="activate"
                @paste="onEvent"
                @animationstart="onEvent"
                @webkitAnimationStart="onEvent"
                @keypress="onEvent"
                @input="onEvent"

                ref="input"
                :name="field.name"
                :autocomplete="field.autocomplete"
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
import { InputTracker } from '../../api/tracking';

export default {
  name: 'RegisterField',
  props: {
    field: Object,
    readonly: Boolean,
  },
  data() {
    return {
      tracker: new InputTracker(this.field.name, this.onAction),
    };
  },
  computed: {

  },
  methods: {
    validate() {
      return this.field.validate();
    },
    activate() {
      if (!this.readonly) {
        this.field.activate();
      }
    },
    onEvent(e) {
      this.tracker.onEvent(e);
    },
    onAction(action) {
      // console.log(tracker);
      console.log(`${action.name} - ${action.type} action with new value '${action.value}'`);
    },
  },
  mounted() {
    this.$refs.input.addEventListener('animationend', this.animend);
  },
  beforeDestroy() {
    this.$refs.input.removeEventListener('animationend', this.animend);
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

@keyframes onAutoFillStart {  from {/**/}  to {/**/}}
@keyframes onAutoFillCancel {  from {/**/}  to {/**/}}
input:-webkit-autofill {
    // Expose a hook for JavaScript when autofill is shown
    // JavaScript can capture 'animationstart' events
    animation-name: onAutoFillStart;

    // Make the background color become yellow really slowly
    transition: background-color 50000s ease-in-out 0s;
}
input:not(:-webkit-autofill) {
    // Expose a hook for JS onAutoFillCancel
    // JavaScript can capture 'animationstart' events
    animation-name: onAutoFillCancel;
}
</style>
