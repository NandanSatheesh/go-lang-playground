package main

type RedirectRepository interface {
	Find(Code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
