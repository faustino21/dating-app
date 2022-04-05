package httpRequest

type MemberIdReq struct {
	MemberId string `json:"member_id" binding:"required"`
}
