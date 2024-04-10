package main

type Sav struct {
	ID            int    // The id of the savage and also the primary key.
	OwnerID       int    // The id of the owner of the savage.
	Points        int    // The points of the savage.
	Updated       bool   // If the savage needs to be updated.
	Location      int    // The location of the savage. Y * maxX + X
	FirstName     string // The first name of the savage.
	LastName      string // The last name of the savage.
	Generation    int    // The generation of the savage.
	Age           int    // The age of the savage.
	Sex           uint8  // 0 male / 1 female
	Pregnant      int8   // days till birth count down. -1 not pregnant.
	MotherID      int    // The id of the mother of the savage.
	FatherID      int    // The id of the father of the savage.
	Hunger        uint8  // The hunger of the savage.
	HungerMax     uint8  // The max hunger of the savage.
	Thirst        uint8  // The thirst of the savage.
	ThirstMax     uint8  // The max thirst of the savage.
	Health        uint8  // The health of the savage.
	HealthMax     uint8  // The max health of the savage.
	Strength      uint8  // The strength of the savage.
	Intelligence  uint8  // The intelligence of the savage.
	Charisma      uint8  // The charisma of the savage.
	Wisdom        uint8  // The wisdom of the savage.
	Dexterity     uint8  // The dexterity of the savage.
	Constitution  uint8  // The constitution of the savage
	Friendly      int8   // How friendly the savage is.
	Sociable      int8   // How sociable the savage is.
	Assertive     int8   // How assertive the savage is.
	Outgoing      int8   // How outgoing the savage is.
	Energetic     int8   // How energetic the savage is.
	Talkative     int8   // How talkative the savage is.
	Articulate    int8   // How articulate the savage is.
	Affectionate  int8   // How affectionate the savage is.
	FunLoving     int8   // How fun loving the savage is.
	ProneToAction int8   // How prone to action the savage is.
	Gregarious    int8   // How gregarious the savage is.
	Quiet         int8   // How quiet the savage is.
	Reserved      int8   // How reserved the savage is.
	Thoughtful    int8   // How thoughtful the savage is.
	Passive       int8   // How passive the savage is.
	Shy           int8   // How shy the savage is.
	Trusting      int8   // How trusting the savage is.
	Polite        int8   // How polite the savage is.
	Blunt         int8   // How blunt the savage is.
	Rude          int8   // How rude the savage is.
	Antagonistic  int8   // How antagonistic the savage is.
	Cruel         int8   // How cruel the savage is.
	Hostile       int8   // How hostile the savage is.
	Distrustful   int8   // How distrustful the savage is.
}

type savtosav struct {
	ID1           int
	ID2           int
	Friendly      int8 // How friendly the savage is.
	Sociable      int8 // How sociable the savage is.
	Assertive     int8 // How assertive the savage is.
	Outgoing      int8 // How outgoing the savage is.
	Energetic     int8 // How energetic the savage is.
	Talkative     int8 // How talkative the savage is.
	Articulate    int8 // How articulate the savage is.
	Affectionate  int8 // How affectionate the savage is.
	FunLoving     int8 // How fun loving the savage is.
	ProneToAction int8 // How prone to action the savage is.
	Gregarious    int8 // How gregarious the savage is.
	Quiet         int8 // How quiet the savage is.
	Reserved      int8 // How reserved the savage is.
	Thoughtful    int8 // How thoughtful the savage is.
	Passive       int8 // How passive the savage is.
	Shy           int8 // How shy the savage is.
	Trusting      int8 // How trusting the savage is.
	Polite        int8 // How polite the savage is.
	Blunt         int8 // How blunt the savage is.
	Rude          int8 // How rude the savage is.
	Antagonistic  int8 // How antagonistic the savage is.
	Cruel         int8 // How cruel the savage is.
	Hostile       int8 // How hostile the savage is.
	Distrustful   int8 // How distrustful the savage is.
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

type tlog struct {
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
