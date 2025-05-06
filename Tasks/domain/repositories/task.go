package repositories

import (
	"context"
	"fmt"
	"knovel/tasks/domain/common"
	"knovel/tasks/domain/entities"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, pagesize, offset int) ([]*entities.Task, error)
	GetTasksByUserId(ctx context.Context, userId string) ([]*entities.Task, error)
	CreateTask(ctx context.Context, name, userId, description, status string) (int64, error)
	UpdateTaskStatus(ctx context.Context, taskId int, status string) (int64, error)
	AssignTask(tx context.Context, taskId int, userId string) error
}

type TaskRepositoryInstance struct {
	dbContext common.DbContext
	tableName string
}

var _ TaskRepository = (*TaskRepositoryInstance)(nil)

func NewTaskRepository(dbContext common.DbContext) TaskRepository {
	return &TaskRepositoryInstance{
		dbContext: dbContext,
		tableName: "task",
	}
}

func (repo *TaskRepositoryInstance) GetTasksByUserId(ctx context.Context, userId string) ([]*entities.Task, error) {
	const query = "SELECT id, name, description, userId, status FROM %s WHERE userid = $1"
	rows, err := repo.dbContext.QueryContext(ctx, repo.table(query), userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*entities.Task, 0)
	for rows.Next() {
		var task = new(entities.Task)
		if err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.UserId, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// AssignTask implements TaskRepository.
func (repo *TaskRepositoryInstance) AssignTask(ctx context.Context, taskId int, userId string) error {
	const query = "UPDATE  %s SET userid = $1 WHERE id = $2"
	_, err := repo.dbContext.ExecContext(ctx, repo.table(query), userId, taskId)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TaskRepositoryInstance) GetTasks(ctx context.Context, pagesize, offset int) ([]*entities.Task, error) {
	if offset < 0 {
		offset = 0
	}
	const query = "SELECT id, name, description, userId, status FROM %s ORDER BY id LIMIT $1 OFFSET $2"

	rows, err := repo.dbContext.QueryContext(context.Background(), repo.table(query), pagesize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks = make([]*entities.Task, 0)
	for rows.Next() {
		task := &entities.Task{}
		if err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.UserId, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (repo *TaskRepositoryInstance) CreateTask(ctx context.Context, name, userId, description, status string) (int64, error) {
	const query = "INSERT INTO %s(name, userId, description, status) VALUES($1, $2, $3, $4)"

	results, err := repo.dbContext.ExecContext(ctx, repo.table(query), name, userId, description, status)
	if err != nil {
		return 0, err
	}
	return results.RowsAffected()
}

func (repo *TaskRepositoryInstance) UpdateTaskStatus(ctx context.Context, taskid int, status string) (int64, error) {

	const query = "UPDATE  %s SET status = $1 WHERE id = $2"
	results, err := repo.dbContext.ExecContext(ctx, repo.table(query), status, taskid)
	if err != nil {
		return 0, err
	}
	return results.RowsAffected()
}

func (repo *TaskRepositoryInstance) table(query string) string {
	return fmt.Sprintf(query, repo.tableName)
}
