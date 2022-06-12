package targets

import "net"

type BypassWebsite struct {
	Host string
	IPs  []net.IPAddr
}

var BypassIPs = []BypassWebsite{
	{
		Host: "12-kanal.ru", IPs: []net.IPAddr{
			*mustResolveIPAddr("109.194.112.233"),
		},
	},
	{
		Host: "360tv.ru", IPs: []net.IPAddr{
			*mustResolveIPAddr("95.213.164.20"),
		},
	},
	{
		Host: "47news.ru", IPs: []net.IPAddr{
			*mustResolveIPAddr("77.244.221.63"),
		},
	},
	{
		Host: "donbasstoday.ru", IPs: []net.IPAddr{
			*mustResolveIPAddr("193.233.15.62"),
		},
	},
	{
		Host: "www.donnews.ru", IPs: []net.IPAddr{
			*mustResolveIPAddr("185.179.190.159"),
		},
	},
	{
		Host: "fine-news.ru", IPs: []net.IPAddr{
			*mustResolveIPAddr("213.139.208.57"),
		},
	},
	{
		Host: "mosregtoday.ru", IPs: []net.IPAddr{
			*mustResolveIPAddr("178.154.213.69"),
		},
	},
	{
		Host: "novorossiia.ru", IPs: []net.IPAddr{
			*mustResolveIPAddr("87.236.16.191"),
		},
	},
	{
		Host: "pizzahut.ru", IPs: []net.IPAddr{
			*mustResolveIPAddr("185.98.84.80"),
		},
	},
	{
		Host: "russia-insider.com", IPs: []net.IPAddr{
			*mustResolveIPAddr("158.69.116.70"),
		},
	},
	{
		Host: "sbarro-pizza.ru", IPs: []net.IPAddr{
			*mustResolveIPAddr("185.165.123.36"),
		},
	},
	{
		Host: "www.voltairenet.org", IPs: []net.IPAddr{
			*mustResolveIPAddr("65.108.13.235"),
		},
	},
	{
		Host: "wpristav.com", IPs: []net.IPAddr{
			*mustResolveIPAddr("193.109.246.53"),
		},
	},
}

func mustResolveIPAddr(address string) (ip *net.IPAddr) {
	ip, err := net.ResolveIPAddr("ip", address)
	if err != nil {
		panic(err)
	}
	return
}
