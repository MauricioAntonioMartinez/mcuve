

type Udemy struct {
	Domain string
}


func (u *Udemy) GetBaseUrl() string {

}

func (y *Udemy) KeywordPath(query string) string {
	return "/results?search_query="+strings.Replace(query," ","+",-1)
}