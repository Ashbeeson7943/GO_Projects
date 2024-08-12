package task

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID               int
	TASK_TITLE       string
	TASK_DETAIL      string
	CREATED_TIME     time.Time
	IS_COMPLETED     bool
	COMPLETED_REASON string
	COMPLETED_TIME   time.Time
}

var taskList = []Task{}

func SaveTask(t Task) {
	file, err := os.OpenFile("tasks.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write(t.convertToCSVRow())
}

func GetTask(id []string) Task {
	//TODO: this
	return Task{}
}

func GetTasks() []Task {
	LoadTasks()
	return taskList
}

func CompleteTask(id []string) {
	//TODO: this
}

func (t *Task) convertToCSVRow() []string {
	id := strconv.Itoa(t.ID)
	ct := t.CREATED_TIME.String()
	comt := t.COMPLETED_TIME.String()
	ic := strconv.FormatBool(t.IS_COMPLETED)
	csvRow := []string{id, t.TASK_TITLE, t.TASK_DETAIL, ct, ic, t.COMPLETED_REASON, comt}
	return csvRow
}

func updateTask() {
	// Write the CSV data
	file, err := os.Create("tasks.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	// this defines the header value and data values for the new csv file
	headers := []string{"id", "task_title", "task_details", "created_time", "completed_reason", "completed_time"}
	data := taskList
	writer.Write(headers)
	for _, row := range data {
		writer.Write(row.convertToCSVRow())
	}
}

func LoadTasks() {
	taskList = nil
	file, err := os.Open("tasks.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	headers := true
	// Print the CSV data
	for _, row := range data {
		ts := []string{}
		if !headers {
			for _, col := range row {
				ts = append(ts, fmt.Sprintf("%s,", col))
			}
			id, _ := strconv.Atoi(ts[0])
			ic, _ := strconv.ParseBool(ts[4])
			ct, _ := time.Parse(time.ANSIC, ts[3])
			comt, _ := time.Parse(time.ANSIC, ts[6])
			t := Task{
				ID:               id,
				TASK_TITLE:       ts[1],
				TASK_DETAIL:      ts[2],
				CREATED_TIME:     ct,
				IS_COMPLETED:     ic,
				COMPLETED_REASON: ts[5],
				COMPLETED_TIME:   comt,
			}
			taskList = append(taskList, t)
		}
		headers = false
	}
}
