package utils

import "github.com/jackc/pgx/v5/pgtype"

func ToText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{}
	}
	return pgtype.Text{
		String: *s,
		Valid:  true,
	}
}

func ToInt2(v *int16) pgtype.Int2 {
	if v == nil {
		return pgtype.Int2{}
	}
	return pgtype.Int2{
		Int16: *v,
		Valid: true,
	}
}

func ToBool(v *bool) pgtype.Bool {
	if v == nil {
		return pgtype.Bool{}
	}
	return pgtype.Bool{
		Bool:  *v,
		Valid: true,
	}
}
