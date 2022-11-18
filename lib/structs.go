package lib

type TodoList struct {
	Id       int    `json:"id"`
	ListName string `json:"listName" binding:"required"`
	Todos    []Todo `json:"todos"`
}

type Todo struct {
	Id       int    `json:"id"`
	Title    string `json:"title" binding:"required"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
	DueDate  string `json:"dueDate"`
	Notes    string `json:"notes"`
}
