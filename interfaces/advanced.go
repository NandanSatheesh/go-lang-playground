package main

type Advanced struct {
}

func (r *Advanced) implementThis(String string) {

}

func (r *Advanced) getAdvanced() BasicInterface {
	return &Advanced{}
}
