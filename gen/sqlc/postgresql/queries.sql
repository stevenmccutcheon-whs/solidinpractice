-- name: CreateService :one
INSERT INTO services (
    service_name,
    repo_url,
    service_language_id,
    service_dependency_id,
    archived
    )
VALUES ( $1, $2, $3, $4, $5 )
ON CONFLICT (service_name) DO UPDATE 
SET service_name = EXCLUDED.service_name,
    repo_url = EXCLUDED.repo_url,
    service_language_id = EXCLUDED.service_language_id,
    service_dependency_id = EXCLUDED.service_dependency_id,
    archived = EXCLUDED.archived
RETURNING *;

-- name: CreateServiceDependency :one
INSERT INTO service_dependencies (dependencies) VALUES ($1)
ON CONFLICT (dependencies) DO UPDATE 
SET dependencies = EXCLUDED.dependencies
RETURNING *;

-- name: CreateServiceLanguage :one
INSERT INTO service_languages (language_name, language_version) VALUES ($1, $2)
ON CONFLICT (language_name, language_version) DO UPDATE 
SET language_name = EXCLUDED.language_name,
    language_version = EXCLUDED.language_version
RETURNING *;

-- name: DoesLanguageExist :one
SELECT language_id FROM service_languages WHERE language_name = $1 AND language_version = $2;

-- name: GetServiceByName :one
SELECT s.*, d.*, l.*
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE s.service_name = $1;

-- name: AllServicesByLanguage :many
SELECT s.*, d.*, l.*
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE l.language_name = ANY(@languages::VARCHAR[]);

-- name: AllServicesByLanguageVersion :many
SELECT s.*, d.*, l.*
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE l.language_version = $1;

-- name: AllServicesByDependency :many
SELECT s.*, d.*, l.*
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE d.dependencies LIKE '%' || $1 || '%';

-- name: AllServicesByLanguageAndDependency :many
SELECT s.*, d.*, l.*
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id
WHERE l.language_name = ANY(@languages::VARCHAR[])
AND d.dependencies LIKE '%' || @dependencies || '%';

-- name: AllServices :many
SELECT s.*, d.*, l.*
FROM services s
LEFT JOIN service_dependencies d ON s.service_dependency_id = d.dependency_id
LEFT JOIN service_languages l ON s.service_language_id = l.language_id;