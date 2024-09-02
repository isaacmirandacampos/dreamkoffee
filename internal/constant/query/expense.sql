-- name: CreateExpense :one
INSERT INTO expenses (
  user_id, description, value, paid_at, payment_at, note
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetExpense :one
SELECT * FROM expenses WHERE id = $1 and deleted_at is null;

-- name: ListExpenses :many
SELECT * FROM expenses where deleted_at is null ORDER BY id desc;

-- name: UpdateExpense :one
UPDATE expenses
SET 
  description = $1, 
  value = $2, 
  paid_at = $3, 
  payment_at = $4, 
  note = $5,
  updated_at = now()
WHERE id = $6 and deleted_at is null RETURNING *;

-- name: DeleteExpense :one
UPDATE expenses
SET deleted_at = now()
WHERE id = $1 and deleted_at is null RETURNING *;

-- name: GetLastExpense :one
SELECT * FROM expenses where deleted_at is null ORDER BY id desc LIMIT 1;