package main

import "fmt"

type Task struct {}
func (t *Task) Run() {
  fmt.Println("run")
}

type Runner struct {
  name string
}

func (r *Runner) Name() string {
  return r.name
}

func (r *Runner) Run(t Task)  {
  t.Run()
}

func (r *Runner) RunAll(ts []Task) {
  for _, t := range ts {
    r.Run(t)
  }
}

type RunCounter struct {
  Runner
  count int
}

func NewRunCounter(name string) *RunCounter {
  return &RunCounter{Runner{name}, 0}
}

func (r *RunCounter) Run(t Task)  {
  r.count++
}

func (r *RunCounter) RunAll(ts []Task) {
  r.count += len(ts)
}

func (r *RunCounter) Count() int {
  return r.count
}

func main()  {
  runCounter := NewRunCounter("hello")
  task := Task{}
  tasks := []Task {task, task, task, task}
  runCounter.RunAll(tasks)
  fmt.Println("counter", runCounter)
}
