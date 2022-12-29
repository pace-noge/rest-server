package taskstore

import (
	"fmt"
	"sync"
	"time"

	"github.com/pace-noge/rest-server/internal/domain"
)

type TaskStore struct {
	sync.Mutex

	tasks  map[int]domain.Task
	nextId int
}

func New() *TaskStore {
	ts := &TaskStore{}
	ts.tasks = make(map[int]domain.Task)
	ts.nextId = 0
	return ts
}

func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int {
	ts.Lock()
	defer ts.Unlock()

	task := domain.Task{
		Id:   ts.nextId,
		Text: text,
		Due:  due,
	}
	task.Tags = make([]string, len(tags))
	copy(task.Tags, tags)

	ts.tasks[ts.nextId] = task
	ts.nextId++
	return task.Id
}

func (ts *TaskStore) GetTask(id int) (domain.Task, error) {
	ts.Lock()
	defer ts.Unlock()

	t, ok := ts.tasks[id]
	if ok {
		return t, nil
	} else {
		return domain.Task{}, fmt.Errorf("task with id=%d not found", id)
	}
}

func (ts *TaskStore) DeleteTask(id int) error {
	ts.Lock()
	defer ts.Unlock()

	if _, ok := ts.tasks[id]; !ok {
		return fmt.Errorf("task with id=%d not found", id)
	}

	delete(ts.tasks, id)
	return nil
}

func (ts *TaskStore) DeleteAllTasks() error {
	ts.Lock()
	defer ts.Unlock()

	ts.tasks = make(map[int]domain.Task)
	return nil
}

func (ts *TaskStore) GetAllTasks() []domain.Task {
	ts.Lock()
	defer ts.Unlock()

	allTasks := make([]domain.Task, 0, len(ts.tasks))
	for _, task := range ts.tasks {
		allTasks = append(allTasks, task)
	}
	return allTasks
}

func (ts *TaskStore) GetTaskByTag(tag string) []domain.Task {
	ts.Lock()
	defer ts.Unlock()

	var tasks []domain.Task
taskloop:
	for _, task := range ts.tasks {
		for _, taskTag := range task.Tags {
			if taskTag == tag {
				tasks = append(tasks, task)
				continue taskloop
			}
		}
	}
	return tasks
}

func (ts *TaskStore) GetTaskByDueDate(year int, month time.Month, day int) []domain.Task {
	ts.Lock()
	defer ts.Unlock()

	var tasks []domain.Task
	for _, task := range ts.tasks {
		y, m, d := task.Due.Date()
		if y == year && m == month && d == day {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
