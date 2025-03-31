package task

import (
	"encoding/json"
	"os"
	"path"
	"time"

	"backeseduardo.com/task-tracker/internal"
)

type Status string

const (
	STATUS_DONE        Status = "done"
	STATUS_TODO        Status = "todo"
	STATUS_IN_PROGRESS Status = "in-progress"
)

var StatusMap = map[string]Status{
	"done":        STATUS_DONE,
	"todo":        STATUS_TODO,
	"in-progress": STATUS_IN_PROGRESS,
}

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Updated     bool      `json:"updated"`
}

type DB struct {
	Seq   int    `json:"seq"`
	Tasks []Task `json:"tasks"`
}

func (d *DB) NextSeq() int {
	d.Seq += 1
	return d.Seq
}

func (d *DB) Save() {
	b, err := json.Marshal(d)
	internal.Check(err)
	os.WriteFile(filePath(), b, 0664)
}

func LoadDB() *DB {
	var d DB
	b, err := os.ReadFile(filePath())
	if err == nil && len(b) > 0 {
		err := json.Unmarshal(b, &d)
		internal.Check(err)
	}
	return &d
}

func filePath() string {
	cwd, err := os.Getwd()
	internal.Check(err)
	return path.Join(cwd, "db.json")
}
