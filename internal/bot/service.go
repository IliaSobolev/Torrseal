package bot

import "regexp"

// Validate magnet link
func validateMagnetLink(link string) bool {
	// idk work or not this patterns
	magnetPattern := []string{
		`^magnet:\?xt=urn:btih:[a-zA-Z0-9]{32,40}(&dn=.+)?(&tr=.+)*$`,
		`^magnet:\?xt=urn:(btih|sha1|sha256):[a-zA-Z0-9]{32,64}(&dn=.+)?(&tr=.+)*$`,
		`^magnet:\?(xt=urn:(btih|sha1|sha256|md5):[a-zA-Z0-9]{32,64}|dn=.+|kt=.+|tr=.+|xs=.+|as=.+|mt=.+)+$`,
	}
	for _, pattern := range magnetPattern {
		re := regexp.MustCompile(pattern)
		if re.MatchString(link) {
			return true
		}
	}
	return false
}
