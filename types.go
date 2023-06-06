package main

type Sav struct {
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
	Attribute1   uint8  // How horny.
	Attribute2   uint8  // How mean.
	Attribute3   uint8  // The attributes of the savage.
}

type distance struct {
	ID1      int
	ID2      int
	Distance int
}

type user struct {
	ID       int
	Username string
	Email    string
	Password string
}

type log struct {
	ID      int
	Date    int
	Who     int
	Message string
}

type gamedb struct {
	ID  int
	Day int
}

type birthrecord struct {
	ID       int
	Date     int
	ChildID  int
	MotherID int
	FatherID int
}
