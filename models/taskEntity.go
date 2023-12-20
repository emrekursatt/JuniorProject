package models

type TaskEntity struct {
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func GetCode(t *TaskEntity) string {
	return t.Code
}

func (t TaskEntity) SetCode(title string) {
	t.Title = title
}

func GetTitle(task *TaskEntity) string {
	return task.Title
}

func (t *TaskEntity) SetTitle(title string) {
	t.Title = title
}

func GetDescription(t *TaskEntity) string {
	return t.Description
}

func (t *TaskEntity) SetDescription(description string) {
	t.Description = description
}

func GetStatus(t *TaskEntity) string {
	var statusString string

	if t.Status == 0 {
		statusString = "Passive"
	}
	statusString = "Active"
	return statusString
}
func (t *TaskEntity) SetStatus(status int) {
	t.Status = status
}
