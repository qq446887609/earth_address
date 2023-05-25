package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type VendorInfo struct {
	MACAddress  string
	CompanyName string
	Address     string
	ZipCode     string
	Country     string
}

func main() {
	macAddress := "e8:4e:06:86:02:5e"
	vendorInfo := findVendorInfo(macAddress, "mac_address.txt")
	fmt.Println(vendorInfo)
}

func findVendorInfo(macAddress, filePath string) *VendorInfo {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isMACAddress(line) {
			info := extractVendorInfo(line)
			if info.MACAddress == macAddress {
				return info
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func isMACAddress(line string) bool {
	macPattern := `([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`
	match, _ := regexp.MatchString(macPattern, line)
	return match
}

func extractVendorInfo(line string) *VendorInfo {
	vendorInfo := &VendorInfo{}

	macPattern := `([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`
	macRegex := regexp.MustCompile(macPattern)
	vendorInfo.MACAddress = macRegex.FindString(line)

	companyPattern := `\t+(.*?)\s+\(hex\)`
	companyRegex := regexp.MustCompile(companyPattern)
	companyName := companyRegex.FindStringSubmatch(line)
	vendorInfo.CompanyName = strings.TrimSpace(companyName[1])

	addressPattern := `\((.*?)\)`
	addressRegex := regexp.MustCompile(addressPattern)
	address := addressRegex.FindStringSubmatch(line)
	vendorInfo.Address = strings.TrimSpace(address[1])

	zipCodePattern := `\d{5}`
	zipCodeRegex := regexp.MustCompile(zipCodePattern)
	zipCode := zipCodeRegex.FindString(line)
	vendorInfo.ZipCode = zipCode

	countryPattern := `\b[A-Z]{2}\b`
	countryRegex := regexp.MustCompile(countryPattern)
	country := countryRegex.FindString(line)
	vendorInfo.Country = country

	return vendorInfo
}
