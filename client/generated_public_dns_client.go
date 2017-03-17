package client

import "github.com/rancher/go-rancher/client"

type RancherDNSClient struct {
    client.RancherClient
    client.RancherBaseClient
	
    ApiVersion ApiVersionOperations
    RootDomainInfo RootDomainInfoOperations
    DnsRecord DnsRecordOperations
}

func constructDNSClient(rancherBaseClient *client.RancherBaseClientImpl) *RancherDNSClient {
	dnsClient := &RancherDNSClient{
		RancherBaseClient: rancherBaseClient,
	}


    dnsClient.ApiVersion = newApiVersionClient(dnsClient)
    dnsClient.RootDomainInfo = newRootDomainInfoClient(dnsClient)
    dnsClient.DnsRecord = newDnsRecordClient(dnsClient) 


	return dnsClient
}

func NewRancherDNSClient(opts *client.ClientOpts) (*RancherDNSClient, error) {
	rancherBaseClient := &client.RancherBaseClientImpl{
		Types: map[string]client.Schema{},
	}
    dnsClient := constructDNSClient(rancherBaseClient)
        
    err := client.SetupRancherBaseClient(rancherBaseClient, opts)
    if err != nil {
        return nil, err
    }

    return dnsClient, nil
}
