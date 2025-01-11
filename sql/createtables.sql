create table attack(
 id serial primary key,
 "name"	varchar(50) NOT NULL,
 "cost" varchar(50) ARRAY NOT NULL,
 "convertedEnergyCost" integer NOT NULL,
 "damage" varchar(50) NOT NULL,
 "description" varchar(200) NOT NULL
);

create table set_images(
 id serial primary key,
 "symbol" varchar(100),
 "logo" varchar(100)
);

create table set_legalities(
 id serial primary key,
 "standard"	varchar(50) NOT NULL,
 "expanded" varchar(50) NOT NULL,
 "unlimited" varchar(50) NOT NULL
);

create table card_images(
 id serial primary key,
 "small"	varchar(100),
 "large" varchar(100)
);

create table card_legalities(
 id serial primary key,
 "standard"	varchar(50) NOT NULL,
 "expanded" varchar(50) NOT NULL,
 "unlimited" varchar(50) NOT NULL
);

create table ancient_trait(
 id serial primary key,
 "name"	varchar(100),
 "description" varchar(200)
);

create table ability(
 id serial primary key,
 "name"	varchar(100),
 "description" varchar(200),
 "type" varchar(100)
);

create table weakness(
 id serial primary key,
 "type" varchar(100),
 "value" varchar(100)
);

create table resistance(
 id serial primary key,
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
  "printedTotal" integer NOT NULL,
  "total" integer NOT NULL,
  "ptcgoCode" varchar(50) NOT NULL,
  "releaseDate" varchar(50) NOT NULL,
  "updatedAt" varchar(50) NOT NULL,
  set_legalities_id integer references "set_legalities" (id),
  set_images_id integer references "set_images" (id)
);

create table "card"(
 id serial primary key,
 "name"	varchar(200) NOT NULL,
 	"supertype" varchar(50) NOT NULL,
  "subtypes" varchar(100) ARRAY NOT NULL,
  "level" varchar(50),
  "hp" varchar(50) NOT NULL,
  "types" varchar(100) ARRAY NOT NULL,
  "evolvesFrom" varchar(100),
  "evolvesTo" varchar(100) ARRAY NOT NULL,
  "rules" varchar(200) ARRAY NOT NULL,
  "retreatCost" varchar(50) ARRAY NOT NULL,
  "convertedRetreatCost" integer NOT NULL,
  "number" varchar(50) NOT NULL,
  "artist" varchar(50) NOT NULL,
  "rarity" varchar(50) NOT NULL,
  "flavorText" varchar(50) NOT NULL,
  "nationalPokedexNumbers" integer ARRAY NOT NULL,
  ancient_trait_id integer references "ancient_trait" (id),
  set_id integer references "set" (id),
  card_legalities_id integer references "card_legalities" (id),
  card_images_id integer references "card_images" (id)
);

create table "card_ability"(
	card_id integer references card,
  ability_id integer references ability,
  primary key (card_id, ability_id)
);

create table "card_weakness"(
	card_id integer references card,
  weakness_id integer references weakness,
  primary key (card_id, weakness_id)
);

create table "card_resistance"(
	card_id integer references card,
  resistance_id integer references resistance,
  primary key (card_id, resistance_id)
);

create table "card_attack"(
	card_id integer references card,
  attack_id integer references attack,
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
values ('test1', 'test1@email.com');

insert into "attack" ("name", "cost", "convertedEnergyCost", "damage", "description")
values ('attack1', '{"colorless","colorless"}', 2, '40', 'Placeholder description.'),
('attack2', '{"Lightning","colorless","colorless"}', 3, '100', 'Placeholder description.');

insert into "set_images" ("symbol", "logo")
values ('set-symbol1.png', 'set-logo1.png'),
('set-symbol2.png', 'set-logo2.png');

insert into "set_legalities" ("standard", "expanded", "unlimited")
values ('legal', 'legal', 'legal');

insert into "card_images" ("small", "large")
values ('card-smallimage1.png', 'card-largeimage1.png'),
('card-smallimage2.png', 'card-largeimage2.png');

insert into "card_legalities" ("standard", "expanded", "unlimited")
values ('legal', 'legal', 'legal');

insert into "ancient_trait" ("name", "description")
values ('ancienttrait1', 'Placeholder description.'),
('ancienttrait2', 'Placeholder description.');

insert into "ability" ("name", "description", "type")
values ('ability1', 'Placeholder description.', 'Poke-Power'),
('ability2', 'Placeholder description.', 'Ability'),
('ability3', 'Placeholder description.', 'Poké-Body');

insert into "weakness" ("type", "value")
values ('Fighting', '+20'), ('Water', 'x2');

insert into "resistance" ("type", "value")
values ('Metal', '-20');

insert into "set" ("name", "series", "printedTotal", "total", "ptcgoCode", "releaseDate", "updatedAt", "set_legalities_id", "set_images_id")
values ('set1','series1',100,150,'s1','2025/01/01','2025/01/01',1,1),
('set2','series2',100,150,'s2','2024/01/01','2024/01/01',1,2);

insert into "card" ("name","supertype","subtypes","level","hp","types","evolvesFrom","evolvesTo","rules","retreatCost","convertedRetreatCost","number","artist","rarity","flavorText","nationalPokedexNumbers","ancient_trait_id","set_id","card_legalities_id","card_images_id")
values ('card1','Pokemon','{"Basic"}',null,'50','{"Lightning"}',null,'{"card2"}','{"Placeholder rules."}','{"Colorless","Colorless"}',2,1,'artist','Uncommon','Placeholder text.','{1}',null,1,1,1),
('card2','Pokemon','{"Stage 1"}',null,'100','{"Lightning"}','card1','{}','{"Placeholder rules."}','{"Lightning","Colorless","Colorless"}',3,2,'artist','Rare','Placeholder text.','{2}',null,1,1,2);

insert into "card_ability" ("card_id", "ability_id")
values (2, 2);

insert into "card_attack" ("card_id", "attack_id")
values (1, 1), (2, 2);

insert into "card_resistance" ("card_id", "resistance_id")
values (1, 1), (2, 1);

insert into "card_weakness" ("card_id", "weakness_id")
values (1, 1), (2, 1);

insert into "collection" ("name", "user_id")
values ('testcollection1', 1);

insert into "collection_card" ("card_id", "collection_id")
values (1, 1), (2, 1);