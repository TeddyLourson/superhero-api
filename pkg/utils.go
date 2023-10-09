package utils

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func UUIDFromString(val string) uuid.UUID {
	return uuid.MustParse(val)
}
func StringToText(val string) pgtype.Text {
	return pgtype.Text{String: val, Valid: true}
}
func Int32ToInt4(val int32) pgtype.Int4 {
	return pgtype.Int4{Int32: val, Valid: true}
}
func UInt32ToInt4(val uint32) pgtype.Int4 {
	return pgtype.Int4{Int32: int32(val), Valid: true}
}
func TextToNullString(val pgtype.Text) *wrapperspb.StringValue {
	if !val.Valid {
		return nil
	}
	return wrapperspb.String(val.String)
}
func UInt32ToNullUInt32(val pgtype.Uint32) *wrapperspb.UInt32Value {
	if !val.Valid {
		return nil
	}
	return wrapperspb.UInt32(val.Uint32)
}
func Int4ToNullUInt32(val pgtype.Int4) *wrapperspb.UInt32Value {
	if !val.Valid {
		return nil
	}
	return wrapperspb.UInt32(uint32(val.Int32))
}
func TimestamptzToString(val pgtype.Timestamptz) *string {
	if !val.Valid {
		return nil
	}
	valStr := val.Time.String()
	return &valStr
}
