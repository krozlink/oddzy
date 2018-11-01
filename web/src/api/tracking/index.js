import AWS from 'aws-sdk';

AWS.config.region = process.env.VUE_APP_AWS_REGION;

const PUSH_INTERVAL = 2000;
const IDLE_INTERVAL = 30000;

const Kinesis = new AWS.Kinesis({
  apiVersion: '2013-12-02',
});


let lastSend = 0;
let timer = 0;

function Send() {
  lastSend = new Date().getTime();

  Kinesis.putRecords({
    Records: this.$store.state.tracker.actions,
    StreamName: process.env.VUE_APP_AWS_KINESIS_STREAM,
  }, (err, data) => {
    if (err) console.error(err);
    else console.log(data);
  });
}

function Pause() {
  clearInterval(timer);
  timer = 0;
}

function Trigger() {
  if (this.$store.state.tracker.actions.length > 0) {
    Send();
  } else if (new Date().getTime() - lastSend > IDLE_INTERVAL) {
    Pause();
  }
}

function Initialise() {
  timer = setInterval(Trigger, PUSH_INTERVAL);
}

function Update() {
  if (timer === 0) {
    Initialise();
  }
}


export default {
  Update,
};
