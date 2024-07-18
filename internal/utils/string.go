package utils

// GetStrValOr get s value if exist else def value else empty string
func GetStrValOr(s *string, def *string) string {
	if s != nil {
		return *s
	}
	if def != nil {
		return *def
	}
	return ""
}
