package repo

import (
	"algtutor/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CurriculumRepo struct {
	dbPool *pgxpool.Pool
}

func NewCurriculumRepo(pool *pgxpool.Pool) *CurriculumRepo {
	return &CurriculumRepo{pool}
}

func (cr CurriculumRepo) CreateCourse(c domain.Course) error {
	ctx := context.Background()
	insertQuery := `
		INSERT INTO course(title)
		VALUES ($1)
	`
	_, err := cr.dbPool.Exec(
		ctx,
		insertQuery,
		c.Title,
	)
	if err != nil {
		// TODO: custom error handling and stuff
		return err
	}
	return nil
}

func (cr CurriculumRepo) CreateUnit(u domain.Unit) error {
	ctx := context.Background()
	insertQuery := `
		INSERT INTO unit(course_id, title, order_index)
		VALUES ($1, $2, $3)
	`
	_, err := cr.dbPool.Exec(
		ctx,
		insertQuery,
		u.CourseId,
		u.Title,
		u.OrderIndex,
	)
	if err != nil {
		// TODO: custom error handling and stuff
		return err
	}
	return nil
}
