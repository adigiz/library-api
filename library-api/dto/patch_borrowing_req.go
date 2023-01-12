package dto

type PatchBorrowingReq struct {
	Status string `binding:"required,eq=RETURNED"`
}
