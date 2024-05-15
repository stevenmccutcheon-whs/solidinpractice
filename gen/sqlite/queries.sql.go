// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"
)

const allServices = `-- name: AllServices :many
SELECT s.service_id, s.service_name, s.repo_url, s.service_language_id, s.service_dependency_id, s.archived, s.repo_created_at, s.repo_updated_at, s.created_at, s.updated_at, d.dependency_id, d.dependencies, d.created_at, d.updated_at, l.language_id, l.language_name, l.language_version, l.created_at, l.updated_at
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
`

type AllServicesRow struct {
	ServiceID           int64
	ServiceName         string
	RepoUrl             string
	ServiceLanguageID   sql.NullInt64
	ServiceDependencyID sql.NullInt64
	Archived            bool
	RepoCreatedAt       time.Time
	RepoUpdatedAt       time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DependencyID        sql.NullInt64
	Dependencies        sql.NullString
	CreatedAt_2         sql.NullTime
	UpdatedAt_2         sql.NullTime
	LanguageID          sql.NullInt64
	LanguageName        sql.NullString
	LanguageVersion     sql.NullString
	CreatedAt_3         sql.NullTime
	UpdatedAt_3         sql.NullTime
}

func (q *Queries) AllServices(ctx context.Context) ([]AllServicesRow, error) {
	rows, err := q.db.QueryContext(ctx, allServices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AllServicesRow
	for rows.Next() {
		var i AllServicesRow
		if err := rows.Scan(
			&i.ServiceID,
			&i.ServiceName,
			&i.RepoUrl,
			&i.ServiceLanguageID,
			&i.ServiceDependencyID,
			&i.Archived,
			&i.RepoCreatedAt,
			&i.RepoUpdatedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DependencyID,
			&i.Dependencies,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.LanguageID,
			&i.LanguageName,
			&i.LanguageVersion,
			&i.CreatedAt_3,
			&i.UpdatedAt_3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const allServicesByDependency = `-- name: AllServicesByDependency :many
SELECT s.service_id, s.service_name, s.repo_url, s.service_language_id, s.service_dependency_id, s.archived, s.repo_created_at, s.repo_updated_at, s.created_at, s.updated_at, d.dependency_id, d.dependencies, d.created_at, d.updated_at, l.language_id, l.language_name, l.language_version, l.created_at, l.updated_at
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE d.dependencies LIKE '%' || ? || '%'
`

type AllServicesByDependencyRow struct {
	ServiceID           int64
	ServiceName         string
	RepoUrl             string
	ServiceLanguageID   sql.NullInt64
	ServiceDependencyID sql.NullInt64
	Archived            bool
	RepoCreatedAt       time.Time
	RepoUpdatedAt       time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DependencyID        sql.NullInt64
	Dependencies        sql.NullString
	CreatedAt_2         sql.NullTime
	UpdatedAt_2         sql.NullTime
	LanguageID          sql.NullInt64
	LanguageName        sql.NullString
	LanguageVersion     sql.NullString
	CreatedAt_3         sql.NullTime
	UpdatedAt_3         sql.NullTime
}

func (q *Queries) AllServicesByDependency(ctx context.Context, dollar_1 sql.NullString) ([]AllServicesByDependencyRow, error) {
	rows, err := q.db.QueryContext(ctx, allServicesByDependency, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AllServicesByDependencyRow
	for rows.Next() {
		var i AllServicesByDependencyRow
		if err := rows.Scan(
			&i.ServiceID,
			&i.ServiceName,
			&i.RepoUrl,
			&i.ServiceLanguageID,
			&i.ServiceDependencyID,
			&i.Archived,
			&i.RepoCreatedAt,
			&i.RepoUpdatedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DependencyID,
			&i.Dependencies,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.LanguageID,
			&i.LanguageName,
			&i.LanguageVersion,
			&i.CreatedAt_3,
			&i.UpdatedAt_3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const allServicesByLanguage = `-- name: AllServicesByLanguage :many
SELECT s.service_id, s.service_name, s.repo_url, s.service_language_id, s.service_dependency_id, s.archived, s.repo_created_at, s.repo_updated_at, s.created_at, s.updated_at, d.dependency_id, d.dependencies, d.created_at, d.updated_at, l.language_id, l.language_name, l.language_version, l.created_at, l.updated_at
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE l.language_name IN ?
`

type AllServicesByLanguageRow struct {
	ServiceID           int64
	ServiceName         string
	RepoUrl             string
	ServiceLanguageID   sql.NullInt64
	ServiceDependencyID sql.NullInt64
	Archived            bool
	RepoCreatedAt       time.Time
	RepoUpdatedAt       time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DependencyID        sql.NullInt64
	Dependencies        sql.NullString
	CreatedAt_2         sql.NullTime
	UpdatedAt_2         sql.NullTime
	LanguageID          sql.NullInt64
	LanguageName        sql.NullString
	LanguageVersion     sql.NullString
	CreatedAt_3         sql.NullTime
	UpdatedAt_3         sql.NullTime
}

func (q *Queries) AllServicesByLanguage(ctx context.Context, languageName string) ([]AllServicesByLanguageRow, error) {
	rows, err := q.db.QueryContext(ctx, allServicesByLanguage, languageName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AllServicesByLanguageRow
	for rows.Next() {
		var i AllServicesByLanguageRow
		if err := rows.Scan(
			&i.ServiceID,
			&i.ServiceName,
			&i.RepoUrl,
			&i.ServiceLanguageID,
			&i.ServiceDependencyID,
			&i.Archived,
			&i.RepoCreatedAt,
			&i.RepoUpdatedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DependencyID,
			&i.Dependencies,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.LanguageID,
			&i.LanguageName,
			&i.LanguageVersion,
			&i.CreatedAt_3,
			&i.UpdatedAt_3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const allServicesByLanguageAndDependency = `-- name: AllServicesByLanguageAndDependency :many
SELECT s.service_id, s.service_name, s.repo_url, s.service_language_id, s.service_dependency_id, s.archived, s.repo_created_at, s.repo_updated_at, s.created_at, s.updated_at, d.dependency_id, d.dependencies, d.created_at, d.updated_at, l.language_id, l.language_name, l.language_version, l.created_at, l.updated_at
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE l.language_name IN /*SLICE:languages*/?
AND d.dependencies LIKE '%' || ?2 || '%'
`

type AllServicesByLanguageAndDependencyParams struct {
	Languages    []string
	Dependencies sql.NullString
}

type AllServicesByLanguageAndDependencyRow struct {
	ServiceID           int64
	ServiceName         string
	RepoUrl             string
	ServiceLanguageID   sql.NullInt64
	ServiceDependencyID sql.NullInt64
	Archived            bool
	RepoCreatedAt       time.Time
	RepoUpdatedAt       time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DependencyID        sql.NullInt64
	Dependencies        sql.NullString
	CreatedAt_2         sql.NullTime
	UpdatedAt_2         sql.NullTime
	LanguageID          sql.NullInt64
	LanguageName        sql.NullString
	LanguageVersion     sql.NullString
	CreatedAt_3         sql.NullTime
	UpdatedAt_3         sql.NullTime
}

func (q *Queries) AllServicesByLanguageAndDependency(ctx context.Context, arg AllServicesByLanguageAndDependencyParams) ([]AllServicesByLanguageAndDependencyRow, error) {
	query := allServicesByLanguageAndDependency
	var queryParams []interface{}
	if len(arg.Languages) > 0 {
		for _, v := range arg.Languages {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:languages*/?", strings.Repeat(",?", len(arg.Languages))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:languages*/?", "NULL", 1)
	}
	queryParams = append(queryParams, arg.Dependencies)
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AllServicesByLanguageAndDependencyRow
	for rows.Next() {
		var i AllServicesByLanguageAndDependencyRow
		if err := rows.Scan(
			&i.ServiceID,
			&i.ServiceName,
			&i.RepoUrl,
			&i.ServiceLanguageID,
			&i.ServiceDependencyID,
			&i.Archived,
			&i.RepoCreatedAt,
			&i.RepoUpdatedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DependencyID,
			&i.Dependencies,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.LanguageID,
			&i.LanguageName,
			&i.LanguageVersion,
			&i.CreatedAt_3,
			&i.UpdatedAt_3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const allServicesByLanguageVersion = `-- name: AllServicesByLanguageVersion :many
SELECT s.service_id, s.service_name, s.repo_url, s.service_language_id, s.service_dependency_id, s.archived, s.repo_created_at, s.repo_updated_at, s.created_at, s.updated_at, d.dependency_id, d.dependencies, d.created_at, d.updated_at, l.language_id, l.language_name, l.language_version, l.created_at, l.updated_at
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE l.language_version = ?
`

type AllServicesByLanguageVersionRow struct {
	ServiceID           int64
	ServiceName         string
	RepoUrl             string
	ServiceLanguageID   sql.NullInt64
	ServiceDependencyID sql.NullInt64
	Archived            bool
	RepoCreatedAt       time.Time
	RepoUpdatedAt       time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DependencyID        sql.NullInt64
	Dependencies        sql.NullString
	CreatedAt_2         sql.NullTime
	UpdatedAt_2         sql.NullTime
	LanguageID          sql.NullInt64
	LanguageName        sql.NullString
	LanguageVersion     sql.NullString
	CreatedAt_3         sql.NullTime
	UpdatedAt_3         sql.NullTime
}

func (q *Queries) AllServicesByLanguageVersion(ctx context.Context, languageVersion string) ([]AllServicesByLanguageVersionRow, error) {
	rows, err := q.db.QueryContext(ctx, allServicesByLanguageVersion, languageVersion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AllServicesByLanguageVersionRow
	for rows.Next() {
		var i AllServicesByLanguageVersionRow
		if err := rows.Scan(
			&i.ServiceID,
			&i.ServiceName,
			&i.RepoUrl,
			&i.ServiceLanguageID,
			&i.ServiceDependencyID,
			&i.Archived,
			&i.RepoCreatedAt,
			&i.RepoUpdatedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DependencyID,
			&i.Dependencies,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.LanguageID,
			&i.LanguageName,
			&i.LanguageVersion,
			&i.CreatedAt_3,
			&i.UpdatedAt_3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createService = `-- name: CreateService :one
INSERT OR REPLACE INTO services (
    service_name,
    repo_url,
    service_language_id,
    service_dependency_id,
    archived
    )
VALUES ( ?, ?, ?, ?, ? ) RETURNING service_id, service_name, repo_url, service_language_id, service_dependency_id, archived, repo_created_at, repo_updated_at, created_at, updated_at
`

type CreateServiceParams struct {
	ServiceName         string
	RepoUrl             string
	ServiceLanguageID   sql.NullInt64
	ServiceDependencyID sql.NullInt64
	Archived            bool
}

func (q *Queries) CreateService(ctx context.Context, arg CreateServiceParams) (Service, error) {
	row := q.db.QueryRowContext(ctx, createService,
		arg.ServiceName,
		arg.RepoUrl,
		arg.ServiceLanguageID,
		arg.ServiceDependencyID,
		arg.Archived,
	)
	var i Service
	err := row.Scan(
		&i.ServiceID,
		&i.ServiceName,
		&i.RepoUrl,
		&i.ServiceLanguageID,
		&i.ServiceDependencyID,
		&i.Archived,
		&i.RepoCreatedAt,
		&i.RepoUpdatedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createServiceDependency = `-- name: CreateServiceDependency :one
INSERT OR REPLACE INTO service_dependencies (dependencies) VALUES (?) RETURNING dependency_id, dependencies, created_at, updated_at
`

func (q *Queries) CreateServiceDependency(ctx context.Context, dependencies string) (ServiceDependency, error) {
	row := q.db.QueryRowContext(ctx, createServiceDependency, dependencies)
	var i ServiceDependency
	err := row.Scan(
		&i.DependencyID,
		&i.Dependencies,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createServiceLanguage = `-- name: CreateServiceLanguage :one
INSERT OR REPLACE INTO service_languages (language_name, language_version) VALUES (?, ?) RETURNING language_id, language_name, language_version, created_at, updated_at
`

type CreateServiceLanguageParams struct {
	LanguageName    string
	LanguageVersion string
}

func (q *Queries) CreateServiceLanguage(ctx context.Context, arg CreateServiceLanguageParams) (ServiceLanguage, error) {
	row := q.db.QueryRowContext(ctx, createServiceLanguage, arg.LanguageName, arg.LanguageVersion)
	var i ServiceLanguage
	err := row.Scan(
		&i.LanguageID,
		&i.LanguageName,
		&i.LanguageVersion,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const doesLanguageExist = `-- name: DoesLanguageExist :one
SELECT language_id FROM service_languages WHERE language_name = ? AND language_version = ?
`

type DoesLanguageExistParams struct {
	LanguageName    string
	LanguageVersion string
}

func (q *Queries) DoesLanguageExist(ctx context.Context, arg DoesLanguageExistParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, doesLanguageExist, arg.LanguageName, arg.LanguageVersion)
	var language_id int64
	err := row.Scan(&language_id)
	return language_id, err
}

const getServiceByName = `-- name: GetServiceByName :one
SELECT s.service_id, s.service_name, s.repo_url, s.service_language_id, s.service_dependency_id, s.archived, s.repo_created_at, s.repo_updated_at, s.created_at, s.updated_at, d.dependency_id, d.dependencies, d.created_at, d.updated_at, l.language_id, l.language_name, l.language_version, l.created_at, l.updated_at
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE s.service_name = ?
`

type GetServiceByNameRow struct {
	ServiceID           int64
	ServiceName         string
	RepoUrl             string
	ServiceLanguageID   sql.NullInt64
	ServiceDependencyID sql.NullInt64
	Archived            bool
	RepoCreatedAt       time.Time
	RepoUpdatedAt       time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DependencyID        sql.NullInt64
	Dependencies        sql.NullString
	CreatedAt_2         sql.NullTime
	UpdatedAt_2         sql.NullTime
	LanguageID          sql.NullInt64
	LanguageName        sql.NullString
	LanguageVersion     sql.NullString
	CreatedAt_3         sql.NullTime
	UpdatedAt_3         sql.NullTime
}

func (q *Queries) GetServiceByName(ctx context.Context, serviceName string) (GetServiceByNameRow, error) {
	row := q.db.QueryRowContext(ctx, getServiceByName, serviceName)
	var i GetServiceByNameRow
	err := row.Scan(
		&i.ServiceID,
		&i.ServiceName,
		&i.RepoUrl,
		&i.ServiceLanguageID,
		&i.ServiceDependencyID,
		&i.Archived,
		&i.RepoCreatedAt,
		&i.RepoUpdatedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DependencyID,
		&i.Dependencies,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
		&i.LanguageID,
		&i.LanguageName,
		&i.LanguageVersion,
		&i.CreatedAt_3,
		&i.UpdatedAt_3,
	)
	return i, err
}