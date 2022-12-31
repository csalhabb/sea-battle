package ip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type IP struct {
	ip   string
	port uint16
}

// SplitIpAndPort This function split an ip "192.168.0.1:8080" to a string "192.168.0.1" and a port 8080 as uint16.
func SplitIpAndPort(str string) (string, uint16) {
	split := strings.Split(str, ":")
	ip, port := split[0], split[1]

	ui16, err := strconv.ParseUint(port, 10, 64)
	ui := uint16(ui16)

	if err != nil {
		panic(err)
	}

	return ip, ui

}

// This function add an association between a provided IP and a provided username.
func addAlias(aliases *map[string]IP, ip string, username string) {
	realIp, port := SplitIpAndPort(ip)
	ipStruct := IP{
		ip:   realIp,
		port: port,
	}
	(*aliases)[username] = ipStruct
}

// This function displays all the associations betweens IP and usernames.
func displayAliases(aliases *map[string]IP) {
	for key, value := range *aliases {
		fmt.Printf("%s (%s:%d)\n", key, value.ip, value.port)
	}
}

// This function displays the associated IP of the username provided.
func displayAlias(aliases *map[string]IP, username string) {
	for key, value := range *aliases {
		if key == username {
			fmt.Printf("%s (%s:%d)\n", key, value.ip, value.port)
		}
	}
}

// This function remove the associated IP of the username provided.
func removeAlias(aliases *map[string]IP, username string) {
	for key, _ := range *aliases {
		if key == username {
			delete(*aliases, username)
			fmt.Println(username + " has been deleted.")
		}
	}
}

// This function returns the IP of a provided username, returning IP and PORT.
func getIpOf(username string, aliases *map[string]IP) (string, uint16) {
	for key, value := range *aliases {
		if key == username {
			return value.ip, value.port
		}
	}
	return "", 0
}

type User struct {
	Username string
	Ip       string
	Port     uint16
}

// This function allows to store every alias in a json file
func SaveAlias(aliases *map[string]IP) {

	user := []User{}
	for key, value := range *aliases {
		user_1 := User{Username: key, Ip: value.ip, Port: value.port}
		user = append(user, user_1)
	}
	//package this data as json data
	finalJson, err := json.MarshalIndent(user, "", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
	//fmt.Println(user)
	_ = ioutil.WriteFile("alias.json", finalJson, 0644)

}

func ReceiveAlias(aliases *map[string]IP) {
	users := []User{}
	file, _ := os.ReadFile("alias.json")
	_ = json.Unmarshal(file, &users)
	for indexUser := range users {
		ip := users[indexUser].Ip
		port := users[indexUser].Port
		ipStruct := IP{
			ip:   ip,
			port: port,
		}
		(*aliases)[users[indexUser].Username] = ipStruct
	}
}

/*
func testAliases(aliases *map[string]IP) {
	addAlias(aliases, "192.168.0.1:55542", "Noam")
	i, p := getIpOff("Noam", aliases)
	fmt.Printf("%s:%d\n", i, p)
	displayAliases(aliases)
	displayAlias(aliases, "Noam")
	removeAlias(aliases, "Noam")
}
*/

func GetAlias() map[string]IP {
	aliases := make(map[string]IP)
	//addAlias(&aliases, "1.1.1.1:1234", "charbel")
	//addAlias(&aliases, "1.1.0.1:1274", "thibault")
	//addAlias(&aliases, "1.10.1.1:1284", "noam")
	//SaveAlias(&aliases)
	return aliases
}
