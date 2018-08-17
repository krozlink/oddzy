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

function formatTimeRemaining(totalSeconds) {
  const remaining = totalSeconds;
  const isNegative = remaining < 0;

  const absSeconds = Math.abs(remaining);
  const days = Math.floor(absSeconds / 86400);
  const hours = Math.floor((absSeconds - (days * 86400)) / 3600);
  const minutes = Math.floor((absSeconds - (hours * 3600)) / 60);
  const seconds = absSeconds % 60;

  const prefix = isNegative ? '-' : '';

  if (days > 0) {
    return `${prefix}${days}d ${hours}h`;
  } else if (hours > 1) {
    return `${prefix}${hours}h ${minutes}m`;
  } else if (minutes > 5) {
    return `${prefix}${minutes}m`;
  } else if (absSeconds < 60 && absSeconds >= 0) {
    return `${prefix}${seconds}s`;
  }

  return `${prefix}${minutes}m ${seconds}s`;
}


export default {
  todayDate,
  tomorrowDate,
  overmorrowDate,
  startOfDay,
  formatDate,
  getDayString,
  getMonthString,
  formatTimeRemaining,
};
