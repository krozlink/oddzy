import api from '../../api/racing';


const getters = {

};

const actions = {
  getRacingSchedule({ commit }) {
    commit('setLoadingStatus', null);
    api.readSchedule()
      .then((result) => {
        commit('updateRaceSchedule', { data: result.data });
        commit('setLoadingStatus', 'successful');
      })
      .catch(() => commit('setLoadingStatus', 'failed'));
  },
};

const mutations = {

  updateRaceSchedule(state, { data }) {
    Object.values(data.meetings).forEach((m) => {
      state.meetings[m.meeting_id] = { ...m, date: data.date };
    });

    Object.values(data.races).forEach((r) => {
      state.races[r.race_id] = r;
    });
  },

  setLoadingStatus(state, status) {
    state.loadingStatus = status;
  },
};

const state = {
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
