package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type CiudadanoPostgresRepository struct {
	db *pgxpool.Pool
}

func NewCiudadanoPostgresRepository(db *pgxpool.Pool) *CiudadanoPostgresRepository {
	return &CiudadanoPostgresRepository{db: db}
}

func (r *CiudadanoPostgresRepository) Create(ctx context.Context, c *entities.Ciudadano) (int, error) {
	const q = `
		INSERT INTO usuario (nombre, email, alias, password, role_id, eliminado, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, FALSE, $6, $6)
		RETURNING user_id
	`

	var id int
	err := r.db.QueryRow(
		ctx,
		q,
		c.Alias,
		c.Email,
		c.Alias,
		c.Password,
		core.CIUDADANO,
		c.CreatedAt,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CiudadanoPostgresRepository) GetByID(ctx context.Context, id int) (*entities.Ciudadano, error) {
	const q = `
		SELECT user_id, email, alias, password, created_at
		FROM usuario
		WHERE user_id = $1 AND role_id = $2 AND eliminado = FALSE
	`

	var c entities.Ciudadano
	err := r.db.QueryRow(ctx, q, id, core.CIUDADANO).Scan(
		&c.ID,
		&c.Email,
		&c.Alias,
		&c.Password,
		&c.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *CiudadanoPostgresRepository) List(ctx context.Context) ([]entities.Ciudadano, error) {
	const q = `
		SELECT user_id, email, alias, password, created_at
		FROM usuario
		WHERE role_id = $1 AND eliminado = FALSE
		ORDER BY user_id DESC
	`

	rows, err := r.db.Query(ctx, q, core.CIUDADANO)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []entities.Ciudadano

	for rows.Next() {
		var c entities.Ciudadano
		if err := rows.Scan(
			&c.ID,
			&c.Email,
			&c.Alias,
			&c.Password,
			&c.CreatedAt,
		); err != nil {
			return nil, err
		}
		out = append(out, c)
	}

	return out, rows.Err()
}

func (r *CiudadanoPostgresRepository) Update(ctx context.Context, c *entities.Ciudadano) error {
	const q = `
		UPDATE usuario
		SET email = $1,
		    alias = $2,
		    nombre = $2,
		    password = $3,
		    updated_at = NOW()
		WHERE user_id = $4 AND role_id = $5 AND eliminado = FALSE
	`

	cmdTag, err := r.db.Exec(ctx, q, c.Email, c.Alias, c.Password, c.ID, core.CIUDADANO)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return errors.New("ciudadano no encontrado")
	}

	return nil
}

func (r *CiudadanoPostgresRepository) Delete(ctx context.Context, id int) error {
	const q = `
		UPDATE usuario
		SET eliminado = TRUE, updated_at = NOW()
		WHERE user_id = $1 AND role_id = $2 AND eliminado = FALSE
	`

	cmdTag, err := r.db.Exec(ctx, q, id, core.CIUDADANO)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return errors.New("ciudadano no encontrado")
	}

	return nil
}

func (r *CiudadanoPostgresRepository) FindByEmail(ctx context.Context, email string) (*entities.Ciudadano, error) {
	const q = `
		SELECT user_id, email, alias, password, created_at
		FROM usuario
		WHERE email = $1 AND role_id = $2 AND eliminado = FALSE
	`

	var c entities.Ciudadano
	err := r.db.QueryRow(ctx, q, email, core.CIUDADANO).Scan(
		&c.ID,
		&c.Email,
		&c.Alias,
		&c.Password,
		&c.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &c, nil
}

func (r *CiudadanoPostgresRepository) FindByAlias(ctx context.Context, alias string) (*entities.Ciudadano, error) {
	const q = `
		SELECT user_id, email, alias, password, created_at
		FROM usuario
		WHERE alias = $1 AND role_id = $2 AND eliminado = FALSE
	`

	var c entities.Ciudadano
	err := r.db.QueryRow(ctx, q, alias, core.CIUDADANO).Scan(
		&c.ID,
		&c.Email,
		&c.Alias,
		&c.Password,
		&c.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &c, nil
}

func (r *CiudadanoPostgresRepository) FindByEmailOrAlias(ctx context.Context, value string) (*entities.Ciudadano, error) {
	const q = `
		SELECT user_id, email, alias, password, created_at
		FROM usuario
		WHERE (email = $1 OR alias = $1) AND role_id = $2 AND eliminado = FALSE
	`

	var c entities.Ciudadano
	err := r.db.QueryRow(ctx, q, value, core.CIUDADANO).Scan(
		&c.ID,
		&c.Email,
		&c.Alias,
		&c.Password,
		&c.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &c, nil
}
