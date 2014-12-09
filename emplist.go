func main() {
	pemKeyBytes, _ := ioutil.ReadFile("admin.pem")
	iss := "955697674329329-70smumqrlcnvv4i7oi4msji9u2emtmi2@developer.gserviceaccount.com"
	scope := "https://www.googleapis.com/auth/admin.directory.user"
	t := jwt.NewToken(iss, scope, pemKeyBytes)
	t.ClaimSet.Prn = "dherbst@dramafever.com"   // this is who you are impersonating

	transport, _ := jwt.NewTransport(t)
	httpClient := transport.Client()
	svc, err := admin.New(httpClient)
	usersvc := admin.NewUsersService(svc)
	fmt.Printf("usersvc=%v\n", usersvc)

	listcall := usersvc.List()
	listcall.Domain("dramafever.com")

	users, err := listcall.Do()

	numUsers := len(users.Users)
	fmt.Printf("len(users)=%v err=%v\n", numUsers, err)
