package core

type Config struct {
	//上传或者下载 POST / GET
	Option string
	//用户认证信息，例如账户为usr 密码为pwd usr:pwd
	Auth       string
	RemoteUrl  string
	Process    int
	RemoteDir  string
	LocalDir   string
	Repository string
	Usr        string
	Pwd        string
}
type Body struct {
	Items             []Asset
	ContinuationToken string
}
type Component struct {
	Id         string
	Repository string
	Format     string
	Group      string
	Name       string
	Version    string
	Assets     *[]Asset
}

type Asset struct {
	DownloadUrl string
	Path        string
	Id          string
}
