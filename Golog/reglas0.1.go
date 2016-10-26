package main

import (
  "fmt"
  "bufio"
  "os"
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

  minimo_convocatoria(2).

  convocatoria(X) :-  minimo_convocatoria(Y), Y@<X.

  tiempo_maximo_titulos(especializacion, 1).
  tiempo_maximo_titulos(maestria, 2).
  tiempo_maximo_titulos(doctorado, 5).

  tiempo_maximo_superado(X,Y) :- tiempo_maximo_titulos(X,Z), Z@<Y.

  numero_titulos_minimo(1).

  numero_titulos_valido(X) :- numero_titulos_minimo(Y), Y@=<X.

`)

reader := bufio.NewReader(os.Stdin)
fmt.Println("Escriba la experiencia en años: ")
experiencia,err := reader.ReadString('\n')

if(err == nil){
  fmt.Printf("\nEXPERIENCIA = "+experiencia+"\n")
  if m.CanProve(`experiencia_valida(asistente,`+experiencia+`).`) {
    fmt.Printf("Puede ser tipo asistente\n")
  }else{
    fmt.Printf("No puede ser tipo asistente\n")
  }
  if m.CanProve(`experiencia_valida(auxiliar,`+experiencia+`).`) {
    fmt.Printf("Puede ser tipo auxiliar\n")
  }else{
    fmt.Printf("No puede ser tipo auxiliar\n")
  }
  if m.CanProve(`experiencia_valida(asociado,`+experiencia+`).`) {
    fmt.Printf("Puede ser tipo asociado\n")
  }else{
    fmt.Printf("No puede ser tipo asociado\n")
  }
  if m.CanProve(`experiencia_valida(titular,`+experiencia+`).`) {
    fmt.Printf("Puede ser tipo titular\n")
  }else{
    fmt.Printf("No puede ser tipo titular\n")
  }
}

fmt.Println("\nEscriba el numero de publicaciones: ")
publicaciones,err := reader.ReadString('\n')

if(err == nil){
  fmt.Printf("\n# PUBLICACIONES = "+publicaciones+"\n")
  if m.CanProve(`publicaciones_validas(asistente,`+publicaciones+`).`) {
    fmt.Printf("Puede ser tipo asistente\n")
  }else{
    fmt.Printf("No puede ser tipo asistente\n")
  }
  if m.CanProve(`publicaciones_validas(auxiliar,`+publicaciones+`).`) {
    fmt.Printf("Puede ser tipo auxiliar\n")
  }else{
    fmt.Printf("No puede ser tipo auxiliar\n")
  }
  if m.CanProve(`publicaciones_validas(asociado,`+publicaciones+`).`) {
    fmt.Printf("Puede ser tipo asociado\n")
  }else{
    fmt.Printf("No puede ser tipo asociado\n")
  }
  if m.CanProve(`publicaciones_validas(titular,`+publicaciones+`).`) {
    fmt.Printf("Puede ser tipo titular\n")
  }else{
    fmt.Printf("No puede ser tipo titular\n")
  }
}

fmt.Println("\nEscriba el numero de presentados a convocatoria: ")
presentados,err := reader.ReadString('\n')

if(err == nil){
  fmt.Printf("\nPRESENTADOS A CONVOCATORIA = "+presentados+"\n")
  if m.CanProve(`convocatoria(`+presentados+`).`) {
    fmt.Printf("No se repite convocatoria\n")
  }else{
    fmt.Printf("Se repite convocatoria\n")
  }
}

fmt.Println("\nEscriba el tiempo de espera para presentacion de titulos (años): ")
espera,err := reader.ReadString('\n')

if(err == nil){
  fmt.Printf("\nTIEMPO DE ESPERA (PRESENTACION TITULOS) = "+espera+"\n")
  if m.CanProve(`tiempo_maximo_superado(especializacion,`+espera+`).`) {
    fmt.Printf("Tiempo superado para especializacion\n")
  }else{
    fmt.Printf("Tiempo no superado para especializacion\n")
  }
  if m.CanProve(`tiempo_maximo_superado(maestria,`+espera+`).`) {
    fmt.Printf("Tiempo superado para maestria\n")
  }else{
    fmt.Printf("Tiempo no superado para maestria\n")
  }
  if m.CanProve(`tiempo_maximo_superado(doctorado,`+espera+`).`) {
    fmt.Printf("Tiempo superado para doctorado\n")
  }else{
    fmt.Printf("Tiempo no superado para doctorado\n")
  }
}

fmt.Println("\nEscriba el numero de titulos: ")
titulos,err := reader.ReadString('\n')

if(err == nil){
  fmt.Printf("\nNUMERO DE TITULOS = "+titulos+"\n")
  if m.CanProve(`numero_titulos_valido(`+titulos+`).`) {
    fmt.Printf("Numero de titulos valido\n")
  }else{
    fmt.Printf("Numero de titulos no valido\n")
  }
}

}
