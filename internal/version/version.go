package version

import "fmt"

var Version = "dev"

func Banner() string {
	return fmt.Sprintf(`   ______ ____   __  ___ ______ ______
  / ____// __ \ /  |/  // ____//_  __/
 / /    / / / // /|_/ // __/    / /
/ /___ / /_/ // /  / // /___   / /
\____/ \____//_/  /_//_____/  /_/

COMET version %s`, Version)
}
