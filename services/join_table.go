package services

func (h *Handler) JoinTableAsQueryBuilder() bool {
	return h.db.JoinTableAsQueryBuilder() != nil
}
