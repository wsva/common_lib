package wcl_crypto_test

import (
	"crypto/x509"
	"fmt"
	"net"
	"testing"

	"github.com/wsva/common_lib/wcl_crypto"
)

func TestCA(T *testing.T) {
	rootconfig := wcl_crypto.CertConfig{
		CertConfigBase: wcl_crypto.CertConfigBase{
			CommonName:   "OW",
			Organization: []string{"OW"},
		},
		PublicKeyAlgorithm: x509.ECDSA,
	}
	rootcrt, rootkey, err := wcl_crypto.NewCertificateAuthority(&rootconfig)
	if err != nil {
		fmt.Println("root", err)
		return
	}
	err = wcl_crypto.WriteCertAndKey("certs", "GMKarRoot", rootcrt, rootkey)
	if err != nil {
		fmt.Println("root write", err)
		return
	}
	ccconfig := wcl_crypto.CertConfig{
		CertConfigBase: wcl_crypto.CertConfigBase{
			CommonName: "10.0.0.1",
			AltNames: wcl_crypto.AltNames{
				IPs: []net.IP{net.ParseIP("10.0.0.1")},
			},
			Usages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		},
		PublicKeyAlgorithm: x509.ECDSA,
	}
	cccrt, cckey, err := wcl_crypto.NewCertAndKey(rootcrt, rootkey, &ccconfig)
	if err != nil {
		fmt.Println("control center", err)
		return
	}
	err = wcl_crypto.WriteCertAndKey("./certs", "GMKarControlCenter", cccrt, cckey)
	if err != nil {
		fmt.Println("control center write", err)
		return
	}
	clientconfig := wcl_crypto.CertConfig{
		CertConfigBase: wcl_crypto.CertConfigBase{
			CommonName: "ow",
			AltNames: wcl_crypto.AltNames{
				IPs: []net.IP{net.ParseIP("10.0.0.1")},
			},
			Usages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		},
		PublicKeyAlgorithm: x509.ECDSA,
	}
	clientcrt, clientkey, err := wcl_crypto.NewCertAndKey(rootcrt, rootkey, &clientconfig)
	if err != nil {
		fmt.Println("control center", err)
		return
	}
	err = wcl_crypto.WriteCertAndKey("./certs", "Client10.0.0.1", clientcrt, clientkey)
	if err != nil {
		fmt.Println("control center write", err)
		return
	}
	fmt.Println("success")
}
