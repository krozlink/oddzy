class Bet {
  constructor(data) {
    this.bet_id = data.betId;
    this.runner_number = data.number;
    this.runner_name = data.name;
    this.meeting_name = data.meetingName;
    this.meeting_number = data.meetingNumber;
    this.price = data.price;
    this.bet_type = data.betType;
    this.amount = data.amount;
    this.selection_id = data.selectionId;
  }
}

export default Bet;
