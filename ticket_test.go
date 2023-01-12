package tambola_generator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateTickets(t *testing.T) {
	ticket, err := GenerateTickets(2)
	fmt.Println(ticket)
	assert.Nil(t, err)
	assert.NotNil(t, ticket)
}

func TestGenerateTicketsWithError(t *testing.T) {
	ticket, err := GenerateTickets(101)
	assert.Nil(t, ticket)
	assert.NotNil(t, err)
}
