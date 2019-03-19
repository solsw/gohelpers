package dotnethelper

import (
	"os"
	"runtime"

	"github.com/solsw/gohelpers/oshelper"
)

// DotNETInstalled checks if .NET Core is installed under Linux.
func DotNETInstalled() bool {
	switch runtime.GOOS {
	default:
		return false
	case "linux":
		res, _ := oshelper.DirExists("/usr/share/dotnet/shared/Microsoft.NETCore.App")
		if res {
			return true
		}
		res, _ = oshelper.DirExists(os.ExpandEnv("$HOME/.dotnet/shared/Microsoft.NETCore.App"))
		return res
	}
}
