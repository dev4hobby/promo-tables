package services

func (h *Handler) JoinTable() bool {
	return h.db.JoinTable() != nil
}
