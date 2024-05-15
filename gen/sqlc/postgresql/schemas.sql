
CREATE TABLE IF NOT EXISTS service_languages (
    language_id SERIAL PRIMARY KEY,
    language_name TEXT NOT NULL,
    language_version TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(language_name, language_version)
);

CREATE TABLE IF NOT EXISTS service_dependencies (
    dependency_id SERIAL PRIMARY KEY,
    dependencies TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS services (
    service_id SERIAL PRIMARY KEY,
    service_name TEXT NOT NULL UNIQUE,
    repo_url TEXT NOT NULL,
    service_language_id INTEGER,
    service_dependency_id INTEGER,
    archived BOOLEAN NOT NULL DEFAULT FALSE,
    repo_created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    repo_updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (service_language_id) REFERENCES service_languages(language_id),
    FOREIGN KEY (service_dependency_id) REFERENCES service_dependencies(dependency_id) ON DELETE CASCADE
);