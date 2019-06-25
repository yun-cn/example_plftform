package actions_test

func (as *ActionSuite) Test_HomeHandler() {
	as.LoadFixture("current user")

	req := as.JSON("/")
	loginCurrentUser("user@example.co.jp", req)

	res := req.Get()

	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}
