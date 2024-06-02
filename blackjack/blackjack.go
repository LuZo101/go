package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king", "queen", "jack", "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	pSum := ParseCard(card1) + ParseCard(card2)
	dSum := ParseCard(dealerCard)
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	if pSum == 21 {
		if dSum >= 10 {
			return "S"
		} else {
			return "W"
		}
	}

	if pSum >= 17 && pSum <= 20 {
		return "S"
	}
	if pSum >= 12 && pSum <= 16 {
		if dSum >= 7 {
			return "H"
		} else {
			return "S"
		}
	} else {
		return "H"
	}
}
