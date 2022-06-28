package targets

// Source: https://twitter.com/FedorovMykhailo/status/1497642156076511233

var TargetDNSServers = map[string]struct{}{
	"95.173.148.51:53":  {}, // 95.* => Federal Guard Service of the Russian Federation
	"95.173.148.50:53":  {},
	"178.124.194.92:53": {}, // bf.belta.by
	"86.57.255.226:53":  {}, // gate.belta.by
	"86.57.255.235:53":  {}, // llm2.belta.by
	"93.84.114.38:53":   {}, // ns.belta.by
	"194.54.153.122:53": {}, // ns1.crimea-post.ru
	"194.54.153.125:53": {}, // ns2.crimea-post.ru
	"213.24.76.2:53":    {}, // ns1.fsb.ru
	"82.209.224.62:53":  {}, // ns1.president.gov.by
	"82.209.224.126:53": {}, // ns2.president.gov.by
	"82.209.224.61:53":  {}, // ns3.president.gov.by
	"82.209.224.125:53": {}, // ns4.president.gov.by
	"95.173.128.77:53":  {}, // ns.gov.ru
	// 	"95.173.130.4:53":    {}, // ns1.duma.gov.ru
	// 	"95.173.130.5:53":    {}, // ns2.duma.gov.ru
	"84.47.152.226:53":   {}, // ns.expert.ru
	"93.95.103.178:53":   {}, // ns3.expert.ru
	"185.12.92.62:53":    {}, // ns1.journal-neo.org
	"74.119.194.66:53":   {}, // ns2.journal-neo.org
	"217.69.139.112:53":  {}, // ns1.mail.ru
	"94.100.180.138:53":  {}, // ns2.mail.ru
	"185.30.176.202:53":  {}, // ns3.mail.ru
	"93.157.56.215:53":   {}, // ns2.mid.ru
	"93.157.56.216:53":   {}, // ns3.mid.ru
	"93.157.9.250:53":    {}, // ns1.mid-dnr.su
	"212.66.32.66:53":    {}, // ns2.mid-dnr.su
	"87.242.66.130:53":   {}, // ns1.mil.ru
	"87.242.66.230:53":   {}, // ns2.mil.ru
	"87.242.66.95:53":    {}, // ns4.mil.ru
	"213.243.109.4:53":   {}, // ns.ng.ru
	"92.242.127.42:53":   {}, // ns1.pravdnr.ru
	"31.133.58.15:53":    {}, // ns2.pravdnr.ru
	"81.19.73.8:53":      {}, // ns2.rambler.ru
	"81.19.83.8:53":      {}, // ns3.rambler.ru
	"81.19.73.9:53":      {}, // ns4.rambler.ru
	"81.19.83.9:53":      {}, // ns5.rambler.ru
	"89.253.223.17:53":   {}, // ns.rf-smi.ru
	"89.253.223.16:53":   {}, // ns1.rf-smi.ru
	"81.26.145.150:53":   {}, // ns1.rg.ru
	"194.190.23.69:53":   {}, // ns2.rg.ru
	"194.190.37.16:53":   {}, // ns3.rg.ru
	"194.190.37.15:53":   {}, // ns4.rg.ru
	"194.190.23.68:53":   {}, // ns5.rg.ru
	"195.93.247.59:53":   {}, // ns7.sputniknews.ru
	"195.93.247.60:53":   {}, // ns8.sputniknews.ru
	"178.248.236.20:53":  {}, // ns9.sputniknews.ru
	"178.248.233.32:53":  {}, // ns10.sputniknews.ru
	"185.216.132.201:53": {}, // ns1.syrianfinance.gov.sy
	"194.28.239.11:53":   {}, // ns1.tass.ru
	"194.28.239.12:53":   {}, // ns2.tass.ru
	"77.238.100.195:53":  {}, // ns1.tsargrad.tv
	"188.64.160.163:53":  {}, // ns2.tsargrad.tv
}

// We need reliable IP addresses
// just like we would've been in Russia/Belarus
var ReferenceDNSServersForHTTP = []string{
	// https://dns.yandex.com/
	"77.88.8.1:53",
	"77.88.8.2:53",
	"77.88.8.3:53",
	"77.88.8.7:53",
	"77.88.8.8:53",
	"77.88.8.88:53",
}
