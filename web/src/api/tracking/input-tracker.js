const ACTION_TYPE = {
  PRESS: 'press',
  AUTOFILL: 'autofill',
  CLEAR: 'clear',
  PASTE: 'paste',
};

const INPUT_TIMEOUT = 10;

function parseAction({ name, events }) {
  const clearEvents = events.filter(e => e.constructor.name === 'InputEvent'
    && (e.inputType === 'deleteContentBackward' || e.inputType === 'deleteWordBackward'));
  const pressEvents = events.filter(e => e.constructor.name === 'KeyboardEvent');
  const pasteEvents = events.filter(e => e.constructor.name === 'ClipboardEvent');
  const autoFillEvents = events.filter(e => e.constructor.name === 'Event');

  let type = '';
  let val = '';

  if (clearEvents.length > 0) {
    type = ACTION_TYPE.CLEAR;
    val = clearEvents[0].target.value;
  } else if (pressEvents.length > 0) {
    type = ACTION_TYPE.PRESS;
    val = pressEvents[0].target.value;
  } else if (pasteEvents.length > 0) {
    type = ACTION_TYPE.PASTE;
    val = pasteEvents[0].target.value;
  } else if (autoFillEvents.length > 0) {
    type = ACTION_TYPE.AUTOFILL;
    val = autoFillEvents[0].target.value;
  }

  if (type === '') return null;
  return {
    name,
    type,
    value: val,
  };
}

class InputTracker {
  constructor(name, finaliser) {
    this.name = name;
    this.events = [];
    this.finaliser = finaliser;

    this.timeout = null;
  }

  onEvent(event) {
    if (this.timeout) {
      clearTimeout(this.timeout);
      this.active = false;
    }

    const context = this;
    this.timeout = setTimeout(() => {
      context.active = false;
      context.timeout = null;
      const action = parseAction(context);
      if (action) {
        context.finaliser(action);
      }
      context.events = [];
    }, INPUT_TIMEOUT);
    this.active = true;
    this.events.push(event);
  }
}

export default InputTracker;
