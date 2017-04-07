CREATE TABLE comments (
  id SERIAL NOT NULL PRIMARY KEY,
  post_id INT NOT NULL,
  body TEXT NOT NULL
);

ALTER TABLE     comments
ADD CONSTRAINT  fk_post_id_posts_id
FOREIGN KEY     (post_id)
REFERENCES      posts (id)
ON DELETE       RESTRICT;
