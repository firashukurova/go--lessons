package main

//Описание: Реализуйте структуру Task с полями title и completed. Реализуйте структуру TaskList с срезом задач и методы для добавления задачи и
//получения количества завершённых задач.
//Методы:
//AddTask(title string)
//CompletedTasks() int

type Task struct {
	Title     string
	Completed bool
}

type TasksList struct {
	Tasks []Task
}

func (r *TasksList) AddTask(title string) {
	r.Tasks = append(r.Tasks, Task{Title: title, Completed: false})
}

func (r *TasksList) CompletedTasks() int {
	total := 0
	for _, task := range r.Tasks {
		if task.Completed {
			total++

		}
	}
	return total
}

func main() {
	
}
