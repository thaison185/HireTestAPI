package repositories

type ReportRepository struct {
	BaseRepository
}

func NewReportRepository(base BaseRepository) *ReportRepository {
	return &ReportRepository{BaseRepository: base}
}
