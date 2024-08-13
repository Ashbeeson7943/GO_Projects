package task

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
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

func saveTasks(tl []Task) {
	file, err := os.OpenFile("tasks.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"id,task_title,task_details,created_time,is_completed,completed_reason,completed_time"})
	for _, task := range tl {
		writer.Write(task.convertToCSVRow())
	}

}

func getTask(id string) []Task {
	var tasks []Task
	for _, task := range taskList {
		tId, _ := strconv.Atoi(id)
		if tId == task.ID {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func ViewTask(id []string) {
	LoadTasks()
	var ID = strings.Join(id[:], " ")
	displayTasks(getTask(ID), false)
}

func displayTasks(tasks []Task, allTasks bool) {
	var formattedTasks []string
	for _, task := range tasks {
		var splitTask []string
		if !task.IS_COMPLETED {
			splitTask = task.convertToCSVRow()
			var out = ""
			for _, i := range splitTask {
				out += strings.Trim(i, ",") + "\t\t\t\t\t\t\t\t\t\t\t\t"
			}
			formattedTasks = append(formattedTasks, out)
		} else if allTasks {
			splitTask = task.convertToCSVRow()
			var out = ""
			for _, i := range splitTask {
				out += strings.Trim(i, ",") + "\t\t\t\t\t\t\t\t\t\t\t\t"
			}
			formattedTasks = append(formattedTasks, out)
		}
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, task := range formattedTasks {
		fmt.Fprintln(w, task)
	}
	w.Flush()
}

func ViewTasks(at bool) {
	LoadTasks()
	displayTasks(taskList, at)
}

func CompleteTask(args []string) {
	LoadTasks()
	os.Remove("tasks.csv")
	//var ID = strings.Join(args[:], " ")
	ID := args[0]
	var reason string
	if len(args) > 1 {
		reason = args[1]
	} else {
		reason = "Completed(DEFAULT)"
	}
	var newTL []Task
	for _, task := range taskList {
		tId, _ := strconv.Atoi(ID)
		if tId == task.ID {
			task.IS_COMPLETED = true
			task.COMPLETED_REASON = reason
			task.COMPLETED_TIME = time.Now()
		}
		newTL = append(newTL, task)
	}
	saveTasks(newTL)
}

func (t *Task) convertToCSVRow() []string {
	var csvRow []string
	id := strconv.Itoa(t.ID)
	ct := t.CREATED_TIME.String()
	if t.IS_COMPLETED {
		comt := t.COMPLETED_TIME.String()
		ic := strconv.FormatBool(t.IS_COMPLETED)
		csvRow = []string{id, t.TASK_TITLE, t.TASK_DETAIL, ct, ic, t.COMPLETED_REASON, comt}
	} else {
		csvRow = []string{id, t.TASK_TITLE, t.TASK_DETAIL, ct}
	}
	return csvRow
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
	for _, row := range data {
		ts := []string{}
		if !headers {
			for _, col := range row {
				ts = append(ts, fmt.Sprintf("%s", col))
			}
			id, wtf := strconv.Atoi(ts[0])
			if wtf != nil {
				log.Fatalf("error: %v", wtf)
			}
			ct, _ := time.Parse(time.ANSIC, ts[3])
			if len(ts) > 4 {
				ic, _ := strconv.ParseBool(ts[4])
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
			} else {
				t := Task{
					ID:           id,
					TASK_TITLE:   ts[1],
					TASK_DETAIL:  ts[2],
					CREATED_TIME: ct,
				}
				taskList = append(taskList, t)
			}
		}
		headers = false
	}
}
