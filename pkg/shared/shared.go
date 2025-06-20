package shared

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2/widget"
)

var Window GuiWindow

type JobData struct {
	Location       string
	JobTitle       string
	CompanyName    string
	Description    string
	DatePosted     string
	Salary         int
	WorkFromHome   string
	Qualifications string
	Links          string
	Country        string
	// LatLong        LatLong
}

type GuiWindow struct {
	ListWidget           *widget.List
	KeywordEntryWidget   *widget.Entry
	LocationEntryWidget  *widget.Entry
	MinSalaryEntryWidget *widget.Entry
	DetailsWidget        *widget.Label
	JobDataGui           *[]JobData
	SelectedJobDetails   string
	Filters              FilterEntries
	// Server               *http.Server
}

type FilterEntries struct {
	KeywordEntry      string
	LocationEntry     string
	MinSalaryEntry    string
	WorkFromHomeEntry bool
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckErrorWarn(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
