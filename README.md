# savages

Baswd on my memory of Little Savages 2 by Matt Squirrell.  
It has been a very long time since I played the 'game', but I have really fond memories of it.

## notes
- each savage is on a 1 meter square
  - if distance <= 10 meter there is a chance of an encounter. The closer the higher.
  - if sex is diff then there is a chance of mating if compatible. 
    - Charisma should be a part of this. 
- There is a random chance of finding water. There is no fixed location of water (yet).
- There is a random chance of finding food. 


need to work on pregnant  
will need a code cleanup first

Help would be very ***nice***!


- Children stay with the mother until the age of 10.  
  - There is a chance if the mother's owner is not the same as the child's owner that the mother will harm the child.
- Need personality modifiers.
  - Kindness / Aggression
  - Good / Bad hunter gatherer
  - yadda yadda 


Log db needs to be made to keep track of each action a savage does per day.

User db needs to be made.

Game state db needs ro be made.

Web interface for interacting with players savages.

main struct for the savages
	ID           int    // The id of the savage and also the primary key.  
	OwnerID      int    // The id of the owner of the savage.  
	Updated      bool   // If the savage needs to be updated.  
	Location     int    // The location of the savage. Y * maxX + X  
	FirstName    string // The first name of the savage.  
	LastName     string // The last name of the savage.  
	Age          int    // The age of the savage.  
	Sex          uint8  // 0 male / 1 female  
	Pregnant     int8   // days till birth count down. -1 not pregnant.  
	MotherID     int    // The id of the mother of the savage.  
	FatherID     int    // The id of the father of the savage.  
	Hunger       uint8  // The hunger of the savage.  
	HungerMax    uint8  // The max hunger of the savage.  
	Thirst       uint8  // The thirst of the savage.  
	ThirstMax    uint8  // The max thirst of the savage.  
	Health       uint8  // The health of the savage.  
	HealthMax    uint8  // The max health of the savage.  
	Strength     uint8  // The strength of the savage.  
	Intelligence uint8  // The intelligence of the savage.  
	Charisma     uint8  // The charisma of the savage.  
	Wisdom       uint8  // The wisdom of the savage.  
	Dexterity    uint8  // The dexterity of the savage.  
	Constitution uint8  // The constitution of the savage.  

  

