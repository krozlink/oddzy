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

function parseDateString(str) {
  // format D(D)/M(M)/(YY)YY
  const dateFormat = /^\d{1,4}[.|/|-]\d{1,2}[.|/|-]\d{1,4}$/;

  if (dateFormat.test(str)) {
    // remove any leading zeros from date values
    const s = str.replace(/0*(\d*)/gi, '$1');
    const dateArray = s.split(/[.|/|-]/);

    // correct month value
    dateArray[1] = parseInt(dateArray[1], 10) - 1;

    // convert day / year to integers
    dateArray[0] = parseInt(dateArray[0], 10);
    dateArray[2] = parseInt(dateArray[2], 10);

    // correct year value
    if (dateArray[2].length < 4) {
      // correct year value
      dateArray[2] = (parseInt(dateArray[2], 10) < 50) ? 2000 + parseInt(dateArray[2], 10) : 1900 + parseInt(dateArray[2], 10);
    }

    const testDate = new Date(dateArray[2], dateArray[1], dateArray[0]);
    if (testDate.getDate() !== dateArray[0] || testDate.getMonth() !== dateArray[1] || testDate.getFullYear() !== dateArray[2]) {
      return null;
    }
    return testDate;
  }
  return null;
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

function pad(n, width) {
  return n.length >= width ? n : new Array((width - n.length) + 1).join('0') + n;
}

function formatTime(unix) {
  const d = new Date(unix);
  return `${d.getHours()}:${pad(d.getMinutes().toString(), 2)}`;
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
  parseDateString,
  getDayString,
  getMonthString,
  formatTimeRemaining,
  formatTime,
};
