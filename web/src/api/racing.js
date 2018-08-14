import axios from 'axios';
import date from './date-helper';

function readSchedule(scheduleDate) {
  const param = scheduleDate === '' ? date.formatDate(date.todayDate()) : scheduleDate;
  return axios.get(`${process.env.VUE_APP_API}/racing/schedule?date=${param}`);
}

export default {
  readSchedule,
};
