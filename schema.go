package nitwit

type Messages struct {
	Response Response  `json:"response"`
	Symbol   Symbol    `json:"symbol"`
	Cursor   Cursor    `json:"cursor"`
	Messages []Message `json:"messages"`
}
type Response struct {
	Status int `json:"status"`
}
type Symbol struct {
	ID     int    `json:"id"`
	Symbol string `json:"symbol"`
	Title  string `json:"title"`
}
type Cursor struct {
	More  bool `json:"more"`
	Since int  `json:"since"`
	Max   int  `json:"max"`
}
type User struct {
	ID             int      `json:"id"`
	Username       string   `json:"username"`
	Name           string   `json:"name"`
	AvatarURL      string   `json:"avatar_url"`
	AvatarURLSsl   string   `json:"avatar_url_ssl"`
	Identity       string   `json:"identity"`
	Classification []string `json:"classification"`
}
type Source struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}
type Symbols struct {
	ID     int    `json:"id"`
	Symbol string `json:"symbol"`
	Title  string `json:"title"`
}
type Message struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	CreatedAt string    `json:"created_at"`
	User      User      `json:"user"`
	Source    Source    `json:"source"`
	Symbols   []Symbols `json:"symbols"`
}
