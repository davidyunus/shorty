create table "url" (
	"id" serial primary key not null,
	"url" varchar(255) not null,
	"shortcode" varchar(255) not null,
	"redirect_count" int not null default 0,
	"start_date" timestamptz not null,
	"last_seen_date" timestamptz not null 
);