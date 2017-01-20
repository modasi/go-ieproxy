package ieproxy

// StaticProxyConf containes the Windows configuration for static proxy
type StaticProxyConf struct {
	// Is the proxy active?
	Active bool
	// Proxy address for each scheme (http, https)
	// "" (empty string) is the fallback proxy
	Protocols map[string]string
	// Addresses not to be browsed via the proxy (comma-separated, like linux)
	NoProxy string
}

// AutomaticProxyConf contains the Windows configuration for automatic proxy
type AutomaticProxyConf struct {
	URL string // url of the .pac file
}

// WindowsProxyConf gathers the Windows configuration for proxy
type WindowsProxyConf struct {
	Static    StaticProxyConf    // static configuration
	Automatic AutomaticProxyConf // automatic configuration
}

// GetConf retrieves the proxy configuration from the Windows Regedit
func GetConf() WindowsProxyConf {
	return WindowsProxyConf{}
}

// OverrideEnvWithStaticProxy writes new values to the
// `http_proxy`, `https_proxy` and `no_proxy` environment variables.
// The values are taken from the Windows Regedit (should be called in `init()` function - see example)
func OverrideEnvWithStaticProxy() {
}
