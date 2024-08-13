package task

func GetNextTaskID() int {
	LoadTasks()
	var newId int
	if len(taskList) >= 1 {
		newId = taskList[len(taskList)-1].ID + 1
	} else {
		newId = 1
	}
	return newId
}
