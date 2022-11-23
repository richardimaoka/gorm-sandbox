//go:build tools
// +build tools

package tools

import (
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/gorm"
)
