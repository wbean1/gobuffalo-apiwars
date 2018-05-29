currencies:
- gold:
    use: 
	- buying items
    obtain_by:
	- opening chests
	- selling items
- energy:
    use:
    - fighting mobs
    - exploring dungeons
    obtain_by:
    - passive regen (based on avg team lvl)
    - resting (accelerated regen)
    - eating (using) food (instant regen)

items:
- chest (lvl)
    description: contains gold, other items
    actions: 
	- open_chest
	- sell_item
	obtain_by:
	- fighting mobs
	- exploring dungeons
- equipment (lvl)
    description: can be equiped for stat enhancement
    actions:
	- equip
	- sell_item
	- forge_into
	obtain_by:
	- fighting mobs
	- opening chests
- food (energy_value)
    description: can be consumed for energy
    actions:
	- use_item
	- sell_item
- scrolls (stat, amount, duration)
    description: can be consumed for stat buff
    actions: 
	- use_item
	- sell_item

actions: 
- explore_zone (CRUD of EXPLORATION entity)
    description: finds mobs, or dungeons, based on party avg lvl
- fight_mob (CRUD of FIGHT entity)
    description: finds items & gains experience, if successful
- use_item (CRUD on ITEM entity)
    description: uses an item (see item descriptions)
- sell_item (CRUD on ITEM entity)
    description: sells an item for gold (see item descriptions)
- forge_into (CRUD on EQUIPMENT entity)
	description: crushes one equipment into another, adding exp to the target
- rest (CRUD on PARTY entity)
    description: accellerates energy regen
- set_party (CRUD on PARTY entity)
    description: defines which characters make up the active party
- equip_to_char (CRUD on CHARACTER entity)
    description: defines which equipment is active on a character

party -- array of 5 characters

DATA DEFINITION:

`user`:
	description: |
		the human player's construct with a corresponding userid & emailaddress.
        all currencies accumulated and actions performed are specific to a user.
	attributes:
	-	id
	-	emailaddress
	-	gold
	-	energy

`inventory`:
	description: |
		the human player's list of acquired equipment, characters, items
	attributes:
	-	user_id
	-	type [equipment, character, item]
	-	type_id

`character`:
	description: |
		the virtual personas the player has recruited to be a part of the api-slaughtering
        team.  The player can arrange characters into a team and send the team into battle.
		Characters gain experience for successful Fights.  Once a character reaches it's
		experience_max, it gains a level, and its base stats increase slightly depending on
		character type.
	attributes:
	- 	id
	-	name (generated)
	-	class [Support, Mage, Warrior, Hunter, Rogue]
	-	level [1-999999]
	-	tier [bronze, silver, gold, diamond] (how do you raise in tier?)
	-	experience
	-	experience_max
	-	atk_base
	-	crit_base
	-	def_base
	-	hitpoints_base
	-	armor_base

`party`:
	description: |
		a collection of 5 characters that make up your active team.  each user
		only has one party (current limitation), and it is the default party for
		the user for all party-based actions.
	attribues:
	-	user_id
	-	char1_id
	-	char2_id
	-	char3_id
	-	char4_id
	-	char5_id
	-	status [active, resting]

`equipment`:
	description: |
		the weapons, armor, and artifacts your characters equip to prepare theem for API
		battle.  Equipment provide stat increases which are vitally important for progressing.
		Equipment gain experience by 
	attributes:
	-	id
	-	name (generated)
	-	slot [helm, primary_hand, off_hand, chest, legs, arms, boots, belt, trinket]
	-	primary stat modifier [atk, crit, def, hp, armor]
	-	primary stat value
	-	secondary stat modifier [atk, crit, def, hp, armor]
	-	secondary stat value
	-	level [1-999999]
	-	tier [bronze, silver, gold, diamond] (how do you raise in tier?)
	-	experience
	-	experience_max

`zone`: |
	description: |
		the virtual areas the player has unlocked to explore and loot.
	attributes:
	-	name
	-	type [normal, dungeon]
	-	level [1-999999]
	-	tier [bronze, silver, gold, diamond]
	-	explores_max
	-	explores_count

`enemy`: |
	description: |
		the adversaries the player finds when exploring zones.  
		the target of fights.
	attributes:
	-	name
	-	type [humanoid, animal, demon]
	-	level [1-999999]
	-	hitpoints
	-	atk
	-	crit
	-	def
	-	armor


/* (NOT MVP)
`quest`: |
	description: |
		time-based objectives for the player to complete to receive a reward.
		obtained by exploring zones, time starts when accepted, not obtained.
	attributes:
	-	id
	-	type [fights, explorations, items]
	-	level [1-999999]
	-	duration
	-	count_objective
	-	count
	-	accepted (bool)
	-	accepted_date
	-	end_date
	-	completed (bool)
	-	receive_type [item, equipment]
	-	receive_id
*/
	

`exploration`: |
	description: |
		the event and outcome from a player sending a party to a zone to explore.
	attributes:
	- id
	- zone_id
	- party_id
	- found_type [character, enemy, quest, zone]
	- found_id

`fight`: |
	description: |
		the event and outcome from a player sending a party to encounter an enemy.
	attributes:
	- id
	- enemy_id
	- party_id
	- result [win, lose]
	- found_type [item, equipment]
	- found_id
	- experience_gain


API DESIGN:
CREATE (POST)
READ   (GET)
UPDATE (PUT)
DELETE (DELETE)

/users         	- 	(GET)  lists the users you have created (only 1 by definition)
/users/#### 	- 	(GET)  lists the details of a user
		 			(DELETE) deletes the user
/parties/		-	(GET)	lists the parties you have created (only 1 by definition)
/parties/####	-	(GET)	lists the characters in the party, and party total stats
				-	(PUT)	update the characters in the party, or status (resting)
/characters                - (GET)  lists the characters you have available
/characters/####           - (GET)  lists status details about specific character
/characters/####/equipment - (GET)	lists the equipment on a specific character
                           - (PUT)  update the equipment on a specific character
/characters/####/items
/items/
/items/####
/equipment/
/equipment/####
/zones/
/zones/####
/enemies/
/enemies/####

pseudo actions:
/explorations/ - finds zones/enemys
/fights/ - battles enemys