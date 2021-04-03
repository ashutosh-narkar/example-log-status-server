package authz

default allow = false

allow {
    m := input.message
    m == "world"
}
