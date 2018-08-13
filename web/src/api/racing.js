import axios from 'axios';

function readSchedule() {
  console.log(`${process.env.VUE_APP_API}/racing/schedule`);
  return axios.get(`${process.env.VUE_APP_API}/racing/schedule?date=2018-08-11`);
}

export default {
  readSchedule,
};
