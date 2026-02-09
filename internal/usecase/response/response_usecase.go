package response

import (
	"context"
	"errors"
	"time"

	"Skillture_Form/internal/domain/entities"
	"Skillture_Form/internal/domain/enums"
	repo "Skillture_Form/internal/repository/interfaces"

	"github.com/google/uuid"
)

// ResponseUsecase handles all business logic related to form responses
// It orchestrates validation, domain rules, and persistence
type ResponseUsecase struct {
	formRepo      repo.FormRepository
	formFieldRepo repo.FormFieldRepository
	responseRepo  repo.ResponseRepository
	answerRepo    repo.ResponseAnswerRepository
	vectorRepo    repo.ResponseAnswerVectorRepository
}

// NewResponseUsecase creates a new instance of ResponseUsecase
// All required repositories are injected via dependency injection
func NewResponseUsecase(formRepo repo.FormRepository, formFieldRepo repo.FormFieldRepository, responseRepo repo.ResponseRepository, answerRepo repo.ResponseAnswerRepository, vectorRepo repo.ResponseAnswerVectorRepository) *ResponseUsecase {
	return &ResponseUsecase{
		formRepo:      formRepo,
		formFieldRepo: formFieldRepo,
		responseRepo:  responseRepo,
		answerRepo:    answerRepo,
		vectorRepo:    vectorRepo,
	}
}

// Submit handles the full lifecycle of submitting a form response:
// - validates the response
// - validates the form state
// - persists the response
// - persists answers
// - persists optional vectors (bulk insert)
func (u *ResponseUsecase) Submit(
	ctx context.Context,
	response *entities.Response,
	answers []*entities.ResponseAnswer,
	vectors []*entities.ResponseAnswerVector,
) error {

	// ===== Domain validation =====
	if err := response.IsValid(); err != nil {
		return err
	}

	form, err := u.formRepo.GetByID(ctx, response.FormID)
	if err != nil {
		return err
	}

	fields, err := u.formFieldRepo.List(ctx, repo.FormFieldFilter{
		FormID: &response.FormID,
	})
	if err != nil {
		return err
	}

	if err := validateSubmit(form, fields, response, answers); err != nil {
		return err
	}

	// ===== Prepare data =====
	if response.ID == uuid.Nil {
		response.ID = uuid.New()
	}
	response.Status = enums.ResponseSubmitted
	response.SubmittedAt = time.Now()

	for _, a := range answers {
		if a.ID == uuid.Nil {
			a.ID = uuid.New()
		}
		a.ResponseID = response.ID
		a.CreatedAt = time.Now()

		if err := a.IsValid(); err != nil {
			return err
		}
	}

	for _, v := range vectors {
		if v.ID == uuid.Nil {
			v.ID = uuid.New()
		}
		v.CreatedAt = time.Now()

		if err := v.IsValid(); err != nil {
			return err
		}
	}

	// ===== Transaction =====
	return u.responseRepo.WithTx(ctx, func(tx repo.ResponseRepository) error {

		if err := tx.Create(ctx, response); err != nil {
			return err
		}

		for _, a := range answers {
			if err := u.answerRepo.WithTxRepo(tx).Create(ctx, a); err != nil {
				return err
			}
		}

		if len(vectors) > 0 {
			if err := u.vectorRepo.WithTxRepo(tx).CreateBulk(ctx, vectors); err != nil {
				return err
			}
		}

		return nil
	})
}

// GetByID retrieves a single response by its ID
func (u *ResponseUsecase) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*entities.Response, error) {

	if id == uuid.Nil {
		return nil, errors.New("response id is required")
	}

	return u.responseRepo.GetByID(ctx, id)
}

// ListByForm returns all responses associated with a specific form
func (u *ResponseUsecase) ListByForm(
	ctx context.Context,
	formID uuid.UUID,
) ([]*entities.Response, error) {

	if formID == uuid.Nil {
		return nil, errors.New("form id is required")
	}

	return u.responseRepo.ListByFormID(ctx, formID)
}

// Delete removes a response by its ID
// It first checks existence before deletion
func (u *ResponseUsecase) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {

	if id == uuid.Nil {
		return errors.New("response id is required")
	}

	// Ensure response exists before deleting
	_, err := u.responseRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return u.responseRepo.Delete(ctx, id)
}
