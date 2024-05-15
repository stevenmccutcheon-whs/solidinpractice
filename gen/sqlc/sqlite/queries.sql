-- name: CreateService :one
INSERT OR REPLACE INTO services (
    service_name,
    repo_url,
    service_language_id,
    service_dependency_id,
    archived
    )
VALUES ( ?, ?, ?, ?, ? ) RETURNING *;

-- name: CreateServiceDependency :one
INSERT OR REPLACE INTO service_dependencies (dependencies) VALUES (?) RETURNING *;

-- name: CreateServiceLanguage :one
INSERT OR REPLACE INTO service_languages (language_name, language_version) VALUES (?, ?) RETURNING *;

-- name: DoesLanguageExist :one
SELECT language_id FROM service_languages WHERE language_name = ? AND language_version = ?;

-- name: GetServiceByName :one
SELECT s.*, d.*, l.*
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE s.service_name = ?;

-- name: AllServicesByLanguage :many
SELECT s.*, d.*, l.*
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE l.language_name IN ?;

-- name: AllServicesByLanguageVersion :many
SELECT s.*, d.*, l.*
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE l.language_version = ?;

-- name: AllServicesByDependency :many
SELECT s.*, d.*, l.*
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE d.dependencies LIKE '%' || ? || '%';

-- name: AllServicesByLanguageAndDependency :many
SELECT s.*, d.*, l.*
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE l.language_name IN sqlc.slice(languages)
AND d.dependencies LIKE '%' || sqlc.arg(dependencies) || '%';

-- name: AllServices :many
SELECT s.*, d.*, l.*
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id;