package domain

type Todo struct {
	Title string
	Done  bool
}

type TodoPageList struct {
	PageTitle string
	Todos     []Todo
}
