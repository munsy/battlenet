package battlenet

func init() {
	// Set region global to "us" if it's not already set.
	if "" == region {
		region = "us"
	}
}
