-- name: GetMovie :one 
SELECT * FROM movies 
WHERE id = $1 LIMIT 1;

-- name: ListMovies :many
SELECT * FROM movies 
WHERE genres && $3 
ORDER BY runtime
LIMIT $1 OFFSET $2;

-- name: CreateMovie :one
INSERT INTO movies (title, year, runtime , genres) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateMovie :one
UPDATE movies 
SET title = CASE WHEN @title <> '' THEN @title::VARCHAR(50) ELSE title END,
    year = CASE WHEN @year <> 0 THEN @year::INTEGER ELSE year END,
    runtime = CASE WHEN @runtime > 0 THEN @runtime::INTEGER ELSE runtime END
WHERE id = @id
RETURNING *;

-- name: DeleteMovie :exec
UPDATE movies
SET deleted_at = now()
WHERE id = $1;  