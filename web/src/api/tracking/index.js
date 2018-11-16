import AWS from 'aws-sdk';

AWS.config.region = process.env.VUE_APP_AWS_REGION;
AWS.config.credentials = new AWS.CognitoIdentityCredentials({
  IdentityPoolId: process.env.VUE_APP_AWS_COGNITO_IDENTITY_POOL,
  Logins: {},
});

const PUSH_INTERVAL = 2000;

const Firehose = new AWS.Firehose();

class Tracker {
  constructor(store) {
    this.store = store;
    this.lastSend = 0;
    this.timer = 0;
    this.lastAction = 0;
  }

  createRecords(session, actions) {
    const subset = actions.slice(this.lastAction, 500);

    return subset.map((a) => {
      const data = JSON.stringify({
        timestamp: a.timestamp,
        area: a.area,
        type: a.type,
        data: a.data,
        session,
      });

      return {
        Data: `${data}\n`,
      };
    });
  }

  Send() {
    this.lastSend = new Date().getTime();
    const records = this.createRecords(this.store.state.tracker.session, this.store.state.tracker.actions);
    Firehose.putRecordBatch({
      Records: records,
      DeliveryStreamName: process.env.VUE_APP_AWS_KINESIS_FIREHOSE_STREAM,
    }, (err, data) => {
      if (err) console.error(err);
      else {
        this.lastAction += records.length;
        console.log(data);
      }
    });
  }

  Start() {
    const t = this;
    this.timer = setInterval(() => {
      if (t.store.state.tracker.actions.length > t.lastAction) {
        t.Send();
      }
    }, PUSH_INTERVAL);
  }
}


export default Tracker;
