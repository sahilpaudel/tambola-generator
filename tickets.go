package tambola_generator

import (
	"errors"
	"github.com/sahilpaudel/tambola-generator/ticket"
)

func GenerateTickets(numberOfTickets int) ([][3][9]int, error) {

	if numberOfTickets > 100 {
		return nil, errors.New("ticket limit breached only 100 tickets allowed")
	}

	tickets := make([][3][9]int, numberOfTickets)

	for i := 0; i < numberOfTickets; i++ {
		tickets = append([][3][9]int{ticket.Generate()}, tickets...)
	}

	return tickets[:numberOfTickets], nil
}
