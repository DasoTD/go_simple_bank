-- name: CreateEntry :one
INSERT INTO entries (
    amount,
    acccount_id
    ) VALUES(
        $1, $2
    )
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries WHERE id= $1;  

-- nmae: ListEntry :many
SELECT * FROM entries ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateEntry :one
UPDATE entries
  set amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;