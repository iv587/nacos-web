package nacos

import (
	"nacos-web/config"
	"net/http/httputil"
	"net/url"
	"sync"
)

var Proxy = struct {
	sync.RWMutex
	index  int
	proxys []*httputil.ReverseProxy
}{
	index:  0,
	proxys: make([]*httputil.ReverseProxy, 0, 0),
}

func CreateProxy(endpoint []config.NacosEndpoint) error {
	for _, e := range endpoint {
		remote, err := url.Parse(e.Addr)
		if err != nil {
			return err
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		Proxy.proxys = append(Proxy.proxys, proxy)
	}
	return nil
}

func balanceProxy() *httputil.ReverseProxy {
	Proxy.Lock()
	defer Proxy.Unlock()
	if Proxy.index > 100 {
		Proxy.index = Proxy.index - 100
	}
	curr := Proxy.index % len(Proxy.proxys)
	Proxy.index = Proxy.index + 1
	return Proxy.proxys[curr]
}
