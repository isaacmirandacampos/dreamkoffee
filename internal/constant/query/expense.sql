
-- name: CreateExpense :one
INSERT INTO expenses (
  name, price
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetExpense :one
SELECT * FROM expenses WHERE id = $1 and deleted_at is null;

-- name: ListExpenses :many
SELECT * FROM expenses where deleted_at is null ORDER BY id desc;

-- name: UpdateExpense :one
UPDATE expenses
SET name = $1, price = $2, updated_at = now()
WHERE id = $3 and deleted_at is null RETURNING *;
