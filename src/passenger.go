package src

type passenger struct {
	card card
	fee  fee
}

func NewPassenger(card card, fee fee) passenger {
	return passenger{card, fee}
}
