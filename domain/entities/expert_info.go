package entities

type ExpertInfo struct {
	ID int 							`json:"id" db:"id"`
	AuthorityNameCharacter string 	`json:"authority_name_character" db:"authority_name_character"`
	Date int 						`json:"date" db:"date"`
	Conclusion string				`json:"conclusion" db:"conclusion"`
}