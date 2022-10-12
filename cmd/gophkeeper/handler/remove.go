package handler

func (h *Handler) RemoveCredentialsByLogin(login string) error {
	return h.service.RemoveCredentialsByLogin(login)
}

func (h *Handler) RemoveTextByName(name string) error {
	return h.service.RemoveTextByName(name)
}

func (h *Handler) RemoveBinaryDataByName(name string) error {
	return h.service.RemoveBinaryDataByName(name)
}

func (h *Handler) RemoveBankCardByNumber(number string) error {
	return h.service.RemoveBankCardByNumber(number)
}
