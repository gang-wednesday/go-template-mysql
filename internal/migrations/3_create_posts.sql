-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE posts (
			id varchar(50) PRIMARY KEY,
			title text NOT NULL,
			content text  NOT NULL,
            author_id varchar(50) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP,
			deleted_at TIMESTAMP
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE posts;