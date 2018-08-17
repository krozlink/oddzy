import axios from 'axios';
import date from './date-helper';

function readSchedule(scheduleDate) {
  const param = scheduleDate === '' ? date.formatDate(date.todayDate()) : scheduleDate;
  return axios.get(`${process.env.VUE_APP_API}/racing/schedule?date=${param}`);
}

function readRaceCard(raceId) {
  return axios.get(`${process.env.VUE_APP_API}/racing/racecard?race_id=${raceId}`);
}

function raceNameURL(name) {
  return name.toLowerCase().replace(' ', '-');
}

export default {
  readSchedule,
  readRaceCard,
  raceNameURL,
};
