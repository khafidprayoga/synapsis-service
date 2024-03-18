package commonHelper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	synapsisv1 "github.com/khafidprayoga/synapsis-service/gen/synapsis/v1"
	"google.golang.org/grpc/metadata"
	"strings"
	"time"
)

type InterServiceDispatcher interface {
	GenerateOutContext(origin context.Context) (context.Context, error)
	GenerateOutContextWithTimeout(origin context.Context, timeout time.Duration, sinceRequest bool) (context.Context, context.CancelFunc, error)
	GenerateOutContextWithDeadline(origin context.Context, endAt *time.Time) (context.Context, context.CancelFunc, error)
}

type RequestContextCommonField struct {
	RequestId       *uuid.UUID
	RequestAt       *time.Time
	TokenClaimsData *synapsisv1.Claims
	MetaData        metadata.MD
}

// GenerateOutContextWithTimeout is a function to propagate new outgoing context with timeout
//
// if the sinceRequest is true, the timeout will be calculated from the requestAt factory field
// if the sinceRequest is false, the timeout will be calculated from now (when call this func)
// if the timeout is 0, the error will be raised for atomicity
func (r RequestContextCommonField) GenerateOutContextWithTimeout(
	origin context.Context, timeout time.Duration, sinceRequest bool,
) (context.Context, context.CancelFunc, error) {

	if timeout == 0 {
		return nil, nil, errors.New("timeout is required and must be greater than 0 nano second")
	}

	outCtx, errBind := r.GenerateOutContext(origin)
	if errBind != nil {
		return nil, nil, errBind
	}

	if sinceRequest {
		if r.RequestAt == nil {
			return nil, nil, errors.New("requestAt is required, please set the requestAt first when since request is true")
		}

		outWithTimeoutCtx, cancelFn := context.WithDeadline(outCtx, r.RequestAt.Add(timeout))
		return outWithTimeoutCtx, cancelFn, nil
	}

	outWithTimeoutCtx, cancelFn := context.WithTimeout(outCtx, timeout)
	return outWithTimeoutCtx, cancelFn, nil
}

// GenerateOutContextWithDeadline is a function to propagate new outgoing context with deadline
//
// if the endAt is nil or zero, the error will be raised for atomicity
func (r RequestContextCommonField) GenerateOutContextWithDeadline(
	origin context.Context, endAt *time.Time,
) (context.Context, context.CancelFunc, error) {
	if endAt == nil || endAt.IsZero() {
		return nil, nil, errors.New("endAt is required")
	}

	outCtx, errBind := r.GenerateOutContext(origin)
	if errBind != nil {
		return nil, nil, errBind
	}

	now := time.Now()
	if endAt.Before(now) {
		return nil, nil, fmt.Errorf("endAt must be greater than now, or after: %v", now)
	}

	wrappedCtx, cancelFn := context.WithDeadline(outCtx, *endAt)
	return wrappedCtx, cancelFn, nil
}

// GenerateOutContext is a function to extract rpc metadata to new outgoing context
func (r RequestContextCommonField) GenerateOutContext(origin context.Context) (context.Context, error) {
	if r.MetaData == nil {
		return nil, errors.New("metadata is required, please set the metadata first")
	}

	if r.RequestAt == nil {
		return nil, errors.New("requestAt is required, please set the requestAt first")
	}

	if r.MetaData.Len() == 0 {
		return nil, errors.New("metadata is empty, please set the metadata first (WARNING)")
	}

	if r.TokenClaimsData == nil {
		return nil, errors.New("token is empty (WARNING)")
	}
	outCtx := metadata.NewOutgoingContext(origin, r.MetaData)

	return outCtx, nil
}

func GetRequestCommonField(
	ctx context.Context,
) (reqCtx *RequestContextCommonField, err error) {

	var (
		requestCtxData = &RequestContextCommonField{}
		requestIdKey   = "requestid"
		requestAtKey   = "requestat"
	)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata from context")
	}

	// Set Extracted Metadata
	requestCtxData.MetaData = md
	// Check requestId
	{
		if md.Get(requestIdKey) != nil {
			id := md.Get(requestIdKey)[0]
			generatedUUID, errGenUUID := uuid.Parse(id)
			if errGenUUID != nil {
				return nil, errGenUUID
			}
			requestCtxData.RequestId = &generatedUUID

		} else {
			id := uuid.New()
			requestCtxData.RequestId = &id
		}
	}

	// Check requestAt
	{
		if md.Get(requestAtKey) != nil {
			parsedReqAt, errGetReqAt := time.Parse(
				time.RFC3339,
				md.Get(requestAtKey)[0],
			)
			if errGetReqAt != nil {
				return nil, errGetReqAt
			}
			requestCtxData.RequestAt = &parsedReqAt
		} else {
			at := time.Now().UTC()
			requestCtxData.RequestAt = &at
		}
	}

	// Get Token
	{
		tokenExtract := md.Get("claims")
		if len(tokenExtract) == 0 {
			return requestCtxData, errors.New("missing token from context")
		}

		claimsValues := []byte(
			strings.Replace(tokenExtract[0], "&", "", -1),
		)

		errUn := json.Unmarshal(claimsValues, &requestCtxData.TokenClaimsData)
		if errUn != nil {
			// also return the metadata
			// bcz is possible to return the only metadata or common field
			// without the token claims
			return requestCtxData, errUn
		}
	}

	return requestCtxData, nil
}
