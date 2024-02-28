-- name: GetMovie :one 
SELECT * FROM movies 
WHERE id = $1 LIMIT 1;

-- name: ListMovies :many
SELECT * FROM movies 
WHERE genres && $3 
ORDER BY runtime
LIMIT $1 OFFSET $2;

-- name: CreateMovie :exec
INSERT INTO movies (title, year, runtime , genres) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateMovie :exec
UPDATE movies 
SET title = CASE WHEN @update_title::boolean THEN @title::VARCHAR(50) ELSE title END,
    year = CASE WHEN @update_year::boolean THEN @year::INTEGER ELSE year END,
    runtime = CASE WHEN @update_runtime::boolean THEN @runtime::INTEGER ELSE runtime END
WHERE id = @id
RETURNING *;

-- name: DeleteMovie :exec
UPDATE movies
SET deleted_at = now()
WHERE id = $1;  