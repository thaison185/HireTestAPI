package repositories

type CampaignRepository struct {
	BaseRepository
}

func NewCampaignRepository(base BaseRepository) *CampaignRepository {
	return &CampaignRepository{BaseRepository: base}
}
