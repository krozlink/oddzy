import Vue from 'vue';
import api from '../../api/racing';

const SCHEDULE_CACHE_TIME = 10 * 1000; // 10 seconds

const getters = {
  filterMeetings: state => (type, date, local) => Object.values(state.meetings).filter(m => m.race_type === type && m.date === date && m.local === local),
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
};

const mutations = {

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
};

const state = {
  scheduleAges: {},
  races: { },
  meetings: { },
  selections: { },
  loadingStatus: null,
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
