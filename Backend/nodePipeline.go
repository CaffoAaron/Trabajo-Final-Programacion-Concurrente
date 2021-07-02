package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type knnNode struct {
	Distancia float64
	x         int
	y         int
	estado    string
}

type Respuesta struct {
	Mensaje string
}

type ConsultaBono struct {
	Casado                         bool `json:"casado"`
	Hijos                          bool `json:"hijos"`
	CarreraUniversitaria           bool `json:"carrera_universitaria"`
	CasaPropia                     bool `json:"casa_propia"`
	OtroPrestamo                   bool `json:"otro_prestamo"`
	Mas_4_Años                     bool `json:"mas_de_4_Años_como_empresa"`
	Mas_1_Local                    bool `json:"mas_de_1_Local"`
	Mas_10_Empleados               bool `json:"mas_de_10_Empleados"`
	PagoIgv_6_Meses                bool `json:"Pago_de_Igv_Ultimos_6_Meses"`
	DeclaronConfidencialPatrimonio bool `json:"declaron_confidencial_patrimonio"`

	PuntajePersonal int
	PuntajeEmpresa  int
	Estado          string
}

//variables globales
var Dataset = [1000]ConsultaBono{}
var eschucha_funcion bool
var remotehost string
var chCont chan int
var n, min, valorUsuario int

func getEstado(p *ConsultaBono) {
	contPersonas := 0
	contEmpresa := 0

	if p.Casado == true {
		contPersonas += 3
	}
	if p.Hijos == false {
		contPersonas += 1
	}
	if p.CarreraUniversitaria == true {
		contPersonas += 3
	}
	if p.CasaPropia == true {
		contPersonas += 4
	}
	if p.OtroPrestamo == false {
		contPersonas += 2
	}
	if p.Mas_4_Años == true {
		contEmpresa += 2
	}
	if p.Mas_1_Local == true {
		contEmpresa += 4
	}
	if p.Mas_10_Empleados == true {
		contEmpresa += 4

	}
	if p.PagoIgv_6_Meses == true {
		contEmpresa += 1
	}
	if p.DeclaronConfidencialPatrimonio == true {
		contEmpresa += 1
	}

	p.PuntajeEmpresa = contPersonas
	p.PuntajePersonal = contEmpresa

	if p.PuntajeEmpresa+p.PuntajePersonal > 15 {
		p.Estado = "Pre-Aprobado"
	}
	if p.PuntajeEmpresa+p.PuntajePersonal <= 15 {
		p.Estado = "Denegado"
	}
}

func enviar(num int) { //enviar el numero mayor al host remoto
	conn, _ := net.Dial("tcp", remotehost)
	defer conn.Close()
	//envio el número
	fmt.Fprintf(conn, "%d\n", num)

}

func enviar_Principal(num int) { //enviar el numero mayor al host remoto
	conn, _ := net.Dial("tcp", "localhost:8000")
	defer conn.Close()
	//envio el número
	fmt.Fprintf(conn, "%d\n", num)

}

func manejador_respueta(conn net.Conn) bool {
	defer conn.Close()
	eschucha_funcion = false
	bufferIn := bufio.NewReader(conn)
	numStr, _ := bufferIn.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	numero, _ := strconv.Atoi(numStr)
	strNumero := strconv.Itoa(numero)
	if strNumero[1] == 49 {
		return true
	} else {
		return false
	}
}

func manejador_fin(conn net.Conn) {
	defer conn.Close()
	var usuario = ConsultaBono{}
	var respuesta = Respuesta{}
	//recuperar el número
	bufferIn := bufio.NewReader(conn)
	numStr, _ := bufferIn.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	numero, _ := strconv.Atoi(numStr)

	strNumero := strconv.Itoa(numero)
	///49 ==== 1
	///48 ==== 0
	if strNumero[1] == 49 {
		usuario.Casado = true
	} else {
		usuario.Casado = false
	}
	if strNumero[2] == 49 {
		usuario.Hijos = true
	} else {
		usuario.Hijos = false
	}
	if strNumero[3] == 49 {
		usuario.CarreraUniversitaria = true
	} else {
		usuario.CarreraUniversitaria = false
	}
	if strNumero[4] == 49 {
		usuario.CasaPropia = true
	} else {
		usuario.CasaPropia = false
	}
	if strNumero[5] == 49 {
		usuario.OtroPrestamo = true
	} else {
		usuario.OtroPrestamo = false
	}
	if strNumero[6] == 49 {
		usuario.Mas_4_Años = true
	} else {
		usuario.Mas_4_Años = false
	}
	if strNumero[7] == 49 {
		usuario.Mas_1_Local = true
	} else {
		usuario.Mas_1_Local = false
	}
	if strNumero[8] == 49 {
		usuario.Mas_10_Empleados = true
	} else {
		usuario.Mas_10_Empleados = false
	}
	if strNumero[9] == 49 {
		usuario.PagoIgv_6_Meses = true
	} else {
		usuario.PagoIgv_6_Meses = false
	}
	if strNumero[10] == 49 {
		usuario.DeclaronConfidencialPatrimonio = true
	} else {
		usuario.DeclaronConfidencialPatrimonio = false
	}

	getEstado(&usuario)
	RespuestaKnn := knn(&usuario)
	if RespuestaKnn == true {
		respuesta.Mensaje = "Usted esta preaprobado para el bono independiente"
		enviar_Principal(11)
	} else {
		respuesta.Mensaje = "Usted no esta apto para el bono independiente"
		enviar_Principal(10)
	}
}

func manejador(conn net.Conn) {
	defer conn.Close()
	var usuario = ConsultaBono{}
	var respuesta = Respuesta{}
	//recuperar el número
	bufferIn := bufio.NewReader(conn)
	numStr, _ := bufferIn.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	numero, _ := strconv.Atoi(numStr)
	valorUsuario = numero
	strNumero := strconv.Itoa(numero)
	///49 ==== 1
	///48 ==== 0
	if strNumero[1] == 49 {
		usuario.Casado = true
	} else {
		usuario.Casado = false
	}
	if strNumero[2] == 49 {
		usuario.Hijos = true
	} else {
		usuario.Hijos = false
	}
	if strNumero[3] == 49 {
		usuario.CarreraUniversitaria = true
	} else {
		usuario.CarreraUniversitaria = false
	}
	if strNumero[4] == 49 {
		usuario.CasaPropia = true
	} else {
		usuario.CasaPropia = false
	}
	if strNumero[5] == 49 {
		usuario.OtroPrestamo = true
	} else {
		usuario.OtroPrestamo = false
	}
	if strNumero[6] == 49 {
		usuario.Mas_4_Años = true
	} else {
		usuario.Mas_4_Años = false
	}
	if strNumero[7] == 49 {
		usuario.Mas_1_Local = true
	} else {
		usuario.Mas_1_Local = false
	}
	if strNumero[8] == 49 {
		usuario.Mas_10_Empleados = true
	} else {
		usuario.Mas_10_Empleados = false
	}
	if strNumero[9] == 49 {
		usuario.PagoIgv_6_Meses = true
	} else {
		usuario.PagoIgv_6_Meses = false
	}
	if strNumero[10] == 49 {
		usuario.DeclaronConfidencialPatrimonio = true
	} else {
		usuario.DeclaronConfidencialPatrimonio = false
	}

	getEstado(&usuario)
	RespuestaKnn := knn(&usuario)
	if RespuestaKnn == true {
		respuesta.Mensaje = "Usted esta preaprobado para el bono independiente"
		enviar(valorUsuario)
	} else {
		respuesta.Mensaje = "Usted no esta apto para el bono independiente"
		enviar_Principal(10)
	}
}

func calculaDistanciaAsincrono(chDistancia chan float64, chY chan int, chEstado chan string, chX chan int, x int, y int, p ConsultaBono) {
	absX := math.Abs(float64(x - p.PuntajeEmpresa))
	absY := math.Abs(float64(y - p.PuntajePersonal))
	distancia := math.Sqrt(math.Pow(absX, 2) + math.Pow(absY, 2))

	chDistancia <- distancia
	chY <- p.PuntajeEmpresa
	chX <- p.PuntajePersonal
	chEstado <- p.Estado
}

func knn(usuario *ConsultaBono) bool {
	var knnNodes = [100]knnNode{}
	chDistancia := make(chan float64)
	chY := make(chan int)
	chX := make(chan int)
	chEstado := make(chan string)
	for i := 0; i < 100; i++ {
		go calculaDistanciaAsincrono(chDistancia, chY, chEstado, chX, usuario.PuntajeEmpresa, usuario.PuntajePersonal, Dataset[i])
		knnNodes[i].Distancia = <-chDistancia
		knnNodes[i].y = <-chY
		knnNodes[i].x = <-chX
		knnNodes[i].estado = <-chEstado
	}
	log.Println(knnNodes)
	for i := 1; i < 100; i++ {
		for j := 0; j < 100-i; j++ {
			if knnNodes[j].Distancia > knnNodes[j+1].Distancia {
				knnNodes[j], knnNodes[j+1] = knnNodes[j+1], knnNodes[j]
			}
		}
	}
	log.Println(knnNodes)
	count := 0
	for i := 0; i < 6; i++ {
		if knnNodes[i].estado == "Pre-Aprobado" {
			count++
		}
	}
	if count >= 3 {
		log.Println("Usted esta preaprobado para el bono independiente")
		return true
	} else {
		log.Println("Usted no esta apto para el bono independiente")
		return false
	}
}

func LeerDataSetFromGit() {
	response, err := http.Get("https://raw.githubusercontent.com/CaffoAaron/DataSet-Programaci-n-Concurrente-y-Distribuida/master/bono_Independiente_trabajaperu.csv") //use package "net/http"
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()
	reader := csv.NewReader(response.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		log.Println(nil)
	}
	//fmt.Println(data)

	for i, row := range data {

		Casado, _ := strconv.ParseBool(row[0])
		Dataset[i].Casado = Casado

		Hijos, _ := strconv.ParseBool(row[1])
		Dataset[i].Hijos = Hijos

		CarreraUniversitaria, _ := strconv.ParseBool(row[2])
		Dataset[i].CarreraUniversitaria = CarreraUniversitaria

		CasaPropia, _ := strconv.ParseBool(row[3])
		Dataset[i].CasaPropia = CasaPropia

		OtroPrestamo, _ := strconv.ParseBool(row[4])
		Dataset[i].OtroPrestamo = OtroPrestamo

		Mas_4_Años, _ := strconv.ParseBool(row[5])
		Dataset[i].Mas_4_Años = Mas_4_Años

		Mas_1_Local, _ := strconv.ParseBool(row[6])
		Dataset[i].Mas_1_Local = Mas_1_Local

		Mas_10_Empreados, _ := strconv.ParseBool(row[7])
		Dataset[i].Mas_10_Empleados = Mas_10_Empreados

		PagoIgv_6_Meses, _ := strconv.ParseBool(row[8])
		Dataset[i].PagoIgv_6_Meses = PagoIgv_6_Meses

		DeclaronConfidencialPatrimonio, _ := strconv.ParseBool(row[9])
		Dataset[i].DeclaronConfidencialPatrimonio = DeclaronConfidencialPatrimonio
	}
	for i := 0; i < 1000; i++ {
		getEstado(&Dataset[i])
	}
	log.Println(Dataset)
}

func mostrarDataset(res http.ResponseWriter, req *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	log.Println("Llamada al endpoint /dataset")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	res.Header().Set("Access-Control-Expose-Headers", "Authorization")
	jsonBytes, _ := json.MarshalIndent(Dataset, "", "\t")
	io.WriteString(res, string(jsonBytes))
}

func realizarKnn(res http.ResponseWriter, req *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	log.Println("Llamada al endpoint /knn")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	res.Header().Set("Access-Control-Expose-Headers", "Authorization")
	var usuario = ConsultaBono{}
	var respuesta = Respuesta{}
	temp := "1"
	//body, _ := ioutil.ReadAll(req.Body)
	casado := req.FormValue("casado")
	hijos := req.FormValue("hijos")
	carrera_universitaria := req.FormValue("carrera_universitaria")
	casa_propia := req.FormValue("casa_propia")
	otro_prestamo := req.FormValue("otro_prestamo")

	mas_de_4_Años_como_empresa := req.FormValue("mas_de_4_Años_como_empresa")
	mas_de_1_Local := req.FormValue("mas_de_1_Local")
	mas_de_10_Empleados := req.FormValue("mas_de_10_Empleados")
	Pago_de_Igv_Ultimos_6_Meses := req.FormValue("Pago_de_Igv_Ultimos_6_Meses")
	declaron_confidencial_patrimonio := req.FormValue("declaron_confidencial_patrimonio")
	

	if casado == "No" {
		temp = temp + "0"
		usuario.Casado = false
	} else {
		temp = temp + "1"
		usuario.Casado = true
	}
	if hijos == "No" {
		temp = temp + "0"
		usuario.Hijos = false
	} else {
		temp = temp + "1"
		usuario.Hijos = true

	}
	if carrera_universitaria == "No" {
		temp = temp + "0"
		usuario.CarreraUniversitaria = false

	} else {
		temp = temp + "1"
		usuario.CarreraUniversitaria = true
	}
	if casa_propia == "No" {
		temp = temp + "0"
		usuario.CasaPropia = false
	} else {
		temp = temp + "1"
		usuario.CasaPropia = true
	}
	if otro_prestamo == "No" {
		temp = temp + "0"
		usuario.OtroPrestamo = false
	} else {
		temp = temp + "1"
		usuario.OtroPrestamo = true
	}

	if mas_de_4_Años_como_empresa == "No" {
		temp = temp + "0"
		usuario.Mas_4_Años = false
	} else {
		temp = temp + "1"
		usuario.Mas_4_Años = true
	}
	if mas_de_1_Local == "No" {
		temp = temp + "0"
		usuario.Mas_1_Local = false
	} else {
		temp = temp + "1"
		usuario.Mas_1_Local = true
	}
	if mas_de_10_Empleados == "No" {
		temp = temp + "0"
		usuario.Mas_10_Empleados = false
	} else {
		temp = temp + "1"
		usuario.Mas_10_Empleados = true
	}
	if Pago_de_Igv_Ultimos_6_Meses == "No" {
		temp = temp + "0"
		usuario.PagoIgv_6_Meses = false
	} else {
		temp = temp + "1"
		usuario.PagoIgv_6_Meses = true
	}
	if declaron_confidencial_patrimonio == "No" {
		temp = temp + "0"
		usuario.DeclaronConfidencialPatrimonio = false
	} else {
		temp = temp + "1"
		usuario.DeclaronConfidencialPatrimonio = true
	}

	//log.Println("response Body:", string(body))

	getEstado(&usuario)

	conn, _ := net.Dial("tcp", "localhost:8001")
	defer conn.Close()

	i, _ := strconv.Atoi(temp)
	log.Println(i)

	fmt.Fprintf(conn, "%d\n", i)

	ln, _ := net.Listen("tcp", "localhost:8000")
	defer ln.Close()
	eschucha_funcion = true
	resultado := false
	for eschucha_funcion == true {
		//manejador de conexiones
		conn, _ := ln.Accept()
		resultado = manejador_respueta(conn)
	}
	if resultado == true {
		respuesta.Mensaje = "Usted esta preaprobado para el bono independiente"
		jsonBytes, _ := json.MarshalIndent(respuesta, "", "\t")
		io.WriteString(res, string(jsonBytes))
	} else {
		respuesta.Mensaje = "Usted no esta apto para el bono independiente"
		jsonBytes, _ := json.MarshalIndent(respuesta, "", "\t")
		io.WriteString(res, string(jsonBytes))
	}

}

func handleRequest() {

	http.HandleFunc("/dataset", mostrarDataset)
	http.HandleFunc("/knn", realizarKnn)
	log.Fatal(http.ListenAndServe(":9200", nil))

}

func main() {
	bufferIn := bufio.NewReader(os.Stdin)
	LeerDataSetFromGit()
	//tipo de nodo
	log.Print("Ingrese el tipo de nodo (i:inicio -n:intermedio - f:final): ")
	tipo, _ := bufferIn.ReadString('\n')
	tipo = strings.TrimSpace(tipo)

	if tipo == "i" {
		handleRequest()
	}
	if tipo == "n" {
		//establecer el identificador del host local (IP:puerto)
		log.Print("Ingrese el puerto local: ")
		puerto, _ := bufferIn.ReadString('\n')
		puerto = strings.TrimSpace(puerto)
		localhost := ("localhost:" + puerto)

		//establecer el identificador del host remoto (IP:puerto)
		log.Print("Ingrese el puerto remoto:")
		puerto, _ = bufferIn.ReadString('\n')
		puerto = strings.TrimSpace(puerto)
		remotehost = ("localhost:" + puerto)

		//Cantidad de numero a recibir x nodo

		//canal para el contador
		chCont = make(chan int, 1) //canal asincrono
		chCont <- 0

		//establecer el modo escucha del nodo
		ln, _ := net.Listen("tcp", localhost)
		defer ln.Close()
		for {
			//manejador de conexiones
			conn, _ := ln.Accept()
			go manejador(conn)
		}
	}
	if tipo == "f" {
		//establecer el identificador del host local (IP:puerto)
		log.Print("Ingrese el puerto local: ")
		puerto, _ := bufferIn.ReadString('\n')
		puerto = strings.TrimSpace(puerto)
		localhost := ("localhost:" + puerto)

		//establecer el identificador del host remoto (IP:puerto)

		//Cantidad de numero a recibir x nodo

		//canal para el contador
		chCont = make(chan int, 1) //canal asincrono
		chCont <- 0

		//establecer el modo escucha del nodo
		ln, _ := net.Listen("tcp", localhost)
		defer ln.Close()
		for {
			//manejador de conexiones
			conn, _ := ln.Accept()
			go manejador_fin(conn)
		}
	}

}
