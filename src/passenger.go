package src

type passenger struct {
	card card
	fee  fee
}

func newPassenger(card card, fee fee) passenger {
	return passenger{card, fee}
}
