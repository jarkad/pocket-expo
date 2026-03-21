-- SPDX-FileCopyrightText: 2026 Jarkad <jarkad@jarkad.net.eu.org>
--
-- SPDX-License-Identifier: EUPL-1.2

-- name: DeleteAll :exec
DELETE FROM counter;

-- name: Create :exec
INSERT INTO counter (count)
VALUES (?);

-- name: Get :one
SELECT count
FROM counter;

-- name: Set :exec
UPDATE counter
SET count = ?;
