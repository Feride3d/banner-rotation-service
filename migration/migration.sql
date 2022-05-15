CREATE TABLE "banner" (
	"id" TEXT NOT NULL,
	"description" TEXT NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE "slot" (
	"id" TEXT NOT NULL,
	"description" TEXT NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE "group" (
	"id" TEXT NOT NULL,
	"description" TEXT NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE "banner_location" (
	"banner_id" TEXT NOT NULL
    "slot_id" TEXT NOT NULL,
	
);

CREATE TABLE "clicks" (
	"slot_id" TEXT NOT NULL,
	"banner_id" TEXT NOT NULL,
	"group_id" TEXT NOT NULL,
	"time" TEXT NOT NULL
);

CREATE TABLE "displays" (
	"slot_id" TEXT NOT NULL,
	"banner_id" TEXT NOT NULL,
	"group_id" TEXT NOT NULL,
	"time" TEXT NOT NULL
);