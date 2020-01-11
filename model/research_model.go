package model

import (
	"fmt"
	"time"
)

// API Explain Strings
var (
	IpNsZoneCountType = "ip_ns_zone_count"
	ActiveIPsType     = "active_ips"
)

type ResearchIpNsZoneCount struct {
	Type         *string             `json:"type"`
	Link         string              `json:"link"`
	IP           string              `json:"ip"`
	ZoneNSCounts []ResearchZoneCount `json:"zone_counts"`
}

type ResearchZoneCount struct {
	Zone    string  `json:"zone"`
	Count   int64   `json:"count"`
	Percent float64 `json:"percent"`
}

// ActiveIPs  Struct that lists addresses for a given date
type ActiveIPs struct {
	Type    *string   `json:"type"`
	Link    string    `json:"link"`
	Date    time.Time `json:"date"`
	IPv4IPs []string  `json:"ipv4_ips"`
	IPv6IPs []string  `json:"ipv6_ips"`
}

// GenerateMetaData generates metadata recursively for ActiveIPs API
func (aip *ActiveIPs) GenerateMetaData() {
	aip.Type = &ActiveIPsType
	y, m, d := aip.Date.Date()
	aip.Link = fmt.Sprintf("/research/active_ips/%04d-%02d-%02d", y, m, d)
}

// GenerateMetaData generates metadata recursively of member models
func (ipzc *ResearchIpNsZoneCount) GenerateMetaData() {
	ipzc.Type = &IpNsZoneCountType
	ipzc.Link = "/research/ipnszonecount/" + ipzc.IP
}
