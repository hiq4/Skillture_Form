package entities_test

import (
	"testing"
	"time"

	"Skillture_Form/internal/domain/entities"
	"Skillture_Form/internal/domain/enums"

	"github.com/google/uuid"
)

func TestResponse_IsValid(t *testing.T) {
	validFormID := uuid.New()

	tests := []struct {
		name     string
		response *entities.Response
		wantErr  bool
	}{
		{
			name: "valid response",
			response: &entities.Response{
				ID:     uuid.New(),
				FormID: validFormID,
				Respondent: map[string]any{
					"email": "user@test.com",
					"name":  "Test User",
				},
				Status:      enums.ResponseSubmitted,
				SubmittedAt: time.Now(),
			},
			wantErr: false,
		},
		{
			name: "missing FormID",
			response: &entities.Response{
				ID:     uuid.New(),
				FormID: uuid.Nil,
				Respondent: map[string]any{
					"email": "user@test.com",
				},
				Status: enums.ResponseSubmitted,
			},
			wantErr: true,
		},
		{
			name: "empty Respondent",
			response: &entities.Response{
				ID:         uuid.New(),
				FormID:     validFormID,
				Respondent: map[string]any{},
				Status:     enums.ResponseSubmitted,
			},
			wantErr: true,
		},
		{
			name: "invalid Status",
			response: &entities.Response{
				ID:     uuid.New(),
				FormID: validFormID,
				Respondent: map[string]any{
					"email": "user@test.com",
				},
				Status: enums.ResponseStatus(99),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.response.IsValid()
			if (err != nil) != tt.wantErr {
				t.Errorf("Response.IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
