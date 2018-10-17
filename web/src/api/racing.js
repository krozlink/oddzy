import axios from 'axios';
import date from './date-helper';

import test from '../test_data/racing';

function readSchedule(scheduleDate) {
  const param = scheduleDate === '' ? date.formatDate(date.todayDate()) : scheduleDate;

  if (process.env.VUE_APP_SERVERLESS_ONLY) {
    return Promise.resolve({
      data: test.RaceSchedule(param),
    });
  }
  return axios.get(`${process.env.VUE_APP_API}/racing/schedule?date=${param}`);
}

function readRaceCard(raceId) {
  if (process.env.VUE_APP_SERVERLESS_ONLY) {
    return Promise.resolve({
      data: test.RaceCard(raceId),
    });
  }

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
