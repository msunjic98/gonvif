package gonvif

import (
	"fmt"
	"log"
	"net"

	"github.com/beevik/guid"
)

func GenerateRequestBuffer() []byte {
	g := guid.New().String()
	messageID := "urn:uuid:" + g
	request := []byte(`
	<Envelope xmlns="http://www.w3.org/2003/05/soap-envelope" xmlns:dn="http://www.onvif.org/ver10/network/wsdl">
		<Header>
			<wsa:MessageID xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing">` + messageID + `</wsa:MessageID>
			<wsa:To xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing">urn:schemas-xmlsoap-org:ws:2005:04:discovery</wsa:To>
			<wsa:Action xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe</wsa:Action> 
		</Header>
		<Body>
			<Probe xmlns="http://schemas.xmlsoap.org/ws/2005/04/discovery" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
				<Types>dn:NetworkVideoTransmitter</Types>
				<Scopes />
			</Probe>
		</Body>
	</Envelope>`)
	return request
}

func SendBufferToUDP(request []byte) {
	conn, err := net.Dial("udp", "239.255.255.250:3702")
	if err != nil {
		fmt.Printf("error %v", err)
	}
	conn.Write(request)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])
}
