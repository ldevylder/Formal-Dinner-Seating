package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

type person struct {
	firstname string
	lastname  string
	tables    []int
}

//function to randomly shuffle the student list
func shuffle(vals []person) []person {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]person, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}

// table numbers
var table = rand.Intn(31)

func main() {
//initial number of people at each table
	tableFill1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	tableFill2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	tableFill3 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	csvFile, _ := os.Open("peopleAndTables.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var students = []person{}

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		//add all students on the list to a slice
		students = append(students, person{
			firstname: line[1],
			lastname:  line[0],
		})
	}

	// first seating
	//shuffle students
	students = shuffle(students)
	table := 1

	for studentIndex := range students {
		if studentIndex < 8 {
			// assign first eight students to kitchen crew
			students[studentIndex].tables = append(students[studentIndex].tables, 0)
			fmt.Println("Kitchen Crew: ", students[studentIndex])
			continue
		} else if studentIndex < 8+31 {
			// assign next thrity-one students to groups of waiters
			students[studentIndex].tables = append(students[studentIndex].tables, studentIndex-7)
			fmt.Println("Waiter:", students[studentIndex])
			continue
		} else if studentIndex > 75 {
			//students after 8 kitchen crew members and 31 waiters
			// assign nine people to first four tables
			students[studentIndex].tables = append(students[studentIndex].tables, table)
			tableFill1[table]++
			if tableFill1[table] > 9 {
				table++
			}
			continue
		} else {
			// set rest of students on the list at each table, eight per table
			students[studentIndex].tables = append(students[studentIndex].tables, table)
			tableFill1[table]++
			if tableFill1[table] > 8 {
				table++
			}
		}
	}

	fmt.Println(students)

	// second seating
	students = shuffle(students)
	table = 1

	for studentIndex := range students {
		if studentIndex < 8 {
			// assign first eight students to kitchen crew
			students[studentIndex].tables = append(students[studentIndex].tables, 0)
			fmt.Println("Kitchen Crew: ", students[studentIndex])
			continue
		} else if studentIndex < 8+31 {
			// assign next thrity-one students to groups of waiters
			students[studentIndex].tables = append(students[studentIndex].tables, studentIndex-7)
			fmt.Println("Waiter:", students[studentIndex])
			continue
		} else if studentIndex > 75 {
			// assign nine people to first four tables
			students[studentIndex].tables = append(students[studentIndex].tables, table)
			tableFill2[table]++
			if tableFill2[table] > 9 {
				table++
			}
			continue
		} else {
			// set rest of students on the list at each table, eight per table
			students[studentIndex].tables = append(students[studentIndex].tables, table)
			tableFill2[table]++
			if tableFill2[table] > 8 {
				table++
			}
		}
	}
	fmt.Println(students)

	// third seating
	students = shuffle(students)
	table = 1

	for studentIndex := range students {
		if studentIndex < 8 {
			// assign first eight students to kitchen crew
			students[studentIndex].tables = append(students[studentIndex].tables, 0)
			fmt.Println("Kitchen Crew: ", students[studentIndex])
			continue
		} else if studentIndex < 8+31 {
			// assign next thrity-one students to groups of waiters
			students[studentIndex].tables = append(students[studentIndex].tables, studentIndex-7)
			fmt.Println("Waiter:", students[studentIndex])
			continue
		} else if studentIndex > 75 {
			// assign nine people to first four tables
			students[studentIndex].tables = append(students[studentIndex].tables, table)
			tableFill3[table]++
			if tableFill3[table] > 9 {
				table++
			}
			continue
		} else {
			// set rest of students on the list at each table, eight per table
			students[studentIndex].tables = append(students[studentIndex].tables, table)
			tableFill3[table]++
			if tableFill3[table] > 8 {
				table++
			}
		}
	}
	fmt.Println(students)
}
