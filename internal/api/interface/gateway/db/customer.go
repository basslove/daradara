package db

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/domain/repository"
	"github.com/jmoiron/sqlx"
	"time"
)

func NewCustomerRepository(client *sqlx.DB) repository.CustomerRepository {
	return &CustomerRepository{Repository: Repository{client: client}}
}

type CustomerRepository struct {
	Repository
}

func (r *CustomerRepository) FindOneByName(ctx context.Context, name string) (*model.Customer, error) {
	values := make([]*model.Customer, 0)

	m := model.Customer{}
	qb := sq.Select("*").From(m.TableName()).PlaceholderFormat(sq.Dollar)
	qb = qb.Where("name = ?", name)

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("CustomerRepository FindOneByName: %w", err)
	}
	stmt, err := r.client.Preparex(sql)
	if err != nil {
		return nil, fmt.Errorf("CustomerRepository FindOneByName: %w", err)
	}
	if err = stmt.SelectContext(ctx, &values, args...); err != nil {
		return nil, fmt.Errorf("CustomerRepository FindByName: %w", err)
	}

	var value *model.Customer
	if len(values) > 0 {
		value = values[0]
	}

	return value, nil
}

func (r *CustomerRepository) FindOneByEmail(ctx context.Context, email string) (*model.Customer, error) {
	values := make([]*model.Customer, 0)

	m := model.Customer{}
	qb := sq.Select("*").From(m.TableName()).PlaceholderFormat(sq.Dollar)
	qb = qb.Where("email = ?", email)

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("CustomerRepository FindOneByEmail: %w", err)
	}
	stmt, err := r.client.Preparex(sql)
	if err != nil {
		return nil, fmt.Errorf("CustomerRepository FindOneByEmail: %w", err)
	}
	if err = stmt.SelectContext(ctx, &values, args...); err != nil {
		return nil, fmt.Errorf("CustomerRepository FindOneByEmail: %w", err)
	}

	var value *model.Customer
	if len(values) > 0 {
		value = values[0]
	}

	return value, nil
}

func (r *CustomerRepository) Create(ctx context.Context, c *model.Customer) (int64, error) {
	var newID int64

	sql := fmt.Sprintf("INSERT INTO %s (email, crypted_password, name, gender, generation, display_name, birthday, phone_number, introduction, image_url, allow_plan_displayed, is_valid) VALUES (:email, :crypted_password, :name, :gender, :generation, :display_name, :birthday, :phone_number, :introduction, :image_url, :allow_plan_displayed, :is_valid) RETURNING id;", c.TableName())
	rows, err := r.client.NamedQueryContext(ctx, sql, c)
	if err != nil {
		return 0, fmt.Errorf("CustomerRepository Create: %w", err)
	}
	if rows.Next() {
		rows.Scan(&newID)
	}

	return newID, nil
}

func (r *CustomerRepository) Update(ctx context.Context, m *model.Customer) error {
	m.UpdatedAt = time.Now()

	sql := fmt.Sprintf("UPDATE %s SET email=:email, crypted_password=:crypted_password, name=:name, gender=:gender, generation=:generation, display_name=:display_name, birthday=:birthday, phone_number=:phone_number, introduction=:introduction, image_url=:image_url, allow_plan_displayed=:allow_plan_displayed, is_valid=:is_valid, last_accessed_at=:last_accessed_at, last_logged_in_at=:last_logged_in_at, updated_at=:updated_at WHERE id = :id", m.TableName())
	_, err := r.client.NamedQueryContext(ctx, sql, m)
	if err != nil {
		return fmt.Errorf("CustomerRepository Update: %w", err)
	}

	return nil
}
