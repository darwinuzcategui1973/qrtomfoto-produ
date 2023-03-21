package conformacion

import (
	"fmt"
	"net"
	"time"
	"strconv"
	"reflect"

	_ "embed"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed conformacion.yaml
var licenciaFile []byte

/*
sistema: Backend de proudctos
version: 1.00.00
status: Licenciado/demo
nombre: Nombre de Empresa
rif: J-11642590-0
usuario: 10
views: 100
tiempov: 1d
*/
type LicenciaConfig struct {
	Sistema string `yaml:"sistema"`
	Version string `yaml:"version"`
	Status  string `yaml:"status"`
	Nombre  string `yaml:"nombre"`
	Rif     string `yaml:"rif"`
	Usuario int    `yaml:"usuario"`
	Views   int    `yaml:"views"`
	Tiempov string `yaml:"tiempov"`
}

/*
	type Settings struct {
		Port string         `yaml:"port"`
		DB   DatabaseConfig `yaml:"database"`
	}
*/
func New() (*LicenciaConfig, error) {
	var lic LicenciaConfig

	err := yaml.Unmarshal(licenciaFile, &lic)
	if err != nil {
		return nil, err
	}

	return &lic, nil
}

func GetMac() string {

	//----------------------
	// Get the local machine IP address
	// https://www.socketloop.com/tutorials/golang-how-do-I-get-the-local-ip-non-loopback-address
	//----------------------

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	var currentIP, currentNetworkHardwareName string

	for _, address := range addrs {

		// check the address type and if it is not a loopback the display it
		// = GET LOCAL IP ADDRESS
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//fmt.Println("Current IP address : ", ipnet.IP.String())
				currentIP = ipnet.IP.String()

			}
		}
	}

	//fmt.Println("------------------------------")
	//fmt.Println("We want the interface name that has the current IP address")
	//fmt.Println("MUST NOT be binded to 127.0.0.1 ")
	//fmt.Println("------------------------------")

	// get all the system's or local machine's network interfaces

	interfaces, _ := net.Interfaces()
	for _, interf := range interfaces {

		if addrs, err := interf.Addrs(); err == nil {
			for index, addr := range addrs {
				fmt.Println("[", index, "]", interf.Name, ">", addr)

				// only interested in the name with current IP address
				if strings.Contains(addr.String(), currentIP) {
					//fmt.Println("Use name : ", interf.Name)
					currentNetworkHardwareName = interf.Name
				}
			}
		}
	}

	fmt.Println("------------------------------")

	// extract the hardware information base on the interface name
	// capture above
	netInterface, err := net.InterfaceByName(currentNetworkHardwareName)

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	name := netInterface.Name
	macAddress := netInterface.HardwareAddr

	fmt.Println("Hardware name : ", name)
	fmt.Println("MAC address : ", macAddress)

	// verify if the MAC address can be parsed properly
	hwAddr, err := net.ParseMAC(macAddress.String())

	if err != nil {
		fmt.Println("No able to parse MAC address : ", err)
		//os.Exit(-1)
		return err.Error()
	}

	fmt.Printf("Physical hardware address : %s \n", hwAddr.String())
	return hwAddr.String()

}

func GetSistema() string {
	var lic LicenciaConfig

	/*

		var lic LicenciaConfig

		err := yaml.Unmarshal(licenciaFile, &lic)
		if err != nil {
			return nil, err
		}

		return &lic, nil


	*/
	err := yaml.Unmarshal(licenciaFile, &lic)
	if err != nil {
		return err.Error()
	}

	//fmt.Println("---------------sistema : ", lic.Sistema)
	//fmt.Println("---------------usuario : ", lic.Nombre)

	return lic.Sistema
}

func GetDatos() *LicenciaConfig {
	var lic LicenciaConfig

	err := yaml.Unmarshal(licenciaFile, &lic)
	if err != nil {
		return nil
	}

	//fmt.Println("**********sistema : ", lic.Sistema)
	//fmt.Println("************usuario : ", lic.Nombre)

	return &lic
}

func GetFecha() time.Time {

	intDia := 0
	fecha := time.Now()
	tiempov:= GetDatos().Tiempov
	
	intDia, err := strconv.Atoi(tiempov)

	if err != nil {
		fmt.Println(intDia, err, reflect.TypeOf(intDia))
		intDia = 30
	}

    fechaVenc:= fecha.AddDate(0, 0, intDia) 

	return fechaVenc
}

