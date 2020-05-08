package hrvst

// Global package flags

// PREFIX is a string all flags should start with:
// all this flags are skipped on converting Arguments into URL values
const PREFIX = "flag_"

// GET_ALL defines argument for `getResourceList` responses (paginated or flat)
const GET_ALL = PREFIX + "get_all"
