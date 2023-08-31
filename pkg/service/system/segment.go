package segment

type Segment struct {
	ID      int     `json:"id"`
	Name    string  `json:"name" binding:"required"`
	Percent float64 `json:"percent"`
}

type UserSegment struct {
	ID          int
	NameSegment []string `json:"name" binding:"required"`
	UserID      int      `json:"userid" binding:"required"`
}

type TableSlugs struct {
	ID int
}

type ListNames struct {
	Name string `json:"name"`
}
