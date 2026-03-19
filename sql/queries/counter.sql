-- SPDX-FileCopyrightText: 2026 Jarkad <jarkad@jarkad.net.eu.org>
--
-- SPDX-License-Identifier: EUPL-1.2

-- name: Reset :exec
DELETE FROM counter;
INSERT INTO counter VALUES (0);

-- name: Get :one
SELECT count
FROM counter;

-- name: Increment :one
UPDATE counter
SET count = count + 1
RETURNING count;

-- name: Decrement :one
UPDATE counter
SET count = count - 1
RETURNING count;
