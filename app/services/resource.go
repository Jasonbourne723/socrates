package services

type ResourceService struct {
}

type IResourceSerivce interface {
}

func NewResourceService() IResourceSerivce {
	return &ResourceService{}
}
