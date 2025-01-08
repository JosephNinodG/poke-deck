create table attack(
 id serial primary key,
 "name"	varchar(50) NOT NULL,
 "cost" varchar(50) ARRAY NOT NULL,
 convertedEnergyCost integer NOT NULL,
  damage varchar(50) NOT NULL,
  description varchar(200) NOT NULL
);

create table images(
 id serial primary key,
 small	varchar(100),
 "large" varchar(100),
 "symbol" varchar(100),
 "logo" varchar(100)
);

create table legalities(
 id serial primary key,
 standard	varchar(50) NOT NULL,
 expanded varchar(50) NOT NULL,
 unlimited varchar(50) NOT NULL
);

create table traits(
 id serial primary key,
 "name"	varchar(100),
 "description" varchar(200),
 "type" varchar(100),
 "value" varchar(100)
);

create table "user"(
 id serial primary key,
 "name"	varchar(200) NOT NULL,
 "email" varchar(200) NOT NULL
);

create table "set"(
 id serial primary key,
 "name"	varchar(200) NOT NULL,
 "series" varchar(50) NOT NULL,
  printedTotal integer NOT NULL,
  total integer NOT NULL,
  ptcgoCode varchar(50) NOT NULL,
  releaseDate varchar(50) NOT NULL,
  updatedAt varchar(50) NOT NULL,
  legalities_id integer references "legalities" (id),
  images_id integer references "images" (id)
);

create table "card"(
 id serial primary key,
 "name"	varchar(200) NOT NULL,
 	supertype varchar(50) NOT NULL,
  subtypes varchar(100) ARRAY NOT NULL,
  "level" varchar(50) NOT NULL,
  "hp" varchar(50) NOT NULL,
  "types" varchar(100) ARRAY NOT NULL,
  evolvesFrom varchar(100) NOT NULL,
  evolvesTo varchar(100) ARRAY NOT NULL,
  rules varchar(200) ARRAY NOT NULL,
  retreatCost varchar(50) ARRAY NOT NULL,
  convertedRetreatCost integer NOT NULL,
  "number" varchar(50) NOT NULL,
  "artist" varchar(50) NOT NULL,
  "rarity" varchar(50) NOT NULL,
  "flavorText" varchar(50) NOT NULL,
  "nationalPokedexNumbers" integer ARRAY NOT NULL,
  ancientTrait_id integer references "traits" (id),
  set_id integer references "set" (id),
  legalities_id integer references "legalities" (id),
  images_id integer references "images" (id)
);

create table "card_ability"(
	card_id integer references card,
  ability_id integer references traits,
  primary key (card_id, ability_id)
);

create table "card_weakness"(
	card_id integer references card,
  weakness_id integer references traits,
  primary key (card_id, weakness_id)
);

create table "card_resistance"(
	card_id integer references card,
  resistance_id integer references traits,
  primary key (card_id, resistance_id)
);

create table "card_attack"(
	card_id integer references card,
  attack_id integer references traits,
  primary key (card_id, attack_id)
);

create table "collection"(
 id serial primary key,
 "name"	varchar(200) NOT NULL,
 user_id integer references "user" (id)
);

create table "collection_card"(
	card_id integer references card,
  collection_id integer references collection,
  primary key (card_id, collection_id)
);

insert into "user" ("name", "email")
values ('test1', 'test1@email.com')

insert into "attack" ("name", "cost", "convertedenergycost", "damage", "description")
values ('attack1', '{"colorless","colorless"}', 2, '40', 'Placeholder description.')

insert into "attack" ("name", "cost", "convertedenergycost", "damage", "description")
values ('attack2', '{"Grass","colorless","colorless"}', 3, '60', 'Placeholder description.')

insert into "images" ("small", "large", "symbol", "logo")
values ('smallimage1.png', 'largeimage1.png', 'symbol1.png', 'logo1.png')

insert into "images" ("small", "large", "symbol", "logo")
values ('smallimage2.png', 'largeimage2.png', 'symbol2.png', 'logo2.png')

insert into "legalities" ("standard", "expanded", "unlimited")
values ('legal', 'legal', 'legal')