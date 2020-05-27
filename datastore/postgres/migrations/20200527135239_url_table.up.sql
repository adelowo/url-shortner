CREATE TABLE IF NOT EXISTS url(
	id SERIAL,
	code CHAR(6) NOT NULL UNIQUE,
	redirect_url TEXT NOT NULL UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
