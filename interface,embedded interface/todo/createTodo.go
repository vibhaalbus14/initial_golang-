package todo

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// defining structure
type Todo struct {
	Name      string    `json:"Created By"`
	Reminder  string    `json:"Reminder"`
	CreatedAt time.Time `json:" Todo Created At"` //from time package, Time object type
}

// constructor creation
func NewTodo(name, Reminder string) (Todo, error) {
	if name == "" || Reminder == "" {
		return Todo{}, errors.New("all the values need to be entered")
	}
	return Todo{Name: name, Reminder: Reminder, CreatedAt: time.Now()}, nil
}

func Input() (Todo, error) {
	var name string
	var Reminder string
	fmt.Print("Enter Name : ")
	reader1 := bufio.NewReader(os.Stdin)
	//read this string till /n is encounterd
	name, err1 := reader1.ReadString('\n')
	if err1 != nil {
		return Todo{}, errors.New("reading name isnt successful")
	}
	name = strings.TrimSuffix(name, "\n")
	name = strings.TrimSuffix(name, "\r")

	fmt.Print("Enter Reminder : ")
	reader2 := bufio.NewReader(os.Stdin)
	//read this string till /n is encounterd
	Reminder, err2 := reader2.ReadString('\n')
	if err2 != nil {
		return Todo{}, errors.New("reading Reminder isnt successful")
	}
	Reminder = strings.TrimSuffix(Reminder, "\n")
	Reminder = strings.TrimSuffix(Reminder, "\r")
	return NewTodo(name, Reminder)

}

// function to display
func (n Todo) Display() {
	fmt.Print("Name is : ", n.Name)
	fmt.Println("Reminder is : ", n.Reminder)
	fmt.Println("Creared At : ", n.CreatedAt)
}

// function to save to file
func (n Todo) Save() error {
	//write to a file with file name=Name
	//for this convert the Name appropriately

	newName := strings.ReplaceAll(n.Name, " ", "_") //it creates a  new copy
	finalName := strings.ToLower(newName) + ".json" //.json extension

	//data to be written in file is made json
	//we need to pass a datatype to marshal function, and only the exportable variable are visible to it
	finalReminder, err1 := json.Marshal(n)
	if err1 != nil {
		return errors.New("error while converting to json")
	}
	//create a file with this name
	err := os.WriteFile(finalName, finalReminder, 0644) //0644=> only owners can read and write and others can view
	return err

}
