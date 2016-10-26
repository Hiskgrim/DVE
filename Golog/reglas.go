package main

import (
  "fmt"
  . "golog"
)

func main(){

m := NewMachine().Consult(`

  tipo_profesor(asistente).
  tipo_profesor(auxiliar).
  tipo_profesor(asociado).
  tipo_profesor(titular).


  experiencia_minima(asistente,0).
  experiencia_minima(auxiliar,2).
  experiencia_minima(asociado,6).
  experiencia_minima(titular,10).

  publicaciones_minimas(asistente,0).
  publicaciones_minimas(auxiliar,0).
  publicaciones_minimas(asociado,1).
  publicaciones_minimas(titular,1).

  experiencia_valida(X,Y) :- experiencia_minima(X,Z), Z@=<Y.
  publicaciones_validas(X,Y) :- publicaciones_minimas(X,Z), Z@=<Y.

  tiempo(completo,5,8).
  tiempo(medio,5,4).

  minimo_convocatoria(1).

  convocatoria(X) :-  minimo_convocatoria(Y), Y@<X.

  tiempo_maximo_titulos(especializacion, 1).
  tiempo_maximo_titulos(maestria, 2).
  tiempo_maximo_titulos(doctorado, 5).

  tiempo_maximo_superado(X,Y) :- tiempo_maximo_titulos(X,Z), Z@<Y.

  numero_titulos_minimo(1).

  numero_titulos_valido(X) :- numero_titulos_minimo(Y), Y@=<X.

`)

}
