package segment

type Segment struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

type UserSegment struct {
	ID          int
	NameSegment []string `json:"name" binding:"required"`
	UserID      int      `json:"userid" binding:"required"`
}

type ListNames struct {
	Name string `json:"name"`
}
