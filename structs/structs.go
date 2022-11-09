package structs

type TodoList struct {
	Id       int    `json:"id"`
	ListName string `json:"listName"`
	Todos    []Todo `json:"todos"`
}

type Todo struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
	DueDate  string `json:"dueDate"`
	Notes    string `json:"notes"`
}
