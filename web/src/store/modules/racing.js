import Vue from 'vue';
import api from '../../api/racing';
import test from '../../api/test-prices';

const SCHEDULE_CACHE_TIME = 30 * 1000; // 10 seconds
const RACECARD_CACHE_TIME = 30 * 1000; // 10 seconds

const getters = {
  filterMeetings: state => (type, date, local) => Object.values(state.meetings).filter(m => m.race_type === type && m.date === date && m.local === local),
  getRacesForMeeting: state => (meetingId) => {
    const ids = state.meetings[meetingId].race_ids;
    return ids.map(id => state.races[id]);
  },
};

const actions = {
  getRacingSchedule({ commit, state }, date) {
    // if the schedule for this date is less than 30 seconds old then do not re-retrieve the data
    if (state.scheduleAges[date] && state.scheduleAges[date] > new Date(new Date() - SCHEDULE_CACHE_TIME).getTime()) {
      commit('setLoadingStatus', 'successful');
    } else {
      commit('setLoadingStatus', null);
      api.readSchedule(date)
        .then((result) => {
          commit('updateRaceSchedule', { data: result.data, date });
          commit('setLoadingStatus', 'successful');
        })
        .catch(() => commit('setLoadingStatus', 'failed'));
    }
  },
  getRaceCard({ commit, state }, raceId) {
    if (state.races[raceId] && state.races[raceId].lastRead > new Date(new Date() - RACECARD_CACHE_TIME).getTime()) {
      commit('setLoadingStatus', 'successful');
    } else {
      commit('setLoadingStatus', null);
      api.readRaceCard(raceId)
        .then((result) => {
          commit('updateRaceCard', { data: result.data, raceId });
          commit('setLoadingStatus', 'successful');
        })
        .catch(() => commit('setLoadingStatus', 'failed'));
    }
  },
  updatePrice({ commit }, update) {
    commit('updatePrice', update);
  },
};

const mutations = {
  updateRaceCard(state, { data, raceId }) {
    Vue.set(state.races, raceId, { ...data, lastRead: new Date().getTime() });
  },

  updateRaceSchedule(state, { data, date }) {
    Object.values(data.meetings).forEach((m) => {
      const local = m.country === 'Australia' || m.country === 'New Zealand';

      Vue.set(state.meetings, m.meeting_id, { ...m, date: data.date, local });
    });

    Object.values(data.races).forEach((r) => {
      Vue.set(state.races, r.race_id, r);
    });

    Vue.set(state.scheduleAges, date, new Date().getTime());
  },

  setLoadingStatus(state, status) {
    state.loadingStatus = status;
  },

  updatePrice(state, update) {
    const price = { ...state.prices[update.selection_id] };
    if (update.type === 'win') {
      if (price.win.price < update.price) {
        price.win.change = 'increase';
      } else if (price.win.price > update.price) {
        price.win.change = 'decrease';
      } else {
        price.win.change = 'none';
      }
      price.win.price = update.price;
    } else {
      if (price.place.price < update.price) {
        price.place.change = 'increase';
      } else if (price.place.price > update.price) {
        price.place.change = 'decrease';
      } else {
        price.place.change = 'none';
      }
      price.place.price = update.price;
    }

    Vue.set(state.prices, update.selection_id, price);
  },
};

function initialTestPrices() {
  const prices = {};

  for (let i = 1; i <= 12; i += 1) {
    prices[i] = {
      win: {
        price: test.WinPrices[i - 1],
        change: 'none',
      },
      place: {
        price: test.PlacePrices[i - 1],
        change: 'none',
      },
    };
  }

  return prices;
}

const state = {
  scheduleAges: {},
  races: { },
  meetings: { },
  selections: { },
  prices: initialTestPrices(),
  loadingStatus: null,
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
