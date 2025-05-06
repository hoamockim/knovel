package config

type PresentationConfig struct {
	TaskClientKey string
}

var pc PresentationConfig

func SetConfig(TaskClientKey string) {
	pc.TaskClientKey = TaskClientKey
}

func GetTaskClientKey() string {
	return pc.TaskClientKey
}
