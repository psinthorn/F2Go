package service

func GetAbout(id int64) (*domain.About, *utils.ApplicationError) {
	return *domain.GetAbout(id)
}
