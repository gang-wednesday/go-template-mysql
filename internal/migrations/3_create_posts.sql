-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE posts (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title text NOT NULL,
			content text  NOT NULL,
            author_id int NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP,
			deleted_at TIMESTAMP,
			CONSTRAINT fk__authors_posts FOREIGN KEY (author_id) REFERENCES authors(id)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE posts;