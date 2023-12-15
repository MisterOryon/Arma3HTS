package tasks

// TaskQueue works like stack to manage sending and deleting files on the remote server.
type TaskQueue struct {
	tasks      []*Task
	needUpload int64
}

// AddTask add task in the queue.
func (tm *TaskQueue) AddTask(newTask *Task) {
	tm.tasks = append(tm.tasks, newTask)
	tm.needUpload += newTask.NeedUpload
}

// Next returns the next task in the queue otherwise we return nil if the queue is empty.
// Works like stack it is the last element entered which will be the first to be returned by this function.
func (tm *TaskQueue) Next() *Task {
	if len(tm.tasks) == 0 {
		return nil
	}

	currantTask := tm.tasks[len(tm.tasks)-1]

	tm.needUpload -= currantTask.NeedUpload
	tm.tasks = tm.tasks[:len(tm.tasks)-1]

	return currantTask
}

// AsNext returns ture if there are any tasks left in the queue otherwise we returns false.
func (tm *TaskQueue) AsNext() bool {
	return len(tm.tasks) > 0
}

// Len returns the number of pending tasks.
func (tm *TaskQueue) Len() int {
	return len(tm.tasks)
}

// NeedUpload returns the number of bytes to upload.
func (tm *TaskQueue) NeedUpload() int64 {
	return tm.needUpload
}
