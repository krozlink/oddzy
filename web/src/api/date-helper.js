function startOfDay(date) {
  return new Date(date.getFullYear(), date.getMonth(), date.getDate());
}

function formatDate(date) {
  return date.toISOString().split('T')[0];
}

function todayDate() {
  return startOfDay(new Date());
}


export default {
  dateString: () => formatDate(new Date()),
  todayDate,
  startOfDay,
  formatDate,
};
