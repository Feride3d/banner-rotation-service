-- +goose Up
-- +goose StatementBegin
CREATE TABLE banners (
	"id" TEXT NOT NULL,
	"description" TEXT NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE slots (
	"id" TEXT NOT NULL,
	"description" TEXT NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE groups (
	"id" TEXT NOT NULL,
	"description" TEXT NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE rotation (
	"banner_id" TEXT NOT NULL
    "slot_id" TEXT NOT NULL,
	
);

CREATE TABLE clicks (
	"slot_id" TEXT NOT NULL,
	"banner_id" TEXT NOT NULL,
	"group_id" TEXT NOT NULL,
	"time" TEXT NOT NULL
);

CREATE TABLE displays (
	"slot_id" TEXT NOT NULL,
	"banner_id" TEXT NOT NULL,
	"group_id" TEXT NOT NULL,
	"time" TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists banners;
DROP TABLE if exists slots;
DROP TABLE if exists groups;
DROP TABLE if exists rotation;
DROP TABLE if exists clicks;
DROP TABLE if exists displays;
-- +goose StatementEnd
id TEXT NOT NULL, PRIMARY KEY,
	description TEXT NOT NULL,