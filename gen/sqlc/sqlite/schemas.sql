
CREATE TABLE IF NOT EXISTS services (
    service_id INTEGER PRIMARY KEY AUTOINCREMENT,
    service_name TEXT NOT NULL,
    repo_url TEXT NOT NULL,
    service_language_id INTEGER,
    service_dependency_id INTEGER,
    archived BOOLEAN NOT NULL DEFAULT FALSE,
    repo_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    repo_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (service_language_id) REFERENCES service_languages(language_id),
    FOREIGN KEY (service_dependency_id) REFERENCES service_dependencies(dependency_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS service_languages (
    language_id INTEGER PRIMARY KEY AUTOINCREMENT,
    language_name TEXT NOT NULL,
    language_version TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS service_dependencies (
    dependency_id INTEGER PRIMARY KEY AUTOINCREMENT,
    dependencies TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);