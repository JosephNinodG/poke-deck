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
  "level" varchar(50),
  "hp" varchar(50) NOT NULL,
  "types" varchar(100) ARRAY NOT NULL,
  evolvesFrom varchar(100),
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
values ('attack2', '{"Lightning","colorless","colorless"}', 3, '100', 'Placeholder description.')

insert into "images" ("small", "large", "symbol", "logo")
values ('card-smallimage1.png', 'card-largeimage1.png', null, null)

insert into "images" ("small", "large", "symbol", "logo")
values (null, null, 'set-symbol1.png', 'set-logo1.png')

insert into "images" ("small", "large", "symbol", "logo")
values ('card-smallimage2.png', 'card-largeimage2.png', null, null)

insert into "images" ("small", "large", "symbol", "logo")
values (null, null, 'set-symbol2.png', 'set-logo2.png')

insert into "legalities" ("standard", "expanded", "unlimited")
values ('legal', 'legal', 'legal')

insert into "traits" ("name", "description", "type", "value")
values ('ancienttrait1', 'Placeholder description.', null, null),
('ancienttrait2', 'Placeholder description.', null, null),
('ability1', 'Placeholder description.', 'Poke-Power', null),
('ability2', 'Placeholder description.', 'Ability', null),
('ability3', 'Placeholder description.', 'Poké-Body', null),
(null, 'weakness', 'Fighting', '+20'),
(null, 'weakness', 'Water', 'x2'),
(null, 'resistance', 'Metal', '-20')

insert into "sets" ("name", "series", "printedtotal", "total", "ptcgocode", "releasedate", "updateddate", "legalities_id", "images_id")
values ('set1','series1',100,150,'s1','2025/01/01','2025/01/01',1,2),
('set2','series2',100,150,'s2','2024/01/01','2024/01/01',1,4)

insert into "card" ("name","supertype","subtypes","level","hp","types","evolvesfrom","evolvesto","rules","retreatcost","convertedretreatcost","number","artist","rarity","flavorText","nationalPokedexNumbers","ancienttrait_id","set_id","legalities_id","images_id")
values ('card1','Pokemon','{"Basic"}',null,'50','{"Lightning"}',null,'{"card2"}','{"Placeholder rules."}','{"Colorless","Colorless"}',2,1,'artist','Uncommon','Placeholder text.','{1}',null,1,1,1),
('card2','Pokemon','{"Stage 1"}',null,'100','{"Lightning"}','card1','{}','{"Placeholder rules."}','{"Lightning","Colorless","Colorless"}',3,2,'artist','Rare','Placeholder text.','{2}',null,1,1,2)

insert into "card_ability" ("card_id", "ability_id")
values (2, 4)

insert into "card_attack" ("card_id", "attack_id")
values (1, 1), (2, 2) 

insert into "card_resistance" ("card_id", "resistance_id")
values (1, 8), (2, 8)

insert into "card_weakness" ("card_id", "weakness_id")
values (1, 6), (2, 6)

insert into "collection" ("name", "user_id")
values ("testcollection1", 1)

insert into "collection_card" ("card_id", "collection_id")
values (1, 1), (2, 1)

-- example JOIN:
SELECT c as card, a as attack from card c join card_attack ca on ca.card_id = c.id join attack a on a.id = ca.attack_id