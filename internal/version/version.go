package version

import "fmt"

const Version = "0.1.0"

func Banner() string {
	return fmt.Sprintf(`   ______ ____   __  ___ ______ ______
  / ____// __ \ /  |/  // ____//_  __/
 / /    / / / // /|_/ // __/    / /
/ /___ / /_/ // /  / // /___   / /
\____/ \____//_/  /_//_____/  /_/

COMET version %s`, Version)
}
