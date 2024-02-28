-- name: GetMovie :one 
SELECT * FROM movies 
WHERE id = $1 LIMIT 1;

-- name: ListMovies :many
SELECT * FROM movies 
WHERE 
    CASE cardinality(@filter::VARCHAR(20)[]) 
        WHEN 0 THEN true 
        ELSE genres && @filter 
    END  
ORDER BY CASE WHEN @sort_by = 'asc' or @sort_by = '' THEN runtime END ASC ,
         CASE WHEN @sort_by = 'desc' THEN runtime END DESC
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