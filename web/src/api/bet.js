class Bet {
  constructor(data) {
    this.runner_number = parseInt(data.runner_number, 10);
    this.runner_name = data.runner_name;
    this.meeting_name = data.meeting_name;
    this.meeting_number = parseInt(data.meeting_number, 10);
    this.price = parseFloat(data.price);
    this.bet_type = data.bet_type;
    this.win_type = data.win_type;
    this.amount = parseFloat(data.amount);
    this.selection_id = data.selection_id;

    this.bet_id = `${this.selection_id}|${this.bet_type}|${this.win_type}`;

    this.message = '';
  }

  validate() {
    if (this.amount < 1) {
      return 'Bet size must be at least $1';
    }
    return '';
  }
}

export default Bet;
