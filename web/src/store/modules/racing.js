import api from '../../api/racing';


const getters = {

};

const actions = {
  getRacingSchedule({ commit }) {
    commit('setLoadingStatus', null);
    api.readSchedule()
      .then((result) => {
        commit('updateRaces', { data: result.data });
        commit('setLoadingStatus', 'successful');
      })
      .catch(() => commit('setLoadingStatus', 'failed'));
  },
};

const mutations = {

  updateRaces(state, data) {
    for (let i = 0; i <= data.meetings.length; i += 1) {
      const meeting = data.meetings[i];
      state.races[meeting.meeting_id] = meeting;
    }

    for (let i = 0; i <= data.races.length; i += 1) {
      const race = data.races[i];
      state.races[race.race_id] = race;
    }

    for (let i = 0; i <= data.selections.length; i += 1) {
      const selection = data.selections[i];
      state.selections[selection.selection_id] = selection;
    }
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
