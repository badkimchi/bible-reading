-- name: GetAccount :one
SELECT *
FROM Accounts
WHERE Account_ID = $1 LIMIT 1;

-- name: GetAccountByEmail :one
SELECT *
FROM Accounts
WHERE Email = $1 LIMIT 1;


-- name: CreateAccount :one
INSERT INTO Accounts (Name, Level, Email)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateAccount :exec
UPDATE Accounts
SET Name     = $2,
    Level    = $3,
    Email    = $4
WHERE Account_ID = $1;

-- name: DeleteAccount :exec
DELETE
FROM Accounts
WHERE Account_ID = $1;