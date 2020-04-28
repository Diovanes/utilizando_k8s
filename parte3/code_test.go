package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := greeting("code")
	if resultado != "<b>code</b>" {
		t.Errorf("Resultado esperado: <b>code</b>, obtido: %s", resultado)
	}
}
