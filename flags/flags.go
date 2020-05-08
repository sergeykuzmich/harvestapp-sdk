package flags

// Global package flags

// prefix is a string all flags should start with:
// all this flags are skipped on converting Arguments into URL values
const Prefix = "flag_"

// GetAll defines argument for `getResourceList` responses (paginated or flat)
const GetAll = Prefix + "get_all"
