-- GET ALL CARD INFORMATION
SELECT
		card.id,
		card.name,
    card.supertype,
  	card.subtypes,
  	card.level,
  	card.hp,
  	card.types,
  	card."evolvesFrom",
  	card."evolvesTo",
 	  card.rules,
  	card."retreatCost",
  	card."convertedRetreatCost",
  	card.number,
  	card.artist,
  	card.rarity,
  	card."flavorText",
  	card."nationalPokedexNumbers",
    jsonb_build_object(
        'standard', card_legalities.standard,
        'expanded', card_legalities.expanded,
        'unlimited', card_legalities.unlimited
    ) AS legalities,
    jsonb_build_object(
        'small', card_images.small,
        'large', card_images.large
    ) AS images,
    jsonb_build_object(
        'name', ancient_trait.name,
        'description', ancient_trait.description
    ) AS ancient_trait,
    jsonb_build_object(
      'name', "set".name,
      'series', "set".series,
      'printedTotal', "set"."printedTotal",
      'total', "set".total,
      'ptcgoCode', "set"."ptcgoCode",
      'releaseDate', "set"."releaseDate",
      'updatedAt', "set"."updatedAt",
      'legalities', jsonb_build_object (
      		'standard', set_legalities.standard,
        	'expanded', set_legalities.expanded,
        	'unlimited', set_legalities.unlimited
      ),
      'images', jsonb_build_object(
        	'symbol', set_images.logo,
        	'logo', set_images.logo
    	)
    ) AS card_set,
    COALESCE(
      jsonb_agg(
    		jsonb_build_object(
          'name', ability.name,
          'description', ability.description,
          'type', ability.type
    		)
    	),
      '[]'
    ) AS abilities,
    COALESCE(
      jsonb_agg(
    		jsonb_build_object(
          'name', attack.name,
          'description', attack.description,
          'cost', attack.cost,
          'convertedEnergyCost', attack."convertedEnergyCost",
          'damage', attack.damage
    		)
    	),
      '[]'
    ) AS attacks,
    COALESCE(
      jsonb_agg(
    		jsonb_build_object(
          'type', weakness.type,
          'value', weakness.value
    		)
    	),
      '[]'
    ) AS weaknesses,
    COALESCE(
      jsonb_agg(
    		jsonb_build_object(
          'type', resistance.type,
          'value', resistance.value
    		)
    	),
      '[]'
    ) AS resistances
FROM
    card
JOIN
    card_legalities ON card.card_legalities_id = card_legalities.id
JOIN
    card_images ON card.card_images_id = card_images.id
LEFT JOIN
    ancient_trait ON card.ancient_trait_id = ancient_trait.id
JOIN
    "set" ON card.set_id = "set".id
JOIN
    set_legalities ON "set".set_legalities_id = set_legalities.id
JOIN
    set_images ON "set".set_images_id = set_images.id
LEFT JOIN
		card_ability ON card.id = card_ability.card_id
LEFT JOIN
		ability ON card_ability.ability_id = ability.id
LEFT JOIN
		card_attack ON card.id = card_attack.card_id
LEFT JOIN
		attack ON card_attack.attack_id = attack.id 
LEFT JOIN
		card_weakness ON card.id = card_weakness.card_id
LEFT JOIN
		weakness ON card_weakness.weakness_id = weakness.id  
LEFT JOIN
		card_resistance ON card.id = card_resistance.card_id
LEFT JOIN
		resistance ON card_resistance.resistance_id = resistance.id  
GROUP BY
		card.id, card.name, 
    card_legalities.standard, card_legalities.expanded, card_legalities.unlimited, 
    card_images.small, card_images.large, 
    ancient_trait.name, ancient_trait.description,
    "set".name, "set".series, "set"."printedTotal", "set".total, "set"."ptcgoCode", "set"."releaseDate",  "set"."updatedAt",
    set_legalities.standard, set_legalities.expanded, set_legalities.unlimited,
    set_images.logo, set_images.logo

-- Get all the cards in collection x for user y
SELECT
		card.cardID,
		card.name,
    card.supertype,
  	card.subtypes,
  	card.level,
  	card.hp,
  	card.types,
  	card."evolvesFrom",
  	card."evolvesTo",
 	  card.rules,
  	card."retreatCost",
  	card."convertedRetreatCost",
  	card.number,
  	card.artist,
  	card.rarity,
  	card."flavorText",
  	card."nationalPokedexNumbers",
    jsonb_build_object(
        'standard', card_legalities.standard,
        'expanded', card_legalities.expanded,
        'unlimited', card_legalities.unlimited
    ) AS legalities,
    jsonb_build_object(
        'small', card_images.small,
        'large', card_images.large
    ) AS images,
    jsonb_build_object(
        'name', ancient_trait.name,
        'description', ancient_trait.description
    ) AS ancient_trait,
    jsonb_build_object(
      'name', "set".name,
      'series', "set".series,
      'printedTotal', "set"."printedTotal",
      'total', "set".total,
      'ptcgoCode', "set"."ptcgoCode",
      'releaseDate', "set"."releaseDate",
      'updatedAt', "set"."updatedAt",
      'legalities', jsonb_build_object (
      		'standard', set_legalities.standard,
        	'expanded', set_legalities.expanded,
        	'unlimited', set_legalities.unlimited
      ),
      'images', jsonb_build_object(
        	'symbol', set_images.logo,
        	'logo', set_images.logo
    	)
    ) AS card_set,
    COALESCE(
      jsonb_agg(
    		jsonb_build_object(
          'name', ability.name,
          'description', ability.description,
          'type', ability.type
    		)
    	),
      '[]'
    ) AS abilities,
    COALESCE(
      jsonb_agg(
    		jsonb_build_object(
          'name', attack.name,
          'description', attack.description,
          'cost', attack.cost,
          'convertedEnergyCost', attack."convertedEnergyCost",
          'damage', attack.damage
    		)
    	),
      '[]'
    ) AS attacks,
    COALESCE(
      jsonb_agg(
    		jsonb_build_object(
          'type', weakness.type,
          'value', weakness.value
    		)
    	),
      '[]'
    ) AS weaknesses,
    COALESCE(
      jsonb_agg(
    		jsonb_build_object(
          'type', resistance.type,
          'value', resistance.value
    		)
    	),
      '[]'
    ) AS resistances
FROM
    collection
JOIN
		"user" ON collection.user_id = "user".id
JOIN
		collection_card ON collection.id = collection_card.collection_id
JOIN
    card ON collection_card.card_id = card.id
JOIN
    card_legalities ON card.card_legalities_id = card_legalities.id
JOIN
    card_images ON card.card_images_id = card_images.id
LEFT JOIN
    ancient_trait ON card.ancient_trait_id = ancient_trait.id
JOIN
    "set" ON card.set_id = "set".id
JOIN
    set_legalities ON "set".set_legalities_id = set_legalities.id
JOIN
    set_images ON "set".set_images_id = set_images.id
LEFT JOIN
		card_ability ON card.id = card_ability.card_id
LEFT JOIN
		ability ON card_ability.ability_id = ability.id
LEFT JOIN
		card_attack ON card.id = card_attack.card_id
LEFT JOIN
		attack ON card_attack.attack_id = attack.id 
LEFT JOIN
		card_weakness ON card.id = card_weakness.card_id
LEFT JOIN
		weakness ON card_weakness.weakness_id = weakness.id  
LEFT JOIN
		card_resistance ON card.id = card_resistance.card_id
LEFT JOIN
		resistance ON card_resistance.resistance_id = resistance.id
WHERE
		collection.id = 1 AND "user".id = 1 --REPLACE values of 1 with vars
GROUP BY
		card.id, card.name, 
    card_legalities.standard, card_legalities.expanded, card_legalities.unlimited, 
    card_images.small, card_images.large, 
    ancient_trait.name, ancient_trait.description,
    "set".name, "set".series, "set"."printedTotal", "set".total, "set"."ptcgoCode", "set"."releaseDate",  "set"."updatedAt",
    set_legalities.standard, set_legalities.expanded, set_legalities.unlimited,
    set_images.logo, set_images.logo

-- Get all the cards in collection x for user y as a JSON object
SELECT
	jsonb_build_object(
  	'id', card."cardID",
		'name', card.name,
    'supertype', card.supertype,
  	'subtypes', card.subtypes,
  	'level', card.level,
  	'hp', card.hp,
  	'types', card.types,
  	'evolvesFrom', card."evolvesFrom",
  	'evolvesTo', card."evolvesTo",
 	  'rules', card.rules,
  	'retreatCost', card."retreatCost",
  	'convertedRetreatCost', card."convertedRetreatCost",
  	'number', card.number,
  	'artist', card.artist,
  	'rarity', card.rarity,
  	'flavorText', card."flavorText",
  	'nationalPokedexNumbers', card."nationalPokedexNumbers",
    'legalities', jsonb_build_object(
        'standard', card_legalities.standard,
        'expanded', card_legalities.expanded,
        'unlimited', card_legalities.unlimited
    ),
    'images', jsonb_build_object(
        'small', card_images.small,
        'large', card_images.large
    ),
    'ancient_trait', jsonb_build_object(
        'name', ancient_trait.name,
        'description', ancient_trait.description
    ),
    'set', jsonb_build_object(
      'name', "set".name,
      'series', "set".series,
      'printedTotal', "set"."printedTotal",
      'total', "set".total,
      'ptcgoCode', "set"."ptcgoCode",
      'releaseDate', "set"."releaseDate",
      'updatedAt', "set"."updatedAt",
      'legalities', jsonb_build_object (
      		'standard', set_legalities.standard,
        	'expanded', set_legalities.expanded,
        	'unlimited', set_legalities.unlimited
      ),
      'images', jsonb_build_object(
        	'symbol', set_images.logo,
        	'logo', set_images.logo
    	)
    ),
    'abilities', COALESCE(
      jsonb_agg(
    		jsonb_build_object(
          'name', ability.name,
          'description', ability.description,
          'type', ability.type
    		)
    	),
      '[]'
    ),
    'attacks', COALESCE(
      jsonb_agg(
    		jsonb_build_object(
          'name', attack.name,
          'description', attack.description,
          'cost', attack.cost,
          'convertedEnergyCost', attack."convertedEnergyCost",
          'damage', attack.damage
    		)
    	),
      '[]'
    ),
    'weaknesses', COALESCE(
      jsonb_agg(
    		jsonb_build_object(
          'type', weakness.type,
          'value', weakness.value
    		)
    	),
      '[]'
    ),
    'resistances', COALESCE(
      jsonb_agg(
    		jsonb_build_object(
          'type', resistance.type,
          'value', resistance.value
    		)
    	),
      '[]'
    )
  ) AS card
FROM
    collection
JOIN
		"user" ON collection.user_id = "user".id
JOIN
		collection_card ON collection.id = collection_card.collection_id
JOIN
    card ON collection_card.card_id = card.id
JOIN
    card_legalities ON card.card_legalities_id = card_legalities.id
JOIN
    card_images ON card.card_images_id = card_images.id
LEFT JOIN
    ancient_trait ON card.ancient_trait_id = ancient_trait.id
JOIN
    "set" ON card.set_id = "set".id
JOIN
    set_legalities ON "set".set_legalities_id = set_legalities.id
JOIN
    set_images ON "set".set_images_id = set_images.id
LEFT JOIN
		card_ability ON card.id = card_ability.card_id
LEFT JOIN
		ability ON card_ability.ability_id = ability.id
LEFT JOIN
		card_attack ON card.id = card_attack.card_id
LEFT JOIN
		attack ON card_attack.attack_id = attack.id 
LEFT JOIN
		card_weakness ON card.id = card_weakness.card_id
LEFT JOIN
		weakness ON card_weakness.weakness_id = weakness.id  
LEFT JOIN
		card_resistance ON card.id = card_resistance.card_id
LEFT JOIN
		resistance ON card_resistance.resistance_id = resistance.id
WHERE
		collection.id = 1 AND "user".id = 1 
GROUP BY
		card.id, card.name, 
    card_legalities.standard, card_legalities.expanded, card_legalities.unlimited, 
    card_images.small, card_images.large, 
    ancient_trait.name, ancient_trait.description,
    "set".name, "set".series, "set"."printedTotal", "set".total, "set"."ptcgoCode", "set"."releaseDate",  "set"."updatedAt",
    set_legalities.standard, set_legalities.expanded, set_legalities.unlimited,
    set_images.logo, set_images.logo