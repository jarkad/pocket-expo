-- SPDX-FileCopyrightText: 2026 Jarkad <jarkad@jarkad.net.eu.org>
--
-- SPDX-License-Identifier: EUPL-1.2

-- migrate:up

CREATE TABLE counter (
    count INTEGER NOT NULL
);

-- migrate:down

DROP TABLE counter;
