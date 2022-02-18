package server

func (b *Backend) router() {
	b.app.Use(setBasicInformation())

	b.app.POST("/login", b.login)
	v1Api := b.app.Group("/api/v1", authToken())
	{
		v1Api.POST("/create_user", b.createUser)
	}
}
