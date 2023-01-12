package ticket

import (
	"math/rand"
	"time"
)

type Ticket struct {
	entries [3][9]int
}

func Generate() [3][9]int {
	rand.Seed(time.Now().UnixNano())
	t := &Ticket{entries: [3][9]int{}}
	numbers := make(map[int][]int)
	for i := 0; i < 9; i++ {
		var group []int
		for j := 0; j < 10; j++ {
			group = append(group, i*10+j+1)
		}
		numbers[i] = group
	}

	// take one row values because it will be same for all in the start
	rows := getRowValues(0, t)
	for colIndex := range rows {
		// taking the length of numbers[colIndex] because we are deleting the values once inserted
		// to avoid collision we take only remaining index
		randomIndex := rand.Intn(len(numbers[colIndex]))
		numberToInsert := numbers[colIndex][randomIndex]

		// get a random row which is incomplete
		possibleRow := []int{0, 1, 2}
		var incompleteRows []int

		for _, v := range possibleRow {
			if !IsRowCompleted(v, t) {
				incompleteRows = append(incompleteRows, v)
			}
		}

		// get random row which is not completed from 0,1,2
		randomIncompleteRowIndex := rand.Intn(len(possibleRow))
		randomRow := possibleRow[randomIncompleteRowIndex]

		// if the location is 0 and column is not completed
		if !IsColumnCompleted(colIndex, t) && t.entries[randomRow][colIndex] == 0 {

			// insert the value at the chosen location
			insertValue(t, randomRow, colIndex, numberToInsert)

			// remove the inserted numbers from the grouped number map
			groupToUpdate := numbers[colIndex]
			groupToUpdate = append(groupToUpdate[:randomIndex], groupToUpdate[randomIndex+1:]...)
			numbers[colIndex] = groupToUpdate
		}
	}
	fillRecursively(t, numbers)
	return t.entries
}

func SortTicket(ticket *Ticket) [3][9]int {
	t := &ticket.entries

	for colIndex := 0; colIndex < 9; colIndex++ {
		// check all 3 possible row index is filled
		if t[0][colIndex] != 0 && t[1][colIndex] != 0 && t[2][colIndex] != 0 {
			for i := 0; i < 2; i++ {
				for j := i + 1; j < 3; j++ {
					if t[i][colIndex] > t[j][colIndex] {
						temp := t[i][colIndex]
						t[i][colIndex] = t[j][colIndex]
						t[j][colIndex] = temp
					}
				}
			}
		} else if t[0][colIndex] != 0 && t[1][colIndex] != 0 && t[2][colIndex] == 0 {
			// only first & second row is populated
			if t[0][colIndex] > t[1][colIndex] {
				temp := t[0][colIndex]
				t[0][colIndex] = t[1][colIndex]
				t[1][colIndex] = temp
			}
		} else if t[0][colIndex] != 0 && t[1][colIndex] == 0 && t[2][colIndex] != 0 {
			// only first and third is populated
			if t[0][colIndex] > t[2][colIndex] {
				temp := t[0][colIndex]
				t[0][colIndex] = t[2][colIndex]
				t[2][colIndex] = temp
			}
		} else if t[0][colIndex] == 0 && t[1][colIndex] != 0 && t[2][colIndex] != 0 {
			// only second and third is populated
			if t[1][colIndex] > t[2][colIndex] {
				temp := t[1][colIndex]
				t[1][colIndex] = t[2][colIndex]
				t[2][colIndex] = temp
			}
		}
	}
	return *t
}

func fillRecursively(ticket *Ticket, numbers map[int][]int) {
	rand.Seed(time.Now().UnixNano())
	cValues := getColumnValues(0, ticket)
	rValues := getRowValues(0, ticket)

	for rIndex := range cValues {
		for cIndex := range rValues {
			// get a random index
			randomIndex := rand.Intn(len(numbers[cIndex]))
			numberToInsert := numbers[cIndex][randomIndex]

			//Random chance for generating variants of the ticket
			isSet := rand.Float32() > 0.5

			if !isSet &&
				!IsTicketCompleted(ticket) &&
				!IsRowCompleted(rIndex, ticket) &&
				!IsColumnCompleted(cIndex, ticket) &&
				ticket.entries[rIndex][cIndex] == 0 {
				insertValue(ticket, rIndex, cIndex, numberToInsert)

				// remove the inserted numbers from the grouped number map
				groupToUpdate := numbers[cIndex]
				groupToUpdate = append(groupToUpdate[:randomIndex], groupToUpdate[randomIndex+1:]...)
				numbers[cIndex] = groupToUpdate
			}
		}
	}
	if !IsTicketCompleted(ticket) {
		fillRecursively(ticket, numbers)
	} else {
		SortTicket(ticket)
	}
}

func getNumberOfEntries(ticket *Ticket) int {
	count := 0
	for _, r := range ticket.entries {
		for _, c := range r {
			if c != 0 {
				count++
			}
		}
	}
	return count
}

func IsTicketCompleted(ticket *Ticket) bool {
	return getNumberOfEntries(ticket) == 15
}

func IsRowCompleted(rowIndex int, ticket *Ticket) bool {
	rowValues := getRowValues(rowIndex, ticket)
	count := 0
	for _, v := range rowValues {
		if v != 0 {
			count++
		}
	}
	return count == 5
}

func IsColumnCompleted(columnIndex int, ticket *Ticket) bool {
	colValues := getColumnValues(columnIndex, ticket)
	count := 0
	for _, v := range colValues {
		if v != 0 {
			count++
		}
	}
	return count == 3
}

func getRowValues(rowIndex int, ticket *Ticket) [9]int {
	return ticket.entries[rowIndex]
}

func getColumnValues(columnIndex int, ticket *Ticket) []int {
	var columns []int
	for _, r := range ticket.entries {
		columns = append(columns, r[columnIndex])
	}
	return columns
}

func insertValue(ticket *Ticket, rowIndex int, colIndex int, value int) {
	ticket.entries[rowIndex][colIndex] = value
}
