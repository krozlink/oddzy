import axios from 'axios';

function readSchedule() {
  axios.get(`${process.env.VUE_APP_API}/racing/schedule`);
}

export default {
  readSchedule,
};
