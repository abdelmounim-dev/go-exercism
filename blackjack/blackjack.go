package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
  // I should be using a switch, but that is too easy lol, I'm in love with the map
  // it returns 0 if the key doesn't exist which is perfect for this use case
	return map[string]int{
    "ace":  11,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
    "ten": 10,
    "jack": 10,
    "queen": 10,
    "king": 10,
  }[card]

}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
  var turn string;
  sum :=  ParseCard(card1)+ ParseCard(card2)
  hasSevenOrHigher := func () bool {
    return ParseCard(dealerCard) >= 7
  }

  hasCounterBlackjack := func () bool {
    return ParseCard(dealerCard) >= 10
  }
  switch {
  case sum == 22:
    turn = "P"
    break;
  case sum == 21:
    if hasCounterBlackjack() {
      turn = "S"
      break;
    }
    turn = "W"
    break;
  case sum >= 17:
    turn = "S"
    break
  case sum >= 12 :
    if hasSevenOrHigher() {
      turn = "H"
      break;
    }
    turn = "S"
    break;
  default:
    turn = "H"
  }


	return turn
}
