package templates

type Page struct {
	Top     PageTop
	Content any
	Bottom  PageBottom
}

type PageTop struct {
	Title string
}

type PageBottom struct{}
