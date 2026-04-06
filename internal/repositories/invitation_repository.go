package repositories

type InvitationRepository struct {
	BaseRepository
}

func NewInvitationRepository(base BaseRepository) *InvitationRepository {
	return &InvitationRepository{BaseRepository: base}
}
