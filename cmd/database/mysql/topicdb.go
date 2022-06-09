package mysql

import (
	"back-end/cmd/database/model"
	"back-end/cmd/database/repo"
	"back-end/core/logger"
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TopicDB struct {
	table               string
	connect             *sqlx.DB
	IgnoreInsertColumns []string
}

func NewTopicDB() (repo.TopicRepo, error) {
	db, err := sqlx.Open("mysql", "root:524020@tcp(127.0.0.1:3306)/nckh")
	if err != nil {
		panic(err)
	}
	return &TopicDB{
		table:               "topic",
		connect:             db,
		IgnoreInsertColumns: []string{"id"},
	}, nil
}

func (topic *TopicDB) Close() {
	err := topic.connect.Close()
	if err != nil {
		return
	}
}

func (u *TopicDB) GetTopic(ctx context.Context, condition *repo.GetCondition) ([]*model.Topic, error) {
	ctxLogger := logger.NewContextLog(ctx)

	db := sq.Select("*").From(u.table)
	if condition != nil {
		db = db.Where(sq.Like{"name": fmt.Sprintf("%%%s%%", condition.Key)})
	}
	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query, error: %s", err.Error())
		return nil, err
	}
	var result []*model.Topic
	err = u.connect.Get(&result, query, arg...)
	if err != nil {
		ctxLogger.Errorf("Failed while select, error: %s", err.Error())
		return nil, err
	}
	if len(result) == 0 {
		return nil, sql.ErrNoRows
	}
	return result, nil
}

func (u *TopicDB) CreateTopic(ctx context.Context, create *model.Topic) error {
	ctxLogger := logger.NewContextLog(ctx)

	db := sq.Insert(u.table).
		Columns(GetListColumn(create, u.IgnoreInsertColumns, []string{})...).
		Values(GetListValues(create, u.IgnoreInsertColumns, []string{})...)
	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while select, error: %s", err.Error())
		return err
	}
	_, err = u.connect.Exec(query, arg...)
	if err != nil {
		ctxLogger.Errorf("Failed while insert topic, error: %s", err.Error())
		return err
	}
	return nil
}

func (u *TopicDB) UpdateTopic(ctx context.Context, update *model.Topic) error {
	ctxLogger := logger.NewContextLog(ctx)

	db := sq.Update(u.table).
		Set("name", update.TopicName).
		Set("group_student_id", update.GroupStudentID).
		Set("lecturer_id", update.LecturerID).
		Set("start_day", update.StartDay).
		Set("end_day", update.EndDay).
		Set("allowance", update.Allowance).
		Set("status", update.TopicStatus).
		Where(sq.Eq{"id": update.TopicID})
	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query, error", err.Error())
		return err
	}
	_, err = u.connect.Exec(query, arg...)
	if err != nil {
		ctxLogger.Errorf("Failed while update topic")
		return err
	}
	return nil
}
