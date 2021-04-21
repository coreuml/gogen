package createjournal

import (
	"context"
)

// Inport of CreateJournal
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase CreateJournal
type InportRequest struct {
	BusinessID      string                  `` //
	Date            string                  `` //
	Description     string                  `` //
	JournalType     string                  `` //
	UserID          string                  `` //
	JournalBalances []JournalBalanceRequest `` //
}

type JournalBalanceRequest struct {
	Side        string  `` //
	AccountCode string  `` //
	Nominal     float64 `` //
}

// InportResponse is response payload after running the usecase CreateJournal
type InportResponse struct {
}
