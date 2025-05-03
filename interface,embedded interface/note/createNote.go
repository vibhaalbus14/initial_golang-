package note

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
type Note struct {
	Title     string    `json:"Title"`
	Content   string    `json:"Content"`
	CreatedAt time.Time `json:"Created At"` //from time package, Time object type
}

// constructor creation
func NewNote(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("all the values need to be entered")
	}
	return Note{Title: title, Content: content, CreatedAt: time.Now()}, nil
}

func Input() (Note, error) {
	var title string
	var content string
	fmt.Print("Enter Title : ")
	reader1 := bufio.NewReader(os.Stdin)
	//read this string till /n is encounterd
	title, err1 := reader1.ReadString('\n')
	if err1 != nil {
		return Note{}, errors.New("reading title isnt successful")
	}
	title = strings.TrimSuffix(title, "\n")
	title = strings.TrimSuffix(title, "\r")

	fmt.Print("Enter content : ")
	reader2 := bufio.NewReader(os.Stdin)
	//read this string till /n is encounterd
	content, err2 := reader2.ReadString('\n')
	if err2 != nil {
		return Note{}, errors.New("reading content isnt successful")
	}
	content = strings.TrimSuffix(content, "\n")
	content = strings.TrimSuffix(content, "\r")
	return NewNote(title, content)

}

// function to display
func (n Note) Display() {
	fmt.Print("Title is : ", n.Title)
	fmt.Println("Content is : ", n.Content)
	fmt.Println("Creared At : ", n.CreatedAt)
}

// function to save to file
func (n Note) Save() error {
	//write to a file with file name=title
	//for this convert the title appropriately

	newTitle := strings.ReplaceAll(n.Title, " ", "_") //it creates a  new copy
	finalTitle := strings.ToLower(newTitle) + ".json" //.json extension

	//data to be written in file is made json
	//we need to pass a datatype to marshal function, and only the exportable variable are visible to it
	finalContent, err1 := json.Marshal(n)
	if err1 != nil {
		return errors.New("error while converting to json")
	}
	//create a file with this name
	err := os.WriteFile(finalTitle, finalContent, 0644) //0644=> only owners can read and write and others can view
	return err

}
