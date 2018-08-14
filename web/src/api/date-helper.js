const DAYS = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
const MONTHS = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];

function startOfDay(date) {
  return new Date(date.getFullYear(), date.getMonth(), date.getDate());
}

function formatDate(date) {
  const d = new Date(date);
  let month = `${d.getMonth() + 1}`;
  let day = `${d.getDate()}`;
  const year = d.getFullYear();

  if (month.length < 2) month = `0${month}`;
  if (day.length < 2) day = `0${day}`;

  return [year, month, day].join('-');
}

function todayDate() {
  return startOfDay(new Date());
}

function tomorrowDate() {
  const tomorrow = new Date(new Date().getTime() + (24 * 60 * 60 * 1000));
  return startOfDay(tomorrow);
}

function overmorrowDate() {
  const tomorrow = new Date(new Date().getTime() + (48 * 60 * 60 * 1000));
  return startOfDay(tomorrow);
}

function getDayString(day) {
  return DAYS[day];
}

function getMonthString(month) {
  return MONTHS[month];
}

export default {
  todayDate,
  tomorrowDate,
  overmorrowDate,
  startOfDay,
  formatDate,
  getDayString,
  getMonthString,
};
