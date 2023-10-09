package requests

type RecordName struct {
	Records []Domain `json:"records"`
}

func ParseRecord(tp *TrustPositif) *Domain {
	var record *Domain

	for _, domain := range tp.Values {
		switch domain.Status {
		case "Ada":
			record = &Domain{
				Domain: domain.Domain,
				Status: Blocked,
			}
		case "Tidak Ada":
			record = &Domain{
				Domain: domain.Domain,
				Status: Unblocked,
			}
		}
	}

	return record
}
